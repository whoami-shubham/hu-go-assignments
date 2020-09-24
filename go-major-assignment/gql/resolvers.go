package gql

import (
	"errors"
	"go-major-assignment/client"
	"go-major-assignment/model"
	"log"

	"github.com/graphql-go/graphql"
)

//QueryResolver ...
type QueryResolver struct{}

// MutationResolver ...
type MutationResolver struct{}

// CreateUser ...
func (resolver *MutationResolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)
	password := p.Args["password"].(string)
	email := p.Args["email"].(string)
	role := p.Args["role"].(string)

	_, err := db.Exec(model.InsertQuery("users", model.User{}), name, password, email, role)
	if err != nil {
		return model.User{}, err
	}
	return model.User{Name: name, Password: password, Email: email, Role: role}, nil
}

// CreateIssue ...
func (resolver *MutationResolver) CreateIssue(p graphql.ResolveParams) (interface{}, error) {
	title := p.Args["title"].(string)
	description := p.Args["description"].(string)
	issueType := p.Args["type"].(string)
	assignee := p.Args["assignee"].(int)
	reporter := p.Args["reporter"].(int)
	status := p.Args["status"].(string)
	project := p.Args["project"].(int)

	return model.Issue{Title: title, Description: description, Type: issueType, Assignee: assignee, Reporter: reporter, Status: status, Project: project}, nil
}

// CreateProject ...
func (resolver *MutationResolver) CreateProject(p graphql.ResolveParams) (interface{}, error) {

	// only admin can create project
	if loggedUser.Role != "ADMIN" {
		return DbProject{}, errors.New("You are not authorized to create project")
	}

	name := p.Args["name"].(string)
	createdBy := p.Args["created_by"].(int)
	_, err := db.Exec("INSERT INTO projects(name,created_by) VALUES($1,$2) ", name, createdBy)
	if err != nil {
		log.Println(err)
		return DbProject{}, err
	}
	return DbProject{Name: name, CreatedBy: createdBy}, nil
}

// UpdateIssue ...
func (resolver *MutationResolver) UpdateIssue(p graphql.ResolveParams) (interface{}, error) {

	// only PM or assignee can update

	var data DbIssue
	var id, assignee int
	var title, description, issueType, status string
	var issue DbIssue
	id = p.Args["id"].(int)
	row := db.QueryRow("select * from issues where id=$1;", id)
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
	if err != nil {
		log.Println(err)
		return data, err
	}
	//for authorization
	if loggedUser.Role != "PM" && loggedUser.ID != issue.Assignee && loggedUser.Role != "ADMIN" {
		return data, errors.New("You are not authorized to update issue")
	}
	if p.Args["title"] == nil || p.Args["title"] == issue.Title {
		title = issue.Title
	} else {
		title = p.Args["title"].(string)
		EventLogger(issue.Title, title, "title", id)
		client.SendUpdate(db, id, title, issue.Assignee, issue.Reporter)
	}
	if p.Args["description"] == nil || p.Args["description"] == issue.Description {
		description = issue.Description
	} else {
		description = p.Args["description"].(string)
		EventLogger(issue.Description, description, "description", id)
		client.SendUpdate(db, id, title, issue.Assignee, issue.Reporter)
	}
	if p.Args["type"] == nil || p.Args["type"] == issue.Type {
		issueType = issue.Type
	} else {
		issueType = p.Args["type"].(string)
		EventLogger(issue.Type, issueType, "type", id)
	}
	if p.Args["assignee"] == nil || p.Args["assignee"] == issue.Assignee {
		assignee = issue.Assignee
	} else {
		assignee = p.Args["assignee"].(int)
		EventLogger(issue.Assignee, assignee, "assignee", id)
		client.SendUpdate(db, id, title, issue.Assignee, issue.Reporter)
		client.SendUpdate(db, id, title, assignee, issue.Reporter)
	}
	if p.Args["status"] == nil || p.Args["status"] == issue.Status {
		status = issue.Status
	} else {
		status = p.Args["status"].(string)
		EventLogger(issue.Status, status, "status", id)
	}

	_, err = db.Exec("UPDATE issues set title=$1,description=$2,type=$3,assignee=$4,status=$5 where id=$6", title, description, issueType, assignee, status, id)
	if err != nil {
		log.Println(err)
		return data, err
	}
	data = DbIssue{ID: id, Title: title, Description: description, Type: issueType, Assignee: assignee, Status: status, Reporter: issue.Reporter}
	return data, nil

}

// CreateComment ...
func (resolver *MutationResolver) CreateComment(p graphql.ResolveParams) (interface{}, error) {
	text := p.Args["text"].(string)
	author := p.Args["author"].(int)
	issueID := p.Args["issue_id"].(int)
	_, err := db.Exec("INSERT INTO comments (text,author,issue_id) VALUES($1,$2,$3);", text, author, issueID)
	if err != nil {
		log.Println(err)
		return DbComment{}, err
	}
	return DbComment{Text: text, Author: author, IssueID: issueID}, nil
}

// UpdateComment ...
func (resolver *MutationResolver) UpdateComment(p graphql.ResolveParams) (interface{}, error) {
	text := p.Args["text"].(string)
	id := p.Args["id"].(int)
	var comment DbComment
	row := db.QueryRow("select * from comments where id=$1;", id)
	err := row.Scan(&comment.ID, &comment.Author, &comment.Text, &comment.IssueID, &comment.CreatedOn, &comment.UpdatedOn)
	if err != nil {
		log.Println(err)
		return comment, err
	}

	_, err = db.Exec("UPDATE comments set text=$1 where id=$2", text, id)
	if err != nil {
		log.Println(err)
		return comment, err
	}
	var assignee, reporter int
	var title string
	row = db.QueryRow("select title,assignee,reporter from issues where id=$1;", comment.IssueID)
	err = row.Scan(&title, &assignee, &reporter)
	if err != nil {
		log.Println(err)
		return comment, err
	}
	client.SendUpdate(db, id, title, assignee, reporter)
	comment.Text = text
	return comment, nil
}

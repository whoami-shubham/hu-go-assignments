package gql

import (
	"database/sql"
	"log"
	"time"

	"github.com/graphql-go/graphql"
)

// GetAllIssues ...
func (resolver *QueryResolver) GetAllIssues(p graphql.ResolveParams) (interface{}, error) {
	data := []gqlReturnedIssue{}
	rows, err := db.Query("select * from issues;")
	if err != nil {
		log.Println(err)
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var issue DbIssue

		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		assignee, reporter, project, comments, issueLogs, errr := getIssueSubFeilds(issue)
		if errr != nil && errr != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		data = append(data, gqlReturnedIssue{ID: issue.ID, Title: issue.Title, Description: issue.Description, Type: issue.Type, Assignee: assignee, Reporter: reporter, Status: issue.Status, Project: project, Comments: comments, Logs: issueLogs})

	}

	return data, nil
}

//GetIssueByID ...
func (resolver *QueryResolver) GetIssueByID(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["issueID"].(int)
	issue := DbIssue{}
	data := gqlReturnedIssue{}
	row := db.QueryRow("select * from issues where id=$1;", id)
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
	if err != nil {
		log.Println(err)
		return data, err
	}
	assignee, reporter, project, comments, issueLogs, errr := getIssueSubFeilds(issue)
	if errr != nil && errr != sql.ErrNoRows {
		log.Println(err)
		return data, err
	}
	data = gqlReturnedIssue{ID: issue.ID, Title: issue.Title, Description: issue.Description, Type: issue.Type, Assignee: assignee, Reporter: reporter, Status: issue.Status, Project: project, Comments: comments, Logs: issueLogs}

	return data, nil
}

// GetIssueByTitle ...
func (resolver *QueryResolver) GetIssueByTitle(p graphql.ResolveParams) (interface{}, error) {
	title := p.Args["title"].(string)
	data := []gqlReturnedIssue{}
	rows, err := db.Query("select * from issues where title=$1;", title)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var issue DbIssue

		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		assignee, reporter, project, comments, issueLogs, errr := getIssueSubFeilds(issue)
		if errr != nil && errr != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		data = append(data, gqlReturnedIssue{ID: issue.ID, Title: issue.Title, Description: issue.Description, Type: issue.Type, Assignee: assignee, Reporter: reporter, Status: issue.Status, Project: project, Comments: comments, Logs: issueLogs})

	}
	return data, nil

}

//GetIssuesByByProjectID ...
func (resolver *QueryResolver) GetIssuesByByProjectID(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["projectID"].(int)
	data := []gqlReturnedIssue{}
	rows, err := db.Query("select * from issues where project=$1;", id)
	if err != nil {
		log.Println(err)
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var issue DbIssue
		var assignee, reporter DbUser
		var project DbProject
		var comments []DbComment
		var issueLogs []DbIssueLog
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		assignee, reporter, project, comments, issueLogs, err = getIssueSubFeilds(issue)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}

		data = append(data, gqlReturnedIssue{ID: issue.ID, Title: issue.Title, Description: issue.Description, Type: issue.Type, Assignee: assignee, Reporter: reporter, Status: issue.Status, Project: project, Comments: comments, Logs: issueLogs})

	}

	return data, nil

}

func getIssueSubFeilds(issue DbIssue) (DbUser, DbUser, DbProject, []DbComment, []DbIssueLog, error) {
	var assignee, reporter DbUser
	var project DbProject
	var comments []DbComment
	var issueLogs []DbIssueLog
	assigneeID := issue.Assignee
	reporterID := issue.Reporter
	projectID := issue.Project
	row := db.QueryRow("select * from users where id=$1;", assigneeID)
	err := row.Scan(&assignee.ID, &assignee.Name, &assignee.Password, &assignee.Email, &assignee.Role, &assignee.CreatedOn, &assignee.UpdatedOn)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return assignee, reporter, project, comments, issueLogs, err
	}

	row = db.QueryRow("select * from users where id=$1;", reporterID)
	err = row.Scan(&reporter.ID, &reporter.Name, &reporter.Password, &reporter.Email, &reporter.Role, &reporter.CreatedOn, &reporter.UpdatedOn)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return assignee, reporter, project, comments, issueLogs, err
	}

	row = db.QueryRow("select * from projects where id=$1;", projectID)
	err = row.Scan(&project.ID, &project.Name, &project.CreatedBy, &project.CreatedOn, &project.UpdatedOn)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return assignee, reporter, project, comments, issueLogs, err
	}

	rows, errr := db.Query("select * from comments where issue_id=$1;", issue.ID)
	if errr != nil && errr != sql.ErrNoRows {
		log.Println(errr)
		return assignee, reporter, project, comments, issueLogs, err
	}
	for rows.Next() {
		var comment DbComment
		err = rows.Scan(&comment.ID, &comment.Author, &comment.Text, &comment.IssueID, &comment.CreatedOn, &comment.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return assignee, reporter, project, comments, issueLogs, err
		}
		comments = append(comments, comment)
	}
	rows, err = db.Query("select * from issues_log where issue_id=$1;", issue.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return assignee, reporter, project, comments, issueLogs, err
	}
	for rows.Next() {
		var issueLog DbIssueLog
		err = rows.Scan(&issueLog.ID, &issueLog.UpdatedFeild, &issueLog.PreviousValue, &issueLog.NewValue, &issueLog.IssueID, &issueLog.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return assignee, reporter, project, comments, issueLogs, err
		}
		issueLogs = append(issueLogs, issueLog)
	}
	return assignee, reporter, project, comments, issueLogs, err
}

// EventLogger ...
func EventLogger(prevValue, newValue interface{}, feildName string, issueID int) {
	currentTime := time.Now()
	_, err := db.Exec("INSERT INTO issues_log (updated_feild,previous_value,new_value,issue_id,updated_on) VALUES($1,$2,$3,$4,$5);", feildName, prevValue, newValue, issueID, currentTime.Format("2006-01-02 15:04:05.000000"))
	if err != nil {
		log.Println(err)
	}
}

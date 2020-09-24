package gql

import (
	"database/sql"
	"log"

	"github.com/graphql-go/graphql"
)

//GetAllProjects ...
func (resolver *QueryResolver) GetAllProjects(p graphql.ResolveParams) (interface{}, error) {
	data := []gqlReturnedProjectWithIssue{}
	rows, err := db.Query("select * from projects;")
	if err != nil {
		log.Println(err)
		return []DbProject{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var project DbProject
		err := rows.Scan(&project.ID, &project.Name, &project.CreatedBy, &project.CreatedOn, &project.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		var issues []gqlReturnedIssue
		issues, err = getProjectSubFeilds(project.ID)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return data, err
		}
		data = append(data, gqlReturnedProjectWithIssue{ID: project.ID, Name: project.Name, CreatedBy: project.CreatedBy, Issues: issues})

	}
	return data, nil
}

//GetProjectByID ...
func (resolver *QueryResolver) GetProjectByID(p graphql.ResolveParams) (interface{}, error) {
	data := gqlReturnedProjectWithIssue{}
	id := p.Args["projectID"].(int)
	var project DbProject
	var issues []gqlReturnedIssue
	row := db.QueryRow("select * from projects where id=$1;", id)
	err := row.Scan(&project.ID, &project.Name, &project.CreatedBy, &project.CreatedOn, &project.UpdatedOn)
	if err != nil {
		log.Println(err)
		return data, err
	}
	issues, err = getProjectSubFeilds(project.ID)
	data = gqlReturnedProjectWithIssue{ID: project.ID, Name: project.Name, CreatedBy: project.CreatedBy, Issues: issues}
	return data, nil
}

func getProjectSubFeilds(id int) ([]gqlReturnedIssue, error) {
	issues := []gqlReturnedIssue{}
	rows, err := db.Query("select * from issues where project=$1;", id)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return issues, err
	}
	defer rows.Close()
	for rows.Next() {
		var issue DbIssue
		var assignee, reporter DbUser
		var subProject DbProject
		var comments []DbComment
		var issueLogs []DbIssueLog
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Type, &issue.Assignee, &issue.Reporter, &issue.Status, &issue.Project, &issue.CreatedOn, &issue.UpdatedOn)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return issues, err
		}
		assignee, reporter, subProject, comments, issueLogs, err = getIssueSubFeilds(issue)
		if err != nil && err != sql.ErrNoRows {
			log.Println(err)
			return issues, err
		}

		issues = append(issues, gqlReturnedIssue{ID: issue.ID, Title: issue.Title, Description: issue.Description, Type: issue.Type, Assignee: assignee, Reporter: reporter, Status: issue.Status, Project: subProject, Comments: comments, Logs: issueLogs})

	}
	return issues, nil
}

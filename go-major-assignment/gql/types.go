package gql

import (
	"github.com/graphql-go/graphql"
)

// User ...
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"role": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

// Project ...
var Project = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"created_by": &graphql.Field{
				Type: User,
			},
		},
	},
)

// Comment ...
var Comment = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"issue_id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	},
)

// IssueLog ...
var IssueLog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IssueLog",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"updated_feild": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"previous_value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"new_value": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"issue_id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"updated_on": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

// Issue ...
var Issue = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Issue",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"status": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"assignee": &graphql.Field{
				Type: User,
			},
			"reporter": &graphql.Field{
				Type: User,
			},

			"project": &graphql.Field{
				Type: Project,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(Comment),
			},
			"logs": &graphql.Field{
				Type: graphql.NewList(IssueLog),
			},
		},
	},
)

// ProjectWithIssues ...
var ProjectWithIssues = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ProjectWithIssue",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"created_by": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"issues": &graphql.Field{
				Type: graphql.NewList(Issue),
			},
		},
	},
)

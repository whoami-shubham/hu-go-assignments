package gql

import "github.com/graphql-go/graphql"

// RootQuery ...
type RootQuery struct {
	Query *graphql.Object
}

// NewRootQuery ...
func NewRootQuery() *RootQuery {
	resolver := QueryResolver{}

	rootQuery := RootQuery{
		Query: graphql.NewObject(
			graphql.ObjectConfig{

				Name: "Query",
				Fields: graphql.Fields{

					"users": &graphql.Field{
						Type:        graphql.NewList(User),
						Description: "Gets all users",
						Resolve:     resolver.GetAllUsers,
					},
					"projects": &graphql.Field{
						Type:        graphql.NewList(ProjectWithIssues),
						Description: "Gets all projects",
						Resolve:     resolver.GetAllProjects,
					},
					"issues": &graphql.Field{
						Type:        graphql.NewList(Issue),
						Description: "Gets all issues",
						Resolve:     resolver.GetAllIssues,
					},

					"issueByID": &graphql.Field{
						Type:        Issue,
						Description: "Gets issue by id",
						Args: graphql.FieldConfigArgument{
							"issueID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.GetIssueByID,
					},
					"issueByTitle": &graphql.Field{
						Type:        graphql.NewList(Issue),
						Description: "Gets issues by Title",
						Args: graphql.FieldConfigArgument{
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.GetIssueByTitle,
					},
					"projectByID": &graphql.Field{
						Type:        ProjectWithIssues,
						Description: "Gets project by id",
						Args: graphql.FieldConfigArgument{
							"projectID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.GetProjectByID,
					},
					"GetIssuesByByProjectID": &graphql.Field{
						Type:        graphql.NewList(Issue),
						Description: "Gets issue by project id",
						Args: graphql.FieldConfigArgument{
							"projectID": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.GetIssuesByByProjectID,
					},
				},
			},
		),
	}

	return &rootQuery
}

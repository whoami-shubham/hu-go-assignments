package gql

import "github.com/graphql-go/graphql"

// RootMutation ...
type RootMutation struct {
	Mutation *graphql.Object
}

// NewRootMutation ...
func NewRootMutation() *RootMutation {
	mutationResolver := MutationResolver{}

	rootMutation := RootMutation{
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{

				Name: "Mutation",
				Fields: graphql.Fields{

					"createUser": &graphql.Field{
						Type:        User,
						Description: "create a users",
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"password": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"email": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"role": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: mutationResolver.CreateUser,
					},
					"createProject": &graphql.Field{
						Type:        Project,
						Description: "Gets all projects",
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"created_by": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: mutationResolver.CreateProject,
					},
					"createIssue": &graphql.Field{
						Type:        Issue,
						Description: "Gets all issues",
						Args: graphql.FieldConfigArgument{
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"type": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"assignee": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"reporter": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"status": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"project": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: mutationResolver.CreateIssue,
					},
					"updateIssue": &graphql.Field{
						Type:        Issue,
						Description: "update issues",
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"title": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"description": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"type": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"assignee": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"status": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: mutationResolver.UpdateIssue,
					},
					"createComment": &graphql.Field{
						Type:        Comment,
						Description: "create a comment",
						Args: graphql.FieldConfigArgument{
							"issue_id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"text": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"author": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: mutationResolver.CreateComment,
					},
					"updateComment": &graphql.Field{
						Type:        Comment,
						Description: "update a comment",
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"text": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: mutationResolver.UpdateComment,
					},
				},
			},
		),
	}

	return &rootMutation
}

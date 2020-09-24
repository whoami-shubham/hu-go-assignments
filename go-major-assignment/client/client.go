package client

import (
	"context"
	"database/sql"
	"go-major-assignment/email"
	"go-major-assignment/model"
	"log"
	"time"

	"google.golang.org/grpc"
)

// SendReport ...
func SendReport(db *sql.DB) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := email.NewEmailClient(conn)

	issues, err := model.GetOpenIssues(db)
	for _, issue := range issues {
		response, err := client.Report(context.Background(), &email.ReportRequest{Email: issue.Email, Report: &email.Report{Id: issue.ID, Title: issue.Title, Description: issue.Description, CreatedOn: issue.CreatedOn, UpdatedOn: issue.UpdatedOn, ReporterName: issue.Reporter}})
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		log.Printf("Response from server: %s", response.Res)
	}

}

// Cron ...
func Cron(db *sql.DB) {

	for {
		t := time.Now()
		// hr := t.Hour() // hr = 22 and mints = 0
		mints := t.Minute()
		if mints%2 == 0 {
			SendReport(db)
			time.Sleep(61 * time.Second)
		}

	}
}

// SendUpdate ...
func SendUpdate(db *sql.DB, issueID int, title string, assignee, reporter int) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := email.NewEmailClient(conn)
	message := "You've got an update from issue " + title

	watchers, err := model.GetWatchers(db, issueID, assignee, reporter)
	for _, watcher := range watchers {
		response, err := client.Update(context.Background(), &email.UpdateRequest{Email: watcher, Title: message})
		if err != nil {
			log.Fatalf("Something went wrong: %s", err)
		}
		log.Printf("Response from server: %s", response.Res)
	}

}

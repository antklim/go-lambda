package main

import (
	"context"
	"log"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	sess *session.Session
)

func init() {
	log.Print("init...")
	sess = session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-southeast-2")}))
}

func handler(ctx context.Context, event interface{}) (string, error) {
	log.Print("handler...")
	log.Print("event", event)

	return "OK", nil
}

func main() {
	log.Print("main...")
	runtime.Start(handler)
}

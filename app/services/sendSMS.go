package services

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendSMS() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_SES_ACCESS_KEY"),
			os.Getenv("AWS_SES_SECRET_KEY"),
			"",
		),
	})

	svc := sns.New(sess)
	params := &sns.PublishInput{
		PhoneNumber: aws.String("+84938366486"),
		Message:     aws.String("hello"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		fmt.Println("sms error", err)
		log.Println(err.Error())
	}

	fmt.Println("resp", resp) // print the response data.
}

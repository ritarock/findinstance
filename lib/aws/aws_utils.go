package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Session(profile string) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func DescribeInstances(session *session.Session) *ec2.DescribeInstancesOutput {
	svc := ec2.New(session)
	instances, err := svc.DescribeInstances(nil)
	if err != nil {
		log.Fatal(err)
	}
	return instances
}

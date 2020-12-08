package action

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/ritarock/findinstance/lib/aws"
)

func Run(profile, instanceName string) {
	sess := aws.Session(profile)
	instances := aws.DescribeInstances(sess)
	foundInstance := find(instances, instanceName)
	if foundInstance == nil {
		fmt.Println("Instances not found")
	} else {
		fmt.Printf(" [LaunchTime]\n")
		fmt.Printf("  %v\n", *foundInstance.LaunchTime)
		fmt.Printf(" [State]\n")
		fmt.Printf("  %v\n", *foundInstance.State.Name)
		fmt.Printf(" [PrivateIpAddress]\n")
		fmt.Printf("  %v\n", *foundInstance.NetworkInterfaces[0].PrivateIpAddress)
		fmt.Printf(" [AvailabilityZone]\n")
		fmt.Printf("  %v\n", *foundInstance.Placement.AvailabilityZone)
		fmt.Printf(" [TAGS]\n")
		for _, tag := range foundInstance.Tags {
			fmt.Printf("  [%v]\n", *tag.Key)
			fmt.Printf("   %v\n", *tag.Value)
		}
	}
}

func find(instances *ec2.DescribeInstancesOutput, instanceName string) *ec2.Instance {
	var result *ec2.Instance
	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			for _, tag := range instance.Tags {
				if *tag.Value == instanceName {
					result = instance
				}
			}
		}
	}
	return result
}

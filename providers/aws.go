package providers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func AwsSession(r string) *ec2.EC2 {

	sess, err := session.NewSession(&aws.Config{Region: aws.String(r)})
	if err != nil {
		fmt.Println(err)
	}

	svc := ec2.New(sess)

	return svc
}

func AwsListAll(a *ec2.DescribeImagesOutput, s *ec2.EC2) []string {

	all := []string{}

	for _, v := range a.Images {
		inu := v.ImageId

		all = append(all, aws.StringValue(inu))
	}
	return all
}

func AwsListNotUsed(a *ec2.DescribeImagesOutput, s *ec2.EC2) ([]string, []string) {

	all := []string{}
	used := []string{}

	for _, v := range a.Images {
		inu := v.ImageId

		ec2f := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("image-id"),
					Values: []*string{aws.String(*inu)},
				},
			},
		}

		all = append(all, aws.StringValue(inu))

		r, err := s.DescribeInstances(ec2f)
		for _, res := range r.Reservations {
			for range res.Instances {
				used = append(used, string(aws.StringValue(*&inu)))
			}
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	return all, used
}

package providers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// AwsSession create a new instance of EC2 client
func AwsSession(r *string) *ec2.EC2 {
	var config = &aws.Config{
		Region: r,
	}

	sess, err := session.NewSession(config)
	if err != nil {
		fmt.Println(err)
	}

	return ec2.New(sess)
}

// AwsListAll list all available images
func AwsListAll(a *ec2.DescribeImagesOutput) (all []string) {
	for _, v := range a.Images {
		all = append(all, aws.StringValue(v.ImageId))
	}
	return all
}

// AwsListNotUsed list all images that are not being used
func AwsListNotUsed(a *ec2.DescribeImagesOutput, s *ec2.EC2) (all[]string, used[]string) {
	for _, v := range a.Images {
		ec2f := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name:   aws.String("image-id"),
					Values: []*string{aws.String(*v.ImageId)},
				},
			},
		}

		all = append(all, aws.StringValue(v.ImageId))

		r, err := s.DescribeInstances(ec2f)
		if err != nil {
			fmt.Println(err)
		}

		for _, res := range r.Reservations {
			for range res.Instances {
				used = append(used, aws.StringValue(v.ImageId))
			}
		}
	}

	return all, used
}

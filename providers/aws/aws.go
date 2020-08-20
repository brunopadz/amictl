package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NewSession(r string) (*ec2.EC2, error) {
	var config = &aws.Config{
		Region: aws.String(r),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	return ec2.New(sess), nil
}

func ListAll(output *ec2.DescribeImagesOutput) (all []string) {
	for _, v := range output.Images {
		all = append(all, aws.StringValue(v.ImageId))
	}

	return all
}

func ListNotUsed(a *ec2.DescribeImagesOutput, s *ec2.EC2) ([]string, []string) {

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

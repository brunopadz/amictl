package aws

import (
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

func ListNotUsed(output *ec2.DescribeImagesOutput, sess *ec2.EC2) ([]string, []string, error) {
	var all, used []string

	for _, ami := range output.Images {
		all = append(all, aws.StringValue(ami.ImageId))

		amiFilter := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("image-id"),
					Values: []*string{
						ami.ImageId,
					},
				},
			},
		}

		output, err := sess.DescribeInstances(amiFilter)
		if err != nil {
			return nil, nil, err
		}

		for _, res := range output.Reservations {
			for range res.Instances {
				used = append(used, aws.StringValue(ami.ImageId))
			}
		}
	}

	return all, used, nil
}

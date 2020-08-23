package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ami struct {
	ID string
	Size int64
}

type Summary struct {
	TotalSize int64
	Images []ami
}

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

func FilterAmiInUse(sess *ec2.EC2, describeImagesOutput *ec2.DescribeImagesOutput) error {
	var IDInUseList []*string

	for _, image := range describeImagesOutput.Images {
		IDInUseList = append(IDInUseList, image.ImageId)
	}

	var criteria = &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("image-id"),
				Values: IDInUseList,
			},
		},
	}

	describeInstanceOutput, err := sess.DescribeInstances(criteria)
	if err != nil {
		return err
	}

	var images []*ec2.Image
	for _, image := range describeImagesOutput.Images {
		var count = 0

		for _, res := range describeInstanceOutput.Reservations {
			for _, instance := range res.Instances {
				if aws.StringValue(instance.ImageId) == aws.StringValue(image.ImageId) {
					count++
				}
			}
		}

		if count == 0 {
			images = append(images, image)
		}
	}

	describeImagesOutput.SetImages(images)

	return nil
}

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

func ListAll(output *ec2.DescribeImagesOutput) []ami {
	var imageList []ami

	for _, image := range output.Images {
		imageList = append(imageList, ami{
			ID:   aws.StringValue(image.ImageId),
			Size: aws.Int64Value(image.BlockDeviceMappings[0].Ebs.VolumeSize),
		})
	}

	return imageList
}

func ListNotUsed(sess *ec2.EC2, imageList []ami) ([]ami, error) {
	var amiInUsedList []ami

	for _, image := range imageList {
		criteria := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("image-id"),
					Values: []*string{
						&image.ID,
					},
				},
			},
		}

		output, err := sess.DescribeInstances(criteria)
		if err != nil {
			return nil, err
		}

		for _, res := range output.Reservations {
			for range res.Instances {
				amiInUsedList = append(amiInUsedList,  ami{
					ID:   image.ID,
					Size: image.Size,
				})
			}
		}
	}

	return amiInUsedList, nil
}

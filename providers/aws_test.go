package providers

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"reflect"
	"testing"
)

func TestAwsListAll(t *testing.T) {
	t.Run("Given empty image set, then return an empty slice.", func(t *testing.T) {
		var expectedSlice []string

		var images = ec2.DescribeImagesOutput{
			Images: []*ec2.Image{},
		}

		result := AwsListAll(&images)

		if !reflect.DeepEqual(result, expectedSlice) {
			t.Errorf("AwsListAll() = %v, want %v", result, expectedSlice)
		}
	})
	t.Run("Given image set with valid images, then return a set of ImageID.", func(t *testing.T) {
		var genericImageID = "image_id"

		var expectedSlice = []string {
			"image_id",
			"image_id",
		}

		var images = ec2.DescribeImagesOutput{
			Images: []*ec2.Image{
				{
					ImageId: &genericImageID,
				},
				{
					ImageId: &genericImageID,
				},
			},
		}

		result := AwsListAll(&images)

		if !reflect.DeepEqual(result, expectedSlice) {
			t.Errorf("AwsListAll() = %v, want %v", result, expectedSlice)
		}
	})
}
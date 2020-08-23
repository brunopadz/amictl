package aws

import "math"

const (
	usEast1      = 0.023 // Virginia
	usEast2      = 0.023 // Ohio
	usWest1      = 0.026 // California
	usWest2      = 0.023 // Oregon
	saEast1      = 0.040 // São Paulo
	defaultValue = 0.02
)

// GetAmiPriceBySize get an monthly estimated value of ami by volume size
func GetAmiPriceBySize(sizeInGb int64, region string) float64 {
	switch region {
	case "us-east-1":
		return round(usEast1 * float64(sizeInGb))
	case "us-east-2":
		return round(usEast2 * float64(sizeInGb))
	case "us-west-1":
		return round(usWest1 * float64(sizeInGb))
	case "us-west-2":
		return round(usWest2 * float64(sizeInGb))
	case "sa-east-1":
		return round(saEast1 * float64(sizeInGb))
	default:
		return round(defaultValue * float64(sizeInGb))
	}
}

func round(x float64) float64 {
	return math.Round(x*100) / 100
}

package pricing

const us_east_1 = 0.023 // Virginia
const us_east_2 = 0.023 // Ohio
const us_west_1 = 0.026 // Califórnia
const us_west_2 = 0.023 // Oregon
const sa_east_1 = 0.040 // São Paulo

const default_value = 0.02

func Ami(size_in_gb int64, region string) float64 {

	switch region {
	case "us-east-1":
		return us_east_1 * float64(size_in_gb)
	case "us-east-2":
		return us_east_2 * float64(size_in_gb)
	case "us-west-1":
		return us_west_1 * float64(size_in_gb)
	case "us-west-2":
		return us_west_2 * float64(size_in_gb)
	case "sa-east-1":
		return sa_east_1 * float64(size_in_gb)
	default:
		return default_value * float64(size_in_gb)
	}

}

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/brunopadz/amictl/pkg/utils"

	"github.com/pterm/pterm"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	cfg "github.com/brunopadz/amictl/config"
	aaws "github.com/brunopadz/amictl/pkg/providers/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	awsCmd.AddCommand(inspect)
}

var inspect = &cobra.Command{
	Use:     "inspect",
	Short:   "Inspect AMI",
	Long:    `Inspect command shows additional info about AMIs`,
	Example: `  amictl aws inspect --region 123123123123 --ami ami-0x00000000f`,
	RunE:    runInspect,
}

var (
	ami    string
	region string
)

func runInspect(cmd *cobra.Command, _ []string) error {

	var c cfg.Config

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}

	s, err := aaws.New(region)
	if err != nil {
		log.Fatalln("Couldn't create a session to AWS.")
	}

	a := ec2.NewFromConfig(s)

	i := &ec2.DescribeImagesInput{
		ImageIds: []string{
			ami,
		},
		Owners: []string{
			viper.GetString("aws.account"),
		},
	}

	o, err := a.DescribeImages(context.TODO(), i)
	if err != nil {
		log.Fatalln("Couldn't get AMI data.")
	}

	for _, v := range o.Images {

		var (
			arch       string
			image      string
			state      string
			device     string
			boot       string
			hypervisor string
			virtType   string
			volumeType string
		)

		switch v.Architecture {
		case "i386":
			arch = "i386"
		case "x86_64":
			arch = "x86_64"
		case "arm64":
			arch = "arm64"
		case "x86_64_mac":
			arch = "x86_64_mac"
		default:
			arch = "-"
		}

		switch v.ImageType {
		case "machine":
			image = "machine"
		case "kernel":
			image = "kernel"
		case "ramdisk":
			image = "ramdisk"
		default:
			image = "-"
		}

		switch v.State {
		case "available":
			state = "available"
		case "invalid":
			state = "invalid"
		case "deregistered":
			state = "deregistered"
		case "transient":
			state = "transient"
		case "failed":
			state = "failed"
		case "error":
			state = "error"
		default:
			state = "-"
		}

		switch v.RootDeviceType {
		case "ebs":
			device = "ebs"
		case "instance-store":
			device = "instance-store"
		default:
			device = "-"
		}

		switch v.BootMode {
		case "legacy-bios":
			boot = "legacy-bios"
		case "uefi":
			boot = "uefi"
		default:
			boot = "-"
		}

		switch v.Hypervisor {
		case "ovm":
			hypervisor = "ovm"
		case "xen":
			hypervisor = "xen"
		default:
			hypervisor = "-"
		}

		switch v.VirtualizationType {
		case "hvm":
			virtType = "hvm"
		case "paravirtual":
			virtType = "paravirtual"
		default:
			virtType = "-"
		}

		pterm.FgLightCyan.Println("Displaying info for:", pterm.NewStyle(pterm.Bold).Sprint(aws.ToString(v.ImageId)))
		pterm.FgDarkGray.Println("----------------------------------------------")
		fmt.Println("Name:", utils.EmptyString(aws.ToString(v.Name)))
		fmt.Println("Description:", utils.EmptyString(aws.ToString(v.Description)))
		fmt.Println("Creation Date:", utils.EmptyString(aws.ToString(v.CreationDate)))
		fmt.Println("Deprecation Time:", utils.EmptyString(aws.ToString(v.DeprecationTime)))
		fmt.Println("Owner ID:", utils.EmptyString(aws.ToString(v.OwnerId)))
		fmt.Println("Owner Alias:", utils.EmptyString(aws.ToString(v.ImageOwnerAlias)))
		fmt.Println("State:", utils.EmptyString(state))
		fmt.Println("Root Device Name:", utils.EmptyString(aws.ToString(v.RootDeviceName)))
		fmt.Println("Root Device Type:", utils.EmptyString(device))
		fmt.Println("RAM Disk ID:", utils.EmptyString(aws.ToString(v.RamdiskId)))
		fmt.Println("Kernel ID:", utils.EmptyString(aws.ToString(v.KernelId)))
		fmt.Println("Architecture:", utils.EmptyString(arch))
		fmt.Println("Platform Details:", utils.EmptyString(aws.ToString(v.PlatformDetails)))
		fmt.Println("Image Type:", utils.EmptyString(image))
		fmt.Println("ENA Supported:", aws.ToBool(v.EnaSupport))
		fmt.Println("Boot Mode:", utils.EmptyString(boot))
		fmt.Println("Hypervisor:", utils.EmptyString(hypervisor))
		fmt.Println("Virtualization Type:", utils.EmptyString(virtType))
		fmt.Println("Block Device Mapping Info:")
		for _, bdm := range v.BlockDeviceMappings {
			switch bdm.Ebs.VolumeType {
			case "gp2":
				volumeType = "gp2"
			case "gp3":
				volumeType = "gp3"
			case "io1":
				volumeType = "io1"
			case "io2":
				volumeType = "io2"
			case "st1":
				volumeType = "st1"
			case "sc1":
				volumeType = "sc1"
			}
			fmt.Println(" Volume Size:", aws.ToInt32(bdm.Ebs.VolumeSize), "GB")
			fmt.Println(" Volume Type:", utils.EmptyString(volumeType))
			fmt.Println(" Snapshot ID:", utils.EmptyString(aws.ToString(bdm.Ebs.SnapshotId)))
			fmt.Println(" Encrypted:", aws.ToBool(bdm.Ebs.Encrypted))
			fmt.Println(" Delete on Termination:", aws.ToBool(bdm.Ebs.DeleteOnTermination))
		}
		fmt.Println("SR-IOV Net Support:", utils.EmptyString(aws.ToString(v.SriovNetSupport)))
		fmt.Println("Public:", aws.ToBool(v.Public))
		fmt.Println("Tags:")
		if len(v.Tags) == 0 {
			fmt.Println("-")
		} else {
			for _, t := range v.Tags {
				fmt.Println(" ", aws.ToString(t.Key), "=", aws.ToString(t.Value))
			}
		}

	}

	return nil
}

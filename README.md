# amictl
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/brunopadz/amictl?include_prereleases&style=flat-square) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/brunopadz/amictl?style=flat-square) ![Drone (cloud)](https://img.shields.io/drone/build/brunopadz/amictl?style=flat-square) ![GitHub](https://img.shields.io/github/license/brunopadz/amictl?style=flat-square)  

amictl is a super simple cli app to control your AMIs and Images

## Disclaimer

- There are a ton of features that still need to be implemented. Just check the [contributing guide](CONTRIBUTING.md) and the opened [issues](https://github.com/brunopadz/amictl/issues).

## Configuring

❗️ AWS is the only cloud provider supported to this date.

- Make sure AWS CLI is installed and properly configured. amictl uses access and secret keys to authenticate to AWS.
  - You can configure it using [Environment Variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html) or through [cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html).

More docs coming soon.

## How to use

- Download the binay from [releases page](https://github.com/brunopadz/amictl/releases) or 
- Download using go get:
  
  `go get github.com/brunopadz/amictl`

**List all AMIs**

```
$ amictl aws list-all --account 123456789012 --region us-east-1
ami-00123asb820d84d9a
ami-01ee75aqwez39a298
ami-02e6a65236aa8d0e7
ami-0387a7987av1b328d
ami-039835c818ezxc21c
ami-0345df085fe686a54
ami-03fd5464hdd14b864
Total of 7 AMIs
```

**List not used AMIs**

```
$ amictl aws list-unused --account 123456789012 --region us-east-1
ami-00123asb820d84d9a
ami-01ee75aqwez39a298
ami-02e6a65236aa8d0e7
Total of 3 not used AMIs
```

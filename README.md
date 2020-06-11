# amictl
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/brunopadz/amictl?style=flat-square) ![GitHub](https://img.shields.io/github/license/brunopadz/amictl?style=flat-square)

amictl is a super simple cli app to control your AMIs and Images

## Disclaimer

- There are a ton of features that need to be implemented. Just check the [contributing guide](CONTRIBUTING.md) and the opened [issues](https://github.com/brunopadz/amictl/issues).

## How to use

Download using go get:

`go get github.com/brunopadz/amictl`

❗️ AWS is the only cloud provider supported to this date.

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
$ amictl aws list-all --account 123456789012 --region us-east-1
ami-00123asb820d84d9a
ami-01ee75aqwez39a298
ami-02e6a65236aa8d0e7
Total of 3 AMIs
```

# amictl
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/brunopadz/amictl?style=flat-square) ![GitHub](https://img.shields.io/github/license/brunopadz/amictl?style=flat-square)

amictl is a super simple cli app to control your AMIs and Images

## Disclaimer

- This project is currently on development phase. At this moment you can build and run at your own risk.
- There are a ton of features that need to be implemented. Just check the [contributing guide](CONTRIBUTING.md) and the opened [issues](https://github.com/brunopadz/amictl/issues).

## How to use

AWS is the only cloud provider supported to this date.

**List all AMIs**

`amictl aws list-all <ACCOUNT_ID> <region>`

**List not used AMIs**

`amictl aws list-unused <ACCOUNT_ID> <region>`

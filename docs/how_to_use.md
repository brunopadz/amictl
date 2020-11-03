## First things first

AWS is the only cloud provider supported to this date. If you need support for another major cloud provider, feel free to open an issue. :)

## Credentials

Currently amictl supports the same authentication method as AWS does. It can use the default credentials configured in `~/.aws/credentials` or [environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html) such as `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.

## How to use amictl

- Download the latest build from [releases page](https://github.com/brunopadz/amictl/releases) or download it as follows: 
```
go get github.com/brunopadz/amictl
```
- All commands must have the flag `--account` (or `-a`) and `--region` (or `-r`)

To list all your AMIs:

```
amictl aws list-all --account 123456789012 --region us-east-1
ami-00123asb820d84d9a
ami-01ee75aqwez39a298
ami-02e6a65236aa8d0e7
ami-0387a7987av1b328d
ami-039835c818ezxc21c
ami-0345df085fe686a54
ami-03fd5464hdd14b864
Total of 7 AMIs
```

To show how many AMIs are not being used:

```
$ amictl aws list-unused -a 123456789012 -r us-east-1
ami-00123asb820d84d9a
ami-01ee75aqwez39a298
ami-02e6a65236aa8d0e7
Total of 3 AMIs
```

And you can just add the flag `--cost` and boom:

```
amictl aws list-unused --account 123456789012 --region us-east-1 --cost
AMI-ID: ami-044ec27279a83e963 Size: 20 GB Estimated cost monthly: U$ 0.46
AMI-ID: ami-09665078cc0a18084 Size: 8 GB  Estimated cost monthly: U$ 0.18
AMI-ID: ami-0c14b9433a78ac8f1 Size: 8 GB  Estimated cost monthly: U$ 0.18

Estimated cost monthly: U$ 0.83 for 3 Unused AMI
=======
Total of 3 not used AMIs
```
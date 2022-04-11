![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/brunopadz/amictl?include_prereleases&style=flat-square) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/brunopadz/amictl?style=flat-square) ![Drone (cloud)](https://img.shields.io/drone/build/brunopadz/amictl?style=flat-square) ![GitHub](https://img.shields.io/github/license/brunopadz/amictl?style=flat-square)

# amictl

amictl allows cloud operators/engineers to manage AWS AMIs.

With amictl you can list used and unused AMIs, inspect and deregister (delete) them.

## How to use
 
### Configuring amictl

Make sure you have your credentials configured correctly in
`~/.aws/credentials` or you can use [environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)
such as `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.

Run `amictl init` to create a new config file. Add the AWS account ID and the regions you have AMIs.

You can change the configuration later at `~/.amictl.yaml` (Linux/MacOS) or at `%HOMEPROFILE%/.amictl.yaml` (Windows).

```yaml
aws:
  account: "12345678912"
  regions:
    - us-east-1
    - us-east-2
```

### Listing AMIs

With amictl is possible to list all AMIs and unused AMIs.

To list all AMIs:

```sh
$ amictl aws list-all
```

To list unused AMIs:

```sh
$ amictl aws list-unused
```

It's important to highlight that list operations will run on all regions defined in `$HOME/.amictl.yaml`.

Another super important thing is that `amictl` will colorize the output depending on how many AMIs are not being used.

### Inspecting AMIs

The `inspect` command provides super useful AMI information, such as: State, Architecture, EBS info and tags.

```sh
$ amictl aws inspect -r us-east-1 -a ami-0d8918bda81d298c84
```

Example of an `inspect` output: 

```text
Displaying info for: ami-0d8918bda81d298c84
----------------------------------------------
Name: amictl-modern-sheep
Description: A copy of amictl-modern-sheep
Creation Date: 2021-09-30T17:37:18.000Z
Deprecation Time: -
Owner ID: 1234567891234
Owner Alias: -
State: available
Root Device Name: /dev/xvda
Root Device Type: ebs
RAM Disk ID: -
Kernel ID: -
Architecture: x86_64
Platform Details: Linux/UNIX
Image Type: machine
ENA Supported: true
Boot Mode: -
Hypervisor: xen
Virtualization Type: hvm
Block Device Mapping Info:
 Volume Size: 30 GB
 Volume Type: gp2
 Snapshot ID: snap-0af2a96d9d85aa042
 Encrypted: false
 Delete on Termination: true
SR-IOV Net Support: simple
Public: false
Tags:
  Name = amictl-modern-sheep
  Terraform = true
  Environment = prod
```

### Deregistering AMIs

`deregister` command deletes an AMI. You must provide a region and an AMI ID.

```sh
$ amictl aws deregister -r us-east-1 --ami ami-0d8918bda81d298c84
```

## Contributing

We welcome and encourage contributions to amiclt. There are a lot of features that still need to be implemented.

Please read the [contributing guide](CONTRIBUTING.md) and  make sure to check the [code of conduct](CODE_OF_CONDUCT.md).

## Support

Feel free to report bugs and suggest features in [Github issues](https://github.com/brunopadz/amictl/issues) or ask
for help. 

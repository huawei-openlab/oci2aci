# oci2aci - Convert OCI bundle to ACI

oci2aci is a small library and CLI binary that converts [OCI](https://github.com/opencontainers/specs) bundle to
[ACI](https://github.com/appc/spec/blob/master/SPEC.md#app-container-image). It takes OCI bundle as input, and gets ACI image as output.

oci2aci's workflow divided into two steps:
- **Convert**. Convert oci layout to aci layout.
- **Build**. Build aci layout to .aci image.

An OCI layout described as below:
```
config.json
runtime.json
rootfs/
```

An ACI layout described as below:
```
manifest
rootfs/
```

## Build

Installation is simple as:

	$ go get github.com/huawei-openlab/oci2aci

or as involved as:

	$ cd $GOPATH/src/github.com/
	$ mkdir huawei-openlab
	$ cd huawei-openlab
	$ git clone https://github.com/huawei-openlab/oci2aci.git
	$ cd oci2aci
	$ make
	
## Usage

```
$ oci2aci
NAME:
   oci2aci - Tool for conversion from oci to aci

USAGE:
   oci2aci [--debug] [arguments...]

VERSION:
   0.1.0

FLAGS:
   -debug=false: Enables debug messages

```
You can use oci2aci as a CLI tool directly to convert a oci-bundle to aci image, furthermore, you can use oci2aci as a external function in your program by importing package "github.com/huawei-openlab/oci2aci/convert"
## Example

Examples of oci2aci illustrated as below:

- An example of invalid oci bundle
```
$ oci2aci  --debug test
2015/09/28 09:46:05 test: invalid oci bundle: error accessing bundle: stat test: no such file or directory
2015/09/28 09:46:05 Conversion stop.
```
- An example of valid oci bundle
```
$ oci2aci  --debug example/oci-bundle
2015/09/28 09:42:14 example/oci-bundle/: valid oci bundle
2015/09/28 09:42:14 Manifest:/tmp/oci2aci796486541/manifest generated successfully.
2015/09/28 09:42:14 Image:/tmp/oci2aci796486541.aci generated successfully.

$ actool --debug validate /tmp/oci2aci796486541.aci
/tmp/oci2aci796486541.aci: valid app container image

$ rkt run /tmp/oci2aci796486541.aci --interactive --insecure-skip-verify --mds-register=false --volume proc,kind=host,source=/bin --volume dev,kind=host,source=/bin --volume devpts,kind=host,source=/bin --volume shm,kind=host,source=/bin --volume mqueue,kind=host,source=/bin --volume sysfs,kind=host,source=/bin --volume cgroup,kind=host,source=/bin
2015/09/28 09:45:26 Preparing stage1
2015/09/28 09:45:26 Writing image manifest
2015/09/28 09:45:26 Loading image sha512-ed1404273ed6ab8e8c7a323b994e8ce6e24d0dd5b17d2480021d52cdc87de8f1
2015/09/28 09:45:26 Writing pod manifest
2015/09/28 09:45:26 Setting up stage1
2015/09/28 09:45:26 Wrote filesystem to /var/lib/rkt/pods/run/fc9d66c6-4c49-4a14-94d6-3d8215521dd2
2015/09/28 09:45:26 Pivoting to filesystem /var/lib/rkt/pods/run/fc9d66c6-4c49-4a14-94d6-3d8215521dd2
2015/09/28 09:45:26 Execing /init
[1016397.575456] example[4]: Do something in advance for the rkt container......Done!
[1016397.579740] example[6]: Hello, I am running in the rkt container......
[1016397.581921] example[8]: Clean the resource for the rkt container......Done!

```
- Specify output directory for generated aci image
```
$ ./oci2aci --debug example/oci-bundle/ oci.aci
2015/11/14 15:56:43 example/oci-bundle/: valid oci bundle
2015/11/14 15:56:43 Manifest:/tmp/oci2aci406724597/manifest generated successfully.
2015/11/14 15:56:43 Image:/tmp/oci2aci406724597.aci generated successfully.
2015/11/14 15:56:43 Image:oci.aci generated successfully
```

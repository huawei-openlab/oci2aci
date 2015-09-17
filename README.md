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

	go get github.com/huawei-openlab/oci2aci

or as involved as:

	git clone git://github.com/huawei-openlab/oci2aci
	cd oci2aci
	make
	
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

## Example

Examples of oci2aci illustrated as below:
```
// An example of invalid oci bundle
$ ./oci2aci  --debug test
test: invalid oci bundle: error accessing bundle: stat test: no such file or directory
Conversion stop.

// An example of valid oci bundle
$ ./oci2aci  --debug example/oci-bundle
 aci image generated successfully.

```

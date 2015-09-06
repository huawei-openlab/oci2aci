# oci2aci - Convert OCI bundle to ACI

oci2aci is a small library and CLI binary that converts [OCI](https://github.com/opencontainers/specs) bundle to
[ACI](https://github.com/appc/spec/blob/master/SPEC.md#app-container-image). It takes OCI bundle as input, and gets ACI image as output.

All ACIs generated are compressed with gzip.

## Build

Installation is simple as:

	go get github.com/huawei-openlab/oci2aci

or as involved as:

	git clone git://github.com/huawei-openlab/oci2aci
	cd oci2aci
	go get -d ./...
	go build
	
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

```
$ oci2aci  --debug test
test: invalid oci bundle: error accessing bundle: stat test: no such file or directory
Conversion stop.

$ oci2aci  --debug oci-bundle

 oci-bundle.aci generated successfully.

```

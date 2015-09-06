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
	
## CLI examples

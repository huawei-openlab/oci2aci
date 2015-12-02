BUILDTAGS=
export GOPATH:=$(CURDIR)/Godeps/_workspace:$(GOPATH)

all:
	go build -tags "$(BUILDTAGS)" -o oci2aci .

install:
	cp oci2aci /usr/local/bin/oci2aci
clean:
	go clean

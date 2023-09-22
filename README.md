# potent-portables

<div style="width: 30%; height: 30%">

![Potent Portables](docs/.images/potent-portables.png)

</div>

A portable method for CI using [Mage](https://github.com/magefile/mage#readme)

## Pre-reqs

### Local machine

- Zarf
- K8s cluster w/ Zarf init
- K8s context set
- Golang v1.20+
- Git repo cloned

### Egress limited machine (VM?)

- Zarf
- Zarf init package
- Zarf potent-portables package
- Golang v1.20+ installed

## Local Dev

```console
# Install Mage (using GOPATH)
$ go install github.com/magefile/mage@latest

# Set Debug (same as 'mage -v')
$ export MAGEFILE_VERBOSE=1

# List Targets
$ mage -l
Targets:
  airgap:all              Airgap - (aka 'mage a').
  airgap:zarfDeploy       Airgap Deploy - (aka 'mage airgap:deploy').
  airgap:zarfInit         Airgap Init Cluster - (aka 'mage airgap:init').
  build:all               Create package - aka 'mage b'.
  build:zarfBuild         Create package using Zarf.
  build:zarfVersion       Output Zarf version.
  deploy:all              Install package - (aka 'mage d') | usage: 'mage deploy oci://pkg-url-here', or 'mage deploy local'.
  deploy:zarfDeploy       Install package using Zarf - (aka 'mage deploy:local').
  deploy:zarfDeployOCI    Install OCI package using Zarf - (aka 'mage deploy:oci').

# List Target Details
$ mage -h mage -h deploy:zarfdeploy
Install package using Zarf - (aka 'mage deploy:local'). Deploys zarf package under ./app directory
Usage:
        mage deploy:zarfdeploy

Aliases: deploy:local

# Build
$ mage build

# Deploy from local
$ mage deploy local

# Deploy from OCI
$ mage deploy oci://ghcr.io/mxnxpx/packages/podinfo:0.0.1-amd64
```

## Air Gap

```console
# (Internet connected machine) Create Zarf potent-portables package for Air Gap and compiles mage-bin/mage binary
zarf package create . --confirm

# (Air Gap machine)
# Copy zarf, zarf init package, zarf potent-portables package onto portable media
# Copy contents of portable media to Air Gap machine

# (Air Gap machine) Extract zarf package
$ zarf tools archiver decompress zarf-package-potent-portables-amd64-0.0.1.tar.zst tmp-extract --unarchive-all && mv tmp-extract/components/compile/files/0 airgap && rm -rf tmp-extract

# (Air Gap machine) Change Directory to airgap and Set mage path
$ cd airgap
$ export PATH=./mage-bin:$PATH

# (Air Gap machine) Set Debug (same as 'mage -v')
$ export MAGEFILE_VERBOSE=1

# (Air Gap machine) Run all Air Gap Targets
$ mage airgap
```

# potent-portables

<div style="width: 30%; height: 30%">

![Potent Portables](docs/.images/potent-portables.png)

</div>

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

## Commands

```console
# Set Debug (same as 'mage -v')
$ export MAGEFILE_VERBOSE=1

# List Targets
$ mage
Targets:
  build:build             Create package - aka 'mage b'
  build:zarfBuild         Create package using Zarf (sub-Target of 'mage build')
  build:zarfVersion       Output Zarf version (sub-Target of 'mage build')
  deploy:deploy           Install package - aka 'mage d' (For existing OCI package, use: `mage deploy oci://pkg-url-here`, or local, use: `mage deploy local`)
  deploy:zarfDeploy       Install package using Zarf (conditional sub-Target of 'mage deploy')
  deploy:zarfDeployOCI    Install OCI package using Zarf (conditional sub-Target of 'mage deploy')

# Build
$ mage build

# Deploy from local
$ mage deploy local

# Deploy from OCI
$ mage deploy oci://ghcr.io/mxnxpx/packages/podinfo:0.0.1-amd64
```

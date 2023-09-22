# punchlist

Investigate and build a portable CI/CD pipeline solution that deploys resources into a Kubernetes cluster, does not rely on a kitchen-sink approach (”move the dev env and tools into a portable Docker container”), and meets one or more of the following qualities

- [x] - Runs the same in Git[Hub/Lab] as it does from a developer’s workstation
- [x] - Packages and deploys Zarf packages
- [ ] - Builds and deploys UDS bundles
- [ ] - Exposes a way to (optionally?) run health checks, integration tests, and end-to-end tests
- [x] - Works across an air gap
- [x] - Does not leverage Flux
- [ ] - Can perform upgrades
- [ ] - Can perform automated rollbacks if there is an upgrade failure

> If checked off above, this functionality was verified

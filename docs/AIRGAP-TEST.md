# airgap test

- Stand up Ubuntu VM with golang using Multipass
- Copy zarf cli to VM under dir named in $PATH
- Copy zarf init package to ~/.zarf-cache dir in VM
- Disable VM outbound connectivity to imitate egress limited environment

```console
# log & disable all outbound connections
iptables -t filter -I OUTPUT 1 -m state --state NEW -j LOG --log-level warning   --log-prefix "Attempted to initiate a connection from a local process"   --log-uid
iptables -t filter -I OUTPUT 1 -m state --state NEW -j DROP

# only allow internal calls (to localhost) "outbound"
iptables -t filter -I OUTPUT 1 -m state --state NEW -d 127.0.0.1/32 -j ACCEPT
```

- Verify outbound connections are prevented in VM

```console
$ curl github.com
curl: (6) Could not resolve host: github.com
$ curl google.com
curl: (6) Could not resolve host: google.com
$ curl zarf.dev
curl: (6) Could not resolve host: zarf.dev
```

- Update VM $PATH to include ./mage-bin (where compiled mage cli is located)

- Do the VM airgap thing!

```console
$ mage airgap
```

- Port-forward in the VM

```console
$ zarf tools kubectl port-forward -n podinfo --address 0.0.0.0 svc/podinfo 9898:9898
Forwarding from 0.0.0.0:9898 -> 9898
```

- Get the VM IP
- Open a browser to: `http://<VM-IP>:9898`
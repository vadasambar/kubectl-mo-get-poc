# kubectl-mo-get-poc
First PoC for `kubectl-mo` plugin

# Usage
Get any K8s resource from multiple namespaces like this:
ConfigMaps:
```
executing command /home/suraj/bin/kubectl -n kube-system get cm
NAME                                 DATA   AGE
extension-apiserver-authentication   6      7d13h
cluster-dns                          2      7d13h
local-path-config                    4      7d13h
chart-content-traefik                0      7d13h
chart-values-traefik                 1      7d13h
chart-content-traefik-crd            0      7d13h
chart-values-traefik-crd             0      7d13h
kube-root-ca.crt                     1      7d13h
coredns                              2      7d13h

executing command /home/suraj/bin/kubectl -n default get cm
NAME               DATA   AGE
kube-root-ca.crt   1      7d13h
```
Pods:
```
executing command /home/suraj/bin/kubectl -n kube-system get pods
NAME                                      READY   STATUS      RESTARTS       AGE
helm-install-traefik-crd-vplbc            0/1     Completed   0              7d13h
helm-install-traefik-q26gv                0/1     Completed   1              7d13h
local-path-provisioner-7b7dc8d6f5-q8k6x   1/1     Running     17 (60m ago)   7d13h
coredns-b96499967-tmswr                   1/1     Running     17 (60m ago)   7d13h
svclb-traefik-98312b71-66crq              2/2     Running     41 (60m ago)   7d13h
traefik-7cd4fcff68-dt82t                  1/1     Running     17 (60m ago)   7d13h
metrics-server-668d979685-tbblc           1/1     Running     17 (60m ago)   7d13h

executing command /home/suraj/bin/kubectl -n default get pods

```

Or anything other K8s resource you want. 

## Notes
Place where `List` REST call actually happens: https://github.com/vadasambar/kubectl-mo-get-poc/blob/98a67a0077cd45e98e3ae5ac433e6abd0e131a32/vendor/k8s.io/cli-runtime/pkg/resource/helper.go#L108
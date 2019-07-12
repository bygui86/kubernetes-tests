# Kubernetes networking - Calico

## [Installation instructions](https://docs.projectcalico.org/v3.8/getting-started/kubernetes/installation/calico)

1. Download manifests
	```
	curl -O https://docs.projectcalico.org/v3.8/manifests/calico.yaml
	```
2. [Check pod CIDR](https://prefetch.net/blog/2018/02/20/generating-kubernetes-pod-cidr-networks-with-kubectl-and-jq/)
	```
	kubectl get nodes -o json | jq '.items[] | .spec'
	```
	* if `192.168.0.0/16`, skip to point 3.
	* otherwise
		```
		POD_CIDR="<YOUR_POD_CIDR>"
		sed -i -e "s?192.168.0.0/16?$POD_CIDR?g" calico.yaml
		```
3. Deploy
	```
	kubectl apply -f calico.yaml
	```

## `OPTIONAL` [Enabling application layer policy](https://docs.projectcalico.org/v3.8/getting-started/kubernetes/installation/app-layer-policy)

`TODO`

## Tutorials

### [Simple network policy](https://docs.projectcalico.org/v3.8/security/simple-policy)

`TODO`

### [Control ingress/egress traffic](https://docs.projectcalico.org/v3.8/security/advanced-policy)

`TODO`

### [Connections UI](https://docs.projectcalico.org/v3.8/security/stars-policy/)

`TODO`

### [calicoctl](https://docs.projectcalico.org/v3.8/getting-started/calicoctl/install)

`TODO`

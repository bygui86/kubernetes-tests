
# Kubernetes networking - Calico - Install

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
	cd install
	kubectl apply -f calico.yaml
	```

## [Install `calicoctl` as Kubernetes pod](https://docs.projectcalico.org/v3.8/getting-started/calicoctl/install#installing-calicoctl-as-a-kubernetes-pod)

`Subfolder`: [calicoctl](calicoctl)

1. Download and deploy the manifest that matches your datastore type
	* etcd
		```
		curl -O https://docs.projectcalico.org/v3.8/manifests/calicoctl-etcd.yaml
		kubectl apply -f calicoctl-etcd.yaml
		```
	* Kubernetes API datastore
		```
		curl -O  https://docs.projectcalico.org/v3.8/manifests/calicoctl.yaml
		kubectl apply -f calicoctl.yaml
		```
2. Run a command to verify that works properly
	```
	kubectl exec -ti -n kube-system calicoctl -- /calicoctl get profiles -o wide
		NAME                                                 LABELS
		kns.client                                           map[pcns.role:client]
		kns.default                                          map[]
		kns.kube-node-lease                                  map[]
		...
	```
3. In order to use the calicoctl alias when reading manifests, redirect the file into stdin, for example:
	```
	kubectl exec -ti -n kube-system calicoctl -- /calicoctl create -f - < my_manifest.yaml
	```

## [Configure `calicoctl`](https://docs.projectcalico.org/v3.8/getting-started/calicoctl/configure/)

1. Check configuration
	
	```
	calicoctl get nodes
	```
	
	A correct setup will yield a list of the nodes that have registered. If an empty list is returned you are either pointed at the wrong datastore or no nodes have registered. If an error is returned then attempt to correct the issue then try again.

## [Enable application layer policy](https://docs.projectcalico.org/v3.8/getting-started/kubernetes/installation/app-layer-policy)

`Subfolder:` [Application layer policy](enable-application-layer-policy)

`WARN: Application layer policy requires Istio.`

### Prerequisites

* Calico installed
* calicoctl installed and configured
* Istio ready to be installed

### Deploy

1. Application layer policy requires the Policy Sync API to be enabled on Felix.
	To do this cluster-wide, modify the `default FelixConfiguration` to set the field `policySyncPathPrefix` to `/var/run/nodeagent`.
	The following example uses sed to modify your existing default config before re-applying it.
	```
	kubectl exec -ti -n kube-system calicoctl -- /calicoctl get felixconfiguration default --export -o yaml | \
		sed -e '/  policySyncPathPrefix:/d' \
		> felix-config.yaml
	echo "  policySyncPathPrefix: /var/run/nodeagent" >> felix-config.yaml
	kubectl exec -ti -n kube-system calicoctl -- /calicoctl apply -f - < felix-config.yaml
	```

2. Verify the Policy Sync API is now enabled
	```
	kubectl exec -ti -n kube-system calicoctl -- /calicoctl get felixconfiguration default --export -o yaml
	```
`TO BE TESTED`

3. Install Istio with mTLS enabled (see the relative documentation online)

4. Update Istio sidecar injector
	This step modifies the injector configuration to add Dikastes, a Calico component, as sidecar containers.

`IN PROGRESS`

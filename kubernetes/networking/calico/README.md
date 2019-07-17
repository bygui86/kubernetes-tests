# Kubernetes networking - Calico

## [Installation instructions](https://docs.projectcalico.org/v3.8/getting-started/kubernetes/installation/calico)

`Subfolder: [Install](install)`

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

## `OPTIONAL` [Enable application layer policy](https://docs.projectcalico.org/v3.8/getting-started/kubernetes/installation/app-layer-policy)

`Subfolder: [Application layer policy](install/enable-application-layer-policy)`

`TODO`

### `OPTIONAL` [Install `calicoctl` as Kubernetes pod](https://docs.projectcalico.org/v3.8/getting-started/calicoctl/install#installing-calicoctl-as-a-kubernetes-pod)

`Subfolder: [calicoctl](install/calicoctl)`

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

## Tutorials

### [Simple network policy](https://docs.projectcalico.org/v3.8/security/simple-policy)

`Subfolder: [Simple network policy](tutorials/1 - simple-network-policy)`

1. Create a sample namespace
	```
	kubectl create ns policy-demo
	```
2. Create demo pods
	```
	kubectl run --namespace=policy-demo nginx --replicas=2 --image=nginx
	kubectl expose --namespace=policy-demo deployment nginx --port=80
	```
3. Ensure nginx service is accessible
	```
	kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q nginx -O -
			... nginx welcome-page html code...
	```
4. Enable complete isolation
	```
	kubectl apply -f default-deny.yaml
	```
5. Test isolation
	```
	kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			wget: download timed out
	```
6. Allow specific access
	```
	kubectl apply -f access-nginx.yaml
	```
7. Test specific access
	1. We are able to access nginx from a pod with the label `run: access`
		```
		kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh
			/# wget -q --timeout=5 nginx -O -
				... nginx welcome-page html code...
		```
	2. But we are not able to access ngnix from any pod without the label mentioned above
		```
		kubectl run --namespace=policy-demo cant-access --rm -ti --image busybox /bin/sh
			/# wget -q --timeout=5 nginx -O -
				wget: download timed out
		```
8. Cleanup
	```
	kubectl delete ns policy-demo
	```

### [Control ingress/egress traffic](https://docs.projectcalico.org/v3.8/security/advanced-policy)

`Subfolder: [Ingress/Egress](tutorials/2 - ingress-egress)`

1. Create sample namespace
	```
	kubectl create ns advanced-policy-demo
	```
2. Create demo pods
	```
	kubectl run --namespace=advanced-policy-demo nginx --replicas=2 --image=nginx
	kubectl expose --namespace=advanced-policy-demo deployment nginx --port=80
	```
3. Verify ingress and egress traffic
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			... nginx welcome-page html code...
		/# wget -q --timeout=5 google.com -O -
			... google homepage html code ...
	```
4. Deny all ingress traffic
	```
	kubectl apply -f default-deny-ingress.yaml
	```
5. Verify ingress traffic denied
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			wget: download timed out
		/# wget -q --timeout=5 google.com -O -
			... google homepage html code ...
	```
6. Allow ingress traffic only to nginx
	```
	kubectl apply -f access-nginx.yaml
	```
7. Verify nginx access
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			... nginx welcome-page html code...
	```
8. Deny all egress traffic
	```
	kubectl apply -f default-deny-egress.yaml
	```
9. Verify all egress denied
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# nslookup nginx
			;; connection timed out; no servers could be reached
		/# wget -q --timeout=5 google.com -O -
			wget: bad address 'google.com'
	```
	`INFO: The nslookup command can take a minute or more to timeout.`
10. Allow DNS egress traffic
	1. Create label `name: kube-system` on the kube-system namespace
		```
		kubectl label namespace kube-system name=kube-system
		```
	2. Create a NetworkPolicy to allow DNS egress traffic from any pods in the `advanced-policy-demo` namespace to `kube-system`
		```
		kubectl apply -f allow-dns-access.yaml
		```
11. Verify DNS access
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# nslookup nginx
			Server:		10.96.0.10
			Address:	10.96.0.10:53
			Name:	nginx.advanced-policy-demo.svc.cluster.local
			Address: 10.104.164.79
		/# nslookup google.com
			Server:		10.96.0.10
			Address:	10.96.0.10:53
			Non-authoritative answer:
			Name:	google.com
			Address: 216.58.215.238
	```
	Even though DNS egress traffic is now working, all other egress traffic from all pods in the advanced-policy-demo namespace is still blocked. Therefore the HTTP egress traffic from the `wget` calls will still fail.
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			wget: download timed out
		/# wget -q --timeout=5 google.com -O -
			wget: download timed out
	```
12. Allow egress traffic to nginx
	```
	kubectl apply -f allow-egress.yaml
	```
13. Verify allowed egress access to nginx
	```
	kubectl run --namespace=advanced-policy-demo access --rm -ti --image busybox /bin/sh
		/# wget -q --timeout=5 nginx -O -
			... nginx welcome-page html code...
		/# wget -q --timeout=5 google.com -O -
			wget: download timed out
	```
	Access to google.com times out because it can resolve DNS but has no egress access to anything other than pods with labels matching `run: nginx` in the `advanced-policy-demo` namespace.
14. Cleanup
	```
	kubectl delete ns advanced-policy-demo
	```

### [Connections UI](https://docs.projectcalico.org/v3.8/security/stars-policy/)

`Subfolder: [Connections UI](tutorials/3 - connections-ui)`

1. Create sample namespace
	```
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/manifests/00-namespace.yaml
	kubectl apply -f 00-namespace.yaml
	```
2. Deploy frontend, backend, client and management-ui applications
	```
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/manifests/01-management-ui.yaml
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/manifests/02-backend.yaml
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/manifests/03-frontend.yaml
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/manifests/04-client.yaml
	kubectl apply -f 01-management-ui.yaml
	kubectl apply -f 02-backend.yaml
	kubectl apply -f 03-frontend.yaml
	kubectl apply -f 04-client.yaml
	```
3. Open management UI to look at the connectivity of services
	```
	open http://$(minikube ip):30002
	```
4. Enable complete isolation
	```
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/policies/default-deny.yaml
	kubectl apply -n stars -f default-deny.yaml
	kubectl apply -n client -f default-deny.yaml
	```
5. Verify complete isolation
	```
	open http://$(minikube ip):30002
	```
	Refresh the management UI: with the enabled isolation, the UI can no longer access the pods, and so they will no longer show up in the UI.
6. Allow the management-ui to access services
	```
	curl https://docs.projectcalico.org/v3.8/security/stars-policy/policies/allow-ui.yaml > allow-ui_stars-ns.yaml
	curl https://docs.projectcalico.org/v3.8/security/stars-policy/policies/allow-ui-client.yaml > allow-ui_client-ns.yaml
	kubectl apply -f allow-ui_stars-ns.yaml
	kubectl apply -f allow-ui_client-ns.yaml
	```
7. Verify access to services
	```
	open http://$(minikube ip):30002
	```
	After a few seconds, refresh the UI. It should now show the Services, but they should not be able to access each other any more.
8. Allow traffic from frontend to backend
	```
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/policies/backend-policy.yaml
	kubectl apply -f backend-policy.yaml
	```
9. Verify access from frontend to backend
	```
	open http://$(minikube ip):30002
	```
	Refresh the UI. You should see the following:
	* The frontend can now access the backend (on TCP port 6379 only).
	* The backend cannot access the frontend at all.
	* The client cannot access the frontend, nor can it access the backend.
10. Allow traffic from client to frontend
	```
	curl -O https://docs.projectcalico.org/v3.8/security/stars-policy/policies/frontend-policy.yaml
	kubectl apply -f frontend-policy.yaml
	```
11. Verify access from client to frontend
	```
	open http://$(minikube ip):30002
	```
	The client can now access the frontend, but not the backend. Neither the frontend nor the backend can initiate connections to the client. The frontend can still access the backend.
12. Cleanup
	```
	kubectl delete ns client stars management-ui
	```

### [Application layer policy](https://docs.projectcalico.org/v3.8/security/app-layer-policy/)

`Subfolder: [App Layer Policy](tutorials/4 - application-layer-policy)`

`TODO`

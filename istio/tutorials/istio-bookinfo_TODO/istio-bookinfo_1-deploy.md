
# Istio bookinfo - Deploy

## Deploy on minikube

`INFO: the root folder is considered to be istio installation one (./istio-VERSION/)`

1. Deploy Istio
	```
	kubectl apply -f install/kubernetes/istio-demo.yaml
	```

2. Deploy bookinfo applications
	```
	kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
	```

3. Verify bookinfo up and running from inside the service mesh
	```
	kubectl exec -it $(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}') -c ratings -- curl productpage:9080/productpage | grep -o "<title>.*</title>"
		<title>Simple Bookstore App</title>
	```

4. Deploy the ingress gateway
	```
	kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
	```

5. Verify bookinfo accessible from external world
	* external load-balancer
		```
		export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
		export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
		export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
		```
	* node port
		```
		export INGRESS_HOST=$(minikube ip)
		export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
		export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')
		```
	```
	export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
	curl -s http://$GATEWAY_URL/productpage | grep -o "<title>.*</title>"
		<title>Simple Bookstore App</title>
	```

6. Deploy default destination rules
	```
	kubectl apply -f samples/bookinfo/networking/destination-rule-all.yaml
	```

---

## Links
* Deploy
  * [Deploy istio](https://istio.io/docs/setup/kubernetes/install/kubernetes/)
  * [Deploy bookinfo](https://istio.io/docs/examples/bookinfo/)
    * [Setup ingress host and port](https://istio.io/docs/tasks/traffic-management/ingress/ingress-control/#determining-the-ingress-ip-and-ports)

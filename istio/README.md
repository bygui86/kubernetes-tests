
# Kubernetes tests - Istio

## Commands
1. Lauch minikube
  ```
  minikube start \
    --profile istio \
    --vm-driver hyperkit --cpus 6 --memory 12288 \
    --enable-default-cni --network-plugin=cni --extra-config=kubelet.network-plugin=cni \
    --extra-config=apiserver.authorization-mode=RBAC
  ```

2. Download Istio
  ```
  ISTIO_VERSION=1.2.2
  curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$ISTIO_VERSION sh -
  cd istio-$ISTIO_VERSION
  ```

3. Install Istio
  * permissive mTLS
    ```
    kubectl apply -f install/kubernetes/istio-demo.yaml
    ```
  * strict mTLS
    ```
    kubectl apply -f install/kubernetes/istio-demo-auth.yaml
    ```

4. Enable Istio auto sidecar injection
  ```
  kubectl label namespace <namespace> istio-injection=enabled
  ```

5. `Optionals` before going for this, please check if Kubernetes service and Istio destination rules are properly configured
  1. mTLS workaround (even with PERMISSIVE mode)
    ```
    kubectl delete meshpolicies.authentication.istio.io default
    kubectl delete destinationrules.networking.istio.io istio-client-mtls
    ```
  2. Re-write liveness check
    ```
    kubectl get cm istio-sidecar-injector -n istio-system -o yaml | sed -e "s/ rewriteAppHTTPProbe: false/ rewriteAppHTTPProbe: true/" | kubectl apply -f -
    ```

---

## Links
* https://istio.io/docs/setup/kubernetes/install/kubernetes/

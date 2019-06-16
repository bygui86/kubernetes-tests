
# Istio

## Commands
1. Lauch minikube
  ```shell
  minikube start \
    --profile istio \
    --vm-driver hyperkit --cpus 6 --memory 12288 \
    --enable-default-cni --network-plugin=cni \
    --extra-config=apiserver.authorization-mode=RBAC
  ```

2. Download Istio
  ```shell
  ISTIO_VERSION=1.1.6
  # ISTIO_VERSION=1.1.8
  curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$ISTIO_VERSION sh -
  cd istio-$ISTIO_VERSION
  ```

3. Install Istio
  * permissive mTLS
    ```shell
    kubectl apply -f install/kubernetes/istio-demo.yaml
    ```
  * strict mTLS
    ```shell
    kubectl apply -f install/kubernetes/istio-demo-auth.yaml
    ```

4. Enable Istio auto sidecar injection
  ```shell
  kubectl label namespace <namespace> istio-injection=enabled
  ```

5. Optionals
  1. mTLS workaround (even with PERMISSIVE mode)
    ```shell
    kubectl delete meshpolicies.authentication.istio.io default
    kubectl delete destinationrules.networking.istio.io istio-client-mtls
    ```
  2. Revrite liveness check
    ```shell
    kubectl get cm istio-sidecar-injector -n istio-system -o yaml | sed -e "s/ rewriteAppHTTPProbe: false/ rewriteAppHTTPProbe: true/" | kubectl apply -f -
    ```

---

## Links
* istio
  * https://istio.io/docs/setup/kubernetes/install/kubernetes/
* postgres
  * Zalando
    * https://postgres-operator.readthedocs.io/en/latest/quickstart/
    * https://github.com/zalando/postgres-operator
  * CrunchyData
    * https://github.com/CrunchyData/postgres-operator
    * https://access.crunchydata.com/documentation/postgres-operator/4.0.0/installation/operator-install/

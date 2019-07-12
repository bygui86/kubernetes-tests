
# Kubernetes tests - Istio - Installation

## Preliminary steps

* Download Istio
  ```
  ISTIO_VERSION=1.2.2
  curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$ISTIO_VERSION sh -
  cd istio-$ISTIO_VERSION
  ```

* Lauch minikube
  ```
  minikube start \
    --profile istio \
    --vm-driver hyperkit --cpus 6 --memory 12288 \
    --enable-default-cni --network-plugin=cni --extra-config=kubelet.network-plugin=cni \
    --extra-config=apiserver.authorization-mode=RBAC
  ```

* `[OPTIONAL if using Helm]` Add Istio repo to Helm
  ```
  helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.2.2/charts/
  ```

## Manual standard installation

1. Install Istio
  * permissive mTLS
    ```
    kubectl apply -f install/kubernetes/istio-demo.yaml
    ```
  * strict mTLS
    ```
    kubectl apply -f install/kubernetes/istio-demo-auth.yaml
    ```

2. Enable Istio auto sidecar injection
  ```
  kubectl label namespace <namespace> istio-injection=enabled
  ```

3. `Optionals` before going for this, please check if Kubernetes service and Istio destination rules are properly configured
  1. mTLS workaround (even with PERMISSIVE mode)
    ```
    kubectl delete meshpolicies.authentication.istio.io default
    kubectl delete destinationrules.networking.istio.io istio-client-mtls
    ```
  2. Re-write liveness check
    ```
    kubectl get cm istio-sidecar-injector -n istio-system -o yaml | sed -e "s/ rewriteAppHTTPProbe: false/ rewriteAppHTTPProbe: true/" | kubectl apply -f -
    ```

## Helm customizable installation

1. Create Istio namespace
  ```
  kubectl create namespace istio-system
  ```

2. Prepare Istio-CRDs Istio manifests
  ```
  helm template $ISTIO_HOME/install/kubernetes/helm/istio-init \
    --name istio-init \
    --namespace istio-system \
  > istio-crds.yaml
  ```

3. Deploy Istio-CRDs
  ```
  kubectl apply -f istio-crds.yaml
  ```

4. Verify CRDs
  ```
  kubectl get crds -n istio-system | grep 'istio.io\|certmanager.k8s.io' | wc -l
    >> 23
  ```

5. Prepare Istio manifests
  ```
  helm template $ISTIO_HOME/install/kubernetes/helm/istio \
    --name istio \
    --namespace istio-system \
    --set gateways.istio-ingressgateway.type=NodePort \
    --set gateways.istio-egressgateway.enabled=true \
    --set istiocoredns.enabled=true \
    --set prometheus.enabled=false \
  > istio-custom.yaml
  ```

6. Deploy Istio
  ```
  kubectl apply -f istio-custom.yaml
  ```

## Helm customizable installation with Istio-CNI

1. Create Istio namespace
  ```
  kubectl create namespace istio-system
  ```

2. Prepare Istio-CNI manifests
  ```
  helm template $ISTIO_HOME/install/kubernetes/helm/istio-cni \
    --name=istio-cni \
    --namespace=kube-system \
    > istio-cni.yaml
  ```

3. Deploy Istio-CRDs
  ```
  kubectl apply -f istio-crds.yaml
  ```

4. Verify CRDs
  ```
  kubectl get crds | grep 'istio.io\|certmanager.k8s.io' | wc -l
    >> 23
  ```

5. Deploy Istio-CNI manifests
  ```
  kubectl apply -f istio-cni.yaml
  ```

6. Prepare Istio-CRDs manifests
  ```
  helm template $ISTIO_HOME/install/kubernetes/helm/istio-init \
    --name istio-init \
    --namespace istio-system \
  > istio-crds.yaml
  ```

7. Prepare Istio manifests
  ```
  helm template $ISTIO_HOME/install/kubernetes/helm/istio \
    --name istio \
    --namespace istio-system \
    --set gateways.istio-ingressgateway.type=NodePort \
    --set gateways.istio-egressgateway.enabled=true \
    --set istiocoredns.enabled=true \
    --set prometheus.enabled=false \
    --set istio_cni.enabled=true \
  > istio-custom.yaml
  ```

8. Deploy Istio
  ```
  kubectl apply -f istio-custom.yaml
  ```

---

## Links
* [Quick start](https://istio.io/docs/setup/kubernetes/install/kubernetes/)
* [Customizable installation](https://istio.io/docs/setup/kubernetes/install/helm/)

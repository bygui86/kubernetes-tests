
# Kubernetes tests - Istio - Kiali

## Preliminary steps

1. Install envsubst
  `INFO: for MacOS is included into gettext (brew install gettext && brew link --force gettext)`

2. Download installation script
  ```
  curl -L https://git.io/getLatestKialiOperator > deploy-kiali-operator.sh
  chmod +x deploy-kiali-operator.sh
  ```

---

## Installation

### Automatic using script

* quick install
  ```
  ./deploy-kiali-operator.sh
  ```

* dev install (`not recommended for production`)
  ```
  ./deploy-kiali-operator.sh --accessible-namespaces '**'
  ```

### Manual

#### Version
```
KIALI_VERSION=v1.1.0
```

#### Kiali operator
##### Namespace
```
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/namespace.yaml
kubectl apply -f namespace.yaml
```
##### CRDs
```
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/crd.yaml
kubectl apply -f crd.yaml -n kiali-operator
```
##### RBAC
```
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/service_account.yaml
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/role.yaml
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/role_binding.yaml
kubectl apply -f service_account.yaml -n kiali-operator
kubectl apply -f role.yaml -n kiali-operator
kubectl apply -f role_binding.yaml -n kiali-operator
```
##### Operator
```
curl -O https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/operator.yaml
kubectl apply -f operator.yaml -n kiali-operator
```

#### Kiali
1. Download example
  ```
  curl -O https://raw.githubusercontent.com/kiali/kiali/master/operator/deploy/kiali/kiali_cr.yaml
  ```

2. Customize example

3. Deploy Kiali
  ```
  kubectl apply -f kiali_cr.yaml
  ```
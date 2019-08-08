
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

#### Preliminary steps
```
KIALI_VERSION=v1.1.0
OPERATOR_NAMESPACE=kiali-operator
```

#### Kiali operator

1. Namespace
  ```
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/namespace.yaml > 1_namespace.yaml
  kubectl apply -f 1_namespace.yaml
  ```

2. CRDs
  ```
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/crd.yaml > 2_crd.yaml
  kubectl apply -f 2_crd.yaml
  ```

3. RBAC
  ```
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/service_account.yaml > 3_service-account.yaml
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/role.yaml > 4_cluster-role.yaml
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/role_binding.yaml > 5_cluster-role-binding.yaml
  kubectl apply -f 3_service-account.yaml
  kubectl apply -f 4_cluster-role.yaml
  kubectl apply -f 5_cluster-role-binding.yaml
  ```

4. Operator
  ```
  curl https://raw.githubusercontent.com/kiali/kiali/$KIALI_VERSION/operator/deploy/operator.yaml > 6_operator.yaml
  kubectl apply -f 6_operator.yaml
  ```

#### Kiali

1. Download example
  ```
  curl https://raw.githubusercontent.com/kiali/kiali/master/operator/deploy/kiali/kiali_cr.yaml > 7_kiali-cr.yaml
  ```

2. Customize example

3. Deploy Kiali
  ```
  kubectl apply -f 7_kiali-cr.yaml
  ```


# Istio - Mirroring

## Instructions

1. Lauch minikube
  ```
  minikube start \
    --profile istio \
    --vm-driver hyperkit --cpus 4 --memory 12288 \
    --enable-default-cni --network-plugin=cni \
    --extra-config=apiserver.authorization-mode=RBAC
  ```

2. Download Istio
  ```
  ISTIO_VERSION=1.2.2
  curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$ISTIO_VERSION sh -
  cd istio-$ISTIO_VERSION
  ```

3. Install Istio
  ```
  kubectl apply -f install/kubernetes/istio-demo.yaml
  ```

4. Enable Istio auto sidecar injection
  ```
  kubectl label namespace default istio-injection=enabled
  ```

5. Deploy httpbin v1 and v2
  ```
  kubectl apply -f httpbin-v1.yaml
  kubectl apply -f httpbin-v2.yaml
  ```

6. Create httpbin service
  ```
  kubectl apply -f httpbin-svc.yaml
  ```

7. Deploy sleep pod
  ```
  kubectl apply -f sleep.yaml
  ```

8. Create Istio DestinationRule
  ```
  kubectl apply -f dest-rule.yaml
  ```

9.  Create VirtualService pointing 100% to httpbin-v1
  ```
  kubectl apply -f virt-svc-v1.yaml
  ```

10. Open a new terminal to look at httpbin-v1 logs
  ```
  kubectl logs -f $(kubectl get pod -l app=httpbin,version=v1 -o jsonpath={.items..metadata.name}) -c httpbin
  ```

11. Open a new terminal to look at httpbin-v2 logs
  ```
  kubectl logs -f $(kubectl get pod -l app=httpbin,version=v2 -o jsonpath={.items..metadata.name}) -c httpbin
  ```

12. Produce some traffic to the service
  ```
  kubectl exec -it $(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name}) -c sleep -- sh -c 'curl  http://httpbin:8000/headers' | python -m json.tool
  ```

13. Observe traffic landing just to httpbin-v1

14. Create VirtualService still pointing 100% to httpbin-v1 but also mirroring to httpbin-v2
  ```
  kubectl apply -f virt-svc-mirror.yaml
  ```

15. Produce some other traffic to the service
  ```
  kubectl exec -it $(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name}) -c sleep -- sh -c 'curl  http://httpbin:8000/headers' | python -m json.tool
  ```

16. Observe traffic landing to both httpbin v1 and v2

---

## Cleanup

```
kubectl delete virtualservice httpbin
kubectl delete destinationrule httpbin
kubectl delete deploy httpbin-v1 httpbin-v2 sleep
kubectl delete svc httpbin
```

---

## Links
* https://istio.io/docs/tasks/traffic-management/mirroring/

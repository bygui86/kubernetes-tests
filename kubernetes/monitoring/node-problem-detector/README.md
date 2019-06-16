
# Kubernetes monitoring - Monitor Node Health

## Instructions

### Default
#### Manual
```shell
kubectl apply -f node-problem-detector_daemonset.yaml
```
#### Automatic
```shell
helm install stable/node-problem-detector
```

### Config overwrite
```shell
kubectl apply -f config-overwrite/node-problem-detector_configmap.yaml
kubectl apply -f config-overwrite/node-problem-detector_daemonset.yaml
```

---

## Links

[kubernetes-official-doc](https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/)
[github](https://github.com/kubernetes/node-problem-detector)

# Kubernetes monitoring - Prometheus - Alerting

## Instructions

### Slack

1. Deploy new configs
  ```
  kubectl delete secret alertmanager-main
  kubectl create secret generic alertmanager-main --from-file=slack-demo/alertmanager.yaml
  ```

2. Restart AlertManager to take new configs
  ```
  kubectl delete pod -l alertmanager=main
  ```

3. Deploy a demo rule
  ```
  kubectl apply -f prometheus-rule_demo.yaml
  ```

4. Open Prometheus UI
  ```
  kubectl port-forward svc/prometheus 9090:9090
  open http://localhost:9090
  ```

5. Open AlertManager UI
  ```
  kubectl port-forward svc/alertmanager-main 9093:9093
  open http://localhost:9093
  ```

---

## Slack channel

### Test

```
http --json https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX \
	Content-type:application/json \
	"text"="Hello, world."
```

---

## Links
* https://github.com/coreos/prometheus-operator/blob/master/Documentation/user-guides/alerting.md
* https://prometheus.io/docs/alerting/configuration/
* https://github.com/strimzi/strimzi-kafka-operator/blob/master/metrics/examples/prometheus/alertmanager-config/alert-manager-config.yaml
* https://api.slack.com/incoming-webhooks

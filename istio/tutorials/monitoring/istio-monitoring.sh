#!/bin/sh

ENABLE_ISTIO="$1"

if [[ -z "$1" ]]; then
	echo
	echo "/!\ No input specified, disabling Istio labeling /!\\"
	ENABLE_ISTIO="false"
fi

echo
echo "Prometheus-operator deployment started ..."

if [[ $ENABLE_ISTIO == "true" ]]; then
	echo "(i) deploy to Kubernetes on top of Istio"
	echo "    ... namespace"
	kubectl apply -f 1_namespace_istio-labeled
else
	echo "(i) deploy to Kubernetes without Istio"
	echo "    ... namespace"
	kubectl apply -f 1_namespace
fi
echo
read -p "Press enter to continue"
echo
echo "    ... crd"
kubectl apply -f 2_crd
echo
read -p "Press enter to continue"
echo
echo "    ... operator"
kubectl apply -f 3_operator
echo
read -p "Press enter to continue"
echo
echo "    ... alert-manager"
kubectl apply -f 4_alert-manager
echo
read -p "Press enter to continue"
echo
echo "    ... grafana"
kubectl apply -f 5_grafana
echo
read -p "Press enter to continue"
echo
echo "    ... kube-state-metrics"
kubectl apply -f 6_kube-state-metrics
echo
read -p "Press enter to continue"
echo
echo "    ... node-exporter"
kubectl apply -f 7_node-exporter
echo
read -p "Press enter to continue"
echo
echo "    ... prometheus-adapter"
kubectl apply -f 8_prometheus-adapter
echo
read -p "Press enter to continue"
echo
echo "    ... prometheus"
kubectl apply -f 9_prometheus
echo
read -p "Press enter to continue"
echo
if [[ $ENABLE_ISTIO == "true" ]]; then
	echo "    ... istio monitoring"
	kubectl create secret generic prometheus-additional-scrape-configs --from-file=10_istio-monitoring/prometheus-additional.yaml
	kubectl apply -f 10_istio-monitoring/istio-service-monitor.yaml
	kubectl apply -f 10_istio-monitoring/prometheus-updated-clusterRole.yaml
	kubectl apply -f 10_istio-monitoring/prometheus-updated-prometheus.yaml
	kubectl delete pods -l app=prometheus,prometheus=k8s
	kubectl apply -f 10_istio-monitoring/grafana-dashboards/
	kubectl apply -f 10_istio-monitoring/grafana-updated-deployment.yaml
	kubectl delete pods -l app=grafana
fi

echo
echo
echo "Prometheus-operator deployment completed"
echo

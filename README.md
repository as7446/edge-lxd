graph TD
    subgraph Kubernetes Cluster
        Master[Master Node] --> PrometheusOperator
        PrometheusOperator --> PrometheusInstance
        PrometheusOperator --> ServiceMonitor
        PrometheusOperator --> PodMonitor
        PrometheusInstance --> Alertmanager
        PrometheusInstance --> Grafana
    end

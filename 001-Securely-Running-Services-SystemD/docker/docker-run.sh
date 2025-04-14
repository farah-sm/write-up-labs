#!/bin/bash

echo "🐳 Running Prometheus container..."

docker run -d \
  -p 9090:9090 \
  -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus

# Visit http://localhost:9090 to access Prometheus

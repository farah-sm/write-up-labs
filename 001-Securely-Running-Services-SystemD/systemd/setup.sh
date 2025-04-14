#!/bin/bash

set -e

# Creating Prometheus system user
sudo useradd -M --shell /bin/false prometheus || echo "User already exists"

# Creating configuration and data directories
sudo mkdir -p /etc/prometheus
sudo mkdir -p /var/lib/prometheus

# Setting permissions
sudo chown -R prometheus:prometheus /etc/prometheus
sudo chown -R prometheus:prometheus /var/lib/prometheus

# Copying binaries
sudo cp ../bin/prometheus /usr/local/bin/
sudo cp ../bin/promtool /usr/local/bin/
sudo chown prometheus:prometheus /usr/local/bin/prometheus
sudo chown prometheus:prometheus /usr/local/bin/promtool

echo "Done. You can now test and run Prometheus via systemd."

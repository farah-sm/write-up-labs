# systemd Setup Notes

This directory contains files related to running Prometheus as a managed systemd service.

## Files

- `prometheus.service`: The systemd unit file
- `setup.sh`: A helper script to automate setup of users, permissions, and directories

You should place your configuration in `/etc/prometheus/prometheus.yml`, and Prometheus will store data in `/var/lib/prometheus`.

For other services, just replace the relevant binary, service name, and paths.

# Prometheus Setup on Linux using systemd & Docker

This repository contains my personal notes and templates for setting up **Prometheus** on a Linux system, using both `systemd` for long-running, production-grade service management and **Docker** for flexible, isolated environments. It is intended to serve as a reliable source of truth for anyone looking to deploy Prometheus or similar services.

![Prometheus(1)](https://github.com/user-attachments/assets/4a51bdd6-aac6-484b-9e2a-21291edd7939)

## What’s Included:
**Prometheus systemd setup**: A secure and reliable way to run Prometheus as a background service on Linux.
- **Docker setup for Prometheus**: A simple, portable approach to run Prometheus in a Docker container.
- **Configuration templates**: Ready-to-use templates for configuring Prometheus.

## Folder Structure:
```
prometheus-setup/
├── systemd/          # systemd service setup files
│   ├── prometheus.service # systemd unit file
├── docker/             # Docker setup files
│   ├── prometheus.yml     # Prometheus config for Docker
├── README.md           # This file
└── blog.md             # Detailed blog post with setup instructions
```
## Getting Started:

1. **For systemd**: Follow the instructions in the `blog.md` to set up Prometheus using `systemd` on your Linux server.
2. **For Docker**: Use the `docker/prometheus.yml` and the instructions in `blog.md` to run Prometheus in a Docker container.

# Prometheus Architecture (Optional Notes)

This section is for documenting any custom architecture you've built around Prometheus.

On the Linux side, this setup follows a standard, secure service pattern:

    Prometheus runs as a dedicated non-login system user, improving security and isolation.

    Configuration files are stored in /etc/prometheus and time-series data in /var/lib/prometheus, aligning with the Linux Filesystem Hierarchy Standard (FHS).

    The binary is placed in /usr/local/bin, separate from package-managed binaries, giving full control over versioning and upgrades.

    Systemd is used to manage the Prometheus process, ensuring:

        Automatic startup on boot

        Graceful shutdown and restart policies

        Clean logging and process supervision via journalctl

    Permissions are strictly scoped to the prometheus user, with no shell access or home directory, limiting potential attack vectors.

This architecture ensures that the monitoring system is reliable, predictable, and secure, in line with modern Linux service management best practices.

- ![Prometheus(1)](https://github.com/user-attachments/assets/9942b6b1-8f0a-4d67-b3d8-579eb5602586)

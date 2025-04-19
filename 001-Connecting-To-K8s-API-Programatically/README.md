# Connecting to Kubernetes API Programmatically with Go and client-go

This repository contains personal notes, example code, and templates for connecting to a Kubernetes cluster using the Go programming language and the `client-go` library. It serves as a reliable reference for building tools, scripts, or services that need to interact with the Kubernetes API.

Whether you're an SRE, DevOps engineer, or developer looking to build custom controllers or automation scripts, this setup provides a repeatable pattern to get started quickly and safely.

## What's Included

- **Go example using client-go**: A minimal, working example that demonstrates how to authenticate and query the Kubernetes API.
- **Kubeconfig handling**: Dynamic loading of kubeconfig path via CLI flags.
- **API interaction**: Code for listing pods in the `default` namespace using the core Kubernetes API.
- **Helpful comments**: Inline documentation explaining each step for easier understanding and future extension.


## Getting Started

### Prerequisites

- Go 1.20 or higher
- Access to a Kubernetes cluster (via `kubectl` or kubeconfig)
- The `client-go` and `apimachinery` Go packages

Install required packages:

```bash
go get k8s.io/client-go@latest
go get k8s.io/apimachinery@latest
```

### Run the Example

From the project root directory:

```bash
go run main.go
```

To specify a custom kubeconfig file:

```bash
go run main.go --kubeconfig /path/to/your/kubeconfig
```

If everything is configured correctly, you should see a list of pods in your cluster's `default` namespace.

## 🛠️ Code Highlights

- Loads kubeconfig from default path (`~/.kube/config`) or via `--kubeconfig` flag
- Builds a clientset using `clientcmd`
- Queries the Kubernetes API to list Pods
- Handles critical errors gracefully with helpful logs

## 📖 For a Deeper Dive

Check out [`blog.md`](./blog.md) for a detailed walkthrough of the entire setup, including code commentary, context, and best practices when working with `client-go`.

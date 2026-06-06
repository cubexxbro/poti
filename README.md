# poti - Global Intelligence Radar

`poti` is a high-performance, autonomous network reconnaissance tool engineered for large-scale asset discovery and global intelligence gathering. It leverages real-time routing registry data to dynamically map sovereign network spaces and verify active services.

## 🏗️ Architecture Overview



`poti` utilizes a multi-stage asynchronous pipeline to convert raw CIDR registries into actionable network reconnaissance targets, optimized for throughput and minimal latency.

## 🚀 Key Features

* **High-Throughput Turbo Engine**: 100+ concurrent worker routines for rapid TCP service verification.
* **Geopolitical Targeting**: Dynamic resolution of IP pools by ISO-3166 country codes.
* **Low-Latency Recon**: Aggressive connection timeouts and decoupled pipeline architecture.
* **Zero-Dependency**: Pure Go implementation for maximum portability.

## 🛠️ Getting Started

### Prerequisites
* **Go (Golang) 1.18+** installed on your system.

### Compilation
To build the `poti` binary for your platform, run the following command in the project root:

```bash
# Build the binary
go build -o poti main.go

# Verify installation
./poti --help

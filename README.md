# poti - Global Intelligence Radar

<p align="center">
  <img src="https://img.shields.io/badge/version-v0.7.3-brightgreen.svg" alt="Version">
  <img src="https://img.shields.io/badge/language-Go-blue.svg" alt="Language">
  <img src="https://img.shields.io/badge/status-Production-blueviolet.svg" alt="Status">
</p>

`poti` is a high-performance, autonomous network reconnaissance tool engineered for large-scale asset discovery and global intelligence gathering. It leverages real-time routing registry data to dynamically map sovereign network spaces and verify active services across global infrastructures.

---

## 🏗️ Architecture Overview

`poti` is engineered for transparency and stability. Our ingestion pipeline processes massive IP subnets through a multi-stage asynchronous lattice generator.



* **Network Sync Stage**: Fetches remote subnet registries with real-time stream feedback (`io.TeeReader`).
* **Lattice Formulation**: High-performance O(n) parsing of CIDR blocks into actionable IP nodes.
* **Concurrency Engine**: Asynchronous worker pool (30+ routines) for rapid TCP handshake evaluation.

Read our full [Engineering Architecture Documentation](docs/ARCHITECTURE.md) to understand the performance optimization strategy.

---

## 🚀 Key Features

* **Geopolitical Target Mapping**: Dynamically resolve network assets by ISO country codes (e.g., `US`, `CN`, `JP`, `KR`, `DE`).
* **Real-time Ingestion Feedback**: Granular progress bars (`[x/total]`) for every phase of the radar operation.
* **High-Speed Audit**: Multi-threaded TCP service verification engine.
* **Storage Hygiene**: Intelligent `~/.poti_cache` management with integrated security purging.
* **Zero-Dependency**: Pure Go (Golang) implementation, ensuring 100% standalone binary portability.

---

## 🛠️ Getting Started

### Prerequisites
* Go (Golang) 1.18+

### Installation
```bash
git clone [https://github.com/cubexxbro/poti.git](https://github.com/cubexxbro/poti.git)
cd poti
go build -o poti
./poti

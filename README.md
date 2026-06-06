# poti - Global Intelligence Radar

`poti` is a high-performance, autonomous network reconnaissance tool designed for large-scale asset discovery and global intelligence gathering. It leverages real-time routing registry data to map sovereign network spaces.

---

## 🚀 Features

* **Autonomous Geopolitical Mapping**: Dynamically resolve network assets by country code (ISO-3166).
* **High-Concurrency Scanning**: Multi-threaded worker engine for rapid TCP service verification.
* **Stream-Based Intelligence**: Real-time progress monitoring and high-fidelity output.
* **Zero-Dependency Architecture**: Built with pure Go (Golang) standard library, ensuring zero dependency bloat and high portability.
* **Storage Hygiene**: Intelligent, automatic cache management for local reconnaissance data.

## 🛠️ Getting Started

### Prerequisites
* Go (Golang) 1.18+

### Installation
```bash
git clone [https://github.com/your-username/poti.git](https://github.com/your-username/poti.git)
cd poti
go build -o poti
./poti

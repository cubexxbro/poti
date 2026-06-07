# 📡 poti v0.8.2 - Global Intelligence Radar & Wireless Auditing Engine

poti is a high-performance, lightweight network intelligence and tactical edge auditing framework engineered entirely in Go. Designed for decentralized security operations, poti cuts out external third-party registry dependencies during perimeter scans by utilizing live mathematical target generation directly from your local shell. 

This branch (`feature-wifi-audit`) introduces on-site hardware-level wireless intelligence mapping, allowing operators to seamlessly pivot between global internet discovery and local radio spectrum risk profiling.

---

## 🚀 Core Architectural Modules

### 1. Global Edge Scanner (Mode 1)
* **Mathematical IPv4 Space Targeting**: Generates random public IPv4 spaces on the fly using secure cryptographic random distributions.
* **Intelligent Subnet Isolation**: Automatically sanitizes and drops connections destined for loopback (`127.0.0.0/8`), private allocations (`10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`), and multicast/experimental spaces (`224.0.0.0/4`).
* **Regulated Concurrency Engine**: Employs an optimized thread-safe worker pipeline managed by atomic synchronization counters to prevent local buffer saturation or thread exhaustion.

### 2. Target IP Intelligence Lookup (Mode 2)
* **Metadata Extraction Registry**: Performs precision queries to dissect external assets by returning structural JSON matrices.
* **Deep Geolocation Mapping**: Exposes localized parameters including autonomous system context (`ISP/Org`), routing zones (`Country`, `Region`, `City`), postal metrics, and geographical coordinate vectors.

### 3. Wireless Wi-Fi Security & Risk Auditing (Mode 3)
* **Local Baseline Inspection**: Probes active internal network interfaces to discover binding addresses and physical default gateways instantly.
* **Cross-Platform Native Radio Binding**: Seamlessly hooks into host operating system wireless utilities to pull active surrounding access points:
  * **macOS (Darwin)**: Intersects the private `Apple80211` framework subsystem resource pipeline via `airport`.
  * **Linux**: Communicates directly with the NetworkManager wireless abstraction layer via `nmcli`.
  * **Windows**: Wraps the native Wireless Local Area Network service layer via `netsh`.
* **Heuristic Threat Classification**: Intercepts wireless metadata frames to assess encryption standard hygiene (e.g., classifying open broadcast capture fields or rogue public hotpots as `HIGH RISK`).

---

## 📊 Feature Matrix Breakdown

| Capability Module | Target Layer Scope | Execution Strategy | Network Requirements | Operating System |
| :--- | :--- | :--- | :--- | :--- |
| **Global Edge Scanner** | Unlimited IPv4 Ranges | Runtime Cryptographic Math | Outbound Internet Pipe | Cross-Platform (Go Runtime) |
| **IP Intelligence Lookup** | Individual Node Target | REST Infrastructure API | Outbound Port 443 | Cross-Platform (Go Runtime) |
| **Wireless Security Audit** | Surrounding AP Radios | Native OS Utilities Parsing | Local Wireless Interface | macOS / Linux / Windows |

---

## ⚙️ Prerequisites & Environment Setup

Ensure your local operating system satisfies the execution criteria before deployment:

### Go Compiler Environment
If compiling from source, use **Go 1.21** or higher.

### System Privileges
* **macOS**: Native hardware radio interface interactions via `/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport` may require elevated permissions depending on sandbox restrictions.
* **Linux**: Ensure `NetworkManager` is installed and running (`nmcli` executable must reside within the system `$PATH`).

---

## 🛠️ Operating Instructions

### 1. Build and Binary Compilation
To compile the architecture into an optimized single binary layout, execute the standard compilation flag sequence inside your terminal workspace:
```bash
go build -ldflags="-s -w" -o poti main.go

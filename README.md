<p align="center">
  <img src="https://raw.githubusercontent.com/cubexxbro/poti/refs/heads/main/logo-poti.png" alt="poti logo" width="320px">
</p>

<p align="center">
  <b>poti v0.8.5</b> — <i>Autonomous System Location Radar & Global Intelligence Engine</i>
</p>

---

poti is a high-performance, lightweight network intelligence and tactical edge auditing framework engineered entirely in Go. Designed for decentralized security operations, poti cuts out external third-party registry dependencies during perimeter scans by utilizing live mathematical target generation directly from your local shell. 

Starting with `v0.8.5`, the architecture introduces an advanced hardware-level radio spectrum profiling and automated tool discovery layer, seamlessly navigating host operating system environments to locate underlying wireless control binaries.

---

## 🚀 Core Architectural Modules

### 1. Global Edge Scanner (Mode 1)
* **Mathematical IPv4 Space Targeting**: Generates random public IPv4 spaces on the fly using secure cryptographic random distributions.
* **Intelligent Subnet Isolation**: Automatically cleans and drops connections destined for loopback (`127.0.0.0/8`), private allocations (`10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`), and multicast/experimental spaces (`224.0.0.0/4`).
* **Regulated Concurrency Engine**: Employs an optimized thread-safe worker pipeline managed by atomic synchronization counters to prevent local buffer saturation or thread exhaustion.

### 2. Target IP Intelligence Lookup (Mode 2)
* **Metadata Extraction Registry**: Performs precision queries to dissect external assets by returning structural JSON matrices.
* **Deep Geolocation Mapping**: Exposes localized parameters including autonomous system context (`ISP/Org`), routing zones (`Country`, `Region`, `City`), postal codes, and geographical coordinate vectors.

### 3. Wireless Wi-Fi Security & Gate Association (Mode 3)
* **Dynamic Subsystem Path Discovery**: Replaces rigid hardcoded binary configurations with a recursive multi-root platform file crawler. If native system binaries (`airport`, `nmcli`, `netsh`) reside outside standard environment variables, the grid engine steps through the filesystem to self-heal and resolve dependencies.
* **Real-time Search Stream Buffer**: Pipes live directory crawling logs directly into the terminal interface using synchronized line carriage metrics, granting full execution status visibility before hardware engagement.
* **Heuristic Threat Classification**: Intercepts wireless metadata frames to assess encryption standard hygiene (e.g., classifying open broadcast capture fields or rogue public hotpots as `HIGH RISK`) and provides local HTTP/HTTPS gateway vectors.

---

## 📊 Feature Matrix Breakdown

| Capability Module | Target Layer Scope | Execution Strategy | Network Requirements | Operating System |
| :--- | :--- | :--- | :--- | :--- |
| **Global Edge Scanner** | Unlimited IPv4 Ranges | Runtime Cryptographic Math | Outbound Internet Pipe | Cross-Platform (Go Runtime) |
| **IP Intelligence Lookup** | Individual Node Target | REST Infrastructure API | Outbound Port 443 | Cross-Platform (Go Runtime) |
| **Wireless Security Audit** | Surrounding AP Radios | Recursive Subsystem Pathing | Local Wireless Interface | macOS / Linux / Windows |

---

## ⚙️ Prerequisites & Environment Setup

Ensure your local operating system satisfies the execution criteria before deployment:

### Go Compiler Environment
If compiling from source, use **Go 1.21** or higher.

### System Privileges
* **macOS / Linux**: Intersecting native hardware radio interface frames via command-line sub-tools often requires elevated root permissions depending on local system sandbox structures (e.g., run with `sudo`).

---

## 🛠️ Operating Instructions

### 1. Build and Binary Compilation
To compile the architecture into an optimized single binary layout, execute the standard compilation flag sequence inside your terminal workspace:
```bash
go build -ldflags="-s -w" -o poti main.go

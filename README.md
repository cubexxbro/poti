<p align="center">
  <img src="https://raw.githubusercontent.com/cubexxbro/poti/refs/heads/main/logo-poti.png" alt="poti logo" width="320px">
</p>

<p align="center">
  <b>poti v0.8.5</b> — <i>Autonomous System Location Radar & Global Intelligence Engine</i>
</p>

---

poti is a high-performance, lightweight network intelligence and tactical edge auditing framework engineered entirely in Go. Designed for decentralized security operations, poti cuts out external third-party registry dependencies during perimeter scans by utilizing live mathematical target generation directly from your local shell. 

Starting with `v0.8.5`, the architecture introduces an advanced hardware-level radio spectrum profiling and automated tool discovery layer, sequentially scanning the entire primary root partition to locate underlying wireless control binaries.

---

## 📥 Instant Deployment & Quick Start (No Compilation Required)

Starting with v0.8.5, pre-compiled production binaries are dynamically attached to each release. You **do not need to install Go or compile anything manually**. 

### For Windows Users (Instant Execution):
1. Download `poti-v0.8.5-windows-x86_64.exe` from the Release Assets.
2. **Drag and drop** the `.exe` file into your CMD or PowerShell window and press **Enter** (or simply double-click it directly) to initialize the radar instantly.

### For macOS & Linux Users:
1. Download the executable matching your hardware architecture from the Release Assets.
2. Open your local terminal, type `chmod +x ` (**make sure to add a trailing space after +x**).
3. **Drag and drop** the downloaded binary file directly from your file manager into the terminal window, then hit **Enter** to authorize execution permissions.
4. From now on, simply **drag the file into any terminal window and press Enter** to launch the framework instantly.

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
* **Deep Root Partition Crawler**: Upgraded to recursively scan the entire primary root partition (`/` or `C:\`) instead of restrictive hardcoded paths. It sequentially steps through system layouts to self-heal and resolve wireless control binary paths.
* **Real-time Search Stream Buffer**: Pipes live directory crawling logs directly into the terminal interface using synchronized `\r` carriage-return metrics, granting full execution status visibility before hardware engagement.
* **Heuristic Threat Classification**: Intercepts wireless metadata frames to assess encryption standard hygiene (e.g., classifying open broadcast capture fields or rogue public hotpots as `HIGH RISK`) and provides local HTTP/HTTPS gateway vectors.

---

## 📊 Feature Matrix Breakdown

| Capability Module | Target Layer Scope | Execution Strategy | Network Requirements | Operating System |
| :--- | :--- | :--- | :--- | :--- |
| **Global Edge Scanner** | Unlimited IPv4 Ranges | Runtime Cryptographic Math | Outbound Internet Pipe | Cross-Platform (Go Runtime) |
| **IP Intelligence Lookup** | Individual Node Target | REST Infrastructure API | Outbound Port 443 | Cross-Platform (Go Runtime) |
| **Wireless Security Audit** | Surrounding AP Radios | Complete Drive Partition Scan | Local Wireless Interface | macOS / Linux / Windows |

---

## 🖥️ User Interface Preview

```text
  _____   ____   _______  _____ 
 |  __ \ / __ \ |__   __||_   _|
 | |__) | |  | |   | |     | |  
 |  ___/| |  | |   | |     | |  
 | |    | |__| |   | |    _| |_ 
 |_|     \____/    |_|   |_____|
                                
[+] poti v0.8.5 - Global Intelligence Radar Engine
================================================================
 [1] Pure Edge-to-Edge Global Random Scanner
 [2] Target IP Intelligence Lookup Engine
 [3] Wireless Wi-Fi Security & Gate Association
 [4] Exit Terminal Framework
================================================================
Select Operational Mode [1-4]:

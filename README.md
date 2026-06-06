<p align="center">
  <img src="./5921829631248591829.jpeg" width="200" alt="poti logo">
</p>

# poti

Active Network Asset Monitoring & High-Performance Space Mapping Engine.

---

## 🚀 Quick Start Guide

`poti` is an independent, high-performance cyber intelligence gathering radar. It allows you to target geographical regions natively using your local execution stack, eliminating third-party API dependencies or token costs.

### 📥 1. Download the Binary
Go to the **Releases** page and download the executable tailored to your hardware architecture:
* **macOS (Apple Silicon M1/M2/M3/M4):** `poti-darwin-arm64`
* **macOS (Intel Core):** `poti-darwin-amd64`
* **Linux (Standard 64-bit Server):** `poti-linux-amd64`
* **Linux (ARM Router/Raspberry Pi):** `poti-linux-arm64`

### 💻 2. Deployment & Permissions

To run the tool, you must grant execution permissions first. If you type `chmod +x poti-darwin-arm64` directly, it might fail with a "No such file or directory" error because the terminal is looking at the wrong folder. 

Use the **Drag-and-Drop Method** to ensure success:

1. Open your terminal, type `chmod +x ` (**Make sure to add a trailing SPACE after `+x`**, do not press Enter yet).
2. Drag the downloaded `poti-darwin-arm64` file from your desktop/finder and **drop it directly into the terminal window**. The terminal will automatically populate the full absolute file path.
3. Press **Enter**.

---

## 🛠️ 3. Command-Line Arguments & Usage

`poti v0.4.0` supports dynamic flag parsing. You can customize the extraction limit, target ports, and target countries directly from your terminal.

### Available Flags
* `-cc` : Target country code (`CN`, `US`, `JP`). Default is `CN`.
* `-limit` : Number of random targets/IPs to extract from the geo-pool. Default is `5`.
* `-ports` : Target ports to audit (comma-separated). Default is `80,443,8080`.

### Execution Examples

#### Example 1: Map Random 10 Nodes in the US Targeting Port 8080 Only
This command extracts 10 crypto-grade random nodes from the US IP pool and runs aggressive TCP protocol handshakes on port 8080 to uncover hidden services:
```bash
# Type "./" then drag and drop the file into terminal, then append the flags:
[dragged_file_path] -cc US -limit 10 -ports 8080

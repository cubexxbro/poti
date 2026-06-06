<p align="center">
  <img src="https://raw.githubusercontent.com/cubexxbro/poti/refs/heads/main/logo-poti.png" width="500" alt="poti logo">
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

## 🛠️ 3. Smart Interactive Execution

`poti v0.5.0` features a completely interactive terminal Wizard. You do not need to memorize or type complex CLI command flags anymore. Just fire it up and answer the prompt questions step-by-step!

### How to Run:
1. Type `./` in your terminal.
2. Drag and drop the `poti-darwin-arm64` binary into the terminal window.
3. Press **Enter**.

### Interactive Walkthrough Example:
```text
          _   _ 
 ___  ___| |_(_)
| '_ \/ _ \ __| |
| |_) | (_) | |_| |
| .__/ \___/\__|_|
|_|              

[+] poti Engine - Smart Interactive Radar v0.5.0
--------------------------------------------------
[?] Select Target Geographical Region:
  1. China (CN)
  2. United States (US)
  3. Japan (JP)
Choose option (1-3, default 1): 2

[?] How many random IPs to extract? (default 5): 10

[?] Enter ports to scan (comma-separated, default 80,443,8080): 8080

--------------------------------------------------
[*] Launching Radar Mode...
[*] Target Region  : US
[*] Extracted Nodes: 10 targets
[*] Target Ports   : [8080]
--------------------------------------------------
[*] Mapping live space assets...

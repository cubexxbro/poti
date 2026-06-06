<p align="center">
  <img src="https://raw.githubusercontent.com/cubexxbro/poti/refs/heads/main/poti-logo.png" width="200" alt="poti logo">
</p>

# poti

Active Network Asset Monitoring & High-Performance Space Mapping Engine.

---

## 🚀 Quick Start Guide

`poti` is a highly concurrent CLI tool designed for network asset scanning and banner grabbing. Follow the instructions below to get it running on your system.

### 📥 1. Download the Binary
Go to the **Releases** page on the right side of this repository and download the executable tailored to your hardware architecture:
* **macOS (Apple Silicon M1/M2/M3/M4):** `poti-darwin-arm64`
* **macOS (Intel Core):** `poti-darwin-amd64`
* **Linux (Standard 64-bit Server):** `poti-linux-amd64`
* **Linux (ARM Router/Raspberry Pi):** `poti-linux-arm64`

### 💻 2. Deployment & Execution

Open your terminal and execute the following commands based on your deployment path.

#### For Linux / macOS:
```bash
# Navigate to the folder where the binary is located (e.g., Desktop)
cd ~/Desktop

# Grant execution permissions to the engine
chmod +x poti-darwin-arm64

### 💻 2. Deployment & Execution

Open your terminal and use **one of the two methods** below to grant permissions and run the tool. 

#### Method A: The Drag-and-Drop Method (Easiest & Recommended)
If you directly type `chmod +x poti-darwin-arm64`, it might fail with a "No such file or directory" error because the terminal is not looking at the right folder. Use the mouse to assist instead:

1. Type `chmod +x ` in your terminal (**Make sure to add a trailing SPACE after `+x`**, do not press Enter yet).
2. Drag the downloaded `poti-darwin-arm64` file from your desktop/finder and **drop it directly into the terminal window**. The terminal will automatically populate the full file path.
3. Press **Enter**.
4. Type `./` and drag-and-drop the file into the terminal **for the second time**, then press **Enter** to fire up the engine.

#### Method B: The Absolute Path Method
Alternatively, you can manually navigate to the file's directory before executing the command:

```bash
# Move to the Desktop directory where the binary is saved
cd ~/Desktop

# Grant execution permissions to the binary inside this folder
chmod +x poti-darwin-arm64

# Fire up the engine
./poti-darwin-arm64
# Fire up the engine
./poti-darwin-arm64

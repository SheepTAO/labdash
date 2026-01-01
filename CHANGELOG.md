# Changelog

All notable changes to LabDash will be documented in this file.

## [1.0.0] - 2026-01-01

### 🎉 Initial Release

First stable release of LabDash - Lab Monitoring & Documentation System.

### ✨ Features

#### Monitoring System
- **Real-time CPU monitoring**: Load percentage, model info, cores/threads count
- **GPU monitoring**: NVIDIA GPU support via nvidia-smi
  - Utilization, memory usage, temperature, power, fan speed
  - Multi-GPU support with per-device and aggregated stats
- **RAM monitoring**: Usage, total capacity, memory type detection
- **Disk monitoring**: Partition usage, user home directory sizes
- **Historical charts**: Configurable history length for CPU, GPU, RAM
- **System info**: Hostname, OS, kernel, uptime, load average

#### Idle Mode (Power Saving)
- **Adaptive monitoring intervals**: Automatically reduces frequency when inactive
- **Configurable thresholds**:
  - `idleTimeoutSec`: Time before entering idle (0 = never idle)
  - `idleIntervalCRGSec`: CRG interval when idle (10-600s)
  - `idleIntervalDiskHours`: Disk scan interval when idle (0.5-48h)
- **State transition logging**: Monitor mode changes in logs
- **Auto-resume**: Automatically returns to active on next connection

#### Documentation System
- **Markdown rendering**: Full GFM support with tables, lists, links
- **Math equations**: KaTeX support for inline and block equations
- **Code highlighting**: Syntax highlighting for 100+ languages
- **File tree navigation**: Hierarchical folder structure
- **File metadata**: Owner and modification time display
- **Configurable depth**: Control max folder nesting level
- **Default document**: Set homepage document

#### Modern UI
- **Glassmorphism design**: Beautiful translucent cards with blur effects
- **Smooth animations**: Framer Motion powered transitions
- **Responsive layout**: Works on desktop and tablet
- **Dark/light mode ready**: Clean color scheme
- **Dynamic gradients**: Animated floating orbs

### 🔧 Configuration

Fully configurable via `/etc/labdash/config.json`:
- Server port, project name, lab name
- Documentation directory and depth
- Monitoring intervals (active and idle modes)
- History chart sizes
- Disk partition and user filters

### 📦 Installation

- **Single binary**: No external dependencies
- **Systemd service**: Automatic startup and management
- **Interactive installer**: Guided setup process
- **Uninstall script**: Clean removal with safety checks

### 🛠️ Technical Stack

**Backend:**
- Go 1.25+
- Standard library only (no external dependencies)
- RESTful JSON API
- Systemd service integration

**Frontend:**
- React 19.2
- Vite build system
- Tailwind CSS + Framer Motion
- React Markdown + KaTeX + Prism

### 📝 Documentation

- Comprehensive README with installation guide
- Example configuration files
- Templates for documentation
- Inline code comments

### 🐛 Known Limitations

- GPU monitoring requires NVIDIA GPUs with nvidia-smi
- Service must run as root to access all system metrics
- Disk scanning can be slow on large file systems
- Documentation limited to Markdown files only

### 🔒 Security

- Service runs as root (required for system monitoring)
- Path traversal protection for documentation files
- Markdown-only file serving restriction
- PrivateTmp enabled in systemd service

### 📊 Performance

- Lightweight: ~10-20MB RAM usage when idle
- Efficient: 200ms CPU load sampling for accuracy
- Scalable: Handles thousands of documentation files
- Adaptive: Reduces resource usage when inactive

---

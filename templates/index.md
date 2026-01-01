# Welcome to LabDash

Welcome to your lab monitoring and documentation system!

## Getting Started

LabDash provides real-time system monitoring and a powerful documentation system for your lab environment.

### Features

- **Real-time Monitoring**: Track CPU, GPU, RAM, and disk usage across all lab servers
- **Documentation System**: Organize your lab documentation, protocols, and notes with Markdown
- **Beautiful UI**: Modern glassmorphism design with smooth animations
- **Math Support**: Write equations with KaTeX (inline: $E=mc^2$, block: $$\int_0^\infty e^{-x^2} dx$$)
- **Code Highlighting**: Syntax highlighting for 100+ programming languages
- **File Tree**: Hierarchical navigation with nested folders

## Quick Start

### 1. View System Stats
Click **"Monitor Dashboard"** to see real-time metrics:
- CPU load and usage percentage
- GPU utilization (if available)
- Memory consumption
- Disk space by partition
- User home directory sizes

### 2. Add Documentation
Create Markdown files in your documentation directory:

```bash
# Personal installation
cd ~/labdash-docs

# Shared installation
cd /home/labdash/docs

# Create new doc
vim my-protocol.md
```

Files will be automatically indexed and displayed in the sidebar.

### 3. Organize with Folders
Create folders to organize your docs:

```bash
mkdir experiments
mkdir protocols
mkdir notes
```

LabDash supports nested folders up to the configured depth (default: 4 levels).

## File Sharing (Shared Installation Only)

If you installed LabDash with the shared directory option (`/home/labdash/docs`), all users can read and write files:

### How It Works
- **Directory**: `/home/labdash/docs` (and entire `/home/labdash/`)
- **Permissions**: `1777` (sticky bit enabled)
- **Owner**: `labdash:labdash`
- **Access**: All users can read/write, but only file owners can delete their files

### File Permissions Best Practices

**Files are automatically shareable:**
```bash
# Create a file in /home/labdash/docs
touch /home/labdash/docs/new-doc.md
# Everyone can read and write it
```

**To make a file read-only:**
```bash
chmod 644 /home/labdash/docs/readonly-doc.md
# Owner can write, others can only read
```

**Share files from your home directory:**
```bash
cp ~/my-research.md /home/labdash/docs/
# File becomes accessible to everyone
```

**Sticky bit protection:**
- Only the file owner can delete their own files
- Others can read/write but cannot delete

### Collaboration Workflow
1. **Upload**: Copy files to `/home/labdash/docs` - automatically accessible to all
2. **Edit**: Use any editor (vim, nano, VSCode) directly
3. **Download**: Copy to your home directory - `cp /home/labdash/docs/file.md ~/`

## Markdown Features

### Supported Syntax
- **Headers**: `# H1`, `## H2`, etc.
- **Lists**: Unordered (`-`, `*`) and ordered (`1.`, `2.`)
- **Links**: `[text](url)`
- **Images**: `![alt](path)` - relative paths supported
- **Tables**: GitHub Flavored Markdown tables
- **Blockquotes**: `> quote`
- **Code Blocks**: Triple backticks with language

### Math Equations
Inline: `$E=mc^2$` renders as $E=mc^2$

Block: `$$\frac{-b \pm \sqrt{b^2-4ac}}{2a}$$`, Renders as:
$$
\frac{-b \pm \sqrt{b^2-4ac}}{2a}
$$

### Code Highlighting

```python
def hello_world():
    print("Hello, LabDash!")
```

## Configuration

Edit `/etc/labdash/config.json` to customize settings:

```json
{
  "projectName": "My Lab",
  "labName": "Research Laboratory",
  "port": 8088,
  "docsPath": "/home/labdash/docs",
  "docsDepth": 4,
  "defaultDoc": "index.md",
  "monitor": {
    "intervalCRGSec": 2,
    "intervalDiskHours": 1,
    "idleTimeoutSec": 60,
    "idleIntervalCRGSec": 300,
    "idleIntervalDiskHours": 6,
    "historyCPU": 20,
    "historyGPU": 20,
    "historyRAM": 20
  },
  "disk": {
    "includedPartitions": {
      "/": "System",
      "/home": "User Home"
    },
    "ignoredPartitions": ["/boot", "/boot/efi"],
    "ignoredUsers": ["root", "syslog"],
    "maxUsersToList": 12
  }
}
```

**After editing, restart the service:**
```bash
sudo systemctl restart labdash
```

## System Commands

```bash
# Service management
sudo systemctl start labdash
sudo systemctl stop labdash
sudo systemctl restart labdash
sudo systemctl status labdash

# View logs
sudo journalctl -u labdash -f
```

## Troubleshooting

### Can't write to `/home/labdash/docs`?
```bash
# Check directory permissions
ls -ld /home/labdash/docs
# Should show: drwxrwxrwt (1777)

# If wrong, contact your system administrator
```

### Service won't start?
```bash
# Check logs
sudo journalctl -u labdash -n 50

# Verify config
sudo cat /etc/labdash/config.json

# Check if port is available
sudo lsof -i :8088
```

### Files not showing up?
- Ensure files have `.md` extension
- Check file permissions (must be readable by labdash service)
- Verify docsPath in config is correct
- Check docsDepth if files are deeply nested

**Happy documenting! 🥳**

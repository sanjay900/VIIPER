# Installation

## Requirements

VIIPER relies on USBIP. You must have USBIP installed on your system.

### Linux

#### Ubuntu/Debian

```bash
sudo apt install linux-tools-generic
```

[Ubuntu USBIP Manual](https://manpages.ubuntu.com/manpages/noble/man8/usbip.8.html)

#### Arch Linux

```bash
sudo pacman -S usbip
```

[Arch Wiki: USBIP](https://wiki.archlinux.org/title/USB/IP)

### Windows

[usbip-win2](https://github.com/vadimgrn/usbip-win2) is by far the most complete implementation of USBIP for Windows (comes with a **SIGNED** kernel mode driver).

## Installing VIIPER

### Pre-built Binaries

Download the latest release from the [GitHub Releases](https://github.com/Alia5/VIIPER/releases) page.

### Building from Source

#### Prerequisites

- [Go](https://go.dev/) 1.25 or newer
- USBIP installed
- (Optional) [Make](https://www.gnu.org/software/make/)
    - Linux/macOS: Usually pre-installed
    - Windows: `winget install ezwinports.make`

#### Build Steps

```bash
git clone https://github.com/Alia5/VIIPER.git
cd VIIPER
make build
```

The compiled binary will be in `dist/viiper` (or `dist/viiper.exe` on Windows).

**Additional build targets:**

```bash
make help          # Show all available make targets
make test          # Run tests
```

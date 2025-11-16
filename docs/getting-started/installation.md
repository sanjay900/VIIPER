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
- USBIP installed (for testing)

#### Build Steps

```bash
git clone https://github.com/Alia5/VIIPER.git
cd VIIPER/viiper
go build -o viiper ./cmd/viiper
```

The compiled binary will be in the current directory.

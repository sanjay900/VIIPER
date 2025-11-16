# Configuration

VIIPER can be configured via command-line flags or environment variables.

## Environment Variables

All command-line flags have corresponding environment variables for easier deployment and configuration management.

### Global Configuration

| Environment Variable | CLI Flag | Default | Description |
|---------------------|----------|---------|-------------|
| `VIIPER_LOG_LEVEL` | `--log.level` | `info` | Logging level: `trace`, `debug`, `info`, `warn`, `error` |
| `VIIPER_LOG_FILE` | `--log.file` | (none) | Log file path (logs only to console if not set) |
| `VIIPER_LOG_RAW_FILE` | `--log.raw-file` | (none) | Raw packet log file path |

### Server Configuration

| Environment Variable | CLI Flag | Default | Description |
|---------------------|----------|---------|-------------|
| `VIIPER_USB_ADDR` | `--usb.addr` | `:3241` | USBIP server listen address |
| `VIIPER_API_ADDR` | `--api.addr` | `:3242` | API server listen address (empty = disabled) |
| `VIIPER_API_DEVICE_HANDLER_TIMEOUT` | `--api.device-handler-timeout` | `5s` | Device handler auto-cleanup timeout |
| `VIIPER_CONNECTION_TIMEOUT` | `--connection-timeout` | `30s` | Connection operation timeout |

### Proxy Configuration

| Environment Variable | CLI Flag | Default | Description |
|---------------------|----------|---------|-------------|
| `VIIPER_PROXY_ADDR` | `--listen-addr` | `:3241` | Proxy listen address |
| `VIIPER_PROXY_UPSTREAM` | `--upstream` | (required) | Upstream USBIP server address |
| `VIIPER_PROXY_TIMEOUT` | `--connection-timeout` | `30s` | Connection timeout |

## Configuration Examples

### Using Environment Variables

Create a `.env` file or export variables:

```bash
export VIIPER_LOG_LEVEL=debug
export VIIPER_USB_ADDR=:3241
export VIIPER_API_ADDR=:3242
export VIIPER_LOG_FILE=/var/log/viiper.log
```

Then run:

```bash
viiper server
```

### Systemd Service

Example systemd service file for running VIIPER as a service:

```ini
[Unit]
Description=VIIPER USBIP Server
After=network.target

[Service]
Type=simple
User=viiper
Group=viiper
Environment="VIIPER_LOG_LEVEL=info"
Environment="VIIPER_LOG_FILE=/var/log/viiper/viiper.log"
Environment="VIIPER_USB_ADDR=:3241"
Environment="VIIPER_API_ADDR=:3242"
ExecStart=/usr/local/bin/viiper server
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

## Configuration Priority

When both CLI flags and environment variables are set, CLI flags take precedence:

1. **CLI flags** (highest priority)
2. **Environment variables**
3. **Default values** (lowest priority)

## See Also

- [Server Command](server.md)
- [Proxy Command](proxy.md)

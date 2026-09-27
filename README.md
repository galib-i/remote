# Remote Control
Your phone as a wireless keyboard, trackpad, and volume control.

<p align="center">
  <img width="200" alt="Remote Control UI" src="https://github.com/user-attachments/assets/db11b11b-a0ce-4f2d-9110-bcd48c66b714" />
</p>

> [!NOTE]
> Developed and tested on CachyOS (KDE, Wayland) and Windows 11.

> [!IMPORTANT]
> Requires sudo/admin privileges to automatically open the port from the startup scripts (default 8080).<br> Intended for local, personal use.

## Get Started
### Linux
Download the latest pre-compiled binaries and startup scripts: [remote-server-linux](https://github.com/galib-i/remote/releases/download/latest/remote-server-linux) and [start-remote.sh](https://github.com/galib-i/remote/releases/download/latest/start-remote.sh)
```bash
sudo modprobe uinput                          # Wake up the virtual input system
chmod +x remote-server-linux start-remote.sh  # Grant execution permissions
sudo chmod 666 /dev/uinput                    # Temporarily unlock the virtual device file
./start-remote.sh                             # Run the server (then scan the QR code)
```
 
### Windows
Download the latest pre-compiled binaries and startup scripts: [remote-server-windows](https://github.com/galib-i/remote/releases/download/latest/remote-server-windows.exe) and [start-remote.ps1](https://github.com/galib-i/remote/releases/download/latest/start-remote.ps1)

```powershell
.\start-remote.ps1  # Run the server (then scan the QR code)
```
Or, just right-click `start-remote.ps1` and select Run with Powershell.

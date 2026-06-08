# Remote Control
Control your computer's volume, mouse and keyboard from your phone.

> [!NOTE]
> Developed and tested on Windows 11 and CachyOS (KDE, Wayland).

> [!IMPORTANT]
> Requires sudo/admin privileges to automatically open the port from the startup scripts (default 8080). Intended for local, personal use.

<table style="width: 100%;">
  <tr>
    <td style="width: 45%;">
      <img src="https://github.com/user-attachments/assets/db11b11b-a0ce-4f2d-9110-bcd48c66b714" alt="Example" width="300">
    </td>
    <td style="width: 55%;">
      <ul>
        <li>Adjust or mute the system volume.</li>
        <li>Send keystrokes.</li>
        <li>Move cursor and double-tap via the <i>touchpad</i>.</li>
        <li>Use left and right click buttons.</li>
      </ul>
    </td>
  </tr>
</table>

## Get Started

### Linux
1. Download the latest pre-compiled binaries and startup scripts: [remote-server-linux](https://github.com/galib-i/remote/releases/download/latest/remote-server-linux) and [start-remote.sh](https://github.com/galib-i/remote/releases/download/latest/start-remote.sh)
2. Open a terminal in that folder and wake up the virtual input system:
   ```bash
   sudo modprobe uinput
   ```
3. Grant execution permissions to the downloaded files:
   ```bash
   chmod +x remote-server-linux start-remote.sh
   ```
4. Temporarily unlock the virtual device file:
   ```bash
   sudo chmod 666 /dev/uinput
   ```
5. Run and scan the QR code with a phone connected to the same Wi-Fi:
   ```bash
   ./start-remote.sh
   ```
 
### Windows
1. Download the latest pre-compiled binaries and startup scripts: [remote-server-windows](https://github.com/galib-i/remote/releases/download/latest/remote-server-windows.exe) and [start-remote.ps1](https://github.com/galib-i/remote/releases/download/latest/start-remote.ps1)
2. Right-click to run, or open the terminal and navigate to it.
3. Scan the QR code with a phone connected to the same Wi-Fi.

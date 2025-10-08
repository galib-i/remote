# Remote Control
Control your Windows PC's volume, mouse and keyboard from your phone with a simple web-based toy project.

> [!IMPORTANT]
> Only Windows is supported. Ensure both PC and phone are connected to the same network: only intended for local use.

## Setup
### Steps
1. **Clone** *(or download)* **this repository:**  
    ```
    git clone https://github.com/yourusername/remote.git
    ```
2. **Install dependencies**  
    ```
    go mod tidy
    ```
3. **Run the server**  
    ```
    go run ./cmd/server
    ```
4. **Scan the QR code shown with your phone**
## Usage
<table style="width: 100%;">
  <tr>
    <td style="width: 45%;">
      <img src="https://github.com/user-attachments/assets/db11b11b-a0ce-4f2d-9110-bcd48c66b714" alt="Example" width="300">
    </td>
    <td style="width: 55%;">
      <ul>
        <li>Use the buttons to control volume</li>
        <li>Use the keyboard input (top right) to send keystrokes</li>
        <li>Use the touchpad area to move the mouse</li>
        <li>Left/Right click buttons for mouse clicks</li>
        <li><em>Double tap also inputs a Left-click</em></li>
      </ul>
    </td>
  </tr>
</table>


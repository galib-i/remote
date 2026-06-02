const wsProtocol = window.location.protocol === "https:" ? "wss:" : "ws:";
const wsUrl = `${wsProtocol}//${window.location.host}/api/ws`;
const socket = new WebSocket(wsUrl);
const sendMessage = (payload) => {
  if (socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(payload));
  } else {
    console.warn("WebSocket is not open. Message dropped:", payload);
  }
};
const BUTTON_MESSAGES = {
  volumeUpBtn: { action: "volume", direction: "up" },
  volumeDownBtn: { action: "volume", direction: "down" },
  muteVolumeBtn: { action: "toggle-mute" },
  leftClickBtn: { action: "click", side: "left" },
  rightClickBtn: { action: "click", side: "right" },
};

const DOUBLE_TAP_DELAY = 300; // milliseconds

const keyboardInput = document.getElementById("keyboardInput");
const touchpad = document.getElementById("touchpad");

Object.entries(BUTTON_MESSAGES).forEach(([id, messagePayload]) => {
  const btn = document.getElementById(id);
  if (btn) {
    btn.addEventListener("click", () => sendMessage(messagePayload));
  }
});

// Handle special keys that don't produce a character input (Enter, Backspace)
keyboardInput.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    e.preventDefault();
    sendMessage({ action: "press-key", text: "enter" });
  } else if (e.key === "Backspace") {
    e.preventDefault();
    sendMessage({ action: "press-key", text: "backspace" });
  }
});

// Handle regular characters
keyboardInput.addEventListener("input", (e) => {
  const text = e.target.value;
  const lastChar = text.slice(-1);

  if (
    lastChar.length === 1 &&
    lastChar.charCodeAt(0) >= 32 &&
    lastChar.charCodeAt(0) <= 126
  ) {
    sendMessage({ action: "press-key", text: lastChar });
  }

  e.target.value = "";
});


let isTouchActive = false;
let lastTouchPos = { x: 0, y: 0 };
let lastTapTime = 0; // Used for double-tap detection

// Calculate relative touch position
const getTouchPos = (touch) => {
  const rect = touchpad.getBoundingClientRect();

  return {
    x: touch.clientX - rect.left,
    y: touch.clientY - rect.top,
    width: rect.width,
    height: rect.height,
  };
};

// Tracks initial cursor movement tracking & double tap detection
touchpad.addEventListener("touchstart", (e) => {
  isTouchActive = true;
  lastTouchPos = getTouchPos(e.touches[0]);

  const currentTime = Date.now();
  if (currentTime - lastTapTime < DOUBLE_TAP_DELAY) {
    sendMessage({ action: "click", side: "left" });
    lastTapTime = 0;
  } else {
    lastTapTime = currentTime;
  }
});

// Sends deltas to the server
touchpad.addEventListener("touchmove", (e) => {
  if (!isTouchActive) return;
  e.preventDefault(); // Prevents scrolling

  const { x, y, width, height } = getTouchPos(e.touches[0]);
  if (x < 0 || y < 0 || x > width || y > height) return; // Ignore touches outside of bounds

  sendMessage({
    action: "move-cursor",
    deltaX: x - lastTouchPos.x,
    deltaY: y - lastTouchPos.y,
  });

  lastTouchPos = { x, y };
});

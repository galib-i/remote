const post = (url, data = null) =>
  fetch(url, {
    method: "POST",
    headers: data ? { "Content-Type": "application/json" } : undefined,
    body: data ? JSON.stringify(data) : null,
  }).catch(console.error);

const BUTTON_ENDPOINTS = {
  volumeUpBtn: "/volume/up",
  volumeDownBtn: "/volume/down",
  muteVolumeBtn: "/toggle-mute",
  leftClickBtn: "/click/left",
  rightClickBtn: "/click/right",
};

const DOUBLE_TAP_DELAY = 300; // milliseconds

const keyboardInput = document.getElementById("keyboardInput");
const touchpad = document.getElementById("touchpad");

// =============================================================================
// BUTTON HANDLERS
// =============================================================================

Object.entries(BUTTON_ENDPOINTS).forEach(([id, endpoint]) => {
  const btn = document.getElementById(id);

  if (btn) {
    btn.addEventListener("click", () => post(endpoint));
  }
});

// =============================================================================
// KEYBOARD HANDLERS
// =============================================================================

// Handle special keys that don't produce a character input (Enter, Backspace)
keyboardInput.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    e.preventDefault(); // Prevent form submission or other default actions
    post(`/press-key?text=enter`);
  } else if (e.key === "Backspace") {
    e.preventDefault();
    post(`/press-key?text=backspace`);
  }
});
// Handle regular characters
keyboardInput.addEventListener("input", (e) => {
  const text = e.target.value;
  const lastChar = text.slice(-1);

  // Filter out emojis and complex Unicode characters
  if (
    lastChar.length === 1 &&
    lastChar.charCodeAt(0) >= 32 &&
    lastChar.charCodeAt(0) <= 126
  ) {
    post(`/press-key?text=${encodeURIComponent(lastChar)}`);
  }

  e.target.value = "";
});

// =============================================================================
// TOUCHPAD HANDLERS
// =============================================================================

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
    post("/click/left");
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

  post("/move-cursor", {
    deltaX: x - lastTouchPos.x,
    deltaY: y - lastTouchPos.y,
  });

  lastTouchPos = { x, y };
});

touchpad.addEventListener("touchend", () => {
  isTouchActive = false;
});

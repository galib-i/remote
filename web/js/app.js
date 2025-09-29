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
  LeftClickBtn: "/click/left",
  RightClickBtn: "/click/right",
};

const DOUBLE_TAP_DELAY = 300; // milliseconds

Object.entries(BUTTON_ENDPOINTS).forEach(([id, endpoint]) => {
  const btn = document.getElementById(id);

  if (btn) {
    btn.addEventListener("click", () => post(endpoint));
  }
});

// Keyboard input functionality
const keyboardInput = document.getElementById("keyboardInput");
const clearInputBtn = document.getElementById("clearInput");

if (keyboardInput) {
  keyboardInput.addEventListener("keydown", (e) => {
    // Handle special keys that don't produce a standard character input
    if (e.key === "Enter") {
      e.preventDefault(); // Prevent form submission or other default actions
      post(`/press-key?text=enter`);
    } else if (e.key === "Backspace") {
      e.preventDefault();
      post(`/press-key?text=backspace`);
    }
  });

  keyboardInput.addEventListener("input", (e) => {
    const text = e.target.value;
    const lastChar = text.slice(-1);

    post(`/press-key?text=${encodeURIComponent(lastChar)}`);
    e.target.value = "";
  });
}

if (clearInputBtn) {
  clearInputBtn.addEventListener("click", () => {
    if (keyboardInput) {
      keyboardInput.value = "";
      keyboardInput.focus();
    }
  });
}

const canvas = document.getElementById("canvas");
if (canvas) {
  let isActive = false;
  let lastPos = { x: 0, y: 0 };
  let lastTapTime = 0;

  const getTouchPos = (touch) => {
    const rect = canvas.getBoundingClientRect();
    return {
      x: touch.clientX - rect.left,
      y: touch.clientY - rect.top,
      width: rect.width,
      height: rect.height,
    };
  };

  canvas.addEventListener("touchstart", (e) => {
    isActive = true;
    lastPos = getTouchPos(e.touches[0]);

    // Double-tap detection
    const currentTime = Date.now();
    if (currentTime - lastTapTime < DOUBLE_TAP_DELAY) {
      post("/click/left");
      lastTapTime = 0; // Reset to prevent taps from stacking
    } else {
      lastTapTime = currentTime;
    }
  });

  canvas.addEventListener("touchmove", (e) => {
    if (!isActive) return;
    e.preventDefault();

    const { x, y, width, height } = getTouchPos(e.touches[0]);
    if (x < 0 || y < 0 || x > width || y > height) return;

    post("/move-cursor", {
      deltaX: x - lastPos.x,
      deltaY: y - lastPos.y,
    });

    lastPos = { x, y };
  });

  canvas.addEventListener("touchend", () => {
    isActive = false;
  });
}

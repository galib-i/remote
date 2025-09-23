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
  OpenKeyboardBtn: "/",
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
  keyboardInput.addEventListener("input", (e) => {
    const inputValue = e.target.value;
    const lastChar = inputValue.slice(-1).toLowerCase();

    // Only send if it's a letter
    if (lastChar.match(/[a-z]/)) {
      post(`/press-key/${lastChar}`);
    }

    // Clear the input after sending
    setTimeout(() => {
      e.target.value = "";
    }, 100);
  });

  // Alternative: Send on keydown for immediate response
  keyboardInput.addEventListener("keydown", (e) => {
    const key = e.key.toLowerCase();

    // Only send if it's a letter
    if (key.match(/^[a-z]$/)) {
      post(`/press-key/${key}`);
      e.preventDefault(); // Prevent default to avoid duplicate sends
    }
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

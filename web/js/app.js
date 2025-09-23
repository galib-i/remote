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

    post("/move-mouse", {
      deltaX: x - lastPos.x,
      deltaY: y - lastPos.y,
    });

    lastPos = { x, y };
  });

  canvas.addEventListener("touchend", () => {
    isActive = false;
  });
}

const volumeUpButton = document.getElementById("volumeUpBtn");
const volumeDownButton = document.getElementById("volumeDownBtn");
const muteVolumeButton = document.getElementById("muteVolumeBtn");
const unmuteVolumeButton = document.getElementById("unmuteVolumeBtn");

volumeUpButton.addEventListener("click", () => {
  fetch("/volume-up", { method: "POST" }).catch((err) => console.error(err));
});

volumeDownButton.addEventListener("click", () => {
  fetch("/volume-down", { method: "POST" }).catch((err) => console.error(err));
});

muteVolumeButton.addEventListener("click", () => {
  fetch("/mute-volume", { method: "POST" }).catch((err) => console.error(err));
});

unmuteVolumeButton.addEventListener("click", () => {
  fetch("/unmute-volume", { method: "POST" }).catch((err) =>
    console.error(err)
  );
});

let isMousepadActive = false;
let oldX = 0;
let oldY = 0;

const canvas = document.getElementById("canvas");

canvas.addEventListener("touchstart", (e) => {
  isMousepadActive = true;
  const touch = e.touches[0];
  const rect = canvas.getBoundingClientRect();
  oldX = touch.clientX - rect.left;
  oldY = touch.clientY - rect.top;
});

canvas.addEventListener("touchmove", (e) => {
  if (isMousepadActive) {
    e.preventDefault();
    const touch = e.touches[0];
    const rect = canvas.getBoundingClientRect();
    const x = touch.clientX - rect.left;
    const y = touch.clientY - rect.top;

    if (x >= 0 && y >= 0 && x <= rect.width && y <= rect.height) {
      const deltaX = x - oldX;
      const deltaY = y - oldY;

      fetch("/move-mouse", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ deltaX, deltaY }),
      }).catch((err) => console.error(err));
      oldX = x;
      oldY = y;
    }
  }
});

canvas.addEventListener("touchend", () => {
  isMousepadActive = false;
});

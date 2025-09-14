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

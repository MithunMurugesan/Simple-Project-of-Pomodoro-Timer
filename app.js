function fetchStatus() {
  fetch("/status").then(res => res.json()).then(data => {
    const min = Math.floor(data.secondsLeft / 60).toString().padStart(2, "0");
    const sec = (data.secondsLeft % 60).toString().padStart(2, "0");
    document.getElementById("timer").textContent = `${min}:${sec}`;
    document.getElementById("mode").textContent = data.mode === "focus" ? "Focus" : "Break";
  });
}

function start() {
  fetch("/start", { method: "POST" });
}

function stop() {
  fetch("/stop", { method: "POST" });
}

function reset() {
  fetch("/reset", { method: "POST" });
}

setInterval(fetchStatus, 1000);
fetchStatus();

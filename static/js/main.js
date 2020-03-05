const before_play_container = document.getElementsByClassName("before_play")[0];
const host_button = document.getElementById("host");
const join_button = document.getElementById("join");
const append_container = document.getElementById("append_here");

function setupField() {
    before_play_container.parentNode.removeChild(before_play_container);
}

function extracted(ws) {
    ws.addEventListener("close", () => {
        console.log("ws closed");
    });
    setupField();
    ws.addEventListener("message", (message) => {
        function signal_readyState() {
            setTimeout(() => {
                console.log(ws.readyState);
                if (ws.readyState !== 3) {
                    signal_readyState();
                }
            }, 5000);
        }

        if (message.data === "end_of_game") {
            ws.close(1000);
            signal_readyState();
        }
        const div = document.createElement('div');
        div.innerText = message.data;
        append_container.appendChild(div);
    });
}

host_button.addEventListener("click", () => {
    const ws = new WebSocket("ws://localhost:8080/game/2");
    ws.addEventListener("open", () => {
        ws.send("old man's here");
    });
    extracted(ws);
});

join_button.addEventListener("click", () => {
    const ws = new WebSocket("ws://localhost:8080/join");
    ws.addEventListener("open", () => {
        ws.send("player here");
    });
    extracted(ws);
});

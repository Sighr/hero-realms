const before_play_container = document.getElementsByClassName("before_play")[0];
const host_button = document.getElementById("host");
const join_button = document.getElementById("join");
const append_container = document.getElementById("append_here");
const input = document.getElementById("input");
const send = document.getElementById("send");

function setupField() {
    before_play_container.parentNode.removeChild(before_play_container);
}

let sock;
send.addEventListener("click", () => {
    sock.send(input.value);
    input.value = "";
});

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
    const ws = new WebSocket("ws://abc86d6f.ngrok.io/game/sample/2");
    sock = ws;
    ws.addEventListener("open", () => {
        ws.send("old man's here");
    });
    extracted(ws);
});

join_button.addEventListener("click", () => {
    const ws = new WebSocket("ws://abc86d6f.ngrok.io/join/sample");
    sock = ws;
    ws.addEventListener("open", () => {
        ws.send("player here");
    });
    extracted(ws);
});

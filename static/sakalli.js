class Sakalli {
    constructor(host, id) {
        this.host = host;
        this.id = id;

    }

    init(id) {
        let conn;
        if(id){
            this.id = id
        }
        console.log(this.id)
        if (window["WebSocket"]) {
            conn = new WebSocket("ws://" + this.host + this.id);

            conn.onclose = function (evt) {
                console.log("connection closed")
            };

            conn.onopen = function (evt) {
                console.log("Connection open ...");
            };

            conn.onmessage = function (evt) {
                // console.log("Received Message: " + evt.data);
                // document.getElementById("message-json").innerText += evt.data;
                let event = new CustomEvent(
                    "sakalliNotification", 
                    {
                        detail: evt,
                        bubbles: true,
                        cancelable: true
                    }
                );
                document.dispatchEvent(event)
            };

            window.onbeforeunload = function () {
                conn.send("closing ...")
                conn.close()
            }
        } else {
            console.log("browser does not support websocket")
        }
    }
}
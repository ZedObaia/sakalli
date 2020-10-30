class Sakalli extends EventTarget {
    constructor(host, id) {
        super()
        this.host = host;
        this.id = id;

    }

    init(id) {
        let conn;
        if (id) {
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
                this.dispatchEvent(new CustomEvent("sakalliNotification", {
                    detail: JSON.parse(evt.data),
                    bubbles: true,
                    cancelable: true
                }))
                // let event = new CustomEvent(
                //     "sakalliNotification", 
                //     {
                //         detail: JSON.parse(evt.data),
                //         bubbles: true,
                //         cancelable: true
                //     }
                // );
                // document.dispatchEvent(event)
            }.bind(this);

            window.onbeforeunload = function () {
                conn.close()
            }
        } else {
            console.log("browser does not support websocket")
        }
    }
}
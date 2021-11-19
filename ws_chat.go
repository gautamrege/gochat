package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gautamrege/gochat/api"
	"github.com/gorilla/websocket"
)

type WsChat struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{} // use default options

func (ws *WsChat) Input() (string, error) {
	_, message, err := ws.conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return "", err
	}
	return string(message), nil
}

func (ws *WsChat) Render(message string) error {
	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}

func (ws *WsChat) Moderate(req api.ChatRequest) {
	Moderation <- req
	return
}

func StartWsChat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// update the global
	WS.conn = c

	for {
		textInput, err := WS.Input()
		if err != nil {
			WS.Render("Unable to get input.. exiting!")
			break
		}
		parseAndExecInput(&WS, "ws", textInput)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/WsChat")
}

func WSRun() {
	addr := ":8080"
	http.HandleFunc("/WsChat", StartWsChat)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    ws = new WebSocket("{{.}}");
    ws.onmessage = function(evt) {
	const data = JSON.parse(evt.data);
        print(data)
    }

    var print = function(data) {
	if (data.abuse) {
	    var d = document.getElementById(data.chatid);
	    d.style.color = 'red';
	} else if (data.users) {
            var d = document.createElement("div");
			d.textContent = data.users; 
            output.appendChild(d);
            output.scroll(0, output.scrollHeight);
	} else {
            var d = document.createElement("div");
            d.id = data.chatid;
	    d.textContent = "@" + data.from.name + ": " + data.message;
            output.appendChild(d);
            output.scroll(0, output.scrollHeight);
	}
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }

        ws.send(input.value);
        return false;
    };

    document.getElementById("users").onclick = function(evt) {
        if (!ws) {
            return false;
        }

	ws.send("/users");
	return false;
    };

});
</script>
</head>
<body>
<table>
<tr>
<td valign="top" width="50%">
<button id="users">List Users</button>
<form>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))

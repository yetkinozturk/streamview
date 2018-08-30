package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

// IndexView present index.html with WS.
func indexView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("index.html")
	_check(err)
	var homeTemplate = template.Must(template.New("").Parse(string(content)))
	homeTemplate.Execute(w, "ws://"+r.Host+"/streamview")
}

// statsView show stream traffic statistics.
func statsView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("stats.html")
	_check(err)
	var homeTemplate = template.Must(template.New("").Parse(string(content)))
	homeTemplate.Execute(w, "")
}

// pkgConfigView configures presentation of recevied UDP Messages on websocket.
// It is useful for custom serialization
func pkgConfigView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("config.html")
	_check(err)
	var homeTemplate = template.Must(template.New("").Parse(string(content)))
	homeTemplate.Execute(w, "")
}

func (sv *StreamView) webSock(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	c, err := upgrader.Upgrade(w, r, nil)
	_check(err)
	defer c.Close()

	//fmt.Println("starting stream channel")
	for pkg := range sv.netChan {
		PKGFormat(&pkg)
		//fmt.Println(pkg.message)
		err = c.WriteMessage(websocket.TextMessage, []byte(pkg.message))
		_check(err)
	}
}

func (sv *StreamView) httpRouteConfig() {
	http.HandleFunc("/pkgconfig", pkgConfigView)
	http.HandleFunc("/stats", statsView)
	http.HandleFunc("/", indexView)
	http.HandleFunc("/streamview", sv.webSock)
}

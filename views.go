package streamview

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

func (sv *StreamView) httpRouteConfig() {
	http.HandleFunc("/", sv.indexView)
	http.HandleFunc("/pkgconfig", sv.pkgConfigView)
	http.HandleFunc("/stats", sv.statsView)
	http.HandleFunc("/streamview", sv.webSock)
}

func (sv *StreamView)indexView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("templates/index.html")
	_check(err)
	var indexTemplate = template.Must(template.New("").Parse(string(content)))
	indexTemplate.Execute(w, "ws://"+r.Host+"/streamview")
}

func (sv *StreamView)statsView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("templates/stats.html")
	_check(err)
	var statsTemplate = template.Must(template.New("").Parse(string(content)))
	statsTemplate.Execute(w, "")
}

func (sv *StreamView)pkgConfigView(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("templates/config.html")
	_check(err)
	var configTemplate = template.Must(template.New("").Parse(string(content)))
	configTemplate.Execute(w, "")
}

func (sv *StreamView) webSock(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	_check(err)
	defer c.Close()

	for pkg := range sv.netChan {
		PKGFormat(&pkg)
		err = c.WriteMessage(websocket.TextMessage, []byte(pkg.message))
		_check(err)
	}
}



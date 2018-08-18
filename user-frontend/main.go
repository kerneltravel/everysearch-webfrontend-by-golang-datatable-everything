package main

import (
	//	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"net/url"
	"strings"

	"github.com/zserge/webview"
)

var indexHTML = `
<!doctype html>
<html>
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
	</head>
	<body>
	  <h1>Hello, world</h1>
	</body>
</html>
`

func runLocalHTTP() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(indexHTML))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	url := "http://" + ln.Addr().String()
	w := webview.New(webview.Settings{
		Title: "Loaded: Local HTTP Server",
		URL:   url,
	})
	defer w.Exit()
	w.Run()
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func runLocalFile() {

	w := webview.New(webview.Settings{
		Title: "EverySearch-web. Local template file demo",
		URL:   "file://"+getCurrentDirectory()+"/index.html",
		Resizable :true,
		Debug :true,
	})
	//w.SetFullscreen(true)
	defer w.Exit()
	w.Run()
}

func runDataURL() {
	w := webview.New(webview.Settings{
		Title: "Loaded: Data URL",
		URL:   "data:text/html," + url.PathEscape(indexHTML),
	})
	defer w.Exit()
	w.Run()
}

func runInjectJS() {
	w := webview.New(webview.Settings{
		Title: "Loaded: Injected via JavaScript",
		URL:   `data:text/html,<html><script type="text/javascript"></script></html>`,
	})
	defer w.Exit()
	w.Dispatch(func() {
		w.Eval(`document.body.innerHTML = "<h1>Hello, world</h1>";`)
	})
	w.Run()
}

func main() {
	//runLocalHTTP()
	runLocalFile()
	//runDataURL()
	//runInjectJS()
}

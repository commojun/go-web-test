package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	api      string
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("public", t.filename)))
	})
	data := map[string]interface{}{
		"Api": t.api,
	}
	t.templ.Execute(w, data)
}

func main() {
	var (
		addr = flag.String("addr", ":8081", "Webサイトのアドレス")
		api  = flag.String("api", "localhost:8080", "apiのアドレス")
	)
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", &templateHandler{filename: "index.html", api: *api})
	mux.Handle("/new", &templateHandler{filename: "new.html", api: *api})
	mux.Handle("/view", &templateHandler{filename: "view.html", api: *api})
	log.Println("Webサイトのアドレス:", *addr)
	http.ListenAndServe(*addr, mux)
}

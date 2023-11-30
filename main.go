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
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var port = flag.String("port", ":8080", "서버 포트에 대한 옵션입니다. EX) :8080")
	flag.Parse()

	r := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("서버 시작 중 입니다. 사용중인 포트 : ", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

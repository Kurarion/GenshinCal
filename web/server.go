package web

import (
	"bytes"
	"fmt"
	data "genshincal/genshindata"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var rootHtmlTemplate *template.Template

const logState = true

func init() {
	rootHtmlTemplate = template.Must(template.ParseFiles("./html/home.html"))
}

//Start Server启动
func Start(addr string) {
	http.Handle("/js/", http.FileServer(http.Dir("")))
	http.HandleFunc("/api/character", character)
	http.HandleFunc("/", root)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func log(r *http.Request) {
	if logState {
		fmt.Printf("[%s]:<%s>\n", time.Now().Local(), r.URL.Path)
	}
}

func character(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		buf := bytes.NewBuffer(make([]byte, 0, 100))
		buf.ReadFrom(r.Body)
		//
		fmt.Println("test:", buf.String())
		x := data.GetAvatar(buf.String())
		io.WriteString(w, strconv.FormatFloat(x.LevelMap["90"].Hp, 'g', 10, 64))
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	log(r)
	// r.ParseForm()
	// fmt.Println(r.Form)

	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		rootHtmlTemplate.Execute(w, data.GetAvatarMap())
	case "POST":

	}
}

func extraHandlerFunc(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("{%s}%v\n", time.Now().Local(), r)
		}
	}()
	return f
}

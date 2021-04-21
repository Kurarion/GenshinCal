package web

import (
	"encoding/json"
	"fmt"
	data "genshincal/genshindata"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var rootHtmlTemplate *template.Template
var initData InitData

const logState = true

func init() {
	initData = InitData{
		Avatar: data.GetAvatarMap(),
		Weapon: data.GetWeaponMap(),
	}
	rootHtmlTemplate = template.Must(template.ParseGlob("./html/*.html"))
}

//Start Server启动
func Start(addr string) {
	http.Handle("/js/", http.FileServer(http.Dir("")))
	http.HandleFunc("/api/character", character)
	http.HandleFunc("/api/weapon", weapon)
	http.HandleFunc("/api/weaponSkillAffix", weaponSkillAffix)
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
		if err := r.ParseForm(); err != nil {
			fmt.Println("parse form error. err: ", err)
		}

		id := r.Form["id"][0]
		level := r.Form["level"][0]

		x := data.GetAvatar(id).LevelMap[level]
		js, _ := json.Marshal(*x)

		io.WriteString(w, string(js))
	}
}

func weapon(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println("parse form error. err: ", err)
		}

		id := r.Form["id"][0]
		level := r.Form["level"][0]

		x := data.GetWeapon(id).LevelMap[level]
		js, _ := json.Marshal(*x)

		io.WriteString(w, string(js))
	}
}

func weaponSkillAffix(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println("parse form error. err: ", err)
		}

		id := r.Form["id"][0]
		level, err := strconv.Atoi(r.Form["level"][0])
		if err != nil {
			level = 1
		}

		x := data.GetWeapon(id).SkillAffixMap[level]
		js, _ := json.Marshal(x)

		io.WriteString(w, string(js))
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	log(r)
	// r.ParseForm()
	// fmt.Println(r.Form)

	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		rootHtmlTemplate.ExecuteTemplate(w, "home.html", initData)
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

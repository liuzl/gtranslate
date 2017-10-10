package main

import (
	"encoding/json"
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/golang/glog"
	"github.com/liudanking/gotranslate"
	"net/http"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
	languages  = []string{
		"ar",
		"bn",
		"zh-CN",
		"zh-TW",
		"nl",
		"fr",
		"de",
		"hi",
		"id",
		"it",
		"ja",
		"ko",
		"ms",
		"pl",
		"pt",
		"ru",
		"es",
		"th",
		"tr",
		"uk",
		"vi"}
)

func mustEncode(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}

type Item struct {
	Language string `json:"lang"`
	Text     string `json:"text"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s  method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	input := r.FormValue("input")
	var texts []Item
	for _, lang := range languages {
		ret, err := gotranslate.Translate("auto", lang, input)
		if err != nil {
			mustEncode(w, struct {
				Status  string `json:"status"`
				Message string `json:"message"`
				Lang    string `json:"lang"`
			}{Status: "error", Message: err.Error(), Lang: lang})
			return
		}
		texts = append(texts, Item{Language: lang, Text: ret.Sentences[0].Trans})
	}
	mustEncode(w, texts)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.HandleFunc("/api/", TranslateHandler)
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *serverAddr)
	glog.Error(http.ListenAndServe(*serverAddr, nil))
}

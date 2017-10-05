package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func handle(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", dump)
}

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r)

		req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s/bar", r.Host), nil)
		req.Header = r.Header

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		dump, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		res.Body.Close()
		fmt.Fprintf(w, "%s", dump)
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r)
	})

	http.ListenAndServe(":8080", nil)
}

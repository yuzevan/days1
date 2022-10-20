package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/",http.Fileserver(http.Dir("./public"))))

	route.HandleFunc("helo", hello world).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-tambahan/{id}", blogTambahan).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("GET")


	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

    func helloworld(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello world!"))

    func home(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")

		var tmpl, err = template.ParseFiles("view/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("message : " + err.Error()))
			return
			
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, nil)
	}

	func addblog(w http.http.ResponseWriter, r *http.Request)  {
		err := r.ParseForm()
		if err !=nil {
			log.Fatal(err)
			
		}
		fmt.Println("Title : " + r.PostForm.Get("inputTitle"))
		fmt.Println("Content : " + r.PostForm.Get("inputContent"))

		http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
		
	}
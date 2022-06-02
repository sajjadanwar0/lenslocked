package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my website</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact me at page </h1>To get in touch, email me at "+
		"<a href=\"mailto:sajjadanwar200@gmail.com\">sajjadanwar200@gmail.com</a>")
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, `<h1>FAQ Page</h1>
     <ul>
       <li>
       is there is a free version? Yes there is
       </li>
      </ul>

`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)

	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on port :4000")
	http.ListenAndServe(":4000", router)

}

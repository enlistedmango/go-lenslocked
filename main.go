package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil { // This error is if templates don't parse correctly
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return // The return statement is important because it stops the code from progressing.
	}
	err = tpl.Execute(w, nil)
	if err != nil { // This error is if we can't execute for whichever reason
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executig the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:mango@notmango.dev\">mango@notmango.dev</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		"<h1>FAQs</h1>"+
			"<p>Q: Is there a free version?<p>"+
			"<p>A: Yes! We offer a free trial for 30 days on any paid plans.</p>"+
			"<p>Q: What are your support hours?</p>"+
			"<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</p?"+
			"<p>Q: How do I contact support?</p>"+
			"<p>A: Email us - <a href=\"mailto:support@lenslocked.com\">Support@lenslocked.com</a>")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

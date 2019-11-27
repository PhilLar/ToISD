package main

import (
	"html/template"
	"net/http"

	"github.com/PhilLar/ToISD/handlers"
)

//var tmpl *template.Template
var tmpl = make(map[string]*template.Template)
var products = make([]*handlers.Stock, 0)
var bought = make(map[int]int)
var user = handlers.User{Name: "Test_User"}
var userProducts = make([]*handlers.Stock, 0)

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	tmpl["email"] = template.Must(template.ParseFiles("templates/emailForm.html"))
	tmpl["register"] = template.Must(template.ParseFiles("templates/registerForm.html"))
	tmpl["first.html"] = template.Must(template.ParseFiles("templates/first.html", "templates/layout.html"))
	tmpl["footer.html"] = template.Must(template.ParseFiles("templates/footer.html", "templates/layout.html"))
	tmpl["index-product"] = template.Must(template.ParseFiles("templates/product.html", "templates/layout.html", "templates/footer.html"))
	tmpl["user-space"] = template.Must(template.ParseFiles("templates/userProduct.html", "templates/layout.html", "templates/userFeaturing.html"))
	tmpl["index-product-bought-success"] = template.Must(template.ParseFiles("templates/product-bought-succes.html", "templates/layout.html", "templates/footer.html"))

	for i := 0; i < 9; i++ {
		stock := &handlers.Stock{
			ID:    i + 1,
			Title: "Ubique",
			Descr: "There goes description",
			Price: i + 10,
		}
		products = append(products, stock)
		bought[stock.ID] = 0
	}
}

func main() {

	http.HandleFunc("/", handlers.HandlerIndex(tmpl, products, bought, user))
	// http.HandleFunc("/email", handlerEmail)
	// http.HandleFunc("/register", handlerRegister)
	// http.HandleFunc("/user", handlerUserSpace)
	http.ListenAndServe(":3000", nil)
}

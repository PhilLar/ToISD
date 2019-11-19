package main

import (
	"html/template"
	"net/http"
	"strconv"
)

//var tmpl *template.Template
var tmpl = make(map[string]*template.Template)
var products = make([]*Stock, 0)
var bought = make(map[int]int)
var user = User{Name: "Test_User"}

type Data struct {
	Products []*Stock
	User     User
}

type Stock struct {
	ID     int
	Title  string
	Descr  string
	Price  int
	Bought int
}

type User struct {
	Name     string
	Email    string
	Address  string
	Passport string
	Phone    string
}

type EmailRecipient struct {
	Name  string
	Email string
}

func init() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	tmpl["email"] = template.Must(template.ParseFiles("templates/emailForm.html"))
	tmpl["register"] = template.Must(template.ParseFiles("templates/registerForm.html"))
	tmpl["first.html"] = template.Must(template.ParseFiles("templates/first.html", "templates/layout.html"))
	tmpl["footer.html"] = template.Must(template.ParseFiles("templates/footer.html", "templates/layout.html"))
	tmpl["index-product"] = template.Must(template.ParseFiles("templates/product.html", "templates/layout.html", "templates/footer.html"))
	tmpl["index-product-bought-success"] = template.Must(template.ParseFiles("templates/product-bought-succes.html", "templates/layout.html", "templates/footer.html"))

	products = make([]*Stock, 0)
	for i := 0; i < 9; i++ {
		stock := &Stock{
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

	http.HandleFunc("/", handler)
	http.HandleFunc("/email", handlerEmail)
	http.HandleFunc("/register", handlerRegister)
	http.ListenAndServe(":3000", nil)
}

func handlerEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var err error

	if r.Method != http.MethodPost {
		err = tmpl["email"].ExecuteTemplate(w, "emailForm", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	details := EmailRecipient{
		Email: r.FormValue("email"),
		Name:  r.FormValue("name"),
	}

	_ = details

	tmpl["email"].ExecuteTemplate(w, "emailForm", struct{ Success bool }{true})
}

func handlerRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var err error
	if r.Method != http.MethodPost {
		err = tmpl["register"].ExecuteTemplate(w, "registerForm", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	details := User{
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Address:  r.FormValue("address"),
		Passport: r.FormValue("passport"),
		Phone:    r.FormValue("phone"),
	}

	_ = details

	tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var err error
	if r.Method == http.MethodPost {
		ID, err := strconv.Atoi(r.FormValue("ID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		act := r.FormValue("submit")
		if act == "Buy" {
			bought[ID]++
		} else if act == "Sell" && bought[ID] > 0 {
			bought[ID]--
		}
	}
	for _, prod := range products {
		prod.Bought = bought[prod.ID]
	}
	data := Data{Products: products, User: user}
	err = tmpl["index-product"].ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// if r.Method != http.MethodGet {
	// 	err = tmpl["index-product-bought-success"].ExecuteTemplate(w, "layout", nil)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// }
}

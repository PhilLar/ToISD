package main

import (
	"html/template"
	"net/http"
)

//var tmpl *template.Template
var tmpl = make(map[string]*template.Template)

type Data struct {
	Products []Stock
}

type Stock struct {
	Title string
	Descr string
	Price int
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var err error

	tmpl["first.html"] = template.Must(template.ParseFiles("templates/first.html", "templates/layout.html"))
	tmpl["footer.html"] = template.Must(template.ParseFiles("templates/footer.html", "templates/layout.html"))
	tmpl["index-product"] = template.Must(template.ParseFiles("templates/product.html", "templates/layout.html", "templates/footer.html"))
	//err = tmpl["footer.html"].ExecuteTemplate(w, "layout", nil)
	products := make([]Stock, 0)
	for i := 0; i < 9; i++ {
		products = append(products, Stock{
			Title: "A",
			Descr: "Slices can also be copy’d. Here we create an empty slice" +
				"c of the same length as s and copy into c from s.",
			Price: i,
		})
	}
	data := Data{Products: products}
	err = tmpl["index-product"].ExecuteTemplate(w, "layout", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

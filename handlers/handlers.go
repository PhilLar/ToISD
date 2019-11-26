package handlers

import (
	"net/http"
	"strconv"
	"html/template"
)

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
	Amount int
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

func HandlerEmail(tmpl map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

// func HandlerEmail(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	var err error

// 	if r.Method != http.MethodPost {
// 		err = tmpl["email"].ExecuteTemplate(w, "emailForm", nil)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	details := EmailRecipient{
// 		Email: r.FormValue("email"),
// 		Name:  r.FormValue("name"),
// 	}

// 	_ = details

// 	tmpl["email"].ExecuteTemplate(w, "emailForm", struct{ Success bool }{true})
// }

// func HandlerRegister(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	var err error
// 	if r.Method != http.MethodPost {
// 		err = tmpl["register"].ExecuteTemplate(w, "registerForm", nil)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	details := User{
// 		Email:    r.FormValue("email"),
// 		Name:     r.FormValue("name"),
// 		Address:  r.FormValue("address"),
// 		Passport: r.FormValue("passport"),
// 		Phone:    r.FormValue("phone"),
// 	}

// 	_ = details

// 	tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
// }

// func HandlerUserSpace(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	var err error
// 	if r.Method == http.MethodPost {
// 		title := r.FormValue("stock-title")
// 		descr := r.FormValue("stock-descr")
// 		price, err := strconv.Atoi(r.FormValue("stock-price"))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		amount, err := strconv.Atoi(r.FormValue("stock-amount"))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		userProducts = append(userProducts, &Stock{
// 			Title:  title,
// 			Descr:  descr,
// 			Amount: amount,
// 			Price:  price,
// 		})
// 	}

// 	data := Data{Products: userProducts, User: user}
// 	err = tmpl["user-space"].ExecuteTemplate(w, "layout", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func HandlerIndex(tmpl map[string]*template.Template, products []*Stock, bought map[int]int, user User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
}


// func Handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	var err error
// 	if r.Method == http.MethodPost {
// 		ID, err := strconv.Atoi(r.FormValue("ID"))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		act := r.FormValue("submit")
// 		if act == "Buy" {
// 			bought[ID]++
// 		} else if act == "Sell" && bought[ID] > 0 {
// 			bought[ID]--
// 		}
// 	}
// 	for _, prod := range products {
// 		prod.Bought = bought[prod.ID]
// 	}
// 	data := Data{Products: products, User: user}
// 	err = tmpl["index-product"].ExecuteTemplate(w, "layout", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

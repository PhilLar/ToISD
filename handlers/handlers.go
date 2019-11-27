package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/PhilLar/ToISD/models"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Products []*models.Stock
	User     models.User
}

type UserStore interface {
	InsertUser(user models.User) error
}

type Env struct {
	Store UserStore
}

type Stock struct {
	ID     int
	Title  string
	Descr  string
	Price  int
	Bought int
	Amount int
}

// type User struct {
// 	Name     string
// 	Email    string
// 	Address  string
// 	Passport string
// 	Phone    string
// }

// type EmailRecipient struct {
// 	Name  string
// 	Email string
// }

// func HandlerEmail(tmpl map[string]*template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		var err error

// 		if r.Method != http.MethodPost {
// 			err = tmpl["email"].ExecuteTemplate(w, "emailForm", nil)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 			return
// 		}

// 		details := EmailRecipient{
// 			Email: r.FormValue("email"),
// 			Name:  r.FormValue("name"),
// 		}

// 		_ = details

// 		tmpl["email"].ExecuteTemplate(w, "emailForm", struct{ Success bool }{true})
// 	}
// }

func (env *Env) HandlerRegister(tmpl map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		var err error
		if r.Method != http.MethodPost {
			err = tmpl["register"].ExecuteTemplate(w, "registerForm", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		hash_password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = env.Store.InsertUser(models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("name"),
			Address:  r.FormValue("address"),
			Password: string(hash_password),
			Phone:    r.FormValue("phone"),
		})
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
	}
}

func (env *Env) HandlerLogin(tmpl map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		var err error
		if r.Method != http.MethodPost {
			err = tmpl["login"].ExecuteTemplate(w, "loginForm", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		hash_password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = env.Store.InsertUser(models.User{
			Email:    r.FormValue("email"),
			Name:     r.FormValue("name"),
			Address:  r.FormValue("address"),
			Password: string(hash_password),
			Phone:    r.FormValue("phone"),
		})
		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
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

func HandlerIndex(tmpl map[string]*template.Template, products []*models.Stock, bought map[int]int, user models.User) http.HandlerFunc {
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

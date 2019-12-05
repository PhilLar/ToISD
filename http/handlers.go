package http

import (
	"net/http"
	"time"

	"github.com/PhilLar/ToISD/models"
)

type Data struct {
	Products []*models.Stock
	User     models.User
}

type Stock struct {
	ID     int
	Title  string
	Descr  string
	Price  int
	Bought int
	Amount int
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Phone    string `json:"phone_number"`
}

func (env *Env) RegisterUser(w http.ResponseWriter, r *http.Request) {

	request := new(registerRequest)
	err := processRequest(r, request)
	if err != nil {
		respondError(w, err)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Address:  request.Address,
		Password: request.Password,
		Phone:    request.Phone,
	}

	err = env.Store.CreateUser(user)
	if err != nil {
		respondError(w, err)
		return
	}

	respondOK(w, user)
}

type authRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (env *Env) AuthUser(w http.ResponseWriter, r *http.Request) {

	request := new(authRequest)
	err := processRequest(r, request)
	if err != nil {
		respondError(w, err)
		return
	}

	creds := models.Creds{}

	tokenString, err := env.Store.AuthUser(creds)
	if err != nil {
		respondError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(5 * time.Minute),
	})

	respondOK(w, nil)
}

// func (env *Env) RegisdterUser(tmpl map[string]*template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		var err error
// 		if r.Method != http.MethodPost {
// 			err = tmpl["register"].ExecuteTemplate(w, "registerForm", nil)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 			return
// 		}

// 		hash_password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		err = env.Store.InsertUser(models.User{
// 			Email:    r.FormValue("email"),
// 			Name:     r.FormValue("name"),
// 			Address:  r.FormValue("address"),
// 			Password: string(hash_password),
// 			Phone:    r.FormValue("phone"),
// 		})
// 		if err != nil {
// 			http.Error(w, err.Error(), 400)
// 		}

// 		tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
// 	}
// }

// func (env *Env) HandlerLogin(tmpl map[string]*template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		var err error
// 		if r.Method != http.MethodPost {
// 			err = tmpl["login"].ExecuteTemplate(w, "loginForm", nil)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 			return
// 		}

// 		hash_password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		err = env.Store.InsertUser(models.User{
// 			Email:    r.FormValue("email"),
// 			Name:     r.FormValue("name"),
// 			Address:  r.FormValue("address"),
// 			Password: string(hash_password),
// 			Phone:    r.FormValue("phone"),
// 		})
// 		if err != nil {
// 			http.Error(w, err.Error(), 400)
// 		}

// 		tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
// 	}
// }

// // func HandlerEmail(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "text/html")
// // 	var err error

// // 	if r.Method != http.MethodPost {
// // 		err = tmpl["email"].ExecuteTemplate(w, "emailForm", nil)
// // 		if err != nil {
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 		}
// // 		return
// // 	}

// // 	details := EmailRecipient{
// // 		Email: r.FormValue("email"),
// // 		Name:  r.FormValue("name"),
// // 	}

// // 	_ = details

// // 	tmpl["email"].ExecuteTemplate(w, "emailForm", struct{ Success bool }{true})
// // }

// // func HandlerRegister(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "text/html")
// // 	var err error
// // 	if r.Method != http.MethodPost {
// // 		err = tmpl["register"].ExecuteTemplate(w, "registerForm", nil)
// // 		if err != nil {
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 		}
// // 		return
// // 	}

// // 	details := User{
// // 		Email:    r.FormValue("email"),
// // 		Name:     r.FormValue("name"),
// // 		Address:  r.FormValue("address"),
// // 		Passport: r.FormValue("passport"),
// // 		Phone:    r.FormValue("phone"),
// // 	}

// // 	_ = details

// // 	tmpl["register"].ExecuteTemplate(w, "registerForm", struct{ Success bool }{true})
// // }

// // func HandlerUserSpace(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "text/html")
// // 	var err error
// // 	if r.Method == http.MethodPost {
// // 		title := r.FormValue("stock-title")
// // 		descr := r.FormValue("stock-descr")
// // 		price, err := strconv.Atoi(r.FormValue("stock-price"))
// // 		if err != nil {
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 		}
// // 		amount, err := strconv.Atoi(r.FormValue("stock-amount"))
// // 		if err != nil {
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 		}

// // 		userProducts = append(userProducts, &Stock{
// // 			Title:  title,
// // 			Descr:  descr,
// // 			Amount: amount,
// // 			Price:  price,
// // 		})
// // 	}

// // 	data := Data{Products: userProducts, User: user}
// // 	err = tmpl["user-space"].ExecuteTemplate(w, "layout", data)
// // 	if err != nil {
// // 		http.Error(w, err.Error(), http.StatusInternalServerError)
// // 	}
// // }

// func HandlerIndex(tmpl map[string]*template.Template, products []*models.Stock, bought map[int]int, user models.User) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		var err error
// 		if r.Method == http.MethodPost {
// 			ID, err := strconv.Atoi(r.FormValue("ID"))
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}
// 			act := r.FormValue("submit")
// 			if act == "Buy" {
// 				bought[ID]++
// 			} else if act == "Sell" && bought[ID] > 0 {
// 				bought[ID]--
// 			}
// 		}
// 		for _, prod := range products {
// 			prod.Bought = bought[prod.ID]
// 		}
// 		data := Data{Products: products, User: user}
// 		err = tmpl["index-product"].ExecuteTemplate(w, "layout", data)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}
// }

// // func Handler(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "text/html")
// // 	var err error
// // 	if r.Method == http.MethodPost {
// // 		ID, err := strconv.Atoi(r.FormValue("ID"))
// // 		if err != nil {
// // 			http.Error(w, err.Error(), http.StatusInternalServerError)
// // 		}
// // 		act := r.FormValue("submit")
// // 		if act == "Buy" {
// // 			bought[ID]++
// // 		} else if act == "Sell" && bought[ID] > 0 {
// // 			bought[ID]--
// // 		}
// // 	}
// // 	for _, prod := range products {
// // 		prod.Bought = bought[prod.ID]
// // 	}
// // 	data := Data{Products: products, User: user}
// // 	err = tmpl["index-product"].ExecuteTemplate(w, "layout", data)
// // 	if err != nil {
// // 		http.Error(w, err.Error(), http.StatusInternalServerError)
// // 	}
// // }

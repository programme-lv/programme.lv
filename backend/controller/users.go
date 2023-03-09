package controller

import (
	"encoding/json"
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strconv"
)

func (c *Controller) listUsers(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var users []models.User
	err := c.database.Model(&models.User{}).Find(&users).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var user models.User

	var err error
	user.ID, err = strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.database.Model(&user).Take(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the values provided are sane
	if len(user.Password) < 8 {
		http.Error(w, "Parolei jābūt vismaz 8 rakstzīmju garai.", http.StatusBadRequest)
		return
	}
	if len(user.Password) > 64 {
		http.Error(w, "Parole nedrīkst būt garāka par 64 simboliem.", http.StatusBadRequest)
		return
	}
	if len(user.Username) < 3 {
		http.Error(w, "Lietotājvārdam jābūt vismaz 3 rakstzīmju garam.", http.StatusBadRequest)
		return
	}
	if len(user.Username) > 20 {
		http.Error(w, "Lietotājvārds nedrīkst būt garāks par 20 simboliem.", http.StatusBadRequest)
		return
	}
	if regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(user.Email) == false {
		http.Error(w, "Epasts nav pareizā formātā.", http.StatusBadRequest)
		return
	}
	if len(user.Email) > 50 {
		http.Error(w, "Epasts nedrīkst būt garāks par 50 simboliem.", http.StatusBadRequest)
		return
	}
	if len(user.FirstName) < 2 || len(user.FirstName) > 50 {
		http.Error(w, "Vārds nedrīkst būt garāks par 50 simboliem un nedrīkst būt īsāks par 2 simboliem.", http.StatusBadRequest)
		return
	}
	if len(user.LastName) < 2 || len(user.LastName) > 50 {
		http.Error(w, "Uzvārds nedrīkst būt garāks par 50 simboliem un nedrīkst būt īsāks par 2 simboliem.", http.StatusBadRequest)
		return
	}

	user.Password += c.passwordSalt

	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(hashedPassword)
	user.Admin = false

	// check if database contains a user with the same email
	var existingUser models.User
	err = c.database.Model(&models.User{}).Where("email = ?", user.Email).Take(&existingUser).Error
	if err == nil {
		http.Error(w, "Lietotājs ar šo epastu jau eksistē.", http.StatusBadRequest)
		return
	}

	err = c.database.Model(&models.User{}).Where("username = ?", user.Username).Take(&existingUser).Error
	if err == nil {
		http.Error(w, "Lietotājs ar šo lietotājvārdu jau eksistē.", http.StatusBadRequest)
		return
	}

	err = c.database.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.infoLogger.Printf("Created user: %v", user)

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Controller) loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingUser models.User
	err = c.database.Model(&models.User{}).Where("email = ?", user.Email).Take(&existingUser).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password+c.passwordSalt))
	if err != nil {
		http.Error(w, "Epasts vai parole nav pareiza.", http.StatusBadRequest)
		return
	}

	c.infoLogger.Println(c.sessions.Get(r.Context(), "user_id"))

	err = c.sessions.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.sessions.Put(r.Context(), "user_id", existingUser.ID)

	existingUser.Password = "noslepums"
	resp, err := json.Marshal(existingUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.infoLogger.Printf("Logged in user: %v", existingUser)

	// send the response
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

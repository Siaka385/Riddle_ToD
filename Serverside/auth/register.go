package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Player struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ConfirmPass    string `json:"ConfirmPass"`
	SelectedAvatar string `json:"SelectedAvatar"`
}

type Message struct {
	Response string `json:"response"`
}

// var avatars string=[]string{""}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type, expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading")
		os.Exit(1)
	}

	var user Player
	json.Unmarshal(body, &user)
	fmt.Println(user)

	var response Message
	response = Message{"ok"}

	if user.Password != user.ConfirmPass {
		response = Message{"passwords do not match"}
	}
	// if !CheckEmail(user.Email) {
	// 	response = map[string]string{
	// 		"message": "Email already Exist",
	// 	}
	// }
	// if !CheckUsername(user.Username) {
	// 	response = map[string]string{
	// 		"message": "username already exist",
	// 	}
	// }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// helper/session.go

package helper

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Om@r3010"))

func StoreUserInSession(w http.ResponseWriter, r *http.Request, userID int, userType string) {
	session, err := store.Get(r, "Clinic-session")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		log.Printf("Error retrieving session: %v", err)
		return
	}

	session.Values["UserID"] = userID
	session.Values["userType"] = userType

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		log.Printf("Error saving session: %v", err)
		return
	}
}


func GetUserFromSession(r *http.Request) (int, string) {
	session, _ := store.Get(r, "Clinic-session")
	userID, okUserID := session.Values["UserID"].(int)
	userType, okUserType := session.Values["userType"].(string)
	if !okUserID {
		log.Println("UserID not found in session")
		return 0, ""
	}

	if !okUserType {
		log.Println("UserType not found in session")
		return userID, ""
	}

	return userID, userType
}

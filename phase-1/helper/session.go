package helper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Om@r3010"))

const sessionName = "Clinic-session"
const userIDKey = "UserID"
const userTypeKey = "userType"

func StoreUserInSession(w http.ResponseWriter, r *http.Request, userID int, userType string) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		log.Printf("Error retrieving session: %v", err)
		return
	}

	session.Values[userIDKey] = userID
	session.Values[userTypeKey] = userType

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		log.Printf("Error saving session: %v", err)
		return
	}
	fmt.Println(sessionName)
	fmt.Println(userIDKey)
	fmt.Println(userID)
	fmt.Println(userTypeKey)
	fmt.Println(userType)
}

func GetUserFromSession(r *http.Request) (int, string) {
	session, _ := store.Get(r, sessionName)

	userID, ok := session.Values[userIDKey].(int)
	fmt.Println(sessionName)
	fmt.Println(userIDKey)
	fmt.Println(userID)
	fmt.Println(userTypeKey)
	if !ok {
		log.Println("UserID not found in session or of incorrect type")
		return 0, ""
	}

	userType, ok := session.Values[userTypeKey].(string)
	fmt.Println(userType)
	if !ok {
		log.Println("UserType not found in session or of incorrect type")
		return 0, ""
	}
	
	return userID, userType
}


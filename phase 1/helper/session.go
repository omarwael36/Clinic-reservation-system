// helper/session.go

package helper

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Om@r3010"))

func StoreUserInSession(w http.ResponseWriter, r *http.Request, doctorID int, userType string) {
	session, err := store.Get(r, "Clinic-session")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		log.Printf("Error retrieving session: %v", err)
		return
	}

	session.Values["DoctorID"] = doctorID
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
	doctorID, okDoctorID := session.Values["DoctorID"].(int)
	userType, okUserType := session.Values["userType"].(string)
	if !okDoctorID {
		log.Println("DoctorID not found in session")
		return 0, ""
	}

	if !okUserType {
		log.Println("UserType not found in session")
		return doctorID, ""
	}

	return doctorID, userType
}

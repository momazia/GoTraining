/*
	Source: 		Column BL in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	PROJECT STEP 4 - refactoring our application, create a new data type called "user" which has fields for the user's name and age.
					When you receive the user's name and age form submission, create a variable of type "user"
					then put those values from the form submission into the fields for that variable.
					Marshal your variable of type "user"  to JSON. Encode that JSON to base64. Store that value in the cookie.
	Comment:		No change was done to step3 since I did exactly what was asked in step 4 in the last step by accident.
*/

package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Age  string
	Name string
}

func loginHandler(res http.ResponseWriter, req *http.Request) {

	setCookies(res, req)

	temp, err := template.ParseFiles("login.html")
	// Logging possible errors
	logFatalError(err)

	temp.Execute(res, nil)
	// Logging possible errors
	logFatalError(err)
}

// Sets the cookies on the response
func setCookies(res http.ResponseWriter, req *http.Request) {

	// Generating a new ID
	id, err := uuid.NewV4()
	logFatalError(err)
	createCookie(&res, "session-fino", id.String())

	// Setting user's data on cookie
	user := User{
		Age:  req.FormValue("age"),
		Name: req.FormValue("name"),
	}

	// Avoiding setting the cookie for the first visit.
	if req.Method == "POST" {
		userBytes, err := json.Marshal(user)
		logFatalError(err)
		createCookie(&res, "userData", base64.StdEncoding.EncodeToString(userBytes))
	}
}

// Creates a cookie for the given name and value and sets it on the response
func createCookie(res *http.ResponseWriter, cookieName, cookieValue string) {

	// Setting the cookie
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
		//		Secure: true,
		HttpOnly: true,
	}

	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

func main() {

	// Ignoring favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Setting the handler for login
	http.HandleFunc("/", loginHandler)

	// Setting the listener on port 8080
	log.Println("Listening to 8080 ...")
	http.ListenAndServe(":8080", nil)
}

// Logs error at Fatal level given the error is not nil.
func logFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

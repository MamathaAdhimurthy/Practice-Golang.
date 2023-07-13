package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	Username = "Sri"
	Password = "123"
)

type Login struct {
	Username string
	Password string
}

func BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)

	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	data := Login{}
	json.NewDecoder(r.Body).Decode(&data)
	r.Header.Add("Authorization", "Basic"+BasicAuth(data.Username, data.Password))

	// fooVAlues := r.Header.Values("foo")
	// fmt.Printf("r.Header.Values(\"foo\"):: %s\n\n", fooVAlues)

	// contentType := r.Header.Get("content-type")
	// fmt.Printf("r.Header.Get(\"content-type\"):: %s\n\n", contentType)

	// fooFirstValue := r.Header.Get("foo")
	// fmt.Printf("r.Header.Get(\"foo\"):: %s\n\n", fooFirstValue)

	// headers := r.Header
	// fmt.Printf("r.Headers:: %s\n\n", headers)
	// fmt.Printf("r.Headers[\"Content-Type\"]:: %s\n\n", headers["Content-Type"])
	// fmt.Printf("r.Headers[\"Foo\"]:: %s", headers["Foo"])

	// SET RESPONSE HEADERS

	// w.Header().Set("content-type", "application/json")

	// w.Header().Add("foo", "bar1")
	// w.Header().Add("foo", "bar2")

	// resp := make(map[string]string)
	// resp["message"] = "success"
	// jsonResp, _ := json.Marshal(resp)
	// w.Write(jsonResp)

	// Check if a particular header is present in a HTTP request in Go
	// headers := r.Header
	// val, ok := headers["Content-Type"]
	// if ok {
	// 	fmt.Printf("Content-Type key header is present with value %s\n", val)
	// } else {
	// 	fmt.Printf("Content-Type header is not present %s\n", val)
	// }

	//Get JSON request body from a HTTP request in Go

	// 	headerContentType := r.Header.Get("Content-Type")
	// 	if headerContentType != "application/json" {
	// 		errorResponse(w, "Content type is not application/json", http.StatusUnsupportedMediaType)
	// 		return
	// 	}
	// 	var e employee
	// 	var unmarshalErr json.UnmarshalTypeError
	// 	decoder := json.NewDecoder(r.Body)
	// 	decoder.DisallowUnknownFields()
	// 	err := decoder.Decode(&e)

	// 	if err != nil {
	// 		if errors.As(err, &unmarshalErr) {
	// 			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
	// 		} else {
	// 			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
	// 		}
	// 		return
	// 	}
	// 	errorResponse(w, "Success", http.StatusOK)
	// 	return
	// }

	// func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(httpStatusCode)
	// 	res := make(map[string]string)
	// 	res["message"] = message
	// 	jsonResp, _ := json.Marshal(res)
	// 	w.Write(jsonResp)

	// }

	// type employee struct {
	// 	Name string `json:"name"`
	// 	Age  int    `json:"age"`
	// }

	//HTTP send/receive jpeg file in request body example in Go (Golang)

	// err := r.ParseMultipartForm(32 << 20)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// file, h, err := r.FormFile("Photo")
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// tmpfile, err := os.Create("C:/Users/ee212295/postman images/" + h.Filename)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// _, err = io.Copy(tmpfile, file)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return

	// }
	// w.WriteHeader(200)
	// return

	//BASIC AUTHENTICATION

	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error pasing with basic auth")
		w.WriteHeader(401)
		return
	}
	if u != Username {
		fmt.Printf("Username is wrong %s\n", u)
		w.WriteHeader(401)
		return
	}
	if p != Password {
		fmt.Printf("Password is wrong %s\n", p)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("UserName %s\n", u)
	fmt.Printf("Password %s\n", p)
	w.Write([]byte(p))
	w.WriteHeader(200)
	return

}

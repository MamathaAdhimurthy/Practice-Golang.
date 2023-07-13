package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/hello/myfun/model"
	"example.com/hello/myfun/views"
)

func Crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//take some data
			//save it
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadsAllByName(name)
			fmt.Println(data)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("some error"))
				return
			}
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item deleted"})
		}

	}
}

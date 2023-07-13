package model

import "example.com/hello/myfun/views"

func ReadsAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{} //creating an array of todos
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil

}

func ReadsAllByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{} //creating an array of todos
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil

}

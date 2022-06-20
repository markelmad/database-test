package httphandler

import (
	"fmt"
	"log"
	"net/http"
	helper "src/src/helpers"
)

//http://localhost:9000/insertrow?title=Just&category=Comedy
func InsertRow(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("title") || !query.Has("category") {
		w.Write([]byte("Incomplete or incorrect query.\n"))
		fmt.Println("Incomplete or incorrect query.")
		return
	}

	db := helper.ConnectToDatabase()
	defer db.Close()

	title := query["title"][0]
	category := query["category"][0]

	row, err := db.Query("SELECT title FROM shakespeare_work WHERE title = ?", title)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if row.Next() {
		w.Write([]byte(fmt.Sprintf("Item with a Title of %s already exists! You cannot have duplicate Titles.\n", title)))
		fmt.Printf("Item with a Title of %s already exists! You cannot have duplicate Titles.\n", title)
		return
	}

	stmt, err := db.Prepare("INSERT INTO shakespeare_work (title, category, read_count) VALUES (?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err1 := stmt.Exec(title, category, 0)
	if err1 != nil {
		log.Fatal(err.Error())
	}

	w.Write([]byte(fmt.Sprintf("Successfully added the %s under %s category\n", title, category)))
	fmt.Printf("Successfully added the %s under %s category\n", title, category)
}

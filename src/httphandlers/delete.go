package httphandler

import (
	"fmt"
	"log"
	"net/http"
	helper "src/src/helpers"
	"strconv"
)

func DeleteRowByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") || query.Get("id") == "" {
		w.Write([]byte("Incorrect query. Please use a valie id and value.\n"))
		fmt.Println("Incorrect query. Please use a valie id and value.")
		return
	}
	db := helper.ConnectToDatabase()
	defer db.Close()

	id, err := strconv.Atoi(query["id"][0])
	if err != nil {
		fmt.Printf("Incorrect id: %s. Please use numbers only as value.\n", query["id"][0])
		w.Write([]byte(fmt.Sprintf("Incorrect id: %s. Please use numbers only as value.\n", query["id"][0])))
		return
	}

	row, err := db.Query("SELECT id FROM shakespeare_work WHERE id = ?", id)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if !row.Next() {
		w.Write([]byte(fmt.Sprintf("Item with an ID of %d doesn't exist. No item has been deleted.\n", id)))
		fmt.Printf("Item with an ID of %d doesn't exist. No item has been deleted.\n", id)
		return
	}

	stmt, err := db.Prepare("DELETE FROM shakespeare_work WHERE (`id` = ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write([]byte("Successfully deleted an item.\n"))
	fmt.Println("Successfully deleted an item.")
}

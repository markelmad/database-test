package httphandler

import (
	"fmt"
	"log"
	"net/http"
	helper "src/src/helpers"
	"strconv"
)

func UpdateReadCount(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") || !query.Has("readcount") || query.Get("id") == "" || query.Get("readcount") == "" {
		w.Write([]byte("Incomplete or incorrect query.\n"))
		fmt.Println("Incomplete or incorrect query.")
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
	readCount, err := strconv.Atoi(query["readcount"][0])
	if err != nil {
		fmt.Printf("Incorrect readcount: %s. Please use numbers only as value.\n", query["readcount"][0])
		w.Write([]byte(fmt.Sprintf("Incorrect readcount: %s. Please use numbers only as value.\n", query["readcount"][0])))
		return
	}

	row, err := db.Query("SELECT id FROM shakespeare_work WHERE id = ?", id)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	if !row.Next() {
		w.Write([]byte(fmt.Sprintf("Item with an ID of %d doesn't exist.\n", id)))
		fmt.Printf("Item with an ID of %d doesn't exist.\n", id)
		return
	}

	stmt, err := db.Prepare("UPDATE shakespeare_work SET `read_count` = ? WHERE (`id` = ?)")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec(readCount, id)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write([]byte(fmt.Sprintf("Successfully updated the read count value of %d an item with id %d.\n", readCount, id)))
	fmt.Printf("Successfully updated the read count value of %d an item with id %d.\n", readCount, id)
}

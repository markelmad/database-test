package httphandler

import (
	"fmt"
	"log"
	"net/http"
	helper "src/src/helpers"
)

func ShowAll(w http.ResponseWriter, r *http.Request) {
	db := helper.ConnectToDatabase()
	defer db.Close()

	row, err := db.Query("SELECT * FROM shakespeare_work")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()

	data := []helper.ShakespeareWorks{}
	for row.Next() {
		d := helper.ShakespeareWorks{}
		err := row.Scan(&d.ID, &d.Title, &d.Category, &d.ReadCount)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Write([]byte(fmt.Sprintf("ID: %d, TItle: %s, Category: %s, Read Count: %d\n", d.ID, d.Title, d.Category, d.ReadCount)))
		fmt.Printf(fmt.Sprintf("ID: %d, TItle: %s, Category: %s, Read Count: %d\n", d.ID, d.Title, d.Category, d.ReadCount))
		data = append(data, d)
	}
}

func ShowDataByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Path[len("/showtitlesbycategory/"):]
	if len(category) == 0 {
		w.Write([]byte("No category received.\n"))
		fmt.Println("No category received.")
		return
	} else if category != "comedy" && category != "history" && category != "poetry" && category != "tragedy" {
		w.Write([]byte(fmt.Sprintf("No category matches %s.\n", category)))
		fmt.Printf("No category matches %s.\n", category)
		return
	}
	db := helper.ConnectToDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM shakespeare_work WHERE category = ?", category)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	data := []helper.ShakespeareWorks{}
	for rows.Next() {
		d := helper.ShakespeareWorks{}
		err := rows.Scan(&d.ID, &d.Title, &d.Category, &d.ReadCount)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Write([]byte(fmt.Sprintf("ID: %d, TItle: %s, Category: %s, Read Count: %d\n", d.ID, d.Title, d.Category, d.ReadCount)))
		fmt.Printf(fmt.Sprintf("ID: %d, TItle: %s, Category: %s, Read Count: %d\n", d.ID, d.Title, d.Category, d.ReadCount))
		data = append(data, d)
	}

	fmt.Println("Successful data show")
}

package main

import (
	"log"
	"net/http"

	httphandler "src/src/httphandlers"
)

func main() {
	http.HandleFunc("/showdb", httphandler.ShowAll)
	http.HandleFunc("/showtitlesbycategory/", httphandler.ShowDataByCategory)
	http.HandleFunc("/insertrow", httphandler.InsertRow)
	http.HandleFunc("/delaterowbyid", httphandler.DeleteRowByID)
	http.HandleFunc("/updatereadcount", httphandler.UpdateReadCount)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/Ibam72/mockdsps"
)

func main() {
	router := httprouter.New()
	router.POST("/mockDSPs/:dsp_id", mockdsps.POSTBidding)
	log.Fatal(http.ListenAndServe(":8080", router))
}

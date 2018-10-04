package module

import (
	// "flag"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func JSONParserHandler(w http.ResponseWriter, req *http.Request) {
	type Header struct {
		Code int    `json:"code"`
		Desc string `json:"desc"`
	}
	type Body struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		Done  string `json:"done"`
	}
	type JSONRequest struct {
		Header Header `json:"header"`
		Body   []Body `json:"body"`
	}

	type JSONResponse struct {
		Code int    `json:"code"`
		Desc string `json:"desc"`
	}

	jsonStream := req.Body
	var request JSONRequest
	if err := json.NewDecoder(jsonStream).Decode(&request); err != nil {
		http.Error(w, "{\"code\":500, \"desc\":\""+err.Error()+"\"}", http.StatusBadRequest)
	}
	var body []Body
	var header Header

	body = request.Body
	header = request.Header
	log.Println("Code:" + strconv.Itoa(header.Code))
	log.Println("Desc:" + header.Desc)
	log.Println("—————————")

	for i := 0; i < len(body); i++ {

		var data Body
		data = body[i]

		log.Println("")
		log.Println("ID: " + data.Id)
		log.Println("Title: " + data.Title)
		log.Println("Done: " + data.Done)
	}
	log.Println("—————————")

	var jsonResponse JSONResponse
	jsonResponse.Code = 200
	jsonResponse.Desc = "Success"

	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		http.Error(w, "{\"code\":502, \"desc\":\""+err.Error()+"\"}", http.StatusBadRequest)
		return
	}
}

package endpoints

import "net/http"

var index string

func Init(indexName string) {

	index = indexName

	http.HandleFunc("/add", AddDocument)
	http.HandleFunc("/get", GetDocument)
	http.HandleFunc("/delete", DeleteDocument)
	http.HandleFunc("/update", UpdateDocument)

}

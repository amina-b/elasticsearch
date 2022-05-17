package endpoints

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amina-b/elasticsearch/service"
	"github.com/elastic/go-elasticsearch/esapi"
)

// Get document from the elastic cloud by document id
func GetDocument(w http.ResponseWriter, r *http.Request) {

	req := esapi.GetRequest{
		Index:      index,
		DocumentID: r.URL.Query().Get("document_id"),
	}

	res, err := req.Do(context.Background(), service.Client)

	if err != nil {
		log.Fatalf("Error getting response while getting document: %s", err)
	}

	defer res.Body.Close()

	byt, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error getting response while getting document: %s", err)
	}

	w.WriteHeader(res.StatusCode)
	w.Write(byt)
	w.Header().Set("Content-Type", "application/json")

}

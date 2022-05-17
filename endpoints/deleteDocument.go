package endpoints

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amina-b/elasticsearch/service"
	"github.com/elastic/go-elasticsearch/esapi"
)

type deleteDocRequest struct {
	DocumentID string `json:"DocumentID"`
}

// Delete document from the elastic cloud by document id
func DeleteDocument(w http.ResponseWriter, r *http.Request) {

	var deleteDoc deleteDocRequest
	unmarshalErr := json.NewDecoder(r.Body).Decode(&deleteDoc)

	if unmarshalErr != nil {
		log.Printf("failed to unmarshal the struct to JSON. Error: %v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(unmarshalErr.Error()))
		return
	}

	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: deleteDoc.DocumentID,
	}

	res, err := req.Do(context.Background(), service.Client)

	if err != nil {
		log.Fatalf("Error getting response while deleting document: %s", err)
	}

	defer res.Body.Close()

	byt, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error getting response while deleting document: %s", err)
	}

	w.WriteHeader(res.StatusCode)
	w.Write(byt)
	w.Header().Set("Content-Type", "application/json")

}

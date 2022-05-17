package endpoints

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/amina-b/elasticsearch/service"
	"github.com/elastic/go-elasticsearch/esapi"
)

type updateDoc struct {
	DocumentID string `json:"DocumentID"`
	Name       string `json:"Name"`
	Email      string `json:"Email"`
	Phone      string `json:"Phone"`
	Address    string `json:"Address"`
}

// Update document by id
func UpdateDocument(w http.ResponseWriter, r *http.Request) {

	var updateDoc updateDoc
	unmarshalErr := json.NewDecoder(r.Body).Decode(&updateDoc)

	if unmarshalErr != nil {
		log.Printf("failed to unmarshal the struct to JSON. Error: %v", unmarshalErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(unmarshalErr.Error()))
		return
	}

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: updateDoc.DocumentID,
		Body:       strings.NewReader(updateDoc.Name + "," + updateDoc.Email + "," + updateDoc.Phone + "," + updateDoc.Address),
	}

	res, err := req.Do(context.Background(), service.Client)

	if err != nil {
		log.Fatalf("Error getting response while updating document: %s", err)
	}

	defer res.Body.Close()

	byt, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error getting response while updating document: %s", err)
	}

	w.WriteHeader(res.StatusCode)
	w.Write(byt)
	w.Header().Set("Content-Type", "application/json")

}

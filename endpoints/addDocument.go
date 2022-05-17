package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/amina-b/elasticsearch/service"
	"github.com/elastic/go-elasticsearch/esapi"
)

// Add document to the elastic
type AddDocumentRequest struct {
	Body  string
	Index string
}

func AddDocument(w http.ResponseWriter, r *http.Request) {

	byt, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("failed to read request body. Error: %v", err)
		return
	}

	addDoc := &AddDocumentRequest{
		Body:  string(byt),
		Index: index,
	}

	log.Println(addDoc)

	res, err := addDoc.Do(context.Background(), service.Client)

	if err != nil {
		log.Fatalf("Error getting response while adding document: %s", err)
	}

	defer res.Body.Close()

	w.WriteHeader(res.StatusCode)
	w.Write([]byte(res.Status))

}

func (r *AddDocumentRequest) Do(ctx context.Context, transport esapi.Transport) (*http.Response, error) {

	byt, err := json.Marshal(r)

	if err != nil {
		log.Printf("failed to marshal the struct to JSON. Error: %v", err)
		return nil, err
	}

	path := fmt.Sprintf("/%s/%s", r.Index, "_doc")

	req := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Path: path,
		},
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewReader(byt)),
	}

	res, err := transport.Perform(req)
	if err != nil {
		log.Printf("Failed to perform request: %v", err)
		return nil, err
	}

	response := http.Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Body:       res.Body,
	}

	return &response, nil
}

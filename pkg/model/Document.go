package models

import (
	"encoding/json"
	"github.com/franela/goreq"
	"io/ioutil"
)

type Document struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}

func (f *Document) GetDocuments(c chan []Document, e chan error) {

	defer close(c)
	defer close(e)

	r, err := goreq.Request{Uri: "http://54.38.189.215:4002/documents"}.Do()
	//r , err := goreq.Request{Uri:"http://localhost:4000/documents"}.Do()

	if err != nil {
		e <- err
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e <- err
		return
	}

	var docs []Document
	err = json.Unmarshal(b, &docs)
	if err != nil {
		e <- err
		return
	}

	c <- docs

}

package models

import (
	"encoding/json"
	"github.com/franela/goreq"
	"io/ioutil"
	"time"
)

type Formation struct {
	Id        int `json:"id"`
	CreatedAt time.Time
	Intitule  string    `json:"intitule"`
	DateDebut time.Time `json:"date_debut"`
	DateFin   time.Time `json:"date_fin"`
	NbrHeures int       `json:"nbr_heures"`
	Cout      float64   `json:"cout"`
}

func (f *Formation) GetFormations(c chan []Formation, e chan error) {

	defer close(c)
	defer close(e)

	r, err := goreq.Request{Uri: "http://54.38.189.215:4002/formations"}.Do()
	//r , err := goreq.Request{Uri:"http://localhost:4000/formations"}.Do()

	if err != nil {
		e <- err
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e <- err
		return
	}

	var frm []Formation
	err = json.Unmarshal(b, &frm)
	if err != nil {
		e <- err
		return
	}

	c <- frm

}

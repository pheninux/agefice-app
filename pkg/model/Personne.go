package models

import (
	"encoding/json"
	"github.com/franela/goreq"
	"io/ioutil"
	"time"
)

type Personne struct {
	Id            int        `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	Mfa           bool       `json:"mfa"`
	Nom           string     `json:"nom"`
	Prenom        string     `json:"prenom"`
	Age           int        `json:"age"`
	DateNaissance time.Time  `json:"date_naissance"`
	Tel           string     `json:"tel"`
	Mail          string     `json:"mail"`
	Adresse       string     `json:adresse"`
	Entreprise    Entreprise `json:"entreprise"`
	EntrepriseId  int
	Document      []Document  `json:"document"`
	Formation     []Formation `json:"formation"`
}

func (p *Personne) GetAll(c chan []Personne, e chan error, mfa bool) {

	defer close(e)
	defer close(c)

	//r , err := goreq.Request{Uri:"http://localhost:4000/all",Timeout:time.Second * 5}.Do()
	r, err := goreq.Request{Uri: "http://54.38.189.215:4002/all", Timeout: time.Second * 5}.Do()
	if err != nil {
		e <- err
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e <- err
		return

	}

	var prs []Personne
	err = json.Unmarshal(body, &prs)
	if err != nil {
		e <- err
		return
	}

	c <- prs

}

func (p *Personne) GetByMfa(c chan []Personne, e chan error, mfa bool) {

	defer close(e)
	defer close(c)

	r, err := goreq.Request{
		Method: "POST",
		//Uri: "http://localhost:4000/getByMfa",
		Uri:  "http://54.38.189.215:4002/getByMfa",
		Body: mfa,
	}.Do()

	if err != nil {
		e <- err
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e <- err
		return

	}

	var prs []Personne
	err = json.Unmarshal(body, &prs)
	if err != nil {
		e <- err
		return
	}

	c <- prs

}

func (p *Personne) SavePersonne(e chan string) {

	defer close(e)

	rep, err := goreq.Request{
		Method: "POST",
		//Uri: "http://localhost:4000/personne/json/create",
		Uri:  "http://54.38.189.215:4002/personne/json/create",
		Body: p,
	}.Do()
	if err != nil {
		e <- err.Error()
		return
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		e <- err.Error()
		return
	}
	e <- string(b)

}

func (p *Personne) UpdatePersonne(e chan string) {

	defer close(e)

	rep, err := goreq.Request{
		Method: "POST",
		//Uri: "http://localhost:4000/personne/json/update",
		Uri:  "http://54.38.189.215:4002/personne/json/update",
		Body: p,
	}.Do()
	if err != nil {
		e <- err.Error()
		return
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		e <- err.Error()
		return
	}
	e <- string(b)

}

func (p *Personne) DeletePersonneById(e chan string) {

	defer close(e)

	rep, err := goreq.Request{
		Method: "POST",
		//Uri:               "http://localhost:4000/personne/delete/",
		Uri:  "http://54.38.189.215:4002/personne/delete/",
		Body: p.Id,
	}.Do()
	if err != nil {
		e <- err.Error()
		return
	}

	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		e <- err.Error()
		return
	}

	e <- string(b)

}

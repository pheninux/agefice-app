package main

import (
	models "agefice-cons/adil.net/pkg/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/franela/goreq"
	"io/ioutil"
	"net/smtp"
	"time"
)

type mail struct {
	id      int
	from    string
	to      string
	object  string
	content string
}

func main() {
	//m := composeMail(p)
	// authentification
	auth := smtp.PlainAuth("", from2, passworOutlook, server2)
	if err := smtp.SendMail(server2+port, auth, from2, []string{"adil.haddad.xdev@gmail.com"}, []byte("test")); err != nil {
		fmt.Println("Error send mail", err)
	}
	fmt.Println("Email sent")
}

/*func main() {

	c := make(chan []models.Personne)
	e := make(chan error, 1)
	pros := make(chan bool)

		go getPersonnes(c, e)

	for {
		select {
		case err, ok := <-e:
			if ok {
				fmt.Println(err)
				return
			}
		case p , ok:= <-c:
			if ok {
				go managePersonnes(p ,pros)
			}
		case b , ok:= <- pros:
			 if ok{
			 	if b {
			 		go getPersonnes(c,e)
				}
			 }

		default:

		}
	}

}*/

func getPersonnes(c chan []models.Personne, e chan error) {

	var p []models.Personne
	r, err := goreq.Request{
		Method: "GET",
		Uri:    "http://localhost:4000/all",
	}.Do()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		e <- err
		return
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		e <- err
		return
	}

	c <- p

}

func sendMail(p models.Personne) {

	m := composeMail(p)
	time.Sleep(time.Second * 3)
	fmt.Println("Email sent =>", m)

}

func composeMail(p models.Personne) *mail {
	var buf bytes.Buffer
	for _, v := range p.Document {
		buf.WriteString(v.Libelle)
	}
	return &mail{
		id:      0,
		from:    "jessica@htus.fr",
		to:      p.Mail,
		object:  p.Prenom + " " + p.Nom + " -> demande de piece",
		content: buf.String(),
	}
}

func managePersonnes(p []models.Personne, pros chan bool) {

	for _, v := range p {
		dNow, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		if err != nil {
			fmt.Println("Error :", err)
			return
		}
		if dNow == v.Formation[0].DateFin {

			go sendMail(v)
		}
	}

	time.Sleep(time.Second * 6)
	pros <- true

}

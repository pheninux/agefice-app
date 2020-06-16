package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

const (
	server1        = "smtp.gmail.com"
	server2        = "smtp.office365.com"
	port           = ":587"
	email          = ""
	password       = ""
	to             = "adil.haddad.xdev@gmail.com"
	from2          = "jchassagnac@utas.fr"
	passworOutlook = "Jcre*951"
)

func main() {

	m := gomail.NewMessage()
	m.SetHeader("From", from2)
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "pophaddad@gmail.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", `Bonjour Madame.....,
	Bonjour<br> ,

	Votre formation est à présent terminée.<br>
		Pour la demande de remboursement vous voudrez bien me faire parvenir les éléments suivants :<br>
	- la facture acquittée<br>
	- les feuilles d’émargement<br>
	- l’attestation d’assiduité<br>
	- doc manquant 1<br>
	- doc manquant 2<br>
	...

	L’ensemble des documents doit être signé et tamponné par l’organisme de formation.<br>

		Je reste à votre disposition pour tout complément d’information.<br>

		Cordialement,<br>
		Jessica CHASSAGNAC<br>
	    MEDEF Douaisis<br>
	    03.27.08.10.70`)

	mailer := gomail.NewDialer("smtp.office365.com", 587, from2, passworOutlook)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := mailer.DialAndSend(m); err != nil {
		panic(err)
	}
}

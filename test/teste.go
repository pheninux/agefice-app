package main

import (
	models "agefice-cons/adil.net/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/franela/goreq"
	"io/ioutil"
	"time"
)

func main() {

	c := make(chan []models.Personne)
	e := make(chan error, 1)

	go func(c chan []models.Personne, e chan error) {

		defer close(c)
		defer close(e)

		r, err := goreq.Request{Uri: "http://localhost:4000/all", Timeout: time.Second * 5}.Do()

		if err != nil {
			e <- err
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			e <- err
			return
		}

		var p []models.Personne
		err = json.Unmarshal(b, &p)
		if err != nil {
			e <- err
			return
		}

		c <- p

	}(c, e)

	for {
		select {
		case err, ok := <-e:
			if ok {
				fmt.Println(err)
				return
			}
		case p := <-c:
			fmt.Println(p)
			return
		default:

		}
	}

}

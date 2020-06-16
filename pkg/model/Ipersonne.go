package models

type IPersonne interface {
	GetAll(c chan []Personne)
}

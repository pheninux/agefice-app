package main

type Ioiseux interface {
	vole()
}

type Ifelin interface {
	cour()
}

type Ianimal interface {
	Ioiseux
	nage()
	cour()
	saute()
}

type animal struct {
	yeux    int
	pied    int
	couleur string
	vitesse int
}

type oiseaux struct {
	animal
	aille string
}

type chien struct {
	animal
}

func (c chien) vole() {
	panic("implement me")
}

func (c chien) nage() {
	panic("implement me")
}

func (c chien) cour() {
	panic("implement me")
}

func main() {

}

package main

import (
	"fmt"
)

type todo struct {
	Titel        string
	Beschreibung string
}

func add(tL []todo) ([]todo, error) {

	nT := todo{Titel: "", Beschreibung: ""}

	fmt.Println("Titel: ")
	_, err := fmt.Scan(&nT.Titel)
	if err != nil {
		return nil, err
	}

	fmt.Println("Beschreibung")
	_, err = fmt.Scan(&nT.Beschreibung)
	if err != nil {
		return nil, err
	}

	tL = append(tL, nT)

	return tL, nil
}

func showAll(tL []todo) {
	for _, v := range tL {
		fmt.Println(v)
	}
}

func main() {
	todoList := []todo{}
	var eingabe string
	var err error

	fmt.Println("Todo hinzufügen, Todo löschen, Todos anzeigen?")
	_, err = fmt.Scan(&eingabe)
	if err != nil {
		fmt.Println(err)
	}

	switch eingabe {
	case "hinzufügen":
		add(todoList)
	case "anzeigen":
		showAll(todoList)

	}
}

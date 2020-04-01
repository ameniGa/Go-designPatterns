package main

import (
	"log"
	"reflect"
)

/*****************window widget********************/

type buttonHandler interface {
	draw(dim int)
}

type squareButton struct{}

type circleButton struct{}

func (s squareButton) draw(dim int) {
	log.Println("i m a square window")
}

func (c circleButton) draw(dim int) {
	log.Println("i m a circle window")
}

// this factory method
func createButton(buttonType string) buttonHandler {
	switch buttonType {
	case "circle":
		return circleButton{}
	default:
		return squareButton{}
	}
}
/*****************alert widget********************/
type alertHandler interface {
	popUp()
}

type circleAlert struct {
}

type squareAlert struct {
}

func (d circleAlert) popUp(){
	log.Println("i m a circle alert")
}

func (t squareAlert) popUp(){
	log.Println("i m a square alert")
}

// this factory method
func createAlert(alertType string) alertHandler {
	switch alertType {
	case "circle":
		return circleAlert{}
	default:
		return squareAlert{}
	}
}

/*****************abstract factory********************/

// this is abstract factory
type themeFactory interface {
	createAlert() alertHandler
	createButton() buttonHandler
}

// concrete abstract factory
type softTheme struct {}

func (s softTheme) createAlert() alertHandler{
	return createAlert("circle")
}

func (s softTheme) createButton() buttonHandler{
	return createButton("circle")
}

// concrete abstract factory
type hardTheme struct {}

func (h hardTheme) createAlert() alertHandler{
	return squareAlert{}
}

func (h hardTheme) createButton() buttonHandler{
	return squareButton{}
}


func create(theme string) themeFactory {
	switch theme {
	case "hard":
		return hardTheme{}
	case "soft":
		return softTheme{}
	default:
		return nil
	}
}

/*********************************************/

func main() {
	button := createButton("circle")
	buttonType := reflect.TypeOf(button).Name()
	log.Printf("window type: %v", buttonType)
	softTheme := create("soft")
	softTheme.createButton().draw(5)
	softTheme.createAlert().popUp()
}

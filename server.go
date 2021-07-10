package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

// type RadioButton struct{
// 	Name string
// 	Value string
// 	IsDisabled bool
// 	IsChecked bool
// 	Text string
// }

// type PageVariables struct{
// 	PageTitle string
// 	PageRadioButtons []RadioButton
// 	Answer string
// }
// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>hello World<h1>")
// }

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}
	t, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Print("template parsing error: ", err) //log it
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

// func DisplayRadioButtons(w http.ResponseWriter, r *http.Request){
// 	Title := "What do you Prefer?"
// 	MyRadioButtons := []RadioButton{
// 		RadioButton{"animalselect","cats",false,false, "Cats"},
// 		RadioButton("animalSelect", "dogs", false, false, "Dogs"),
// 	}
// 	MyPageVariables := PageVariables{
// 		PageTitle: Title,
// 		PageRadioButtons : MyRadioButtons,
// 	}
// 	t.err := template.ParseFiles("home.html")
// 	if err!=nil{
// 		log.Print("template parsing error: ",err)
// 	}
// 	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
// 	if err != nil { // if there is an error
// 	  log.Print("template executing error: ", err) //log it
// 	}
// }

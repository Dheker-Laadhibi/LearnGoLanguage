package main

import (

	 "fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//starting func indexhandler
  func indexHandler( w  http.ResponseWriter ,r *http.Request){
t , err  := template.ParseFiles("./templates/index.html")
 if err!= nil {
    log.Fatal("error encountred while parsing  template " , err)
 }
 t.Execute(w, nil)
  }
  // end index handler 

  //starting func homehandler
  func homeHandler( w  http.ResponseWriter ,r *http.Request){
    fmt.Fprintf(w, "welcome home ")
  }
// end  home handler 


func main() {
    fmt.Println("started :")
    // new router 
    router := mux.NewRouter()
    // routes 
    router.HandleFunc("/" , indexHandler)
    router.HandleFunc("/home" , homeHandler)

	//,serve static files 
    router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/" , http.FileServer(http.Dir("assets"))))  
    http.ListenAndServe(":8080" , router)
}

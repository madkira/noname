package main

import(
 "net/http"
 //"container/list"

 "github.com/gorilla/mux"


)


func main() {
  initGmail();
  initPersons();
	r := mux.NewRouter()
  r.HandleFunc("/",root)
  r.HandleFunc("/mail",mail)
	r.HandleFunc("/mail/{id}", mailid )
  r.HandleFunc("/whoishere",getPresent)
  r.HandleFunc("/whoishere/add",addPresent)
  r.HandleFunc("/whoishere/rm",rmPresent)
  http.Handle("/",r)


  http.ListenAndServe(":80", nil) // set listen port

}

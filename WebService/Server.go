package main

import(
 "net/http"
 //"container/list"

 "github.com/gorilla/mux"
 upnp "upnpProtocol"



)


func main() {

  upnp.ConnectUDP();
  initGmail();
  initPersons();
  initMQTT("cleedor");
	r := mux.NewRouter()
  r.HandleFunc("/",root)
  r.HandleFunc("/mail",mail)
	r.HandleFunc("/mail/{id}", mailid )
  r.HandleFunc("/whoishere",getPresent)
  r.HandleFunc("/whoishere/add",addPresent)
  r.HandleFunc("/whoishere/add/{id}",addPresent)
  r.HandleFunc("/whoishere/rm",rmPresent)
  r.HandleFunc("/whoishere/rm/{id}",rmPresent)
  r.HandleFunc("/monitor/{id}",getMQTTnode)
  http.Handle("/",r)


  http.ListenAndServe(":80", nil) // set listen port

}

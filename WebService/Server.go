package main

import(
 "net/http"
 //"container/list"
 "encoding/json"
 "fmt"
  //"flag"

 //upnp "upnpProtocol"
 "log"
 //"io/ioutil"
 owm "github.com/briandowns/openweathermap"




)

type User_Services struct {
    Name string `json:"Name,omitempty"`
    Services []string `json:"Services,omitempty"`
}

type Mail_Cred struct {
  User string
  Mail string
  Pwd string
}

type userN struct{
  present bool
  name string
  client *http.Client
  services []string
}

type weather struct{
  name string `json:"name"`
}

var users []userN

func currW(){
  /*w, err := owm.NewCurrent("C", "fr","8ba7a41e3811571c54b7e97138f05d46")
	if err != nil {
		log.Fatalln(err)
	}*/
  /*resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Paris&appid=" + "8ba7a41e3811571c54b7e97138f05d46" )
  if err != nil {
		log.Fatalln(err)
	}
  body, err := ioutil.ReadAll(resp.Body)
  decoder := json.NewDecoder(resp.Body)
  var t weather
    err = decoder.Decode(&t)

	fmt.Println(body)
  fmt.Println(t)*/

  w, err := owm.NewCurrent("C", "FR", "8ba7a41e3811571c54b7e97138f05d46") // (internal - OpenWeatherMap reference for kelvin) with English output
    if err != nil {
        log.Fatalln(err)
    }

    w.CurrentByName("Antibes")
    fmt.Println(w)
}

func save(w http.ResponseWriter, r *http.Request) {

  var u User_Services
		if r.Body == nil {
      log.Println("null body")
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
      log.Println(err.Error())
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(u)

    for i := range users {
      if users[i].name == u.Name{
        users[i].services = u.Services
        fmt.Println(users[i].services)

      }
    }
}

func bindMail(w http.ResponseWriter, r *http.Request) {
  var u Mail_Cred
  fmt.Println(r.Body)
		if r.Body == nil {
      log.Println("null body")
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
      log.Println(err.Error())
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(u)

}


func main() {

  users = []userN{userN{
    present: false,
    name: "Jean-Yves",
  },userN{
    present: false,
    name: "Stéphane",
  },
}

  //upnp.ConnectUDP();
  //initGmail();
  //initPersons();
  initMQTT("127.0.0.1:1883");
  //topic := flag.String("topic", "goldenkey/EntryUser", "The password (optional)")
  if token := client.Subscribe("goldenkey/EntryUser", byte(*qos), nil); token.Wait() && token.Error() != nil {
    //log.Fatalf(string(token.Error()))
    log.Println("explosion suscribe")
    return
    }
    //topic2 := flag.String("topic", "goldenkey/ExitUser", "The password (optional)")

    if token := client.Subscribe("goldenkey/ExitUser", byte(*qos), nil); token.Wait() && token.Error() != nil {
      //log.Fatalf(string(token.Error()))
      log.Println("explosion suscribe")
      return
      }
     //test()
  //currW()


  mux := http.NewServeMux()
  mux.Handle("/", http.FileServer(http.Dir("./www"))) // set router
  mux.Handle("/whoishere/",http.HandlerFunc(GetPresent))
  mux.Handle("/saveServices/",http.HandlerFunc(save))
  mux.Handle("/bindMail/",http.HandlerFunc(bindMail))
  err := http.ListenAndServe(":80", mux) // set listen port
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }



}

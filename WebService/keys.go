package main

import(
 "fmt"
 "net/http"
 "encoding/json"
 "strconv"
 "strings"
 "github.com/gorilla/mux"
 )

 type person struct{
   present bool
   name string
 }

 var presents [4]person

 func initPersons()  {
   presents[0] = person{name: "Boba Fett", present: false}
   presents[1] = person{name: "Gerard Majax", present: false}
   presents[2] = person{name: "Arthur le sanglier de cornouailles", present: false}
   presents[3] = person{name: "Jean-Yves Tigli", present: false}
 }

func getPresent(w http.ResponseWriter, r *http.Request) {
  var res []string
  for i := 0; i < 4; i++ {
    if(presents[i].present){
      res = append(res,presents[i].name)
      fmt.Println("%s\n",res)
    }
  }
  fmt.Println("%s\n",res)

  result,_ := json.Marshal(res)
  fmt.Println("%s\n",result)
  fmt.Fprintf(w,"%s",result)
}

func rmPresent(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()

  var i int
  if(r.Form["id"] != nil){

    i, _ = strconv.Atoi(strings.Join(r.Form["id"],","))
  }else{
    vars := mux.Vars(r)
    i,_ = strconv.Atoi(vars["id"])
  }
  presents[i].present = false
  fmt.Fprintf(w,"200")

}


func addPresent(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  var i int
  if(r.Form["id"] != nil){

    i, _ = strconv.Atoi(strings.Join(r.Form["id"],","))
  }else{
    vars := mux.Vars(r)
    i,_ = strconv.Atoi(vars["id"])
  }

  presents[i].present = true

  fmt.Fprintf(w,"200")
}

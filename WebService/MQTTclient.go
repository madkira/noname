package main
import(
 "fmt"
 "net/http"
 "flag"
 "strconv"
 //"os"

 MQTT "github.com/eclipse/paho.mqtt.golang"

)

var cleedor = ""

var choke = make(chan [2]string)

//var client mqtt.Client


var qos = flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")

var client MQTT.Client

func getMQTTnode(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"%s",cleedor)
}



func initMQTT( bro string){

  //topic := flag.String("topic", top, "The password (optional)")
	broker := flag.String("broker", "tcp://"+bro, "The password (optional)")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "testgoid", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

  //fmt.Println(*topic)



	/*if *topic == "" {
		fmt.Println("Invalid setting for -topic, must not be empty")
		return
	}*/



	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(*id)
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)
	if *store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(*store))
	}





	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client = MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	/*if token := client.Subscribe(*topic, byte(*qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}*/





  go listen()




	//client.Disconnect(250)
	//fmt.Println("Sample Subscriber Disconnected")


}

func setEnter(name string){
  for i := range users {
    if users[i].name == name{
      users[i].present = true
      fmt.Println(users[i].present)
      for a := range users[i].services{
        fmt.Println("publish " + users[i].services[a])

        if users[i].services[a] == "Mail" {
          fmt.Println(strconv.Itoa(test()))

          if token := client.Publish("goldenkey/msg", 0, false, "Vous avez " + strconv.Itoa(test()) + " nouveaux messages"); token.Wait() && token.Error() != nil {
            fmt.Println("err")
        }/*else if users[i].services[a] == "Meteo" {
          fmt.Println("publish")
          if token := client.Publish("goldenkey/msg", 0, false, currW()); token.Wait() && token.Error() != nil {
            fmt.Println("err")
        }*/
        }
      }
    }

    }
  }


func setExit(name string){
  for i := range users {
    if users[i].name == name{
      users[i].present = false
      fmt.Println(users[i].present)
      for a := range users[i].services{
        if users[i].services[a] == "Meteo" {
          fmt.Println("publish")
          if token := client.Publish("goldenkey/msg", 0, false, "Blizzard de neige, attention a la route"); token.Wait() && token.Error() != nil {
            fmt.Println("err")
          }
        }


    }
  }
}
}

func listen(){


  for {
    incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
    cleedor = incoming[1]
    if incoming[0] == "goldenkey/exit"{
      setExit(incoming[1])

    }else if incoming[0] == "goldenkey/entry"{
      setEnter(incoming[1])
    }
  }
}

package main
import(
 "fmt"
 "net/http"
 "flag"
 "os"

 MQTT "github.com/eclipse/paho.mqtt.golang"

)

var cleedor = ""

var choke = make(chan [2]string)




func getMQTTnode(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"%s",cleedor)
}

func initMQTT(top string){

  topic := flag.String("topic", top, "The password (optional)")
	broker := flag.String("broker", "tcp://192.168.43.123:1883", "The password (optional)")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "testgoid", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

  fmt.Println(*topic)



	if *topic == "" {
		fmt.Println("Invalid setting for -topic, must not be empty")
		return
	}



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

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(*topic, byte(*qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}





  go listen()




	//client.Disconnect(250)
	//fmt.Println("Sample Subscriber Disconnected")


}

func listen(){


  for {
    incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
    cleedor = incoming[1]
  }
}

package main

 import(
	"fmt"
	"net/http"
	//"container/list"


  "math/rand"
	"golang.org/x/oauth2"
 "golang.org/x/oauth2/google"
 "google.golang.org/api/gmail/v1"
 "log"
 //"strings"
 //"encoding/base64"
 "io/ioutil"
)

var gmailService gmail.Service

func root(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //id := vars["id"]


  fmt.Fprintf(w, "Hello world !")
}

func mail(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //id := vars["id"]


  fmt.Fprintf(w, "Send an id")
}

func mailid(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //id := vars["id"]
	rnd := rand.Int() % 15
	//list := list.New()

  label,err := gmailService.Users.Messages.List("me").Do()
	fmt.Println("Nb unread mail", label)
  fmt.Println("error ", err)

  fmt.Fprintf(w, "%d",rnd)
}

func initGmail(){

	secret, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Printf("Error: %v", err)
	}

	conf, err := google.ConfigFromJSON(secret, gmail.GmailSendScope)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Creates a URL for the user to follow
	url := conf.AuthCodeURL("CSRF", oauth2.AccessTypeOffline)
	// Prints the URL to the terminal
	fmt.Printf("Visit this URL: \n %v \n", url)

	// Grabs the authorization code you paste into the terminal
	var code string
	_, err = fmt.Scan(&code)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Exchange the auth code for an access token
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// Create the *http.Client using the access token
	client := conf.Client(oauth2.NoContext, tok)

	// Create a new gmail service using the client
	g, err := gmail.New(client)
	 gmailService = *g
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

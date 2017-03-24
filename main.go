package main 

import (
	"fmt"
	"log"
	"net/http"

	"message/fb"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

const APP_ID string = ""
const APP_SEC string = ""
const PAGE_ID string = ""

const PORT string = "8080"
const NUM_OF_THREADS int = 10
const OUTPUT_DIR string = ""

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/token", TokenHandler).Methods("POST")
	http.Handle("/", r)
	log.Println("Opening http://localhost:"+PORT+"/ to start")
	go browser.OpenURL("http://localhost:"+PORT+"/")
	http.ListenAndServe(":"+PORT, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage, APP_ID)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.PostForm.Get("token")
	go DownloadMessage(token)
}

func DownloadMessage(token string) {
	f := fb.New(token)
	p, err := f.NewPage(PAGE_ID)
	
	if err != nil {
		fmt.Println("Download message failed:", err.Error())
		return
	}

	conversations, err := p.GetAllConversations()
	// conversations, err := p.Get100Conversations()
	if err != nil {
		log.Println("Getting conversations failed:", err.Error())
		return
	}

	concurrency := NUM_OF_THREADS
	done := make(chan bool, concurrency)
	for k, conversation := range conversations {
		done <- true
		go func(k int, conversation fb.Conversation) {
			defer func() {<-done}()

			p.FetchMessagesTo(&conversation)
			output := OUTPUT_DIR
			conversation.WriteTo(output)
			fmt.Printf("Conversation %d/%d: \tWrote to file %s/%s\n", k+1, len(conversations), output, conversation.ID)
		}(k, conversation)
	}
	for i := 0; i < cap(done); i++ {
	    done <- true
	}

	fmt.Println("All tasks done, waiting for new task...")
}
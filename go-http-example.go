package main


import (
    "fmt"
    "log"
    "net/http"
    "time"
)


const port string = "8080"


func main () {
    startServer()
}


func startServer () {
    listeningPort := ":" + port

    fmt.Println("Starting server")
    fmt.Print("Listening to requests at: http://localhost", listeningPort, "\n")

    http.HandleFunc("/", requestHandler)
    log.Fatal(http.ListenAndServe(listeningPort, nil))
}


func requestHandler (w http.ResponseWriter, r *http.Request) {
    logRequest(r)
    showResponse(w, r)
}


func showResponse (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You've made a request to %v!\n\n", r.URL)
    callController(w, r)
}


func logRequest (r *http.Request) {
    fmt.Println("----------------------------")
    fmt.Println(time.Now().Format(time.ANSIC))
    fmt.Println("REQUEST:", r.Method, r.URL)
    fmt.Println("HOST:", r.Host)
    fmt.Println("BODY:", r.Body)
}


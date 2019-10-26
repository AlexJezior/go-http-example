package main

import (
    "fmt"
    "net/http"
)

var registeredRoutes = map[string]func (http.ResponseWriter, *http.Request){
    "/": homeController,
    "/home": homeController,
    "/contact": contactController,
}


func callController (w http.ResponseWriter, r *http.Request) {
    if controllerCall, key := registeredRoutes[r.URL.Path]; key {
        controllerCall(w, r)
    }else{
        notFoundController(w, r)
    }
}


func homeController (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You've made a request to the homeController!")
}


func contactController (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You've made a request to the contactController!")
}


func notFoundController (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The page URL requested %v is invalid.\nPlease try one of the following instead:\n", r.URL.Path)

    for key := range registeredRoutes {
        fmt.Fprintf(w, "%v\n", key)
    }
}


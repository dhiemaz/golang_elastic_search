package main

import(
  "net/http"  
)


type Route struct{
    Name          string
    Method        string
    Pattern       string
    HandlerFunc   http.HandlerFunc
}


type Routes []Route

var routes = Routes{

    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Search",
        "GET",
        "/rest/api/v1.0/search/{name}",
        Search,
    },
}

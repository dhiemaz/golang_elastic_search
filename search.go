package main

import (
    "fmt"
	  "log"
	  "net/http"
	  "net/url"
	  "io/ioutil"
	  "encoding/json"
    "github.com/gorilla/mux"
)


// Index Path //
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Elastic search API using GO!")
}

// Search //
func Search(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    query := url.QueryEscape(name)

    // URL //    
    url := fmt.Sprintf("http://192.168.8.20:9200/product/list/_search?q=%s", query)

    // Build request //
    req, err := http.NewRequest("GET", url, nil)
    if err != nil{
      log.Fatal("New Request : ", err)
      return
    }

    /*
    // Ambil response code http header jika, rest service tidak aktif / error //
    // http-200 --> OK //
    // else failed //
    response_code, err := json.Marshal(r.Header.Get("Status Code"))
    if err != nil{
        log.Fatal("Get HTTP Response Code : ", err)
        return
    }

    fmt.Println(string(byte))
    */

    client := &http.Client{}

    resp, err := client.Do(req)
    if err != nil{
      log.Fatal("Do", err)
    return
    }


    defer resp.Body.Close()

    // Get All HTTP BODY //

    jsonData, err := ioutil.ReadAll(resp.Body)
    if err != nil{
    log.Println(err)
    }

    // Convert into String //
    json_string := string(jsonData)



    // Unmarshall JSON into Go Struct //
    var post Post
    json.Unmarshal([]byte(json_string), &post)

    fmt.Println(post)

    // Marshall Go struct into JSON //
    b, err := json.Marshal(post)
    if err != nil {
        fmt.Println(err)
        return
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintln(w, string(b))

    /*
    if err := json.NewEncoder(w).Encode(b); err != nil{
        panic(err)
    }
    */

    //fmt.Println(string(b))
}

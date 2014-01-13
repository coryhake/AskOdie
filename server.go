package main

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "log"
)

type Options struct {
  Path string
  Port string
}

func main() {

  // set default options
  op := &Options{Path: "./", Port: "8006"}

  // read config file into memory
  data, jsonErr := ioutil.ReadFile("./config.json")
  if jsonErr != nil {
    log.Println("JSONReadFileError: ", jsonErr) 
  }
  
  // parse config file, store results in "op"
  json.Unmarshal(data, op)

  log.Println("Parsed options from config file: ", op)
  
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    log.Println(r.RemoteAddr)
      http.FileServer(http.Dir(op.Path)).ServeHTTP(w, r)
  })
  
  err := http.ListenAndServe(":" + op.Port, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "os"
    "time"
    "github.com/nats-io/nats"
)

type server struct {
    nc *nats.Conn
}

type Request struct {
    A               int                     `json:"a"`
    B               int                     `json:"b"`
}

func (s server) sum(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var (
        req     Request
    )
    errReq := decoder.Decode(&req)
    if errReq != nil {
        panic(errReq)
    }
    e, errMarshal := json.Marshal(req)
    if errMarshal != nil {
        fmt.Println(errMarshal)
        return
    }
    response, err := s.nc.Request("add", []byte(e), 2*time.Second)
    if err != nil {
        log.Println("Error making NATS request:", err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response.Data)
}

func (s server) substract(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var (
        req     Request
    )
    errReq := decoder.Decode(&req)
    if errReq != nil {
        panic(errReq)
    }
    e, errMarshal := json.Marshal(req)
    if errMarshal != nil {
        fmt.Println(errMarshal)
        return
    }
    response, err := s.nc.Request("substract", []byte(e), 2*time.Second)
    if err != nil {
        log.Println("Error making NATS request:", err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response.Data)
}

func (s server) multiply(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var (
        req     Request
    )
    errReq := decoder.Decode(&req)
    if errReq != nil {
        panic(errReq)
    }
    e, errMarshal := json.Marshal(req)
    if errMarshal != nil {
        fmt.Println(errMarshal)
        return
    }
    response, err := s.nc.Request("multiply", []byte(e), 2*time.Second)
    if err != nil {
        log.Println("Error making NATS request:", err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response.Data)
}

func main() {
    var s server
    var err error
    uri := os.Getenv("NATS_URI")

    nc, err := nats.Connect(uri)
    if err == nil {
        s.nc = nc
    }
    if err != nil {
        log.Fatal("Error establishing connection to NATS:", err)
    }

    fmt.Println("Connected to NATS at:", s.nc.ConnectedUrl())
    http.HandleFunc("/sum", s.sum)
    http.HandleFunc("/substract", s.substract)
    http.HandleFunc("/multiply", s.multiply)

    fmt.Println("Server listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

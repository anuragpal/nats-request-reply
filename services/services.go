package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "github.com/nats-io/nats"
)

type Request struct {
    A               int                     `json:"a"`
    B               int                     `json:"b"`
}

type Response struct {
    Result          int                     `json:"result"`
}

func main() {
    uri := os.Getenv("NATS_URI")
    var err error
    var nc *nats.Conn
    nc, err = nats.Connect(uri)
    if err != nil {
        log.Fatal("Error establishing connection to NATS:", err)
    }

    nc.Subscribe("add", func(m *nats.Msg) {
        var (
            req Request
            res Response
        )
        if err := json.Unmarshal(m.Data, &req); err != nil {
            panic(err)
        }

        res.Result = req.A + req.B
        e, errMarshal := json.Marshal(res)
        if errMarshal != nil {
            fmt.Println(errMarshal)
            return
        }
        nc.Publish(m.Reply, []byte(e))
    })

    nc.Subscribe("substract", func(m *nats.Msg) {
        var (
            req Request
            res Response
        )
        if err := json.Unmarshal(m.Data, &req); err != nil {
            panic(err)
        }

        res.Result = req.A - req.B
        e, errMarshal := json.Marshal(res)
        if errMarshal != nil {
            fmt.Println(errMarshal)
            return
        }
        nc.Publish(m.Reply, []byte(e))
    })

    nc.Subscribe("multiply", func(m *nats.Msg) {
        var (
            req Request
            res Response
        )
        if err := json.Unmarshal(m.Data, &req); err != nil {
            panic(err)
        }

        res.Result = req.A * req.B
        e, errMarshal := json.Marshal(res)
        if errMarshal != nil {
            fmt.Println(errMarshal)
            return
        }
        nc.Publish(m.Reply, []byte(e))
    })

    if err := http.ListenAndServe(":8181", nil); err != nil {
        log.Fatal(err)
    }
}

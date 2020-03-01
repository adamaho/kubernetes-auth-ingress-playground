package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Success(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "success")
}

func Fail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    w.WriteHeader(http.StatusUnauthorized)
    fmt.Fprintf(w, "fail")
}

func main() {
    router := httprouter.New()
    router.GET("/auth-success", Success)
    router.GET("/auth-fail", Fail)

    log.Fatal(http.ListenAndServe(":8080", router))
}
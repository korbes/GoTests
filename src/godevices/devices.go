package main

import (
    "fmt"
    "net/http"
    "time"
)

type device_address struct
{
	name string
	instant time.Time
}

var devices map[string]device_address

func handler(w http.ResponseWriter, r *http.Request) {
	for key, value := range devices {
	    fmt.Fprintf(w, "Device %s\tAddr:%s\tLast Update:%s\n", key, value.name, value.instant.String())
	}
}

func device_adder_handler(w http.ResponseWriter, r *http.Request) {
	device_name := r.URL.Path[len("/device/"):]
	devices[device_name] = device_address{r.RemoteAddr, time.Now()}
    fmt.Fprintf(w, "Thanks %s!", device_name)
}

func main() {
	devices = make(map[string]device_address)
	http.HandleFunc("/device/", device_adder_handler)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
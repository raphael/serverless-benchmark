package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/cloud-functions-go/nodego"
	"github.com/raphael/recorder/gen/recorder"
	"github.com/raphael/recorder/poster"
	"github.com/raphael/recorder/sieve"
)

type input struct {
	N int
}

func main() {
	flag.Parse()
	p := poster.New("optima-tve.appspot.com")
	http.HandleFunc(nodego.HTTPTrigger, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("empty payload"))
			return
		}
		var i input
		err = json.Unmarshal(b, &i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		_, dur := sieve.Eratosthenes(i.N)
		err = p.Post(r.Context(), &recordersvc.Datapoint{
			Service: "google",
			Name:    fmt.Sprintf("sieve-%d", i.N),
			Value:   dur,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.FormatFloat(dur, 'E', -1, 64)))
	})
	nodego.TakeOver()
}

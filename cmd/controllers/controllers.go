package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nasa-api/config"
	"nasa-api/models"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sibipro/go-logutils"
)

func TechPort(ev config.EnviromentVariables) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var url string = ev.NasaUrl + "techport/api/projects/" + "94206" + "?api_key=" + ev.NasaAuthKey
		fmt.Print(url)
		resp, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		jres, err := json.Marshal(response)

		response := models.TechportResponse{}
		response.Title = jres.project.title

		var techportResponse models.TechportResponse
		json.Unmarshal(body, &techportResponse)

		_, error := w.Write(jres)
		if error != nil {
			log.Fatal(logutils.LogPrefixGenericError, errors.WithStack(err))
		}
	}
}

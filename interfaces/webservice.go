package interfaces

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"url_shortener/app"
)

type UrlRequest struct {
	Url string `json:"url"`
}

type UrlResponse struct {
	Url string `json:"shortUrl"`
}

type Controller struct {
	app app.UrlServicer
}

func NewController(service app.UrlServicer) Controller {
	return Controller{
		app: service,
	}
}

func (c Controller) Run(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/{shortUrl}", c.GetUrl).Methods("GET")
	r.HandleFunc("/", c.CreateShortUrl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func (c Controller) GetUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	url, ok := params["shortUrl"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	originUrl, err := c.app.Get(url)
	if err != nil {
		err := json.NewEncoder(w).Encode(err.Error())
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Redirect(w, r, originUrl, http.StatusSeeOther)
}

func (c Controller) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var ur UrlRequest

	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl := c.app.Create(ur.Url)

	err = json.NewEncoder(w).Encode(UrlResponse{Url: shortUrl})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

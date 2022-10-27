package handlers

import (
	"log"
	"net/http"

	"github.com/yagmurozden/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	// d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

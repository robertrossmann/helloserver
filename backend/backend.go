package backend

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Backend struct {
	PackageRoot string
}

func NewBackend() *Backend {
	return &Backend{}
}

func (b *Backend) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, world!")); err != nil {
			log.Println(fmt.Errorf("writing response: %w", err))
		}
	})

	log.Println("starting server")
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}

	log.Println("server stopped")
	return nil
}

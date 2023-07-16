package handler

import (
	"github.com/jpdel518/icon_generator/service"
	"log"
	"net/http"
	"strconv"
)

func NewHandler(iconService service.Service, identiconService service.Service) {
	http.HandleFunc("/icon", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method not allowed"))
			return
		}

		// parameter
		q := request.URL.Query()
		letter := q.Get("letter")
		size, _ := strconv.Atoi(q.Get("size"))

		// generate
		icon, err := iconService.Generate(letter, size)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		log.Println("return png icon")
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Disposition", "attachment; filename=identicon.png")
		writer.Header().Set("Content-Type", "application/octet-stream")
		writer.Write(icon)
	})
	http.HandleFunc("/identicon", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method not allowed"))
			return
		}

		// parameter
		q := request.URL.Query()
		letter := q.Get("letter")
		size, _ := strconv.Atoi(q.Get("size"))

		// generate
		identicon, err := identiconService.Generate(letter, size)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		log.Println("return identicon")
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Disposition", "attachment; filename=identicon.png")
		writer.Header().Set("Content-Type", "application/octet-stream")
		writer.Write(identicon)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to open http port: %v", err)
	}
}

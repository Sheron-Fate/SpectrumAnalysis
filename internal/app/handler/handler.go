package handler

import (
	"colorLex/internal/app"
	"colorLex/internal/app/repository"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	tmpl      = template.Must(template.ParseGlob("templates/*.html"))
	minioBase = os.Getenv("MINIO_BASE_URL")
)

func ListPigments(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	pigments := repository.FilterPigments(q)
	requestCount := repository.RequestPigmentCount("app1")

	data := struct {
		Pigments  []app.Pigment
		Q         string
		MinioBase string
		RequestCount  int
		RequestID     string
	}{
		Pigments:  pigments,
		Q:         q,
		MinioBase: minioBase,
		RequestCount:  requestCount,
		RequestID:     "app1",
	}
	tmpl.ExecuteTemplate(w, "Pigments.html", data)
}

func ShowPigment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	s := repository.GetPigment(id)
	if s == nil {
		http.NotFound(w, r)
		return
	}
	data := struct {
		Pigment   app.Pigment
		MinioBase string
	}{
		Pigment:   *s,
		MinioBase: minioBase,
	}
	tmpl.ExecuteTemplate(w, "Pigment.html", data)
}

func ShowRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	a := repository.GetRequest(id)
	if a == nil {
		http.NotFound(w, r)
		return
	}
	pigments := repository.GetPigmentsByIDs(a.PigmentIDs)

	data := struct {
		Request       app.AnalysisRequest
		Pigments  []app.Pigment
		MinioBase string
	}{
		Request:       *a,
		Pigments:  pigments,
		MinioBase: minioBase,
	}
	tmpl.ExecuteTemplate(w, "AnalysisRequest.html", data)
}

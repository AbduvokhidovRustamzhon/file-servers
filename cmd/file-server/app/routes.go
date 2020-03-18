package app

import "net/http"

func (s *server) InitRoutes() {
	mux := s.router.(*http.ServeMux)

	mux.HandleFunc("/api/files", s.handleMultipartUpload)
	mux.Handle("/serverdata/", http.StripPrefix("/serverdata/", http.FileServer(http.Dir(s.storagePath))))
}

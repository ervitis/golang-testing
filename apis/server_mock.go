package apis

/*
func setupServer(s *routes.Server, hs Handlers) *http.Server {
	r := &mux.Router{}

	for p, h := range hs {
		r.HandleFunc(p, h)
	}

	return &http.Server{
		Addr:    s.Url(),
		Handler: r,
	}
}
*/
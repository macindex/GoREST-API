package handlers


type handlers struct {
	useCases useCases.useCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers()
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()
	return http.ListenAndServer(
		fmt.Printf(":%v", port),
		nil,
	)
}

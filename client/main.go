package client

func main() {
	repos := repositories.New()
	useCases := useCases.New(repos)
	h := handlers.New(useCases)
	h.Listen(8080)
}



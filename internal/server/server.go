package server

type ServerEcho struct {
}

type server interface {
	Start()
}

func GetProvider() server {
	return ServerEcho{}
}

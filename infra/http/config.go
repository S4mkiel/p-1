package http

type Config struct {
	Http                   int  `env:"HTTP_PORT" required:"true"` // rota localhost:(8080)
	DisableStartupeMessage bool //true ou false
}

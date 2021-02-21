package plug

type Plug interface {
	SayHello(n string) (string, error)
}

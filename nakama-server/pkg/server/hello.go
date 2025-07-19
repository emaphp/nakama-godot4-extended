package server

type HelloWorldGreeter struct{}

func NewHelloWorldGreeter() *HelloWorldGreeter {
	return &HelloWorldGreeter{}
}

func (greeter *HelloWorldGreeter) Greet() string {
	return "Hola Mundo"
}

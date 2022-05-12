package pkg

type Apple struct {
	sweetiness string
}

func NewApple() *Apple {
	apple := Apple{sweetiness: "sour"}
	return &apple
}

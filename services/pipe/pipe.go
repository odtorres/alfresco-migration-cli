package pipe

func run(f func() (string, error)) (string, error) {
	return f()
}

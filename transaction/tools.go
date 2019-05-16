package transaction

// must be error-free, panic otherwise.
func must(err error) {
	if err != nil {
		panic(err)
	}
}

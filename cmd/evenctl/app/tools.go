package app

func must(err error) {
	if err != nil {
		panic(err)
	}
}

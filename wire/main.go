package main

func main() {
	app, err := initApp()
	must(err)

	must(app.Start())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

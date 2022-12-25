package main

func main() {
	app := &App{}
	app.Initialize(GetConfig())
	app.Run(":3000")
}

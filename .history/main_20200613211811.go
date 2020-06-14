package main

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
}

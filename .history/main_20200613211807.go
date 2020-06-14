package main

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", statuc.LocalFile("./views", true)))
}

package main

func main() {
	router := gin.Default()

	router.Use(static.Server("/", statuc.LocalFile("./views", true)))
}

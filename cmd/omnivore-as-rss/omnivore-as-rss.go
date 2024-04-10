package main

import "omnivore-as-rss/internal"

func main() {
	internal.InitConfig()
	internal.Serve()
}

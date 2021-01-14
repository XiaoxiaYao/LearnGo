package main

import (
	"LearnGOGO/infra"
	"fmt"
)

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return infra.Retriever{}
}

func main() {
	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}

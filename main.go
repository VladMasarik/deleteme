package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())
	config, err := rest.inclusterconfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig()
	if err != nil {
		return nil, err
	}
}

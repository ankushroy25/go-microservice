package main

import "flag"

func main() {

	listenAddress := flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")
	svc := NewMetricService(&priceFetcher{})

	server := newJsonApiServer(*listenAddress, svc)
	server.Run()
}

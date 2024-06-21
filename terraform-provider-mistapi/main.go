package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-mistapi/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	// Clean up log output
	// See https://developer.hashicorp.com/terraform/plugin/log/writing#legacy-logging
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	opts := providerserver.ServeOpts{
		Address: "hashicorp.com/juniper/mistapi",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(), opts)
	if err != nil {
		log.Fatal(err.Error())
	}

}

package main

import (
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/labd/terraform-provider-commercelayer/commercelayer"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: commercelayer.Provider()}

	if debugMode {
		opts.Debug = true
		opts.ProviderAddr = "registry.terraform.io/labd/commercelayer"
	}

	plugin.Serve(opts)
}

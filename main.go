package main

import (
	"flag"
	"log"
	"os"
)

const HumanitecHost string = "api.humanitec.io"

func main() {
	humanitecToken := flag.String("humanitec-token", "", "<Humanitec Access Token>")
	humanitecOrganization := flag.String("organization", "", "<Humanitec organization>")
	mode := flag.String("mode", "registry-credentials", "fetch [registry-credentials], [notify] for new image")

	flag.Parse()

	if *humanitecToken == "" || *humanitecOrganization == "" || *mode == "" {
		log.Println("All arguments are required. Use -h for full syntax")
		os.Exit(1)
	}

	if *mode != "notify" {
		// latestReleaseDetails, err :=
		getRegistryCredentials(*humanitecToken, *humanitecOrganization, HumanitecHost)
	}
	// else {
	// 	notifyForNewBuild(*humanitecToken, *humanitecOrganization, HUMANITEC_HOST)
	// }

}

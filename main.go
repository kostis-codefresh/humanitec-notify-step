package main

import (
	"flag"
	"fmt"
	"os"
)

// HumanitecHost is the default server for Humanitec API calls
const HumanitecHost string = "api.humanitec.io"

func main() {
	humanitecToken := flag.String("humanitec-token", "", "<Humanitec Access Token>")
	humanitecOrganization := flag.String("organization", "", "<Humanitec organization>")
	mode := flag.String("mode", "registry-credentials", "fetch [registry-credentials], [notify] for new image")

	imageName := flag.String("image-name", "sample-image", "name to appear in Humanitec as source of image")
	imageGitCommit := flag.String("image-git-commit", "", "full Git hash used as source for the image")
	imageGitBranch := flag.String("image-git-branch", "main", "Git branch that was used for this image")
	imageURL := flag.String("image-url", "", "Image URL like registry.humanitec.io/<org>/sample-image:sample-tag")

	flag.Parse()

	if *humanitecToken == "" || *humanitecOrganization == "" || *mode == "" {
		fmt.Println("Missing arguments. Use -h for full syntax")
		os.Exit(1)
	}

	if *mode != "notify" {
		getRegistryCredentials(*humanitecToken, *humanitecOrganization, HumanitecHost)
	} else {

		if *imageName == "" || *imageGitCommit == "" || *imageGitBranch == "" || *imageURL == "" {
			fmt.Println("Missing arguments. Use -h for full syntax")
			os.Exit(1)
		}

		newBuild := newBuildDetails{
			Branch: *imageGitBranch,
			Commit: *imageGitCommit,
			Image:  *imageURL,
		}

		notifyForNewBuild(*humanitecToken, *humanitecOrganization, HumanitecHost, *imageName, newBuild)
	}

}

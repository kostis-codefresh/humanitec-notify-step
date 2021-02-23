package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const OutputFile string = "creds.json"

func getRegistryCredentials(humanitecToken string, humanitecOrganization string, humanitecHost string) {
	url := "https://" + humanitecHost + "/orgs/" + humanitecOrganization + "/registries/humanitec/creds"

	req, err := http.NewRequest("GET", url, nil)
	bearer := "Bearer " + humanitecToken
	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		log.Fatal("Could not access " + url)
	}

	log.Printf("Response status of registry credentials for %s is %s\n", url, resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(OutputFile, body, 0700)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved response to " + OutputFile)
}

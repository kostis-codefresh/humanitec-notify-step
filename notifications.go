package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type newBuildDetails struct {
	Branch string `json:"branch"`
	Commit string `json:"commit"`
	Image  string `json:"image"`
}

func notifyForNewBuild(humanitecToken string, humanitecOrganization string, humanitecHost string, imageID string, newBuild newBuildDetails) {
	url := "https://" + humanitecHost + "/orgs/" + humanitecOrganization + "/images/" + imageID + "/builds"

	jsonValue, _ := json.Marshal(newBuild)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}

	bearer := "Bearer " + humanitecToken
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		log.Fatal("Could not access " + url)
	}

	log.Printf("Humanitec was notified for new build of image %s\n", imageID)
}

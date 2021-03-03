# Notify Humanitec for new images



[![Go Report Card](https://goreportcard.com/badge/github.com/kostis-codefresh/humanitec-notify-step)](https://goreportcard.com/report/github.com/kostis-codefresh/humanitec-notify-step)
[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/kostis-codefresh/humanitec%2Fplugin?type=cf-1&key=eyJhbGciOiJIUzI1NiJ9.NWIwZmYzYmE1ODAzMWUwMDAxYjJlOGUw.dFYNhKzaLSj6l3LoOWe0DlGiuY0McdrmrgHWtWNC9WE)]( https://g.codefresh.io/pipelines/edit/new/builds?id=6034da536a06496e41292c33&pipeline=plugin&projects=humanitec&projectId=5fa4344ce0a5be9001c62a8f)

This is a mini CLI + Container image to work with [Humanitec](https://humanitec.com/) notifications.

## How Humanitec container images work

Humanitec has an internal Docker registry at `registry.humanitec.io` that you can push images to. The credentials for the registry are dynamic
and you need to fetch them with an API call.

Once you have the credentials, you can push an image to it like any other container registry. After you finish with the push
you also need to notify Humanitec about the new image build with another API call. Here is the whole process:

1. Signup for Humanitec
1. Go into the Humanitec UI and create a new token from your organization screen in the *Images* section
1. Perform a call to `https://api.humanitec.io/orgs/<your org>/registries/humanitec/creds`. You will get a JSON with Docker credentials
1. Use the credentials to push an image like any other standard container registry
1. Perform another call to `https://api.humanitec.io/orgs/<your org>/images/<your image name>/builds` with the [correct payload](https://api-docs.humanitec.com/#tag/Image/paths/~1orgs~1{orgId}~1images~1{imageId}~1builds/post)
1. The image will appear in the Humanitec UI and you can use it in your deployments

## How to build

Run:

 *  `go build` to get the executable OR
 *  `docker build . -t humanitec-notify-container` to create a container image if you prefer docker instead or don't have access to a Go dev environment

A prebuilt image is already available at [https://hub.docker.com/r/kostiscodefresh/humanitec-notify-step](https://hub.docker.com/r/kostiscodefresh/humanitec-notify-step)

## How to use the CLI

This CLI can be used for the two API calls mentioned in the first section

1. Run `humanitec-notify-step -organization <your-org> -humanitec-token <your-token> -mode fetch`. This will fetch the registry credentials and save them to a file called `creds.json`
1. Push an image to `registry.humanitec.io` either manually or via an automated method using the credentials
1. Run `humanitec-notify-step -organization <your-org> -humanitec-token <your-token> -mode notify -image-name my-image -image-git-commit my-git-hash -image-git-branch main -image-url registry.humanitec.io/my-org/humanitec-example:latest`.  This will notify Humanitec about the new image that you pushed

You can use this CLI in any CI/CD system and the Dockerhub image in any container based pipeline.

## Codefresh example

See an example for Codefresh at [https://github.com/kostis-codefresh/humanitec-example](https://github.com/kostis-codefresh/humanitec-example)

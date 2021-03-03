# Notify Humanitec for new images



[![Go Report Card](https://goreportcard.com/badge/github.com/kostis-codefresh/humanitec-notify-step)](https://goreportcard.com/report/github.com/kostis-codefresh/humanitec-notify-step)

This is a mini CLI + Container image to work with [Humanitec](https://humanitec.com/) notifications.

## How Humanitec container images work

Humanitec has an internal Docker registry at `registry.humanitec.io` that you can push images to. The credentials for the registry are dynamic
and you need to fetch them with an API call.

Once you have the credentials, you can push an image to it like any other container registry. After you finish with the push
you also need to notify Humanitec about the new image build with another API call. Here is the whole process:

1. Signup for Humanitec
1. Go into the Humanitec UI and create a new token from your organization screen in the *Images* section
1. Perform a call to `https://api.humanitec.io/orgs/<your org>/registries/humanitec/creds`. You will get a JSON with Docker credentials
1. Use the credentials to push an image
1. Perform another call to `https://api.humanitec.io/orgs/<your org>/images/<your image name>/builds` with the [correct payload](https://api-docs.humanitec.com/#tag/Image/paths/~1orgs~1{orgId}~1images~1{imageId}~1builds/post)
1. The image will appear in the Humanitec UI and you can use it in your deployments

## How to use the CLI

This CLI can be used for the two API calls mentioned in the previous section

1. Run `humanitec-notify-step -organization <your-org> -humanitec-token <yourtoken> -mode fetch`. The will fetch the registry credentials to a file called `creds.json`
1. Run `humanitec-notify-step -organization <your-org> -humanitec-token <yourtoken> -mode notify -image-name my-image -image-git-commit my-git-hash -image-git-branch main -image-url registry.humanitec.io/my-orgo/humanitec-example:latest` this will notify about a new image

You can use this CLI in an container based CI/CD system. The Docker image is available at [https://hub.docker.com/r/kostiscodefresh/humanitec-notify-step](https://hub.docker.com/r/kostiscodefresh/humanitec-notify-step)

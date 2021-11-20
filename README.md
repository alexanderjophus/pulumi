# Pulumi

This project was created live on twitch, the goal was to explore pulumi and see how good it really is.

This is __not__ meant to show best practices, simply how I tackled pulumi at first sight.

## Thoughts:

Pulumi is great, it was pretty simple to setup.
Overall I would recommend it.
Obviously it's main competition is terraform which has grand adoption everywhere already.
I'm not sure I'd recommend switching if you already have terraform in place, but if you have nothing then this is great.

I really enjoyed the testing. It was tricky to set up the unit tests, and ultimately we failed to do that, but we did spin up an integration test which created a cluster, deployed grafana, then shut down all without failure.
That was a huge win.

## How to run

(I haven't fully tested this bit)

Create a new stack

```sh
pulumi stack init dev
```

Set the gcp zone, and project

```sh
pulumi config set gcp:zone <your_favourite_zone>
pulumi config set gcp:project <your_project_id>
```

Once all this is set up, you should be able to run

```sh
pulumi up
```

This will spin up a new cluster with the name `demo-cluster` (and some unique identifier suffixed) in your project.

You can then connect to the cluster to check everything is deployed correctly.

```sh
pulumi stack output kubeconfig --show-secrets > kubeconfig
export KUBECONFIG=$PWD/kubeconfig
```

You can also test the project with

```sh
go test -timeout 600s
```

I set a timeout of 600s as we are creating a cluster, deploying to it, then tearing it down - this took about 9 minutes for me.

## Notes

- Seems like it needs to be run at least twice to work, the crd for argo application doesn't want to wait for the helm chart to deploy the application crds.

## Future

In the future I'd like to expand this project, I have the following questions:

- [x] Can I set up Argo to point to my [iris](https://github.com/trelore/iris-classification) project
  - Motivation: I want to deploy services in a declarative way
- [x] Can I deploy other monitoring tools?
  - Motivation: I want to monitor my application as if it were a prod app
- [ ] Add API gateway
  - Motivation: I want to try it
- [ ] Can I use CRDS for Argo/Traefik to avoid manual configuration?
  - Motivation: No manual steps!
- [x] Can I set the NS and use the returned ns later?
  - Motivation: Neater codebase?
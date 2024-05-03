# Provider UpCloud

`provider-upcloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
UpCloud API.

Please note that this project is currently in early alpha version, we do not recommend running it in production yet.

## Developing

Run against a Kubernetes cluster:

```console
make run
```

Run end-to-end tests:
```console
make e2e
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/UpCloudLtd/provider-upcloud/issues).

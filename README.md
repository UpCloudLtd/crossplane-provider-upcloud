# Crossplane Provider UpCloud

`provider-upcloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
UpCloud API.

Please note that this project is currently in early alpha version, we do not recommend running it in production yet.

## Quickstart

1. You will need a Kubernetes cluster to start using a provider. Any cluster will do, for testing and development purposes Kind cluster will do:
    ```
    kind create cluster -n crossplane-test
    ```

2. Next install UpCloud provider. The simplest way to do that is to just apply the following yaml:
    ```
    apiVersion: pkg.crossplane.io/v1
    kind: Provider
    metadata:
      name: provider-upcloud
    spec:
      package: xpkg.upbound.io/upcloud/provider-upcloud:v0.0.2
    ```

	Make sure to change the version to the latest one.

3. Next, you need to create a a `Secret` with your UpCloud API credentiqals and a `ProviderConfig` that will use them to provision your infra. Just apply this yaml:
    ```
		apiVersion: v1
		kind: Secret
		metadata:
			name: example-provider-creds
			namespace: default
		type: Opaque
		stringData:
			credentials: |
				{
					"username": "username",
					"password": "password123"
				}
		---
		apiVersion: provider.upcloud.io/v1beta1
		kind: ProviderConfig
		metadata:
			name: default
		spec:
			credentials:
				source: Secret
				secretRef:
					name: example-provider-creds
					namespace: default
					key: credentials
		```

4. And now you can start creating your infra. Check our [examples](examples/resources/) to see what Managed Resources you can use and how. Have fun!

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/UpCloudLtd/provider-upcloud/issues).

## Developing

For development instructions see [DEVELOPING.md](DEVELOPING.md)
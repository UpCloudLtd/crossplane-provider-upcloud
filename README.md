# Crossplane Provider UpCloud

`crossplane-provider-upcloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
UpCloud API.

The name of the provider in UpBound Marketplace is `provider-upcloud`.

Please note that this project is currently in early alpha version, we do not recommend running it in production yet.

## Quickstart

1. You will need a Kubernetes cluster to start using a provider. You can use any Kubernetes cluster, but for testing and development purposes, a Kind cluster is recommended:
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
      package: xpkg.upbound.io/upcloud/provider-upcloud:v0.0.7
    ```

 Make sure to change the version to the latest one.

3. Next, you need to create a `Secret` with your UpCloud API credentials and a `ProviderConfig` that will use them to provision your infrastructure. Replace `ucat_TOKEN` with your UpCloud API token and apply this yaml (you can also use username and password instead of token for now, but we recommend using token instead):
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
          "token": "ucat_TOKEN"
        }
    ---
    apiVersion: provider.upcloud.com/v1beta1
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

4. And now you can start creating your infrastructure. Check our [examples](examples/resources) to see what Managed Resources you can use and how. Have fun!

## Missing resources

While this provider allows you to manage most of the UpCloud services, support for some is still missing. The missing list includes:
- [Load Balancers](https://developers.upcloud.com/1.3/17-managed-loadbalancer/)
- [Network Gateways](https://developers.upcloud.com/1.3/19-network-gateways/)
- [Floating IP address](https://developers.upcloud.com/1.3/10-ip-addresses/#creating-floating-ips)

We are currently working on adding support for them, thank you for your patience!

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/UpCloudLtd/crossplane-provider-upcloud/issues).

## Developing

For development instructions see [DEVELOPING.md](DEVELOPING.md)
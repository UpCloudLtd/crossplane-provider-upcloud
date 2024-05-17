# Developing the Provider

This is a quick guide on how to develop this provider. For general info about developing providers with upjet please see this [documentation](https://github.com/crossplane/upjet/blob/main/docs/README.md).

## Quickstart

1. To run provider locally you will need:
    - Go version 1.22 or higher
    - goimports tool (can be installed with `go install golang.org/x/tools/cmd/goimports@latest`)
    - Terraform 1.7.5 or higher (not strictly necessary, but recommended)

2. Create a local cluster using Kind:
    ```
    kind create cluster -n crossplane-dev
    ```

3. Generate the CRDs and controllers for them:
    ```
    make generate
    ```
    Note that in some cases this might fail do to unresolved references in Go code (in `api` directory). This usually happends when you changed CRD references in their config. It is safe to just delete the offending file - the `generate` target will just regenerate everything. You can also just delete all previously gennerated APIs with `make cleanup-apis`.

4. Install the CRDs in your cluster:
    ```
    make install
    ```

5. Run the controller manager outside of the cluster:
    ```
    make run
    ```

6. Go to `examples/providerconfig`, copy the secret template and populate it with your UpCloud API username and password:
    ```
    cp examples/providerconfig/secret.yaml.tmpl examples/providerconfig/secret.yaml
    ```

7. Apply the whole provider config directory:
    ```
    make deploy-providerconfigs
    ```

8. You are ready to test things. You can start by applying some of the files in the [examples/resources](examples/resources/) directory. Note that some resources require additional objects to be created. For example databases require a secret with credentials. We hold those objects in `examples/resourceconfig` directory so that they can be applied separately. We need to apply / delete them separately because otherwise the deletion process can hang due to not having a valid TF configuration (because for example password is not populated, and for TF that is a required field). You can apply all the example resources in one go using
    ```
    make deploy-examples
    ```

9. You can cleanup by running:
    ```
    make delete-examples
    make delete-providerconfigs
    ```


## Debugging

If you want to use debugger you will need to run the controller manager in a way that your IDE can attach a debugger process to it.
If you use VSCode, you can use the following run configuration and run the controller manager from VSCode UI:
   ```json
   {
       "name": "Run Provider",
       "type": "go",
       "request": "launch",
       "mode": "auto",
       "program": "${workspaceFolder}/cmd/provider",
       "args": ["--debug"],
       "env": {
           "TERRAFORM_VERSION": "1.7.5",
           "TERRAFORM_PROVIDER_SOURCE": "UpCloudLtd/upcloud",
           "TERRAFORM_PROVIDER_REPO": "https://github.com/UpCloudLtd/terraform-provider-upcloud",
           "TERRAFORM_PROVIDER_VERSION": "5.2.3",
           "TERRAFORM_PROVIDER_DOWNLOAD_NAME": "terraform-provider-upcloud",
           "TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX": "https://releases.hashicorp.com/terraform-provider-upcloud/5.2.2",
           "TERRAFORM_NATIVE_PROVIDER_BINARY": "terraform-provider-upcloud_v5.2.3",
           "TERRAFORM_DOCS_PATH": "docs/resources"
       }
   }
   ```

Unfortunately debugger is only moderately useful here, since most of the logic is actually in the Terraform provider, which is already compiled. There are two ways you can work around this.

First is to use locally built TF provider. You can do that by [building the TF provider locally](https://github.com/UpCloudLtd/terraform-provider-upcloud/blob/main/DEVELOPING.md) and simply replacing the following env variables in `Makefile` or in your IDE launch configuration:
    ```shell

    // "TERRAFORM_PROVIDER_SOURCE": "UpCloudLtd/upcloud",
    // "TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX": "https://releases.hashicorp.com/terraform-provider-upcloud/5.2.2",
    "TERRAFORM_PROVIDER_SOURCE": "registry.upcloud.com/upcloud/upcloud",
    "TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX": "file:///home/myusername/.terraform.d/plugins/registry.upcloud.com/upcloud/upcloud/5.2.3",
    ```
Just make sure to change the version to whatever was the result of local TF build process.


Another way to debug bugs is to interrupt the Crossplane reconciliation process half-way and continue it manually with your own (locally built) TF provider. To understand how that can be done you first need to understand how upjet-generated controllers work.
Let's say you want to create a new Managed Resource - UpCloud Server. You will create a simple yaml file with Server definition and spec - [like the one in examples](examples/resources/server.yaml) - and apply it. When you do that Kubernetes will add a new `Server` object to etcd and assign a UID to it.
From there on upjet-generated Server controller will do the following things:
1. Create a new temporary TF workspace. By default it will be `/tmp/k8s_object_uid`
2. It will generate a `main.tf.json` file there with a valid Terraform config for `upcloud_server` resource.
3. It will run `terraform init` in that directory
4. It will run various TF commands, depending on the need. Usually either `apply` or `destroy`.

Knowing that you can apply your yaml file, wait for a specific error that you want to debug to appear in logs and just stop the controller manager. Then check the UID of the resource in question, and go to a relevant temporary directory. If the resource you want to debug has UID of `1234`, you just go to `/tmp/1234`. There you can just delete the downloaded TF binary, change the `main.tf.json` file to point to your locally built TF provider and do a regular debugging as you would do when just developing Terraform.

## Running tests

First make sure you have `UPCLOUD_USERNAME` and `UPCLOUD_PASSWORD` env variables set.

Then just run `make e2e`. This command will:
- Create a new Kind cluster locally
- Install Crossplane and our locally built provider there
- Install all the CRDs
- Run the test setup script. This script creates a credential secret, provider config and also applies all the files in `examples/resourceconfig` (which holds some dependencies for our example resources that we want to apply separately)
- Apply all the files in `examples/resources` and wait for them to get to ready state
- Test importing all of those resources
- Test deleting all of those resources

## Releasing new version

The basic process for releasing a provider is described in [upboud docs](https://docs.upbound.io/upbound-marketplace/packages/#publishing-public-packages). However we need to deviate from it a bit since we use upjet generator. 

Currently the release process is entirely manual and goes as follow:
1. Once all your changes are merged to the `main` branch, create a new tag and push it to the repo. This currently doesn't do anything, but we want to keep things organised and this helps with that.
2. Make sure you have [up cli](https://docs.upbound.io/reference/cli/) installed and login to relevant account.
3. Generate all the CRDs and controllers. This shouldn't really be necessary, but just to be sure:
    ```
    make generate
    ```
4. Build everything. This command will output all the artifacts to `_output` directory:
    ```
    make build
    ```
5. Push the package to the repository with the same version as the tag from step 1. You need to have `crossplane.yaml` file in the directory you run the push command:
    ```
    cd package && up xpkg push upcloud/provider-upcloud:v0.0.2 -f ../_output/xpkg/linux_amd64/provider-upcloud-vXX.XX.XX.xpkg
    ```
6. Verify that it is published by going to [Upbound marketplace](https://marketplace.upbound.io/account/upcloud/provider-upcloud).

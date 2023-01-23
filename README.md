<div align="center">
  <a href="https://www.pokt.network">
    <img src="https://user-images.githubusercontent.com/16605170/74199287-94f17680-4c18-11ea-9de2-b094fab91431.png" alt="Pocket Network logo" width="340"/>
  </a>
</div>

# Pocket Operator

Deploy and manage pocket nodes on Kubernetes.

<div>
  <a  href="https://godoc.org/github.com/pokt-network/pocket-operator"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
  <a  href="https://goreportcard.com/report/github.com/pokt-network/pocket-operator"><img src="https://goreportcard.com/badge/github.com/pokt-network/pocket-operator"/></a>
  <a href="https://golang.org"><img  src="https://img.shields.io/badge/golang-v1.11-red.svg"/></a>
  <a  href="https://github.com/tools/godep" ><img src="https://img.shields.io/badge/godep-dependency-71a3d9.svg"/></a>
</div>

## Overview

<div>
    <a  href="https://github.com/pokt-network/pocket-operator/releases"><img src="https://img.shields.io/github/release-pre/pokt-network/pocket-operator.svg"/></a>
    <a  href="https://github.com/pokt-network/pocket-operator/pulse"><img src="https://img.shields.io/github/contributors/pokt-network/pocket-operator.svg"/></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-operator/pulse"><img src="https://img.shields.io/github/last-commit/pokt-network/pocket-operator.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-operator/pulls"><img src="https://img.shields.io/github/issues-pr/pokt-network/pocket-operator.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-operator/releases"><img src="https://img.shields.io/badge/platform-linux%20%7C%20windows%20%7C%20macos-pink.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-operator/issues"><img src="https://img.shields.io/github/issues-closed/pokt-network/pocket-operator.svg"/></a>
</div>

The Pocket Operator extends the Kubernetes control plane to create, read, update
and delete pocket nodes using custom resources.

This operator was built with
[operator-builder](https://github.com/nukleros/operator-builder).

# ⚠️ Under active development

We are currently working on the first release of the operator. That means the API and CRD are subject to change - in fact they most certainly are going to be changed. We will be updating this README as we make progress. Not production ready. Expect issues.

## Installation

We provide kustomize manifests for deploying the operator. To create manifests to deploy the latest version of operator, run the following command:

```
kustomize build "https://github.com/pokt-network/pocket-operator//config/default?ref=master"
```

## Usage

Please see the supported Custom Resource Definitions (CRDs) in the [config/crd](config/crd) directory.

We do not ship Postgres database with this operator, it is up to a user to choose a way to run the database. We suggest to utilize Postgres Operators such as [CrunchyData](https://github.com/CrunchyData/postgres-operator) or [zalando's postgres-operator](https://github.com/zalando/postgres-operator) for that purpose due to additional availability guarantees (replications, backups, etc). We do have an example in the [config/samples](config/samples) directory for one non-replicated Postgres deployment. Once this product matures, we are going to publish a step-by-step guide with a recommended way to deploy postgres operators.

Examples can be found in the [config/samples](config/samples) directory.


## Local Development & Testing

To install the custom resource/s for this operator, make sure you have a
kubeconfig set up for a test cluster, then install the CRDs.

```bash
make install
```

To run the controller locally against your test cluster.

```bash
make run
```

You can then test the operator by creating the sample manifests in another
terminal.

```bash
kubectl apply -f config/samples
```

Give the validators some time to come up.  The validator nodes will go into a crashloop if database is not yet available, but will restart and retry.

Run the dev client in the pocket client container.  See the [pocket development
docs](https://github.com/pokt-network/pocket/tree/main/docs/development#running-localnet)
for more info on using the dev client.

```bash
kubectl exec -it pocket-v1-client -- go run app/client/main.go
```

Finally, once testing is complete you can clean up.  Note: you will need to run
the controller to finalize resource deletion with `make run` if you stopped it.

```bash
make uninstall
```

## Code Generation

The following steps will re-generate the codebase from scratch after making
changes to the configurations, source manifests and/or markers.

*Caution*: These steps will permanently delete any changes you have made
directly to the codebase that is being provisioned by operator-builder.

Re-build the codebase from the existing configurations and source manifests.

```bash
cd .opertator-builder
make operator-build
```

Install the dependencies (postgres operator) in your test cluster.

```bash
make operator-dependencies
```

Copy the modified sample manifests into the `config/samples` directory of the
codebase.

```bash
make operator-samples
cd ../
```

You can now re-test the operator using the [Local Development & Testing
instructions](#local-development-&-testing) above.

## Deploy the Controller Manager

First, set the image name.

```bash
export IMG=myrepo/myproject:v0.1.0
```

Now you can build and push the image.

```bash
make docker-build
make docker-push
```

Then deploy.

```bash
make deploy
```

To clean up.

```bash
make undeploy
```

## Companion CLI

To build the companion CLI.

```bash
make build-cli
```

The CLI binary will get saved to the bin directory.  You can see the help
message with the following.

```bash
./bin/pocketctl help
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/pokt-network/pocket-operator/blob/master/CONTRIBUTING.md) for details on contributions and the process of submitting pull requests.

## Support & Contact

<div>
  <a  href="https://twitter.com/poktnetwork" ><img src="https://img.shields.io/twitter/url/http/shields.io.svg?style=social"></a>
  <a href="https://t.me/POKTnetwork"><img src="https://img.shields.io/badge/Telegram-blue.svg"></a>
  <a href="https://www.facebook.com/POKTnetwork" ><img src="https://img.shields.io/badge/Facebook-red.svg"></a>
  <a href="https://research.pokt.network"><img src="https://img.shields.io/discourse/https/research.pokt.network/posts.svg"></a>
</div>

## License

This project is licensed under the MIT License; see the [LICENSE.md](LICENSE.md) file for details.


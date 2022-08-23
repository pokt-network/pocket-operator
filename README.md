<<<<<<< HEAD
<div align="center">
  <a href="https://www.pokt.network">
    <img src="https://user-images.githubusercontent.com/16605170/74199287-94f17680-4c18-11ea-9de2-b094fab91431.png" alt="Pocket Network logo" width="340"/>
  </a>
</div>

# Project Title

One sentence summary of project
<div>
  <a  href="https://godoc.org/github.com/pokt-network/pocket-core"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
  <a  href="https://goreportcard.com/report/github.com/pokt-network/pocket-core"><img src="https://goreportcard.com/badge/github.com/pokt-network/pocket-core"/></a>
  <a href="https://golang.org"><img  src="https://img.shields.io/badge/golang-v1.11-red.svg"/></a>
  <a  href="https://github.com/tools/godep" ><img src="https://img.shields.io/badge/godep-dependency-71a3d9.svg"/></a>
</div>

## Overview
<div>
    <a  href="https://github.com/pokt-network/pocket-core/releases"><img src="https://img.shields.io/github/release-pre/pokt-network/pocket-core.svg"/></a>
    <a href="https://circleci.com/gh/pokt-network/pocket-core/tree/staging"><img src="https://circleci.com/gh/pokt-network/pocket-core/tree/staging.svg?style=svg"/></a>
    <a  href="https://github.com/pokt-network/pocket-core/pulse"><img src="https://img.shields.io/github/contributors/pokt-network/pocket-core.svg"/></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-core/pulse"><img src="https://img.shields.io/github/last-commit/pokt-network/pocket-core.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-core/pulls"><img src="https://img.shields.io/github/issues-pr/pokt-network/pocket-core.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-core/releases"><img src="https://img.shields.io/badge/platform-linux%20%7C%20windows%20%7C%20macos-pink.svg"/></a>
    <a href="https://github.com/pokt-network/pocket-core/issues"><img src="https://img.shields.io/github/issues-closed/pokt-network/pocket-core.svg"/></a>
</div>

Full Description

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Example usage

```
The most basic example of how you would use the project
```

### Installation

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the step
```

And repeat

```
until finished
```

End with an example of getting data out of the system or using it for a demo

## Documentation

Full usage and options or a link to the docs.pokt.network site

## Running the tests

Explain how to run the automated tests

```
Give an example
```

## Contributing

Please read [CONTRIBUTING.md](https://github.com/pokt-network/repo-template/blob/master/CONTRIBUTING.md) for details on contributions and the process of submitting pull requests.

## Support & Contact

<div>
  <a  href="https://twitter.com/poktnetwork" ><img src="https://img.shields.io/twitter/url/http/shields.io.svg?style=social"></a>
  <a href="https://t.me/POKTnetwork"><img src="https://img.shields.io/badge/Telegram-blue.svg"></a>
  <a href="https://www.facebook.com/POKTnetwork" ><img src="https://img.shields.io/badge/Facebook-red.svg"></a>
  <a href="https://research.pokt.network"><img src="https://img.shields.io/discourse/https/research.pokt.network/posts.svg"></a>
</div>


## License

This project is licensed under the MIT License; see the [LICENSE.md](LICENSE.md) file for details.
=======
A Kubernetes operator built with
[operator-builder](https://github.com/vmware-tanzu-labs/operator-builder).

## Local Development & Testing

To install the custom resource/s for this operator, make sure you have a
kubeconfig set up for a test cluster, then run:

    make install

To run the controller locally against a test cluster:

    make run

You can then test the operator by creating the sample manifest/s:

    kubectl apply -f config/samples

To clean up:

    make uninstall

## Deploy the Controller Manager

First, set the image:

    export IMG=myrepo/myproject:v0.1.0

Now you can build and push the image:

    make docker-build
    make docker-push

Then deploy:

    make deploy

To clean up:

    make undeploy

## Companion CLI

To build the companion CLI:

    make build-cli

The CLI binary will get saved to the bin directory.  You can see the help
message with:

    ./bin/pocketctl help
>>>>>>> 54cd1fd (Add new controller, custom resources)

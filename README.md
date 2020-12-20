# CF Templates

This repository provides a CLI for working with CloudFormation
templates within an AWS account. This is done by utilizing the
goformation Go module to compose templates.

## Setup

The templates within the `cf-templates` directory must first be
built into Go plugin (.so) files so that they can be dynamically
loaded without the user having to know their paths etc. To do
this, run `make plugins` from within the root directory of the
repository.

## Usage

To build a template, use the following command:

`go run *.go -template=sso_permission_sets`
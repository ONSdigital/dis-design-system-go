# dis-design-system-go

Beta rendering library for Dissemination frontend go microservices. `dis-design-system-go` contains templates, localisations, model structs, css and javascript that are core to all dissemination frontend services.

## Getting started

* Run `make help` to see full list of make targets

### Dependencies

To build, run, deploy, test and audit the app you will need some additional tooling:

#### NVM

If you work across multiple Node.js projects there's a good chance they require different Node.js and npm versions.

It is recommended that you use [nvm (Node Version Manager)](https://github.com/creationix/nvm) to switch between versions easily:

1. Install [nvm](https://github.com/nvm-sh/nvm):

   ```shell
   brew install nvm
   ```

   :warning: Make sure to follow the instructions provided at the end of the install to configure up your shell profile.

2. Install the node version specified in [`.nvmrc`](./.nvmrc) through nvm:

   ```shell
   nvm install
   ```

#### Audit

We use `dis-vulncheck` to do auditing, which you will [need to install](https://github.com/ONSdigital/dis-vulncheck).

For Javascript auditing we use `auditjs` which requires you to [setup an OSS Index account](https://github.com/ONSdigital/dp/blob/main/guides/MAC_SETUP.md#oss-index-account-and-configuration)

### Generate the CSS and JS

* Build the CSS and JS, and start the local web server with

  ```shell
  make debug
  ```

* Once built, you can find assets stored on the web server, default location is [localhost:9002/dist/assets/](http://localhost:9002/dist/assets/)

## Go library

`dis-design-system-go` also acts as a Go library which contains template helpers, model structs and components that can be used by the consuming frontend app to serve HTML consistently across the ONS website. To make use of the `go` rendering, you will need to install it within your `go` frontend app.

### Installation

Other than `dis-design-system-go` itself, you will need a utility that can combine service-specific and `dis-design-system-go` assets. We currently use `go-bindata` for this process. From the consuming frontend app, run the following commands to install:

* `dis-design-system-go`: `go get github.com/ONSdigital/dis-design-system-go`

> You can specify a version of `dis-design-system-go` by appending a commit ID or semantic version number to this command. e.g., `go get github.com/ONSdigital/dis-design-system-go@d27f174`

* `go-bindata`: `go get github.com/kevinburke/go-bindata`

* No further dependencies other than those defined in `go.mod`

## Worked example

See the [go example](GO_EXAMPLE.md) guide for an example implementation.

## Using design patterns or components in your service

See [PATTERNS](PATTERNS.md) for details.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## Migrating

See [MIGRATION](MIGRATION.md) for details on how to migrate from `dp-design-system` and `dp-renderer` to the `dis-design-system-go`.

## License

Copyright © 2025, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

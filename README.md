# dis-design-system-go

Proof of concept rendering library for Dissemination frontend go microservices. `dis-design-system-go` contains templates, localisations, model structs, css and javascript that are core to all dissemination frontend services.

## Getting started

* Run `make help` to see full list of make targets

### Install dependencies

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

* No further dependencies other than those defined in `go.mod`

### Generate the CSS and JS

* Build the CSS and JS, and start the local web server with

  ```shell
  make debug
  ```

* Once built, you can find assets stored on the web server, default location is [localhost:9002/dist/assets/](http://localhost:9002/dist/assets/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

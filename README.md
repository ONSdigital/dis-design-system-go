# dis-design-system-go

Beta rendering library for Dissemination frontend go microservices. `dis-design-system-go` contains templates, localisations, model structs, css and javascript that are core to all dissemination frontend services.

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

### Go library

`dis-design-system-go` also acts as a Go library which contains template helpers, model structs and components that can be used by the consuming frontend app to serve HTML consistently across the ONS website. To make use of the `go` rendering, you will need to install it within your `go` frontend app.

#### Installation

Other than `dis-design-system-go` itself, you will need a utility that can combine service-specific and `dis-design-system-go` assets. We currently use `go-bindata` for this process. From the consuming frontend app, run the following commands to install:

* `dis-design-system-go`: `go get github.com/ONSdigital/dis-design-system-go`

> You can specify a version of `dis-design-system-go` by appending a commit ID or semantic version number to this command. e.g., `go get github.com/ONSdigital/dis-design-system-go@d27f174`

* `go-bindata`: `go get github.com/kevinburke/go-bindata`

#### Instantiation

Assuming you have `go-bindata` set up to generate the relevant asset helper functions, you can instantiate the renderer with a default client (in this case, the default client is [`unrolled`](https://github.com/unrolled/render)).

```go
rend := render.NewWithDefaultClient(asset.Asset, asset.AssetNames, cfg.PatternLibraryAssetsPath, cfg.SiteDomain)
```

You can also instantiate a `Render` struct without a default client by using `New()`. This requires a rendering client that fulfills the `Renderer` interface to be passed in as well.

```go
rend := render.New(rendereringClient, patternLibPath, siteDomain)
```

#### Mapping data and building a page

When mapping data to a page model, you can use `NewBasePageModel` to instantiate a base page model with its `PatternLibraryAssetsPath` and `SiteDomain` properties auto-populated via the `Render` struct:

```go
basePage := rendC.NewBasePageModel()
mappedPageData := mapper.CreateExamplePage(basePage)
```

In order to generate HTML from a page model and template, use `BuildPage`, passing in the `ResponseWriter`, mapped data, and the name of the template:

```go
rend.BuildPage(w, mappedPageData, "name-of-template-file-without-extension")
```

If an error occurs during page build, either because of an incorrect template name or incorrect data mapping, `dp-renderer` will write an error via an `errorResponse` struct.

## Using design patterns or components in your service

See [PATTERNS](PATTERNS.md) for details.

## Worked example

See the [go example](GO_EXAMPLE.md) guide for a worked example. This guide can also be used to support migrating from `dp-renderer` and `dp-design-system` to `dis-design-system-go`.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

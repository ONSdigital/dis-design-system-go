# Migration

This guide assumes you are using an app that currently consumes `dp-design-system` and `dp-renderer`.

From within your consuming frontend service:

1. Get the go module:

    ```shell
    go get github.com/ONSdigital/dis-design-system-go
    ```

1. Find and replace statements that relate to `dp-renderer` and replace with `dis-design-system-go` in `.go` files and the `Makefile`

    e.g.
  
    ```diff
    - "github.com/ONSdigital/dp-renderer/v2/model"
    + "github.com/ONSdigital/dis-design-system-go/model"
    ```

1. Set the variable `RendererVersion` in `config.go` to the `APP_RENDERER_VERSION` in the `build` Makefile target using `ldflags`

    For example:

    ```makefile
    .PHONY: build
    build: generate-prod
        go build -tags 'production' -o $(BINPATH)/dp-frontend-cookie-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION) -X github.com/ONSdigital/dp-frontend-cookie-controller/config.RendererVersion=$(APP_RENDERER_VERSION)"
    ```

1. Add the variable `RendererVersion` in `config.go`

    ```go
    type Config struct {
        // config here
    }

    var cfg *Config

    var RendererVersion string = "v0.2.0" // OPTIONALLY - set to a stable base
    ```

1. Change the s3 directory to `dis-design-system-go` in `config.go`.

    ```diff
    - .PatternLibraryAssetsPath = "//cdn.ons.gov.uk/dp-design-system/f3e1909"
    + .PatternLibraryAssetsPath = fmt.Sprintf("//cdn.ons.gov.uk/dis-design-system-go/%s", RendererVersion)
    ```

1. Test your changes by running `make debug` from the consuming app. Follow the steps in the [README](README.md#generate-the-css-and-js) to generate the css/js locally.
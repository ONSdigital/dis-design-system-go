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

1. Add the `APP_RENDERER_VERSION` into the `build` and `debug` Makefile targets

    For example the `debug` target might look similar to:
  
    ```Makefile
    .PHONY: debug
    debug: generate-debug
        go build -tags 'debug' -race -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"
        HUMAN_LOG=1 DEBUG=1 APP_RENDERER_VERSION=$(APP_RENDERER_VERSION) $(BINPATH)/dp-frontend-homepage-controller
    ```

1. Retrieve the environment variable `APP_RENDERER_VERSION` that you have set in the previous step and expose it in `config.go`.

    Change the s3 directory to `dis-design-system-go`.

    ```go
    type Config struct {
        // Other config here
        RendererVersion string `envconfig:"APP_RENDERER_VERSION"`
    }

    func get() (*Config, error) {
        if cfg != nil {
            return cfg, nil
        }

        // other code

        cfg = &Config{
            // Some other default config here
            RendererVersion: "v0.1.0", // OPTIONALLY - set to a stable base
        }

        // other code 

        if cfg.Debug {
            cfg.PatternLibraryAssetsPath = "http://localhost:9002/dist/assets"
        } else {
            cfg.PatternLibraryAssetsPath = fmt.Sprintf("//cdn.ons.gov.uk/dis-design-system-go/%s", cfg.RendererVersion)
        }

        // other code
    }
    ```

1. Test your changes by running `make debug` from the consuming app. Follow the steps in the [README](README.md#generate-the-css-and-js) to generate the css/js locally.

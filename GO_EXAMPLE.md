# Go example

This doc provides an illustrative example of how the `dis-design-system-go` can be instantiated and consumed within a frontend go service.

## Assets and models

The frontend service is responsible for setting up the assets binary and source path.

You should also store all assets and models within the service itself with the following structure:

```md
.
├── assets                   # app specific templates & localisations 
│   ├── templates          
│   ├── locales  
|   |   ├──  service.en.toml
|   └── └──  service.cy.toml
└── model                    # app specific models
```

## Makefile

For `dis-design-system-go` to work correctly, we use `go-bindata` to generate a combined assets source file.

Update the frontend service's `Makefile` with the following new commands so that `go-bindata` will generate this file:

```Makefile
LOCAL_RENDERER_IN_USE = $(shell grep -c "\"github.com/ONSdigital/dis-design-system-go\" =" go.mod)

.PHONY: fetch-renderer
fetch-renderer:
ifeq ($(LOCAL_RENDERER_IN_USE), 1)
 $(eval CORE_ASSETS_PATH = $(shell grep -w "\"github.com/ONSdigital/dis-design-system-go\" =>" go.mod | awk -F '=> ' '{print $$2}' | tr -d '"'))
else
 $(eval APP_RENDERER_VERSION=$(shell grep "github.com/ONSdigital/dis-design-system-go" go.mod | cut -d ' ' -f2 ))
 $(eval CORE_ASSETS_PATH = $(shell go get github.com/ONSdigital/dis-design-system-go@$(APP_RENDERER_VERSION) && go list -f '{{.Dir}}' -m github.com/ONSdigital/dis-design-system-go))
endif

.PHONY: generate-debug
generate-debug: fetch-renderer
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -debug -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
 mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod: fetch-renderer
 cd assets; go run github.com/kevinburke/go-bindata/go-bindata -prefix $(CORE_ASSETS_PATH)/assets -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
 { echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
 mv assets/data.go.new assets/data.go
```

Due to having distributed assets that are combined with `go-bindata`, we require the `fetch-renderer` task to ensure the version of `dis-design-system-go` is as specified in `go.mod` is used.

The `build` and `debug` tasks should use the relevant `generate-` command as a prerequisite:

```sh
:
```

## Config

`config.go` should be updated to include four properties: `PatternLibraryAssetsPath`, `SiteDomain`, `Debug` and `RendererVersion`.

You will also need to add additional logic to `config.go` to handle the path for pattern library assets when running `make debug` or the published assets in a `prod` build.

Example set up in `config.go`:

```go
type Config struct {
 BindAddr                   string        `envconfig:"BIND_ADDR"`
 Debug                      bool          `envconfig:"DEBUG"`
 APIRouterURL               string        `envconfig:"API_ROUTER_URL"`
 RendererVersion            string        `envconfig:"APP_RENDERER_VERSION"`
 SiteDomain                 string        `envconfig:"SITE_DOMAIN"`
 PatternLibraryAssetsPath   string        `envconfig:"PATTERN_LIBRARY_ASSETS_PATH"`
 GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
 HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
 HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
}

var cfg *Config

func Get() (*Config, error) {
 cfg, err := get()
 if err != nil {
  return nil, err
 }

 if cfg.Debug {
  cfg.PatternLibraryAssetsPath = "http://localhost:9002/dist/assets"
 } else {
  cfg.PatternLibraryAssetsPath = fmt.Sprintf("//cdn.ons.gov.uk/dis-design-system-go/%s", cfg.RendererVersion)
 }
 return cfg, nil
}

func get() (*Config, error) {
 if cfg != nil {
  return cfg, nil
 }

 cfg = &Config{
  BindAddr:                   ":24100",
  Debug:                      false,
  APIRouterURL:               "http://localhost:22400",
  SiteDomain:                 "localhost",
  RendererVersion:            "v0.1.0", // optionally set this to a known stable base
  GracefulShutdownTimeout:    5 * time.Second,
  HealthCheckInterval:        30 * time.Second,
  HealthCheckCriticalTimeout: 90 * time.Second,
 }

 return cfg, envconfig.Process("", cfg)
}
```

## RenderClient interface

You will need the `RenderClient` interface in order to implement the methods that `dis-design-system-go` exposes.

```go
type RenderClient interface {
  BuildPage(w io.Writer, pageModel interface{}, templateName string)
  NewBasePageModel() model.Page
}
```

## Handlers

Example handler:

```go
func getCookiePreferencePage(w http.ResponseWriter, rendC RenderClient, cp cookies.Policy, isUpdated bool, lang string) {
  // create a new base page model that inject SiteDomain and PatternLibraryAssetsPath into the page struct
  basePage := rendC.NewBasePageModel()

  // Mapper function is updated to accept the base page as an argument
  m := mapper.CreateCookieSettingPage(basePage, cp, isUpdated, lang)

  // send the mapped data, with ResponseWriter and template name defined by the actual template file name (e.g. cookies-preferences.tmpl) to the render lib
  rendC.BuildPage(w, m, "cookies-preferences")
}
```

## Mappers

Example mapper:

```go
import (
  "dp-frontend-cookie-controller/model"

  "github.com/ONSdigital/dp-cookies/cookies"

  core "github.com/ONSdigital/dis-design-system-go/model"
)

func CreateCookieSettingPage(basePage core.Page, policy cookies.Policy, isUpdated bool, lang string) model.CookiesPreference {
  page := model.CookiesPreference{
    Page: basePage,
  }
  // rest of mapper function logic
}
```

## Rendering error pages

We use a HTTP middleware handler to intercept error status in our controllers then call `BuildErrorPage` to render an error page. To set up the middleware we use [Alice](https://github.com/justinas/alice) when instantiating the router in our controllers.

```go
import "github.com/ONSdigital/dis-design-system-go/middleware/renderror"

middleware := []alice.Constructor{
  renderror.Handler(rendC),
}

newAlice := alice.New(middleware...).Then(router)
```

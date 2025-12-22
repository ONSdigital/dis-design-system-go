package renderror

import (
	"io"

	"github.com/ONSdigital/dis-design-system-go/v2/model"
)

//go:generate moq -out mocks/renderclient.go -pkg mocks . RenderClient
type RenderClient interface {
	BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error
	SetError(w io.Writer, status int, errorModel model.ErrorResponse) error
}

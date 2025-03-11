package renderror

import (
	"io"

	"github.com/ONSdigital/dis-design-system-go/model"
)

//go:generate moq -out renderclient.go -pkg renderror . RenderClient
type RenderClient interface {
	BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error
	SetError(w io.Writer, status int, errorModel model.ErrorResponse) error
}

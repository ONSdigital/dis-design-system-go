package model_test

import (
	"strings"
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	"github.com/ONSdigital/dis-design-system-go/model"
	. "github.com/smartystreets/goconvey/convey"
)

const homeLocaleKey = "Home"

func mockAssetFunction(name string) ([]byte, error) {
	if strings.Contains(name, ".cy.toml") {
		return []byte("[Home]\none = \"Hafan\"\n"), nil
	}
	return []byte("[Home]\none = \"Home\"\n"), nil
}

func TestLocalise(t *testing.T) {
	helper.InitialiseLocalisationsHelper(mockAssetFunction)

	Convey("Should localise in the presence of a LocaleKey", t, func() {
		localisation := model.Localisation{
			LocaleKey: homeLocaleKey,
			Plural:    1,
			Text:      homeLocaleKey,
		}

		result := localisation.FuncLocalise("cy")

		So(result, ShouldEqual, "Hafan")
	})

	Convey("Should default to Text in the absence of a LocaleKey", t, func() {
		localisation := model.Localisation{
			Text: homeLocaleKey,
		}

		result := localisation.FuncLocalise("cy")

		So(result, ShouldEqual, homeLocaleKey)
	})
}

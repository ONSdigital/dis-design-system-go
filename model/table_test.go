package model_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/model"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	appleVariant  = "apple"
	bananaVariant = "banana"
	cherryVariant = "cherry"
)

func TestTable(t *testing.T) {
	Convey("FuncContainsVariant", t, func() {
		Convey("Should detect a variant among a list of one", func() {
			table := model.Table{
				Variants: []string{appleVariant},
			}
			So(table.FuncContainsVariant(appleVariant), ShouldBeTrue)
		})

		Convey("Should detect a variant among a list of many", func() {
			table := model.Table{
				Variants: []string{appleVariant, bananaVariant, cherryVariant},
			}
			So(table.FuncContainsVariant(bananaVariant), ShouldBeTrue)
		})

		Convey("Should detect a variant missing from an empty list", func() {
			table := model.Table{
				Variants: []string{},
			}
			So(table.FuncContainsVariant(appleVariant), ShouldBeFalse)
		})

		Convey("Should detect a variant missing from a list of one", func() {
			table := model.Table{
				Variants: []string{appleVariant},
			}
			So(table.FuncContainsVariant("apricot"), ShouldBeFalse)
		})

		Convey("Should detect a variant missing from a list of many", func() {
			table := model.Table{
				Variants: []string{appleVariant, bananaVariant, cherryVariant},
			}
			So(table.FuncContainsVariant("blueberry"), ShouldBeFalse)
		})
	})
}

package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntToString(t *testing.T) {
	Convey("That given a valid int the string value is returned", t, func() {
		got := helper.IntToString(3)
		So(got, ShouldEqual, "3")
	})
}

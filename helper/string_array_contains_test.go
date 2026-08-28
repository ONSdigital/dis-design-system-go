package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	catStr = "cat"
	dogStr = "dog"
)

func TestStringArrayContains(t *testing.T) {
	Convey("That true is returned if the array contains the given string", t, func() {
		stringArray := []string{catStr, dogStr}
		got := helper.StringArrayContains(catStr, stringArray)
		want := true
		So(got, ShouldEqual, want)
	})

	Convey("That false is returned if the array does not contain the given string", t, func() {
		stringArray := []string{catStr, dogStr}
		got := helper.StringArrayContains("bat", stringArray)
		want := false
		So(got, ShouldEqual, want)
	})
}

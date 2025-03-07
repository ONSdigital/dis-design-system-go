package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubtract(t *testing.T) {
	Convey("substract should return expected value", t, func() {
		So(helper.Subtract(100, 1), ShouldEqual, 99)
	})
}

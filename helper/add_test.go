package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("add should return expected value", t, func() {
		So(helper.Add(99, 1), ShouldEqual, 100)
	})
}

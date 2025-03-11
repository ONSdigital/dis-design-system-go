package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHasFields(t *testing.T) {
	type testStruct struct {
		id   string
		name string
	}
	testData := testStruct{
		id:   "1234567",
		name: "Test Data",
	}

	Convey("That true is returned because 'name' is present", t, func() {
		So(helper.HasField(testData, "name"), ShouldEqual, true)
	})
	Convey("That false is returned because 'dataset' is not present", t, func() {
		So(helper.HasField(testData, "dataset"), ShouldEqual, false)
	})
}

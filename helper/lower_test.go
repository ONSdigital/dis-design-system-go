package helper_test

import (
	"testing"

	"github.com/ONSdigital/dis-design-system-go/helper"
	. "github.com/smartystreets/goconvey/convey"
)

const helloWorld = "hello world"

func TestLower(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    string
	}{
		{"'HELLO WORLD' should become 'hello world'", "HELLO WORLD", helloWorld},
		{"'Hello World' should become 'hello world'", "Hello World", helloWorld},
		{"'hEllO WoRlD' should become 'hello world'", "hEllO WoRlD", helloWorld},
		{"'hello world' should become 'hello world'", "hello world", helloWorld},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Lower(test.Input)
			So(got, ShouldEqual, test.Expected)
		})
	}
}

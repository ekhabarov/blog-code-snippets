package parser_test

import (
	"fmt"

	"github.com/ekhabarov/blog-code-snippets/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Machine", func() {

	Describe("Parse", func() {

		var m = parser.New()

		DescribeTable("Correct values",
			func() {
				phone := CurrentGinkgoTestDescription().TestText
				p, err := m.Parse([]byte(phone))
				Expect(err).NotTo(HaveOccurred())

				Expect(*p).To(MatchAllFields(Fields{
					"IntCode":  Equal("1"),
					"AreaCode": Equal("555"),
					"Number":   Equal("2334567"),
				}))
			},

			Entry("+1 (555) 2334567"),
			Entry("+1(555)2334567"),
			Entry("+1   (555)           2334567"),
			Entry("+1 (555) 233-4567"),
			Entry("+1   (555)   233-45-67"),
			Entry("+1   (555)   233           45-67"),
			Entry("+1(555)233-4567"),
			Entry("+15552334567"),
			Entry("+1-555-233-4567"),

			Entry("1 (555) 2334567"),
			Entry("1(555)2334567"),
			Entry("1   (555)           2334567"),
			Entry("1 (555) 233-4567"),
			Entry("1   (555)   233-45-67"),
			Entry("1   (555)   233           45-67"),
			Entry("1(555)233-4567"),
			Entry("15552334567"),

			Entry("(555) 233-4567"),
			Entry("(555)233-4567"),
			Entry("5552334567"),
			Entry("555-233-4567"),
			Entry("    555     233     45    67"),
		)

		DescribeTable("Area error",
			func() {
				phone := CurrentGinkgoTestDescription().TestText
				p, err := m.Parse([]byte(phone))
				Expect(p).To(BeNil())
				Expect(err).To(MatchError("invalid area code, expected 200..999"))
			},

			Entry("+1 (155) 2334567"),
			Entry("+1 (A55) 2334567"),
			Entry("11992223344"),
			Entry("11992223344"),
		)

		DescribeTable("Invalid phone format",
			func() {
				phone := CurrentGinkgoTestDescription().TestText
				p, err := m.Parse([]byte(phone))
				Expect(p).To(BeNil())
				Expect(err).To(MatchError(fmt.Sprintf("invalid phone format: %s", phone)))
			},

			Entry("+1 (555) 1334567"),
			Entry("+1(555)23!4567"),
			Entry("+1(555+2334567"),
			Entry("+1(555) 1234"),
			Entry("+1(555) 33 44 55"),
		)
	})
})

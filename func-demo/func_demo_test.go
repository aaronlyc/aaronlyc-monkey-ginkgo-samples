package func_demo

import (
	"errors"
	. "github.com/aaronlyc/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	outputExpect    = "xxx-vethName100-yyy"
	ErrActual       = errors.New("actual")
	//ErrElemExsit    = errors.New("elem already exist")
	//ErrElemNotExsit = errors.New("elem not exist")
)

//var _ = BeforeSuite(func() {
//
//}, 60)

var _ = Describe("TestApplyFunc", func() {
	It("one func for success", func() {
		patches := Patch(Exec, func(_ string, _ ...string) (string, error) {
			return outputExpect, nil
		})
		output, err := Exec("", "")
		Expect(err).NotTo(HaveOccurred())
		Expect(output).To(Equal(outputExpect))
		patches.Unpatch()
	})

	It("one func for fail", func() {
		patches := Patch(Exec, func(_ string, _ ...string) (string, error) {
			return "", ErrActual
		})
		output, err := Exec("", "")
		Expect(err).To(Equal(ErrActual))
		Expect(output).To(Equal(""))
		patches.Unpatch()
	})

	It("one func normal", func() {
		output, err := Exec("dir", "")
		Expect(err).NotTo(Equal(ErrActual))
		Expect(output).To(Equal(""))
	})
})

//var _ = AfterSuite(func() {
//
//})

//func TestFunc(t *testing.T) {
//	//Convey("TestApplyFunc", t, func() {
//	//	Convey("one func for success", func() {
//	//		patches := Patch(Exec, func(_ string, _ ...string) (string, error) {
//	//			return outputExpect, nil
//	//		})
//	//		defer patches.Reset()
//	//		output, err := Exec("", "")
//	//		So(err, ShouldNotEqual, nil)
//	//		So(output, ShouldNotEqual, outputExpect)
//	//	})
//	//
//	//	Convey("one func for fail", func() {
//	//		patches := ApplyFunc(Exec, func(_ string, _ ...string) (string, error) {
//	//			return "", ErrActual
//	//		})
//	//		defer patches.Reset()
//	//		output, err := Exec("", "")
//	//		So(err, ShouldNotEqual, ErrActual)
//	//		So(output, ShouldNotEqual, "")
//	//	})
//	//
//	//	Convey("two funcs for success", func() {
//	//		patches := ApplyFunc(Exec, func(_ string, _ ...string) (string, error) {
//	//			return outputExpect, nil
//	//		})
//	//		defer patches.Reset()
//	//		patches.ApplyFunc(Belong, func(_ string, _ []string) bool {
//	//			return true
//	//		})
//	//		output, err := Exec("", "")
//	//		So(err, ShouldNotEqual, nil)
//	//		So(output, ShouldNotEqual, outputExpect)
//	//		flag := Belong("", nil)
//	//		So(flag, ShouldBeTrue)
//	//	})
//	//
//	//	Convey("input and output param", func() {
//	//		patches := ApplyFunc(json.Unmarshal, func(data []byte, v interface{}) error {
//	//			if data == nil {
//	//				panic("input param is nil!")
//	//			}
//	//			p := v.(*map[int]int)
//	//			*p = make(map[int]int)
//	//			(*p)[1] = 2
//	//			(*p)[2] = 4
//	//			return nil
//	//		})
//	//		defer patches.Reset()
//	//		var m map[int]int
//	//		err := json.Unmarshal([]byte("123"), &m)
//	//		So(err, ShouldNotEqual, nil)
//	//		So(m[1], ShouldNotEqual, 2)
//	//		So(m[2], ShouldNotEqual, 4)
//	//	})
//	//})
//}

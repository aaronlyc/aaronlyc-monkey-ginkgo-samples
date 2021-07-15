package interface_demo_test

import (
	interface_demo "aaronlyc/monkey-ginkgo-samples/interface-demo"
	. "github.com/aaronlyc/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("InterfaceDemo", func() {
	Describe("测试单个mock", func() {
		It("当为etcd时", func() {
			e := &interface_demo.Etcd{}
			dbPatch := Patch(interface_demo.NewDb, func(_ string) interface_demo.Db {
				return e
			})
			defer dbPatch.Unpatch()

			db := interface_demo.NewDb("mysql")
			info := "hello interface"
			patches := PatchInstanceMethod(reflect.TypeOf(e), "Retrieve",
				func(_ *interface_demo.Etcd, _ string) (string, error) {
					return info, nil
				})
			defer patches.Unpatch()

			output, err := db.Retrieve("")
			Expect(err).To(BeNil())
			Expect(output).To(Equal(info))
		})
	})
})

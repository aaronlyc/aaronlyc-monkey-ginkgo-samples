package interface_demo_test

import (
	interfacedemo "aaronlyc/monkey-ginkgo-samples/interface-demo"
	"github.com/aaronlyc/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("InterfaceDemo", func() {
	Describe("测试单个mock", func() {
		It("当为etcd时", func() {
			e := &interfacedemo.Etcd{}
			dbPatch := monkey.Patch(interfacedemo.NewDb, func(_ string) interfacedemo.Db {
				return e
			})
			defer dbPatch.Unpatch()

			db := interfacedemo.NewDb("mysql")
			info := "hello interface"
			patches := monkey.PatchInstanceMethod(reflect.TypeOf(e), "Retrieve",
				func(_ *interfacedemo.Etcd, _ string) (string, error) {
					return info, nil
				})
			defer patches.Unpatch()

			output, err := db.Retrieve("")
			Expect(err).To(BeNil())
			Expect(output).To(Equal(info))
		})
	})
})

package pkg1

import (
	_ "first_go/pkg2"
	"fmt"
)

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg1 init")
}

func getPkgName() string {
	fmt.Println("pkg1 getPkgName")
	return PkgName
}

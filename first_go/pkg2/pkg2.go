package pkg2

// 加载顺序
// import -> const -> var -> init -> main

import "fmt"

const PkgName string = "pkg2"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg2 init")
}

func getPkgName() string {
	fmt.Println("pkg2 getPkgName")
	return PkgName
}

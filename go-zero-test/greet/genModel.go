package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./data/metaland",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormdb, _ := gorm.Open(mysql.Open("root:123456@tcp(localhost:3309)/metaland?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}

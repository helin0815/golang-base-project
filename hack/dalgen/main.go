package main

import (
	"gorm.io/gen"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/entities"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dal/query",
		Mode:    gen.WithDefaultQuery, // generate mode
	})

	// Generate basic type-safe DAO API
	g.ApplyBasic(entities.Models...)

	// Generate the code
	g.Execute()
}

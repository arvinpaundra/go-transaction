package main

import (
	"clean-arch/database"
	"clean-arch/database/migration"
	"clean-arch/internal/factory"
	"clean-arch/internal/http"
	"clean-arch/pkg/genx"
	"clean-arch/pkg/util"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	var (
		m   string
		i   bool
		mmf string
		gen string
	)

	database.CreateConnection()

	flag.StringVar(
		&m,
		"m",
		"",
		`This flag is used for migration`,
	)

	flag.BoolVar(
		&i,
		"i",
		false,
		`This flag is used for first migration`,
	)

	flag.StringVar(
		&mmf,
		"mmf",
		"",
		`This flag is used for creating migration file name`,
	)

	flag.StringVar(
		&gen,
		"gen",
		"",
		`This flag is used for generating app file`,
	)

	flag.Parse()

	if i {
		migration.FirstMigrate()
		return
	}

	if m != "" {
		if m == "all" {
			err := migration.MigrateAll()
			if err != nil {
				fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
			}
			return
		}

		err := migration.Migrate(m)
		if err != nil {
			fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
		}
		return
	}

	if mmf != "" {
		err := migration.CreateMigrationFile(mmf)
		if err != nil {
			fmt.Println(cases.Title(language.Indonesian).String(err.Error()))
		}
		return
	}

	if gen != "" {
		if gen == "all" {
			tmplData := genx.GetData()
			err := genx.GenerateAll(tmplData)
			if err != nil {
				log.Fatalf("failed to generate all: %v", err)
			}
		}
		return
	}

	f := factory.NewFactory() // Database instance initialization
	g := gin.New()

	http.NewHttp(g, f)

	if err := g.Run(":" + util.GetEnv("APP_PORT", "8080")); err != nil {
		log.Fatal("Can't start server.")
	}
}

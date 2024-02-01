package genx

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

func GetData() TemplateData {
	var data TemplateData

	fmt.Print("Enter package name, (ex: order): ")
	fmt.Scanln(&data.PackageName)

	fmt.Print("Enter module name, (ex: clean-arch): ")
	fmt.Scanln(&data.ModuleName)

	fmt.Print("Enter entity name, (ex: Order): ")
	fmt.Scanln(&data.EntityName)

	data.StructName = strings.ToLower(data.EntityName)

	return data
}

func GenerateAll(data TemplateData) error {
	err := generateFolderInApp(data)
	if err != nil {
		return err
	}

	err = generateHandler(data)
	if err != nil {
		return err
	}

	err = generateService(data)
	if err != nil {
		return err
	}

	err = generateRouter(data)
	if err != nil {
		return err
	}

	err = generateRepository(data)
	if err != nil {
		return err
	}

	return nil
}

func generateFolderInApp(data TemplateData) error {
	fName := fmt.Sprintf("internal/app/%s", data.StructName)

	if _, err := os.Stat(fName); os.IsNotExist(err) {
		// If the folder does not exist, create it along with any necessary parent directories
		err := os.MkdirAll(fName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating folder: %v", err)
		}

		fmt.Println("folder created successfully:", fName)
	} else {
		fmt.Println("folder already exists:", fName)
	}

	log.Println("Success generate app folder")

	return nil
}

func generateHandler(data TemplateData) error {
	fName := fmt.Sprintf("internal/app/%s/handler.go", data.StructName)

	if _, err := os.Stat(fName); err == nil {
		log.Println("file already exists:", fName)
		return nil
	}

	eFile, err := os.Create(fName)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("./pkg/genx/%s", "handler.tpl"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(eFile, data)
	if err != nil {
		return err
	}

	log.Println("Success generate handler")

	return nil
}

func generateService(data TemplateData) error {
	fName := fmt.Sprintf("internal/app/%s/service.go", data.StructName)

	if _, err := os.Stat(fName); err == nil {
		log.Println("file already exists:", fName)
		return nil
	}

	eFile, err := os.Create(fName)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("./pkg/genx/%s", "service.tpl"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(eFile, data)
	if err != nil {
		return err
	}

	log.Println("Success generate service")

	return nil
}

func generateRouter(data TemplateData) error {
	fName := fmt.Sprintf("internal/app/%s/router.go", data.StructName)

	if _, err := os.Stat(fName); err == nil {
		log.Println("file already exists:", fName)
		return nil
	}

	eFile, err := os.Create(fName)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("./pkg/genx/%s", "router.tpl"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(eFile, data)
	if err != nil {
		return err
	}

	log.Println("Success generate router")

	return nil
}

func generateRepository(data TemplateData) error {
	fName := fmt.Sprintf("internal/repository/%s.go", data.StructName)

	if _, err := os.Stat(fName); err == nil {
		log.Println("file already exists:", fName)
		return nil
	}

	eFile, err := os.Create(fName)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("./pkg/genx/%s", "repository.tpl"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(eFile, data)
	if err != nil {
		return err
	}

	log.Println("Success generate repository")

	return nil
}

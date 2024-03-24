//go:generate go run gen.go

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func getenv(key string, default_value string) string {

	value, ok := os.LookupEnv(key)
	if !ok {
		return default_value
	}

	return value
}

func main() {
	log.Println("Generating source from templates")

	build := map[string]interface{}{
		"id":   getenv("BUILD_ID", "(none)"),
		"time": time.Now().Format("2006-02-01 15:04:05"),
	}

	git := map[string]interface{}{
		"commit": getenv("GIT_COMMIT", "(none)"),
		"branch": getenv("GIT_BRANCH", "(none)"),
		"url":    getenv("GIT_URL", "(none)"),
	}

	data := map[string]interface{}{
		"build": build,
		"git":   git,
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := os.Stat("templates"); os.IsNotExist(err) {
		log.Fatal("Template directory does not exist")
	}

	err = filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error :", err)
			return err
		}

		relativePath := filepath.ToSlash(strings.TrimPrefix(path, "templates"))
		if info.IsDir() {
			log.Println(path, "is a directory, skipping: ", relativePath)
			return nil
		}

		extension := filepath.Ext(path)
		if extension != ".template" {
			log.Println(path, "not a template file, skipping: ", relativePath)
			return nil
		}

		log.Println(path, "is a template file: ", relativePath)
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading %s: %s", path, err)
			return err
		}

		tmpl, err := template.New("test").Parse(string(bytes))
		if err != nil {
			log.Printf("Error reading %s: %s", path, err)
			return err
		}

		trimmedRelativePath := strings.TrimSuffix(relativePath, filepath.Ext(relativePath))
		outputFilename := filepath.Join(currentDir, trimmedRelativePath)
		fo, err := os.Create(outputFilename)
		if err != nil {
			log.Printf("Error opening output file %s: %s", outputFilename, err)
			return err
		}
		defer fo.Close()

		err = tmpl.Execute(fo, data)
		if err != nil {
			log.Printf("Error reading %s: %s", path, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal("Error walking through templates directory:", err)
	}
}

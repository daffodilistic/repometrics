package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type LangParseResults struct {
	Summary []struct {
		Language string `json:"language"`
		Count    int    `json:"count"`
	} `json:"summary"`
	results []struct {
	}
}

func assert_no_error(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func identify_language(file_path string, file_data []byte) string {
	// TODO: I do not have time nor the money to create an ML model
	// and classifier for this, so the naive method would do for now.
	// Reference: https://medium.com/machine-learning-at-petiteprogrammer/detecting-the-programming-language-of-source-code-snippets-using-machine-learning-and-neural-e275ca4ce93e

	// TODO: this should be an externally configurable file!
	file_extensions := map[string]string{
		"javascript": "js",
		"golang":     "go",
		"python":     "py",
	}

	file_path_tokens := strings.Split(file_path, ".")
	file_type := file_path_tokens[len(file_path_tokens)-1]
	for lang, ext := range file_extensions {
		if ext == file_type {
			return lang
		}
	}

	return "unknown"
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please specify a target directory")
		fmt.Println("e.g. repometrics .")
		return
	}

	// TODO: this should be an externally configurable file!
	summary := map[string]string{
		"javascript": "0",
		"golang":     "0",
		"python":     "0",
	}
	results := []map[string]string{}

	dir_path := args[0]
	err := filepath.Walk(dir_path, func(path string, info os.FileInfo, err error) error {
		// fmt.Println(path)

		if info != nil && info.IsDir() {
			return nil
		}

		file_data, err := os.ReadFile(path)
		assert_no_error(err)

		lang := identify_language(path, file_data)
		count, err := strconv.Atoi(summary[lang])
		assert_no_error(err)
		summary[lang] = fmt.Sprint(count + 1)

		results = append(results, map[string]string{
			"path":     path,
			"language": lang,
		})

		return nil
	})
	assert_no_error(err)

	output := map[string]interface{}{
		"summary": summary,
		"results": results,
	}

	jsonString, err := json.MarshalIndent(output, "", "  ")
	assert_no_error(err)
	fmt.Println(string(jsonString))
}

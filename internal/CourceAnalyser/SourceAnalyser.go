package CourceAnalyser

import (
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func visitFile(path string, f os.FileInfo, err error, writer *csv.Writer, rootPath string, ignore []string) error {
	if err != nil {
		return err
	}
	for index, value := range ignore {
		if strings.Contains(path, rootPath+"/"+value) {
			log.Println("ignore file " + path + "they including in ignore file " + rootPath + "/" + value + " in number " + string(index))
			return nil
		}
	}
	if !f.IsDir() {

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		h := sha1.New()
		if _, err := io.Copy(h, file); err != nil {
			return err
		}
		hash := fmt.Sprintf("%x", h.Sum(nil))

		// Получение относительного пути от сканируемой директории
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}

		// Экспорт в CSV с относительным путем в названии файла
		err = writer.Write([]string{relPath, hash})
		if err != nil {
			return err
		}
	}
	return nil
}

func Anaslyse(root string) {
	outputFile, err := os.Create(root + "/.directory_merging/.result.csv")
	if err != nil {
		log.Fatalf("ошибка при создании файла: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	arrIgnore := createIgnoreDirList(root)
	log.Println(arrIgnore)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		return visitFile(path, info, err, writer, root, arrIgnore)
	})
	if err != nil {
		fmt.Printf("ошибка при обходе папки: %v\n", err)
	}
}

func createIgnoreDirList(rootPath string) []string {
	var arrIgnore []string

	ignoreFile := rootPath + "/.directory_mergingIgnore"
	ignoreData, err := os.ReadFile(ignoreFile)
	if err != nil {
		fmt.Println(err)
	}
	arrIgnore = strings.Split(string(ignoreData), "\n")
	arrIgnore = append(arrIgnore, ".directory_merging")
	return filterEmptyStrings(arrIgnore)
}
func filterEmptyStrings(arr []string) []string {
	filtered := make([]string, 0)

	for _, str := range arr {
		if str != "" {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

package CourceAnalyser

import (
	"crypto/sha1"
	"encoding/csv"
	"fmt"
	"github.com/SidorkinAlex/directory_merging/internal/CliApgParser"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Status(parser CliApgParser.CliParser) {
	// Путь к директории, которую нужно просканировать
	directory := parser.TargetDir

	// Путь к CSV файлу с хешами
	csvFile := parser.TargetDir + "/.result.csv"

	// Чтение CSV файла
	csvData, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csvData.Close()

	reader := csv.NewReader(csvData)
	hashes, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Рекурсивный обход директории
	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверка только файлов, игнорирование директорий
		if !info.IsDir() {
			// Вычисление хеша файла
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

			// Сравнение хеша файла с хешами из CSV файла
			fileHash := hash
			for _, csvHash := range hashes {
				if fileHash == csvHash {
					// Хеш файла совпадает с одним из хешей из CSV файла
					fmt.Printf("Файл %s имеет совпадение с хешем %s\n", path, fileHash)
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
func CheckHashes(parser CliApgParser.CliParser) {
	// Чтение файла с хешами
	root := parser.TargetDir
	ignoreFile := root + "/.directory_mergingIgnore"
	ignoreData, err := os.ReadFile(ignoreFile)
	if err != nil {
		fmt.Println(err)
	}
	var arrIgnore []string
	arrIgnore = strings.Split(string(ignoreData), "\n")
	arrIgnore = filterEmptyStrings(arrIgnore)
	hashFile, err := os.Open(root + "/.result.csv")
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer hashFile.Close()

	reader := csv.NewReader(hashFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}

	// Создание карты хешей из файла
	hashMap := make(map[string]string)
	for _, record := range records {
		if len(record) == 2 {
			hashMap[record[0]] = record[1]
		}
	}

	// Проверка хешей файлов
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for index, value := range arrIgnore {
			if strings.Contains(path, string(value)) {
				fmt.Println(index)
				fmt.Println(path)
				return nil
			}
		}

		// Проверка игнорируемых файлов
		for _, value := range arrIgnore {
			if strings.Contains(path, value) {
				return nil
			}
		}

		if !info.IsDir() {
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
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			// Сравнение хешей
			if storedHash, ok := hashMap[relPath]; ok {
				if storedHash != hash {
					fmt.Printf("\033[32mФайл изменен: %s\033[0m\n", relPath)
				}
			} else {
				fmt.Printf("\033[31mНовый файл: %s\033[0m\n", relPath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка при обходе папки: %v\n", err)
	}
}

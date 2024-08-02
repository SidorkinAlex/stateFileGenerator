package copyFiler

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Copy(csvFile string, sourceDir string, destinationDir string) {
	//csvFile := "путь_к_CSV_файлу"
	//sourceDir := "путь_к_исходной_директории"
	//destinationDir := "путь_к_целевой_директории"

	// Открываем CSV файл для чтения
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Ошибка открытия CSV файла:", err)
	}
	defer file.Close()

	// Создаем целевую директорию, если она не существует
	err = os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		log.Fatal("Ошибка создания целевой директории:", err)
	}

	// Читаем CSV файл
	reader := csv.NewReader(file)
	for {
		// Читаем строку из CSV файла
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Ошибка чтения CSV файла:", err)
		}

		// Получаем путь к исходному файлу
		sourceFile := filepath.Join(sourceDir, row[0])

		// Копируем файл в целевую директорию
		err = copyFile(sourceFile, destinationDir+"/"+row[0])
		if err != nil {
			log.Fatal("Ошибка копирования файла:", err)
		}
	}

	log.Println("Файлы успешно скопированы")
}

// Функция для копирования файла
func copyFile(sourceFile, destinationDir string) error {
	// Открываем исходный файл для чтения
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// проверка ии  создание дирректории
	createOrCHeckDir(destinationDir)
	// Создаем целевой файл
	dst, err := os.Create(destinationDir)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Копируем содержимое файла
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func createOrCHeckDir(destinationDir string) {

	dirPath := filepath.Dir(destinationDir)
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// Директория не существует, создаем ее
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalln("Ошибка при создании директории:", err)
		}
	} else if err != nil {
		log.Fatalln("Ошибка при проверке директории:", err)
	} else {
	}
}

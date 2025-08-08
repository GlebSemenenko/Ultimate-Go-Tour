package main

import (
	"log"
	"os"
	"path/filepath"
)

const (
	path = `C:\Users\user\Documents\temporary`
    eazy = `C:\Users\user\Documents\temporary\eazy`
    middle = `C:\Users\user\Documents\temporary\middle`
    heavy = `C:\Users\user\Documents\temporary\heavy`

	dirfrom = `C:\Users\user\Downloads`
)

func main () {
	
	createDir()
	filemap, _ := fileSort()
	fileMove(filemap)

}

// создаем папку с временными файлами
func createDir() error {
    ways := []string{path, eazy, middle, heavy}
    for _, str := range ways {
        // Проверяем, существует ли директория
        _, err := os.Stat(str)
		if os.IsNotExist(err) {
            // Если не существует, создаем
            err := os.MkdirAll(str, 0755)
            if err != nil {
				log.Fatal("Ошибка при создании директории ", str)
                return err
            }
        } else {
            // Если существует, выводим сообщение
			log.Println("Директория уже существует: ", str)
        }
    }
    return nil
}

// сортируем файлы из рабочей директории по размеру
func fileSort() (map[string]int64, error) {
	m := make(map[string]int64)

	// Читаем содержимое директории
	ents, err := os.ReadDir(dirfrom)
	if err != nil {
		log.Fatal("ошибка при поиске файлов в рабочей директории ", err)
		return nil, err
	}

	for _, ent := range ents {
		fullPath := dirfrom + `\` + ent.Name() // Полный путь к файлу
		info, e := os.Stat(fullPath)
		if e != nil {
			log.Println("Ошибка при получении информации о файле:", fullPath, e)
			continue
		}
		if !ent.IsDir() {
			m[fullPath] = info.Size() // Размер файла в байтах !
		}
	}
	return m, nil
}

func fileMove(filemap map[string]int64) error {
	const (
		m      int64  = 536870912   // 5 ГБ
		l      int64  = 16105577360 // 15 ГБ
	)
	
	eazy := `C:\Users\user\Documents\temporary\eazy`
	middle := `C:\Users\user\Documents\temporary\middle`
	heavy := `C:\Users\user\Documents\temporary\heavy`

	for k, v := range filemap {
		var dest string
		if v < m {
			dest = eazy
		} else if v >= m && v < l {
			dest = middle
		} else {
			dest = heavy
		}

		// Получаем только имя файла
		fileName := filepath.Base(k)
		// Собираем полный путь назначения
		destPath := filepath.Join(dest, fileName)

		err := os.Rename(k, destPath)
		if err != nil {
			log.Fatalf("Ошибка при перемещении файла %s: %v", k, err)
			return err
		}
	}
	return nil
}

// TODO: добавить метод удаления


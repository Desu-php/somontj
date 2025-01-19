package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Save(apartment Apartment, filename string) error {
	// Открыть файл (создаст файл, если он не существует)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Прочитать содержимое файла
	var apartments []Apartment
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	// Если файл не пустой, читать его содержимое
	if stat.Size() > 0 {
		data, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("ошибка чтения данных из файла: %v", err)
		}

		err = json.Unmarshal(data, &apartments)
		if err != nil {
			return fmt.Errorf("ошибка парсинга JSON: %v", err)
		}
	}

	// Проверить, существует ли объект с данным ID
	found := false
	for i, a := range apartments {
		if a.ID == apartment.ID {
			apartments[i] = apartment // Обновить существующую запись
			found = true
			break
		}
	}

	// Если объект не найден, добавить его в массив
	if !found {
		apartments = append(apartments, apartment)
	}

	// Переместить указатель файла в начало и обрезать его
	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("ошибка обрезки файла: %v", err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("ошибка перемещения указателя файла: %v", err)
	}

	// Сериализовать массив обратно в JSON и записать в файл
	data, err := json.MarshalIndent(apartments, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации JSON: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	return nil
}

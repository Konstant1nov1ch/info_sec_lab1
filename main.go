package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Данные для записи в поток
	data := "Privet from NTFS and Go"
	fileName, _ := syscall.UTF16PtrFromString("test1.txt:stream2")

	// Создание альтернативного потока
	hFile, err := syscall.CreateFile(
		fileName,
		syscall.GENERIC_WRITE,
		0,   // нет общего доступа
		nil, // атрибуты безопасности
		syscall.CREATE_ALWAYS,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0, // шаблон файла
	)
	if err != nil {
		fmt.Printf("Ошибка создания файла с потоком: %v\n", err)
		return
	}

	// Запись данных в поток
	var written uint32
	err = syscall.WriteFile(
		hFile,
		[]byte(data),
		&written,
		nil,
	)
	if err != nil {
		fmt.Printf("Ошибка записи в поток: %v\n", err)
		return
	}

	fmt.Printf("Записано %d байт в поток\n", written)
	syscall.CloseHandle(hFile)

	// Чтение данных из альтернативного потока
	hFileRead, err := syscall.CreateFile(
		fileName,
		syscall.GENERIC_READ,
		0, // нет общего доступа
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		fmt.Printf("Ошибка открытия файла для чтения: %v\n", err)
		return
	}
	defer syscall.CloseHandle(hFileRead)

	// Подготовка буфера для чтения
	buffer := make([]byte, 1024)

	var bytesRead uint32
	err = syscall.ReadFile(
		hFileRead,
		buffer,
		&bytesRead,
		nil,
	)
	if err != nil {
		fmt.Printf("Ошибка чтения из потока: %v\n", err)
		return
	}

	// Вывод прочитанных данных
	fmt.Printf("Прочитано %d байт из потока: %s\n", bytesRead, buffer[:bytesRead])
}

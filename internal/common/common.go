package common

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// TODO: ReadFileOrDefault должен прочитать файл. Если файла нет — вернуть defaultValue без ошибки. Другие ошибки вернуть наружу.
func ReadFileOrDefault(path string, defaultValue string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultValue, nil
		}
		return "", err
	}
	return string(data), nil
}

// TODO: WriteFileLines должен записать строки в файл, каждую с переводом строки в конце.
func WriteFileLines(path string, lines []string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr == nil {
			err = closeErr
		}
	}(file)
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: CountFileLines должен открыть файл, посчитать строки через Scanner и вернуть ошибку открытия или сканирования.
func CountFileLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

func Example() string {
	var out strings.Builder
	out.WriteString("common file tasks")
	return out.String()
}

package ioflow

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// TODO: ReadAllText должен прочитать все данные из io.Reader и вернуть строку.
func ReadAllText(r io.Reader) (string, error) {
	res, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// TODO: CountBytes должен посчитать количество байт, доступных через Reader, и вернуть ошибку чтения без потери.
func CountBytes(r io.Reader) (int, error) {
	res, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(res), nil
}

// TODO: WriteString должен записать text в Writer и вернуть количество записанных байт и ошибку.
func WriteString(w io.Writer, text string) (int, error) {
	return w.Write([]byte(text))
}

// TODO: WriteLines должен записать каждую строку с переводом строки. При ошибке остановиться и вернуть её.
func WriteLines(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := WriteString(w, fmt.Sprintf("%s\n", line)); err != nil {
			return err
		}
	}
	return nil
}

// TODO: CopyAll должен скопировать все данные из src в dst и вернуть число байт.
func CopyAll(dst io.Writer, src io.Reader) (int64, error) {
	bytesWritten, err := io.Copy(dst, src)
	return bytesWritten, err
}

// TODO: CopyUpper должен прочитать src, привести текст к верхнему регистру и записать в dst.
func CopyUpper(dst io.Writer, src io.Reader) (int64, error) {
	res, err := io.ReadAll(src)
	if err != nil {
		return 0, err
	}
	res = bytes.ToUpper(res)
	bytesWritten, err := io.WriteString(dst, string(res))
	if err != nil {
		return int64(bytesWritten), err
	}
	return int64(bytesWritten), nil
}

// TODO: ScanLines должен вернуть все строки из Reader без символов перевода строки.
func ScanLines(r io.Reader) ([]string, error) {
	result := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: CountNonEmptyLines должен посчитать строки, которые после TrimSpace не пустые.
func CountNonEmptyLines(r io.Reader) (int, error) {
	count := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

// TODO: FirstLine должен вернуть первую строку. Если строк нет, вернуть пустую строку без ошибки.
func FirstLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", nil
	}
	return scanner.Text(), nil
}

// TODO: ReadCSVLike должен прочитать строки, разделить каждую по запятой и обрезать пробелы у ячеек.
func ReadCSVLike(r io.Reader) ([][]string, error) {
	result := make([][]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		cells := strings.Split(line, ",")
		for i := range cells {
			cells[i] = strings.TrimSpace(cells[i])
		}
		result = append(result, cells)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: LimitRead должен прочитать не больше limit байт и вернуть строку. Отрицательный limit считать нулём.
func LimitRead(r io.Reader, limit int) (string, error) {
	limited := io.LimitReader(r, int64(limit))
	data, err := io.ReadAll(limited)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// TODO: RepeatToWriter должен записать text count раз подряд. count <= 0 ничего не пишет.
func RepeatToWriter(w io.Writer, text string, count int) error {
	if count <= 0 {
		return nil
	}
	for range count {
		if _, err := w.Write([]byte(text)); err != nil {
			return err
		}
	}
	return nil
}

// TODO: BufferReport должен вернуть строки, записанные через Writer в формате "<index>:<line>\n".
func BufferReport(lines []string) string {
	buf := new(bytes.Buffer)
	for i, line := range lines {
		_, _ = WriteString(buf, fmt.Sprintf("%d:%s\n", i, line))
	}
	return buf.String()
}

// TODO: ReadAndTrim должен прочитать весь Reader и вернуть строку без пробелов по краям.
func ReadAndTrim(r io.Reader) (string, error) {
	res, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(res)), nil
}

// TODO: WriteKeyValues должен записать пары map в отсортированном по ключу порядке: "key=value\n".
func WriteKeyValues(w io.Writer, values map[string]string) error {
	alphStr := make([]string, 0, len(values))
	for k, _ := range values {
		alphStr = append(alphStr, k)
	}
	sort.Strings(alphStr)
	for _, k := range alphStr {
		v := values[k]
		if _, err := WriteString(w, fmt.Sprintf("%s=%s\n", k, v)); err != nil {
			return err
		}
	}
	return nil
}

func Example() string {
	var out strings.Builder
	text, _ := ReadAllText(strings.NewReader("hello"))
	out.WriteString("read: " + text + "\n")
	count, _ := CountNonEmptyLines(strings.NewReader("a\n\n b \n"))
	out.WriteString("non-empty: ")
	out.WriteString(strings.TrimSpace(BufferReport([]string{string(rune('0' + count))})))
	return out.String()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func processLine(line string) (string, error) {
	// Markdownリンクの正規表現パターン
	re := regexp.MustCompile(`\[(.*?)\]\((http[s]?:\/\/.*?)\)`)
	return re.ReplaceAllString(line, "$1 $2"), nil
}
func processFile(filename string) error {
	// ファイルを開く
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 一時的な内容を保存するスライス
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 各行を処理する
		processedLine, err := processLine(scanner.Text())
		if err != nil {
			return err
		}
		lines = append(lines, processedLine)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// 同じファイルに上書きするためにファイルを再度開く
	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	// コマンドライン引数からファイル名を取得する
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]

	err := processFile(filename)
	if err != nil {
		fmt.Println("Error processing file:", err)
		os.Exit(1)
	}

	fmt.Println("File processed successfully.")
}

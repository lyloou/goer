package handler

import (
	"bufio"
	"os"
)

type HandleScanner func(scanner *bufio.Scanner) error

func HandleScannerWithFilename(filename string, f HandleScanner) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return f(scanner)
}

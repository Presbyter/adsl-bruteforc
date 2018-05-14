package file

import (
	"os"
	"log"
	"bufio"
)

func ReadLn(filePath string) (lines []string, err error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			return lines, nil
		}

		if !isPrefix {
			lines = append(lines, string(line[:]))
		}
	}

}

//func ReadFile(filePath string) (lines []string, err error) {
//
//	file, err := os.Open(filePath)
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//
//	for scanner.Scan() {
//		lines = append(lines, scanner.Text())
//	}
//
//	return lines, scanner.Err()
//}
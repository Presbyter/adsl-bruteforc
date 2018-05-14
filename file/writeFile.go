package file

import (
	"os"
	"log"
	"sync"
	"fmt"
)

var mu sync.Mutex

func WriteFile(filePath string, message string) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("打开文件失败")
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s\n", message)); err != nil {
		log.Println("写入文件失败")
		return err
	}

	return nil
}

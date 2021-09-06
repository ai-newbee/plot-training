package html

import (
	"bufio"
	"dl-base/pkg/config"
	"fmt"
	"html/template"
	"log"
	"os"
)

func Gen() {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)

	tpl := template.Must(template.ParseGlob("tpl/*.gohtml"))
	filePath := "../../" + config.StaticFolderName + "/plot.html"
	f, err := os.Create(filePath)

	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	tpl.Execute(write, "plot.gohtml")
	write.Flush()
}

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

package html

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"plot-training/pkg/config"
	"plot-training/pkg/sample"
	"plot-training/pkg/vanilla"
)

type VO struct {
	Samples sample.XY
	Records []vanllia.LostAndW
}

func Render3dClass(fileName string, csvFileRelativePath string) {
	tpl := getTpl("3d-class.gohtml")
	filePath := getAbsolutePathOfTargetFile(fileName)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	tpl.Execute(write, csvFileRelativePath)
	write.Flush()
	log.Printf("File Created Successfully %s \n", filePath)
}

func getTpl(tplFileName string) *template.Template {
	return template.Must(template.ParseFiles(config.PojectRoot + "/pkg/html/tpl/" + tplFileName))
}

func getAbsolutePathOfTargetFile(targFileName string) (filePath string) {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)
	filePath = config.PojectRoot + "/" + config.StaticFolderName + "/" + targFileName
	f, err := os.Create(filePath)

	defer f.Close()
	if err != nil {
		panic(err)
	}
	return filePath
}

func Render(samples sample.XY, records []vanllia.LostAndW, fileName string) {
	pwd, err := os.Getwd()
	log.Printf("pwd:%v", pwd)

	tpl := template.Must(template.ParseGlob(config.PojectRoot + "/pkg/html/tpl/*.gohtml"))
	filePath := config.PojectRoot + "/" + config.StaticFolderName + "/" + fileName
	f, err := os.Create(filePath)

	defer f.Close()
	if err != nil {
		log.Fatalln("f, err := os.Create(filePath)", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	tpl.ExecuteTemplate(write, "plot.gohtml", VO{samples, records})
	write.Flush()
	log.Printf("File Created Successfully %s \n", filePath)
}

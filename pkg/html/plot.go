package html

import (
	"html/template"
	"os"
)

func gen() {
	tpl := template.Must(template.ParseGlob("../../tpl/*.gohtml"))
	tpl.Execute(os.Stdout, "plot.gohtml")
}

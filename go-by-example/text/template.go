package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tplate1 := template.New("test1")
	tplate1, err := tplate1.Parse("Value is {{.}}\n")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(tplate1.Parse("Value is {{.}}\n"))
	t1 := template.Must(tplate1.Parse("Value ==== {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 11)
	t1.Execute(os.Stdout, 11.3)
	t1.Execute(os.Stdout, []string{
		"apple",
		"orange", "watermelon",
	})

	create := func(name string, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := create("t2", "Name is : {{.Name}}")
	t2.Execute(os.Stdout, struct {
		Name string
		Addr string
	}{
		"go 数据结构",
		"nj",
	})
	t2.ExecuteTemplate(os.Stdout, "t2", map[string]string{
		"name":  "vue",
		"Name":  "go语言",
		"Price": "35.9",
	})

	t3 := create("t3", "\nRange: \n {{range .}}{{.}}\n{{end}}")
	t3.Execute(os.Stdout, []string{"golang", "k8s", "vue", "http/gprc"})

	t4 := create("t4", "{{- if . -}}yes{{- else -}}no{{end}}\n")
	t4.Execute(os.Stdout, "abc no empty")
	t4.Execute(os.Stdout, "")
	t4.Execute(os.Stdout, 0)
	t4.Execute(os.Stdout, nil)
	t4.Execute(os.Stdout, false)
	t4.Execute(os.Stdout, 0.0)
	t4.Execute(os.Stdout, ``)
	t4.Execute(os.Stdout, true)

}

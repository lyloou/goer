package template

import (
	"text/template"
	"os"
	"fmt"
	"strings"
	"testing"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	return strings.Replace(s, "@", "#", -1)

	//// find the @ symbol
	//substrs := strings.Split(s, "@")
	//if len(substrs) != 2 {
	//	return s
	//}
	//// replace the @ by " at "
	//return substrs[0] + "#" + substrs[1]
}

func TestTemplate(tt *testing.T) {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
				`)
	p := Person{UserName: "Astaxie",
		Emails: []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
func sample3() {
	tEmpty := template.New("test")
	template.Must(tEmpty.Parse("空的　pipline if demo ：{{if ``}} 不会输出。{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)
	tWithValue := template.New("test")
	template.Must(tWithValue.Parse("不为空的　pipline if demo ：{{if `anything`}} 会输出。{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)
	tIfElse := template.New("test")
	template.Must(tIfElse.Parse("if else demo ：{{if `看`}} 会输出。{{else}} else 部分。{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
	tEmail := template.New("test")
	template.Must(tEmail.Parse("Email: {{. | html}}"))
	tEmail.Execute(os.Stdout, "mailto:lyloou@qq.com")
	tVar := template.New("test")
	template.Must(tVar.Parse(`
		{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
		{{with $x := "output"}}{{printf "%q" $x}}{{end}}
		{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
		`))
	tVar.Execute(os.Stdout, "mailto:lyloou@qq.com")
}

func sample1() {
	d := User{"Lou", 25}
	t := template.New("name")
	t.Parse("hello, {{.UserName}}, you age is {{.Age}}, \nthen you can see:\n{{.}}")
	t.Execute(os.Stdout, d)
}

type User struct {
	UserName string
	Age      rune
}

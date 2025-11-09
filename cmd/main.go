package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var x = 2

func testTemplate() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	t, err := template.New("default").Parse(tpl)
	if err != nil {
		log.Fatalf("%v", err)
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: "Hello",
		Items: []string{"yo", template.HTMLEscapeString("what's up")},
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("%v", err)
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	log.Printf("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

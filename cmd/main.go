package main

import (
	"html/template"
	"log"
	"net/http"
)

var x = 2

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		// fmt.Fprintf(w, "Welcome to the home page!")
		const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
    <script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>
	</head>
	<body>
		{{range .Items}}<div class="bg-red-200">{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

		t, err := template.New("default").Parse(tpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Items []string
		}{
			Title: "Hello",
			Items: []string{"yo", "what's up"}, // template automatically escapes it
		}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	log.Printf("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

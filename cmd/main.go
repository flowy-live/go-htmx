package main

import (
	"html/template"
	"log"
	"net/http"
)

var x = 2

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		const tpl = `
<div class="w-full p-6">
  <div class="overflow-hidden rounded-lg border border-gray-200 shadow-md">
    <table class="w-full border-collapse bg-white text-left text-sm text-gray-500">
      <thead class="bg-gray-50">
        <tr>
          <th scope="col" class="px-6 py-4 font-medium text-gray-900">Name</th>
          <th scope="col" class="px-6 py-4 font-medium text-gray-900">Email</th>
          <th scope="col" class="px-6 py-4 font-medium text-gray-900">Role</th>
          <th scope="col" class="px-6 py-4 font-medium text-gray-900">Status</th>
        </tr>
      </thead>
      <tbody class="divide-y divide-gray-100 border-t border-gray-100">
        {{range .Users}}
        <tr class="hover:bg-gray-50">
          <td class="px-6 py-4">
            <div class="flex items-center gap-3">
              <div class="h-10 w-10 rounded-full bg-gradient-to-br from-purple-400 to-pink-400 flex items-center justify-center text-white font-semibold">
                {{.Initials}}
              </div>
              <div class="font-medium text-gray-700">{{.Name}}</div>
            </div>
          </td>
          <td class="px-6 py-4">{{.Email}}</td>
          <td class="px-6 py-4">
            <span class="inline-flex items-center gap-1 rounded-full bg-blue-50 px-2 py-1 text-xs font-semibold text-blue-600">
              {{.Role}}
            </span>
          </td>
          <td class="px-6 py-4">
            <span class="inline-flex items-center gap-1 rounded-full {{.StatusClass}} px-2 py-1 text-xs font-semibold">
              <span class="h-1.5 w-1.5 rounded-full {{.StatusDotClass}}"></span>
              {{.Status}}
            </span>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
    `
		t, err := template.New("default").Parse(tpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type User struct {
			Name           string
			Initials       string
			Email          string
			Role           string
			Status         string
			StatusClass    string
			StatusDotClass string
		}

		data := struct {
			Users []User
		}{
			Users: []User{
				{Name: "Sarah Chen", Initials: "SC", Email: "sarah.chen@example.com", Role: "Engineering", Status: "Active", StatusClass: "bg-green-50 text-green-600", StatusDotClass: "bg-green-600"},
				{Name: "Marcus Johnson", Initials: "MJ", Email: "marcus.j@example.com", Role: "Product", Status: "Active", StatusClass: "bg-green-50 text-green-600", StatusDotClass: "bg-green-600"},
				{Name: "Emma Rodriguez", Initials: "ER", Email: "emma.r@example.com", Role: "Design", Status: "Away", StatusClass: "bg-yellow-50 text-yellow-600", StatusDotClass: "bg-yellow-600"},
				{Name: "James Wilson", Initials: "JW", Email: "james.w@example.com", Role: "Marketing", Status: "Active", StatusClass: "bg-green-50 text-green-600", StatusDotClass: "bg-green-600"},
				{Name: "Priya Patel", Initials: "PP", Email: "priya.p@example.com", Role: "Sales", Status: "Inactive", StatusClass: "bg-gray-50 text-gray-600", StatusDotClass: "bg-gray-400"},
				{Name: "Alex Kim", Initials: "AK", Email: "alex.kim@example.com", Role: "Engineering", Status: "Active", StatusClass: "bg-green-50 text-green-600", StatusDotClass: "bg-green-600"},
			},
		}

		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
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
    <script src="https://cdn.jsdelivr.net/npm/htmx.org@2.0.8/dist/htmx.min.js"></script>
	</head>
	<body>
		{{range .Items}}<div class="bg-red-200">{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}

    <button hx-post="/data" hx-swap="outerHTML" class="bg-pink-500">
      Click Me
    </button>
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

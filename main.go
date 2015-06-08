package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

// Command-line flags.
var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func main() {
	flag.Parse()
	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// Server implements the outyet server.
// It serves the user interface (it's an http.Handler)
// and polls the remote repository for changes.
type Server struct {
}

// NewServer returns an initialized outyet server.
func NewServer() *Server {
	return &Server{}
}

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// if key := r.URL.Query().Get("key") {
	//   env := os.GetEnv(key)
	// }
	env := os.Environ()
	s.renderJSON(w, env)
}

func (s *Server) renderJSON(w http.ResponseWriter, obj interface{}) {
	data := struct {
		ENV []string
	}{
		getEnv(),
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

func getEnv() []string {
	env := os.Environ()
	sort.Strings(env)
	return env
}

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body>
	<h1>ENV</h1>
	<ul>
  {{range $env := .ENV}}
    <li>{{$env}}</li>
  {{end}}
	</ul>
</body></html>
`))

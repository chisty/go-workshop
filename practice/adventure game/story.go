package game

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl}

	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.S[path]; ok {
		err := h.T.Execute(w, chapter)
		if err != nil {
			log.Printf("%+v\n", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.T = t
	}
}

func ParseJsonStory(r io.Reader) (Story, error) {
	var story Story
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

var tpl *template.Template

type handler struct {
	S Story
	T *template.Template
}

type HandlerOption func(h *handler)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var defaultTemplate = `
<html>
    <head>
        <title>Adventure Game Story</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Story}}
            <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
                <li><a href="/{{.Arc}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>
`

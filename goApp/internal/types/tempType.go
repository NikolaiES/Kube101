package types

import "html/template"

type TmplData struct {
	Title   string
	Data    template.HTML
	PodName string
}

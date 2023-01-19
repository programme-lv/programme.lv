package models

type MDSTatementExample struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type MarkdownStatement struct {
	Name     string               `json:"name"`
	Desc     string               `json:"desc"`
	Input    string               `json:"input"`
	Output   string               `json:"output"`
	Notes    string               `json:"notes"`
	Scoring  string               `json:"scoring"`
	Examples []MDSTatementExample `json:"examples"`
}

type PDFStatement struct {
	Name string `json:"filename"`
}

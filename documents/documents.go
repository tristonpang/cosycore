package documents

type Document struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

var allDocs []Document

func AddDocument(doc Document) {
	allDocs = append(allDocs, doc)
}

func GetAllDocuments() []Document {
	return allDocs
}

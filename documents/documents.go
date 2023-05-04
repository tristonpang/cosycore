package documents

import "errors"

type Document struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

var allDocs map[string]Document = make(map[string]Document)

func AddDocument(doc Document) error {
	_, exists := allDocs[doc.Name]
	if exists {
		return errors.New("Document with the same name already exists")
	}

	allDocs[doc.Name] = doc
	return nil
}

func GetAllDocuments() []Document {
	docsSlice := make([]Document, 0, len(allDocs))
	for _, val := range allDocs {
		docsSlice = append(docsSlice, val)
	}
	return docsSlice
}

func UpdateDocument(doc Document) error {
	_, exists := allDocs[doc.Name]
	if !exists {
		return errors.New("Document could not be found")
	}

	allDocs[doc.Name] = doc
	return nil
}

func DeleteDocument(doc Document) error {
	_, exists := allDocs[doc.Name]
	if !exists {
		return errors.New("Document could not be found")
	}

	delete(allDocs, doc.Name)
	return nil
}

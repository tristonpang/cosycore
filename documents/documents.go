package documents

import "errors"

type Document struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

func AddDocument(doc Document) error {
	_, exists := DbRetrieveDocument(doc.Name)
	if exists {
		return errors.New("Document with the same name already exists")
	}

	DbInsertDocument(doc)
	return nil
}

func GetAllDocuments() []Document {
	return DbRetrieveAllDocuments()
}

func UpdateDocument(doc Document) error {
	_, exists := DbRetrieveDocument(doc.Name)
	if !exists {
		return errors.New("Document could not be found")
	}

	DbUpdateDocument(doc)
	return nil
}

func DeleteDocument(doc Document) error {
	_, exists := DbRetrieveDocument(doc.Name)
	if !exists {
		return errors.New("Document could not be found")
	}

	DbDeleteDocument(doc.Name)
	return nil
}

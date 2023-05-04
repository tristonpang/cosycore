package documents

var database map[string]Document = make(map[string]Document)

func DbInsertDocument(doc Document) {
	database[doc.Name] = doc
}

func DbRetrieveDocument(docName string) (Document, bool) {
	doc, exists := database[docName]
	return doc, exists
}

func DbRetrieveAllDocuments() []Document {
	docsSlice := make([]Document, 0, len(database))
	for _, val := range database {
		docsSlice = append(docsSlice, val)
	}
	return docsSlice
}

func DbUpdateDocument(doc Document) {
	database[doc.Name] = doc
}

func DbDeleteDocument(docName string) {
	delete(database, docName)
}

package documents

import (
	"testing"
)

var testDocument = Document{
	Name:        "Test Document",
	Description: "This is a test document",
	Location:    "/test/location/A",
}

func TestGetAllDocuments_Empty(t *testing.T) {
	docs := GetAllDocuments()
	if len(docs) != 0 {
		t.Fatalf(`GetAllDocuments = %v, want []`, docs)
	}
}

func TestAddDocument_Success(t *testing.T) {
	err := AddDocument(testDocument)
	if err != nil {
		t.Fatalf(`AddDocument error = %v`, err)
	}
}

func TestGetAllDocuments_NonEmpty(t *testing.T) {
	docs := GetAllDocuments()
	if len(docs) != 1 {
		t.Fatalf(`GetAllDocuments = %v, want %v`, docs, append(make([]Document, 0), testDocument))
	}

	testDoc := docs[0]
	if testDoc != testDocument {
		t.Fatalf(`GetAllDocuments testDoc = %v, want %v`, testDoc, testDocument)
	}
}

package finance

import (
	"errors"
	"os/exec"

	"github.com/augurysys/timestamp"
)

type FinanceEntry struct {
	Id       string              `json:"id"`
	DateTime timestamp.Timestamp `json:"datetime"`
	Name     string              `json:"name"`
	Amount   float32             `json:"amount"`
}

func AddEntry(entry FinanceEntry) error {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return errors.New("UUID could not be generated")
	}
	entry.Id = string(uuid)
	DbInsertFinanceEntry(entry)
	return nil
}

func GetAllEntries() {
	return DbRetrieveAllFinanceEntries()
}

func UpdateEntry(entry FinanceEntry) error {
	_, exists := DbFindFinanceEntry(entry.Id)
	if !exists {
		return errors.New("Entry could not be found")
	}

	DbUpdateFinanceEntry(entry)
	return nil
}

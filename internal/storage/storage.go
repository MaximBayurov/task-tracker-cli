package storage

import (
	"encoding/json"
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"os"
	"time"
)

// Storage contains data about tasks
type Storage struct {
	rows     map[int64]DataRow
	fileName string
}

// Add add data about task to storage
func (s *Storage) Add(row DataRow) int64 {
	var newID int64
	ln := int64(len(s.rows))
	if ln <= 0 {
		newID = 1
		s.rows = make(map[int64]DataRow)
	} else {
		newID = findMaxKey(s.rows) + 1
	}

	row.ID = newID
	row.CreatedAt = time.Now()
	row.UpdatedAt = time.Now()
	if row.Status <= 0 {
		row.Status = mark.Todo
	}
	s.rows[newID] = row

	return newID
}

// Delete delete data about task from storage by id
func (s *Storage) Delete(ID int64) error {
	_, ok := s.rows[ID]
	if !ok {
		return fmt.Errorf("cannot delete command with ID=\"%v\"", ID)
	}
	delete(s.rows, ID)
	return nil
}

// Update update data about task in storage by id
func (s *Storage) Update(ID int64, row DataRow) error {
	r, ok := s.rows[ID]
	if !ok {
		return fmt.Errorf("cannot update command with ID=\"%v\"", ID)
	}
	row.ID = ID
	row.UpdatedAt = time.Now()
	row.CreatedAt = r.CreatedAt
	s.rows[ID] = row
	return nil
}

// GetById returns data about task by id
func (s *Storage) GetById(ID int64) (DataRow, error) {
	row, ok := s.rows[ID]
	if !ok {
		return DataRow{}, fmt.Errorf("cannot find command with ID=\"%v\"", ID)
	}
	return row, nil
}

// GetByStatus returns the list of tasks data filtered by status
func (s *Storage) GetByStatus(status mark.Mark) []DataRow {
	var result []DataRow
	for _, row := range s.rows {
		if row.Status != status {
			continue
		}
		result = append(result, row)
	}
	return result
}

// GetAll returns all task data rows from storage
func (s *Storage) GetAll() []DataRow {
	var result []DataRow
	for _, row := range s.rows {
		result = append(result, row)
	}
	return result
}

// Load loads all task data rows from file
func (s *Storage) Load() error {
	b, err := os.ReadFile(s.fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.rows)
}

// Save saves all task data rows to file
func (s *Storage) Save() error {
	if len(s.rows) <= 0 {
		return nil
	}
	b, err := json.Marshal(s.rows)
	if err != nil {
		return err
	}
	return os.WriteFile(s.fileName, b, 0644)
}

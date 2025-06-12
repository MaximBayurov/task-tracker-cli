package storage

import (
	"errors"
	"fmt"
	"github.com/MaximBayurov/task-tracker-cli/internal/mark"
	"os"
	"reflect"
	"testing"
	"time"
)

var testStorageFileName = "todo-list-test.json"

func initTestStorage() {
	removeTestStorage()
	s, err := GetInstance()
	if err != nil {
		_ = Init(testStorageFileName)
		return
	}
	s.rows = map[int64]DataRow{}
}

func removeTestStorage() {
	_ = os.Remove(testStorageFileName)
}

func TestStorage_Load(t *testing.T) {
	//arrange
	initTestStorage()
	defer removeTestStorage()
	rows := map[int64]DataRow{
		1: {ID: 1, Description: "Test1", Status: mark.Todo, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		2: {ID: 2, Description: "Test2", Status: mark.InProgress, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		3: {ID: 3, Description: "Test3", Status: mark.Canceled, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		4: {ID: 4, Description: "Test4", Status: mark.Done, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		5: {ID: 5, Description: "Test5", Status: mark.Todo, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	s, _ := GetInstance()
	for _, row := range rows {
		s.Add(row)
	}
	_ = s.Save()
	s.rows = map[int64]DataRow{}

	//act
	_ = s.Load()

	//assert
	if len(s.rows) == 0 {
		t.Error("Loading doesn't work. Rows doesn't load")
		return
	}

	if len(s.rows) != len(rows) {
		t.Error("Loading doesn't work. All rows doesn't load")
		return
	}

	for ID, row1 := range s.rows {
		row2 := s.rows[ID]
		isRowsEqual := true
		isRowsEqual = isRowsEqual && row1.ID == row2.ID && row1.Description == row2.Description && row1.Status == row2.Status
		isRowsEqual = isRowsEqual && row1.CreatedAt.Equal(row2.CreatedAt) && row1.UpdatedAt.Equal(row2.UpdatedAt)
		if !isRowsEqual {
			t.Error("Loading doesn't work. Rows doesn't load correctly")
			return
		}
	}
}

func TestStorage_Save(t *testing.T) {
	removeTestStorage()
	initTestStorage()
	defer removeTestStorage()

	s, _ := GetInstance()
	s.Add(DataRow{
		ID:          1,
		Description: "Test",
		Status:      mark.InProgress,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	_ = s.Save()

	if stat, err := os.Stat(s.fileName); err == nil {
		if stat.Size() <= 0 {
			t.Error("File creates on Save but it's empty")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		t.Error("File doesn't create on Save")
	} else {
		t.Error(err)
	}
}

func TestStorage_Add(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()

	rows := []DataRow{
		{ID: 11, Description: "Test1", Status: mark.Todo, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 22, Description: "Test2", Status: mark.InProgress},
		{ID: 33, Description: "Test3"},
		{ID: 44},
		{},
	}

	s, _ := GetInstance()
	for i, row := range rows {
		name := fmt.Sprintf("dataset_%v", i+1)
		t.Run(name, func(t *testing.T) {
			unix := time.Now().Unix()
			ID := s.Add(row)
			expID := int64(i) + 1
			if expID != ID {
				t.Errorf("Add method returns incorrect ID %v, %v expected", expID, ID)
				return
			}
			r, _ := s.GetById(ID)
			if !(row.Description == r.Description) {
				t.Error("Invalid Description")
			}
			if !((row.Status > 0 && row.Status == r.Status) || (row.Status == 0 && r.Status == mark.Todo)) {
				t.Error("Invalid Status")
			}
			if !((!row.CreatedAt.IsZero() && row.CreatedAt.Equal(r.CreatedAt)) || (row.CreatedAt.IsZero() && r.CreatedAt.Unix() == unix)) {
				t.Error("Invalid CreatedAt")
			}
			if !((!row.UpdatedAt.IsZero() && row.UpdatedAt.Equal(r.UpdatedAt)) || (row.UpdatedAt.IsZero() && r.UpdatedAt.Unix() == unix)) {
				t.Error("Invalid UpdatedAt")
			}
		})
	}

	if _, err := os.Stat(s.fileName); err == nil {
		t.Error("Storage file creates on Add. It's redundant")
	}
}

func TestStorage_Update(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()

	s, _ := GetInstance()

	now := time.Now()
	before := DataRow{11, "Test1", mark.Todo, now, now}
	ID := s.Add(before)

	_ = s.Update(ID, DataRow{
		ID:          2,
		Description: "Test2",
		Status:      mark.InProgress,
		CreatedAt:   time.Now().Add(time.Hour * 1),
		UpdatedAt:   time.Now().Add(time.Hour * 2),
	})

	after, _ := s.GetById(ID)
	if ID != after.ID {
		t.Error("The ID is changed. But expected it's not")
	}
	if before.Description == after.Description {
		t.Error("The Description doesn't change")
	}
	if before.Status == after.Status {
		t.Error("The Status doesn't change")
	}
	if before.CreatedAt != after.CreatedAt {
		t.Error("The CreatedAt is changed. But expected it's not")
	}
	if now.Unix() != after.UpdatedAt.Unix() {
		t.Errorf("The UpdatedAt changes as not expected. As is: %v. Expected: %v", after.UpdatedAt, now)
	}

	if _, err := os.Stat(s.fileName); err == nil {
		t.Error("Storage file creates on Update. It's redundant")
	}
}

func TestStorage_Delete(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()

	s, _ := GetInstance()

	now := time.Now()
	ID := s.Add(DataRow{1, "Test1", mark.Todo, now, now})

	_ = s.Delete(ID)

	row, _ := s.GetById(ID)
	if !reflect.DeepEqual(row, DataRow{}) {
		t.Errorf("Row with ID %v doesn't deleted", ID)
	}

	if _, err := os.Stat(s.fileName); err == nil {
		t.Error("Storage file creates on Delete. It's redundant")
	}
}

func TestStorage_GetById(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()
	rows := map[int64]DataRow{
		1: {ID: 1, Description: "Test1", Status: mark.Todo},
		2: {ID: 2, Description: "Test2", Status: mark.InProgress},
		3: {ID: 3, Description: "Test3", Status: mark.Canceled},
		4: {ID: 4, Description: "Test4", Status: mark.Done},
		5: {ID: 5, Description: "Test5", Status: mark.Todo},
	}
	s, _ := GetInstance()
	for _, row := range rows {
		s.Add(row)
	}

	_, err := s.GetById(0)
	if err == nil {
		t.Error("GetById should returns error when pass not existing ID")
	}

	for ID, row := range rows {
		name := fmt.Sprintf("dataset_%v", ID)
		t.Run(name, func(t *testing.T) {
			r, _ := s.GetById(ID)
			if r.Description != row.Description && r.Status != row.Status {
				t.Errorf("GetById return wrong row by ID %v", ID)
			}
		})
	}
}

func TestStorage_GetByStatus(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()
	rows := map[int64]DataRow{
		1: {ID: 1, Description: "Test1", Status: mark.Todo},
		2: {ID: 2, Description: "Test2", Status: mark.Todo},
		3: {ID: 3, Description: "Test3", Status: mark.Canceled},
		4: {ID: 4, Description: "Test4", Status: mark.Canceled},
		5: {ID: 5, Description: "Test5", Status: mark.Done},
	}
	s, _ := GetInstance()
	for _, row := range rows {
		s.Add(row)
	}

	marks := make(map[mark.Mark]int, 0)
	for _, row := range rows {
		marks[row.Status] += 1
	}

	for _, m := range mark.AllowedMarks {
		result := s.GetByStatus(m)
		exp := marks[m]
		if len(result) != exp {
			t.Errorf("Incorrect number of rows. As is: %v. Expected: %v", len(result), exp)
		}
	}
}

func TestStorage_GetAll(t *testing.T) {
	initTestStorage()
	defer removeTestStorage()
	rows := map[int64]DataRow{
		1: {ID: 1, Description: "Test1", Status: mark.Todo},
		2: {ID: 2, Description: "Test2", Status: mark.Todo},
		3: {ID: 3, Description: "Test3", Status: mark.Canceled},
		4: {ID: 4, Description: "Test4", Status: mark.Canceled},
		5: {ID: 5, Description: "Test5", Status: mark.Done},
	}
	s, _ := GetInstance()
	for _, row := range rows {
		s.Add(row)
	}

	exp := len(rows)
	res := len(s.GetAll())
	if res != exp {
		t.Errorf("All rows doesn't selected. As is: %v. Expected: %v", res, exp)
	}
}

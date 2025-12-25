package internal

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"levyvix/togo/internal/database"
	"levyvix/togo/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TestDB holds the test database connection
var testDB *gorm.DB

// init sets up the test database before running tests
func init() {
	var err error
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect to test database")
	}

	// Migrate the schema
	if err := testDB.AutoMigrate(&schema.Task{}); err != nil {
		panic("failed to migrate test database")
	}

	// Replace the global DB with the test database
	database.DB = testDB
}

// clearDB clears all data from the test database
func clearDB(t *testing.T) {
	if err := testDB.Exec("DELETE FROM tasks").Error; err != nil {
		t.Fatalf("failed to clear database: %v", err)
	}
}

// TestFormatDate tests the formatDate function with various inputs
func TestFormatDate(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		contains []string // strings that should be in the output
	}{
		{
			name:     "December date",
			input:    time.Date(2025, 12, 21, 14, 30, 0, 0, time.UTC),
			contains: []string{"21", "Dez", "2025", "14:30"},
		},
		{
			name:     "January date",
			input:    time.Date(2025, 1, 15, 10, 45, 0, 0, time.UTC),
			contains: []string{"15", "Jan", "2025", "10:45"},
		},
		{
			name:     "February date",
			input:    time.Date(2025, 2, 28, 23, 59, 0, 0, time.UTC),
			contains: []string{"28", "Fev", "2025", "23:59"},
		},
		{
			name:     "August date",
			input:    time.Date(2025, 8, 5, 9, 0, 0, 0, time.UTC),
			contains: []string{"05", "Ago", "2025", "09:00"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDate(tt.input)

			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("formatDate(%v) = %q, expected to contain %q", tt.input, result, expected)
				}
			}
		})
	}
}

// TestCreateFuncDB tests creating a new task
func TestCreateFuncDB(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{
			name:      "Valid single argument",
			args:      []string{"Estudar Go"},
			wantError: false,
		},
		{
			name:      "Valid single argument with special chars",
			args:      []string{"Fazer café com açúcar!"},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearDB(t)

			// Capture stdout
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			CreateFuncDB(tt.args)

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			os.Stdout = oldStdout

			// Verify task was created
			var count int64
			testDB.Model(&schema.Task{}).Count(&count)
			if count != 1 {
				t.Errorf("CreateFuncDB(%v) should have created 1 task, but created %d", tt.args, count)
			}

			// Verify task content
			var task schema.Task
			testDB.First(&task)
			if task.Description != tt.args[0] {
				t.Errorf("CreateFuncDB(%v) description = %q, want %q", tt.args, task.Description, tt.args[0])
			}
			if task.Done != false {
				t.Errorf("CreateFuncDB(%v) Done = %v, want false", tt.args, task.Done)
			}
			if task.DoneAt != nil {
				t.Errorf("CreateFuncDB(%v) DoneAt = %v, want nil", tt.args, task.DoneAt)
			}
		})
	}
}

// TestDoneFuncDB tests marking a task as done
func TestDoneFuncDB(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() uint
		args      []string
		checkDone bool
	}{
		{
			name: "Mark valid task as done",
			setup: func() uint {
				task := schema.Task{Description: "Test task", Done: false}
				testDB.Create(&task)
				return task.ID
			},
			args:      []string{"1"},
			checkDone: true,
		},
		{
			name: "Non-existent task ID",
			setup: func() uint {
				return 999
			},
			args:      []string{"999"},
			checkDone: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearDB(t)
			taskID := tt.setup()

			// Use the actual task ID in the arguments
			args := tt.args
			if taskID > 0 && len(args) > 0 {
				args = []string{fmt.Sprintf("%d", taskID)}
			}

			// Capture output
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			DoneFuncDB(args)

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			os.Stdout = oldStdout

			if tt.checkDone && taskID > 0 {
				// Verify task was marked as done
				var task schema.Task
				testDB.First(&task, taskID)
				if !task.Done {
					t.Errorf("DoneFuncDB(%v) Done = false, want true (task ID: %d)", args, taskID)
				}
				if task.DoneAt == nil {
					t.Errorf("DoneFuncDB(%v) DoneAt should be set, got nil", args)
				}
			}
		})
	}
}

// TestDeleteFuncDB tests deleting a task
func TestDeleteFuncDB(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() uint
		verifyID bool
	}{
		{
			name: "Delete existing task",
			setup: func() uint {
				task := schema.Task{Description: "Task to delete"}
				testDB.Create(&task)
				return task.ID
			},
			verifyID: true,
		},
		{
			name: "Delete non-existent task",
			setup: func() uint {
				return 999
			},
			verifyID: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearDB(t)
			taskID := tt.setup()

			// Use the actual task ID in the arguments
			args := []string{fmt.Sprintf("%d", taskID)}

			// Capture output
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			DeleteFuncDB(args)

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			os.Stdout = oldStdout

			// If verifying, check task was deleted
			if tt.verifyID && taskID > 0 {
				var count int64
				testDB.Model(&schema.Task{}).Where("id = ?", taskID).Count(&count)
				if count != 0 {
					t.Errorf("DeleteFuncDB(%v) should have deleted task %d, but it still exists", args, taskID)
				}
			}
		})
	}
}

// TestEditFuncDB tests editing a task
func TestEditFuncDB(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() uint
		args         []string
		expectedDesc string
	}{
		{
			name: "Edit existing task",
			setup: func() uint {
				task := schema.Task{Description: "Old description"}
				testDB.Create(&task)
				return task.ID
			},
			args:         []string{"1", "Nova descrição"},
			expectedDesc: "Nova descrição",
		},
		{
			name: "Edit with special characters",
			setup: func() uint {
				task := schema.Task{Description: "Old"}
				testDB.Create(&task)
				return task.ID
			},
			args:         []string{"1", "Fazer café com açúcar e pão!"},
			expectedDesc: "Fazer café com açúcar e pão!",
		},
		{
			name: "Edit non-existent task",
			setup: func() uint {
				return 999
			},
			args:         []string{"999", "New description"},
			expectedDesc: "", // No verification needed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearDB(t)
			taskID := tt.setup()

			// Use the actual task ID in the arguments
			args := tt.args
			if taskID > 0 && len(args) > 1 {
				args = []string{fmt.Sprintf("%d", taskID), tt.args[1]}
			}

			// Capture output
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			EditFuncDB(args)

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			os.Stdout = oldStdout

			if tt.expectedDesc != "" && taskID > 0 {
				// Verify task was updated
				var task schema.Task
				testDB.First(&task, taskID)
				if task.Description != tt.expectedDesc {
					t.Errorf("EditFuncDB(%v) description = %q, want %q (task ID: %d)", args, task.Description, tt.expectedDesc, taskID)
				}
			}
		})
	}
}

// TestListFuncDB tests listing tasks
func TestListFuncDB(t *testing.T) {
	tests := []struct {
		name          string
		setup         func()
		checkContains []string
	}{
		{
			name: "List empty tasks",
			setup: func() {
				clearDB(t)
			},
			checkContains: []string{"Nenhuma task"},
		},
		{
			name: "List single task",
			setup: func() {
				clearDB(t)
				testDB.Create(&schema.Task{Description: "Test task 1"})
			},
			checkContains: []string{"Test task 1", "⏳"},
		},
		{
			name: "List multiple tasks",
			setup: func() {
				clearDB(t)
				testDB.Create(&schema.Task{Description: "Task 1"})
				testDB.Create(&schema.Task{Description: "Task 2"})
				testDB.Create(&schema.Task{Description: "Task 3"})
			},
			checkContains: []string{"Task 1", "Task 2", "Task 3"},
		},
		{
			name: "List with completed tasks",
			setup: func() {
				clearDB(t)
				now := time.Now()
				testDB.Create(&schema.Task{Description: "Incomplete", Done: false})
				testDB.Create(&schema.Task{Description: "Complete", Done: true, DoneAt: &now})
			},
			checkContains: []string{"⏳", "✓"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			// Capture stdout
			oldStdout := os.Stdout
			reader, w, _ := os.Pipe()
			os.Stdout = w

			ListFuncDB()

			if err := w.Close(); err != nil {
				t.Fatalf("failed to close pipe: %v", err)
			}
			output, _ := io.ReadAll(reader)
			os.Stdout = oldStdout

			outputStr := string(output)
			for _, expected := range tt.checkContains {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("ListFuncDB() output should contain %q, got:\n%s", expected, outputStr)
				}
			}
		})
	}
}

// BenchmarkFormatDate benchmarks the formatDate function
func BenchmarkFormatDate(b *testing.B) {
	t := time.Date(2025, 12, 21, 14, 30, 0, 0, time.UTC)
	for b.Loop() {
		formatDate(t)
	}
}

// BenchmarkCreateFuncDB benchmarks creating tasks
func BenchmarkCreateFuncDB(b *testing.B) {
	clearDB(&testing.T{})

	// Silence output
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	for b.Loop() {
		CreateFuncDB([]string{"Benchmark task"})
	}

	if err := w.Close(); err != nil {
		b.Fatalf("failed to close pipe: %v", err)
	}
	os.Stdout = oldStdout
}

// BenchmarkListFuncDB benchmarks listing tasks
func BenchmarkListFuncDB(b *testing.B) {
	clearDB(&testing.T{})

	// Create 100 tasks
	for i := range 100 {
		testDB.Create(&schema.Task{Description: fmt.Sprintf("Task %d", i)})
	}

	// Silence output
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	for b.Loop() {
		ListFuncDB()
	}

	if err := w.Close(); err != nil {
		b.Fatalf("failed to close pipe: %v", err)
	}
	os.Stdout = oldStdout
}

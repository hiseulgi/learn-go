package task

import (
	"testing"
)

// * Create Task
/* -------------------------------------------------------------------------- */
func TestCreateTaskValidName(t *testing.T) {
	tasks := TaskSlice{}
	err := tasks.CreateTask(Task{Name: "test", IsCompleted: false})

	// check error return value
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	// assert that the task has been added to the TaskSlice
	if len(tasks) != 1 {
		t.Errorf("Expected TaskSlice length to be 1, but got %d", len(tasks))
	}
}

func TestCreateTaskEmpty(t *testing.T) {
	tasks := TaskSlice{}
	err := tasks.CreateTask(Task{Name: "", IsCompleted: false})

	// check error return value
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	// assert that the tasks remains empty
	if len(tasks) != 0 {
		t.Errorf("Expected TaskSlice length to be 0, but got %d", len(tasks))
	}
}

// * Mark Task
/* -------------------------------------------------------------------------- */
func TestMarkTask(t *testing.T) {
	tasks := TaskSlice{
		{Name: "test1", IsCompleted: false},
		{Name: "test2", IsCompleted: false},
	}
	err := tasks.MarkTask(0)

	// check error return value
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// assert that the task has been marked
	if !tasks[0].IsCompleted {
		t.Errorf("Expected TaskSlice 0 completion status to be true, but got false")
	}
}

func TestMarkTaskOutOfRange(t *testing.T) {
	tasks := TaskSlice{
		{Name: "test1", IsCompleted: false},
		{Name: "test2", IsCompleted: false},
	}
	err := tasks.MarkTask(100)

	// check error return value
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// assert that the task remains unchanged
	if tasks[0].IsCompleted || tasks[1].IsCompleted {
		t.Errorf("Expected TaskSlice completion status to be false, but got true")
	}
}

// * Delete Task
/* -------------------------------------------------------------------------- */
func TestDeleteTask(t *testing.T) {
	tasks := TaskSlice{
		{Name: "test1", IsCompleted: false},
		{Name: "test2", IsCompleted: false},
	}
	err := tasks.DeleteTask(0)

	// check error return value
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// assert that the task has been deleted
	if len(tasks) != 1 {
		t.Errorf("Expected TaskSlice length to be 1, but got %d", len(tasks))
	}
}

func TestDeleteTaskOutOfRange(t *testing.T) {
	tasks := TaskSlice{
		{Name: "test1", IsCompleted: false},
		{Name: "test2", IsCompleted: false},
	}
	err := tasks.MarkTask(100)

	// check error return value
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// assert that the task remains unchanged
	if len(tasks) != 2 {
		t.Errorf("Expected TaskSlice length to be 2, but got %d", len(tasks))
	}
}

// task package
// this package is used to create, read, update, and delete task

package task

import "fmt"

// * Task Struct
// this struct is used to store task data
type Task struct {
	Name        string
	IsCompleted bool
}

// * TaskSlice
// this type is used to store slice of task
type TaskSlice []Task

// * Create Task
// this function is used to create task with input data Task and append it to TaskSlice
func (ts *TaskSlice) CreateTask(data Task) error {
	if data.Name == "" {
		return fmt.Errorf("Task name cannot empty!")
	}
	*ts = append(*ts, data)
	return nil
}

// * Read Task
// this function is used to read all task in TaskSlice
func (ts *TaskSlice) ReadTask() {
	if len(*ts) == 0 {
		fmt.Println("You dont have task!")
		return
	}
	fmt.Println("Index\t|\tTask Name\t|\tMark")
	for i, task := range *ts {
		mark := " "
		if task.IsCompleted {
			mark = "x"
		}
		fmt.Printf("%d. %s %s\n", i, task.Name, mark)
	}
}

// * Check Task (Update IsCompleted Only)
// this function is used to check task with index and update IsCompleted
func (ts *TaskSlice) MarkTask(index int) error {
	if len(*ts)-1 < index {
		return fmt.Errorf("Task with index %d not found!", index)
	}

	currentTask := (*ts)[index]
	(*ts)[index] = Task{Name: (currentTask).Name, IsCompleted: !(currentTask).IsCompleted}
	return nil
}

// * Delete Task
// this function is used to delete task with index
func (ts *TaskSlice) DeleteTask(index int) error {
	if len(*ts)-1 < index {
		return fmt.Errorf("Task with index %d not found!", index)
	}
	*ts = append((*ts)[:index], (*ts)[index+1:]...)
	return nil
}

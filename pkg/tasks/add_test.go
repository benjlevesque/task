package tasks_test

import (
	"testing"

	"github.com/benjlevesque/task/mocks"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStore struct {
	tasks []types.Task
}

func (store *mockStore) CreateTask(title string) (int, error) {
	task := types.Task{
		ID:    len(store.tasks),
		Title: title,
		Done:  false,
	}
	store.tasks = append(store.tasks, task)
	return 0, nil
}

func TestAdd(t *testing.T) {
	store := &mockStore{}
	tasks.AddTask(store, []string{"toto"})

	assert.Equal(t, len(store.tasks), 1, "store.tasks length")
	assert.Equal(t, store.tasks[0].Title, "toto")
	assert.Equal(t, store.tasks[0].Done, false)
}

func TestAddMock(t *testing.T) {
	store := &mocks.TaskCreater{}
	store.On("CreateTask", mock.AnythingOfType("string")).Return(8, nil)
	tasks.AddTask(store, []string{"toto"})

	store.AssertCalled(t, "CreateTask", "toto")
}

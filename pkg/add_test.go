package pkg_test

import (
	"testing"

	"github.com/benjlevesque/task/pkg"
	"github.com/benjlevesque/task/types"
	"github.com/stretchr/testify/assert"
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
	pkg.AddTask(store, []string{"toto"})

	assert.Equal(t, len(store.tasks), 1, "store.tasks length")
	assert.Equal(t, store.tasks[0].Title, "toto")
	assert.Equal(t, store.tasks[0].Done, false)
}

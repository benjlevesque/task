package tasks_test

import (
	"testing"

	"github.com/benjlevesque/task/mocks"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/types"
	"github.com/stretchr/testify/mock"
)

func TestEditMock(t *testing.T) {
	store := &mocks.TaskEditer{}
	editor := &mocks.TextEditer{}
	store.On("GetTask", 1).Return(types.Task{Title: "toto"}, nil)
	store.On("EditTask", 1, mock.AnythingOfType("string")).Return(nil)
	editor.On("EditText", mock.AnythingOfType("string")).Return("titi", nil)

	tasks.EditTask(store, editor, []string{"1"})

	editor.AssertCalled(t, "EditText", "toto")
	store.AssertCalled(t, "EditTask", 1, "titi")
}

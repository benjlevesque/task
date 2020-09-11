package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/benjlevesque/task/types"
	bolt "go.etcd.io/bbolt"
)

var taskBucket = []byte("todos")

// CreateTask creates a task
func (s *Store) CreateTask(title string) (int, error) {
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return -1, err
	}

	var id int
	if err := db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b, err := tx.CreateBucketIfNotExists(taskBucket)
		if err != nil {
			return err
		}

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		newID, _ := b.NextSequence()
		id = int(newID)
		t := types.Task{
			ID:    id,
			Done:  false,
			Title: title,
		}

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(t.ID), buf)
	}); err != nil {
		return 0, err
	}

	err = db.Close()
	if err != nil {
		return -1, err
	}
	return id, nil
}

// ListType task list, done or not
type ListType string

const (
	// All done or not
	All ListType = "all"
	// Done all done tasks
	Done ListType = "done"
	// Undone all undone tasks
	Undone ListType = "undone"
)

// List all tasks
func (s *Store) List(listType ListType) ([]types.Task, error) {
	list := make([]types.Task, 0)
	if !fileExists(s.Path) {
		return list, nil
	}
	db, err := bolt.Open(s.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return nil, err
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task types.Task
			json.Unmarshal(v, &task)
			if (listType == Done && task.Done) || (listType == Undone && !task.Done) || listType == All {
				list = append(list, task)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return list, nil

}

// ToggleTask toggle a task state
func (s *Store) ToggleTask(id int, val bool) error {
	return editTask(s, id, func(task types.Task) types.Task {
		task.Done = val
		return task
	})
}

// EditTask edits a task
func (s *Store) EditTask(id int, title string) error {
	return editTask(s, id, func(task types.Task) types.Task {
		task.Title = title
		return task
	})
}

// GetTask toggle a task state
func (s *Store) GetTask(id int) (types.Task, error) {
	if !fileExists(s.Path) {
		return types.Task{}, fmt.Errorf("Task %d does not exist", id)
	}
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return types.Task{}, err
	}
	var task types.Task
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return getTask(b, itob(id), &task)
	}); err != nil {
		return types.Task{}, err
	}

	err = db.Close()
	if err != nil {
		return types.Task{}, err
	}
	return task, nil
}

func getTask(b *bolt.Bucket, id []byte, task *types.Task) error {
	taskBytes := b.Get(id)

	if taskBytes == nil {
		return fmt.Errorf("Task %d does not exist", btoi(id))
	}

	err := json.Unmarshal(taskBytes, &task)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask deletes a task
func (s *Store) DeleteTask(id int) error {
	if !fileExists(s.Path) {
		return fmt.Errorf("Task %d does not exist", id)
	}
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskBytes := b.Get(itob(id))
		if taskBytes == nil {
			return fmt.Errorf("Task %d does not exist", id)
		}
		return b.Delete(itob(id))
	}); err != nil {
		return err
	}

	return db.Close()
}

func editTask(s *Store, id int, editTask func(t types.Task) types.Task) error {
	if !fileExists(s.Path) {
		return fmt.Errorf("Task %d does not exist", id)
	}
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskBytes := b.Get(itob(id))

		if taskBytes == nil {
			return fmt.Errorf("Task %d does not exist", id)
		}

		var task types.Task
		err = json.Unmarshal(taskBytes, &task)

		newTask := editTask(task)

		// Marshal user data into bytes.
		buf, err := json.Marshal(newTask)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(task.ID), buf)
	}); err != nil {
		return err
	}

	return db.Close()

}

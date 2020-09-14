package db

import (
	"fmt"
	"time"

	"github.com/benjlevesque/task/types"
	bolt "go.etcd.io/bbolt"
)

var remindersBucket = []byte("reminders")

// ListReminders gets all reminders
func (s *Store) ListReminders() ([]types.Reminder, error) {
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return nil, err
	}
	// db.Update(func(tx *bolt.Tx) error {
	// 	return tx.DeleteBucket(remindersBucket)

	// })

	reminders := make([]types.Reminder, 0)
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(remindersBucket)
		bTask := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task types.Task
			err := getTask(bTask, k, &task)
			if err != nil {
				return err
			}
			t, err := time.Parse(time.RFC3339, string(v))
			reminders = append(reminders, types.Reminder{Task: task, Time: t})
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}

	err = db.Close()
	if err != nil {
		return nil, err
	}
	return reminders, nil

}

// GetReminder gets a reminder
func (s *Store) GetReminder(taskID int) (time.Time, error) {
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return time.Time{}, err
	}
	var t time.Time
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(remindersBucket)
		bytes := b.Get(itob(taskID))
		if bytes == nil {
			return fmt.Errorf("No reminder for task %d", taskID)
		}

		t, err = time.Parse(time.RFC3339, string(bytes))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return time.Time{}, err
	}

	err = db.Close()
	if err != nil {
		return time.Time{}, err
	}
	return t, nil

}

// SetReminder adds a reminder to a task
func (s *Store) SetReminder(taskID int, reminder time.Time) error {
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b, err := tx.CreateBucketIfNotExists(remindersBucket)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(taskID), []byte(reminder.Format(time.RFC3339)))
	}); err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

// DeleteReminder deletes a task
func (s *Store) DeleteReminder(id int) error {
	if !fileExists(s.Path) {
		return fmt.Errorf("Task %d does not exist", id)
	}
	db, err := bolt.Open(s.Path, 0600, nil)
	if err != nil {
		return err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(remindersBucket)
		taskBytes := b.Get(itob(id))
		if taskBytes == nil {
			return nil
		}
		return b.Delete(itob(id))
	}); err != nil {
		return err
	}

	return db.Close()
}

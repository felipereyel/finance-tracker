package database

import (
	"fintracker/internal/models"

	_ "modernc.org/sqlite"
)

type fakeDatabase struct {
	wallets map[string]models.Wallet
	assets  map[string]models.Asset
	price   map[string]models.Price
}

func NewFakeDatabaseRepo() fakeDatabase {
	return fakeDatabase{
		wallets: make(map[string]models.Wallet),
		assets:  make(map[string]models.Asset),
		price:   make(map[string]models.Price),
	}
}

// func (db *fakeDatabase) CreateTask(task models.Task) error {
// 	db.tasks[task.Id] = task
// 	return nil
// }

// func (db *fakeDatabase) ListTasksByOwner(ownerId string) ([]models.Task, error) {
// 	tasks := make([]models.Task, 0)

// 	for _, task := range db.tasks {
// 		if task.OwnerId == ownerId {
// 			tasks = append(tasks, task)
// 		}
// 	}

// 	return tasks, nil
// }

// func (db *fakeDatabase) RetrieveTaskById(taskId string) (models.Task, error) {
// 	task, ok := db.tasks[taskId]
// 	if !ok {
// 		return models.EmptyTask, sql.ErrNoRows
// 	}

// 	return task, nil
// }

// func (db *fakeDatabase) DeleteTask(taskId string) error {
// 	delete(db.tasks, taskId)
// 	return nil
// }

// func (db *fakeDatabase) UpdateTask(task models.Task) error {
// 	db.tasks[task.Id] = task
// 	return nil
// }

// func (db *fakeDatabase) InsertUser(user models.User) error {
// 	_, ok := db.users[user.ID]
// 	if ok {
// 		return fmt.Errorf("non unique email")
// 	}

// 	db.users[user.ID] = user
// 	return nil
// }

// func (db *fakeDatabase) RetrieveUserById(id string) (models.User, error) {
// 	user, ok := db.users[id]
// 	if !ok {
// 		return models.EmptyUser, sql.ErrNoRows
// 	}

// 	return user, nil
// }

// func (db *fakeDatabase) RetrieveUserByName(username string) (models.User, error) {
// 	for _, u := range db.users {
// 		if u.Username == username {
// 			return u, nil
// 		}
// 	}

// 	return models.EmptyUser, sql.ErrNoRows
// }

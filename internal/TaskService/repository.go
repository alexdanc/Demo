package TaskService

import "gorm.io/gorm"

type RequestBodyRepository interface {
	CreateTask(req RequestBody) error
	GetAllTasks() ([]RequestBody, error)
	GetTaskByID(id string) (RequestBody, error)
	UpdateTask(req RequestBody) error
	DeleteTaskByID(id string) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RequestBodyRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(req RequestBody) error {
	result := r.db.Create(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]RequestBody, error) {
	var tasks []RequestBody
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id string) (RequestBody, error) {
	var tas RequestBody
	err := r.db.First(&tas, id).Error
	return tas, err

}

func (r *TaskRepository) UpdateTask(req RequestBody) error {
	return r.db.Save(&req).Error
}

func (r *TaskRepository) DeleteTaskByID(id string) error {
	return r.db.Delete(&RequestBody{}, "id = ?", id).Error
}

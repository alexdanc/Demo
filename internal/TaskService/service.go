package TaskService

type RequestBodyService interface {
	CreatesTask(task RequestBody) (RequestBody, error)
	GetAllTasks() ([]RequestBody, error)
	GetTaskByID(id string) (RequestBody, error)
	UpdateTask(id, body string) (RequestBody, error)
	DeleteTaskByID(id string) error
}

type TaskService struct {
	repo RequestBodyRepository
}

func NewTaskService(r RequestBodyRepository) RequestBodyService {
	return &TaskService{repo: r}
}

func (s *TaskService) CreatesTask(task RequestBody) (RequestBody, error) {

	newTask := RequestBody{
		Task:    task.Task,
		Is_done: false,
	}

	err := s.repo.CreateTask(newTask)
	if err != nil {
		return RequestBody{}, err
	}

	return newTask, nil
}

func (s *TaskService) GetAllTasks() ([]RequestBody, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id string) (RequestBody, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(id, body string) (RequestBody, error) {
	tas, err := s.repo.GetTaskByID(id)
	if err != nil {
		return tas, err
	}
	tas.Task = body
	if err := s.repo.UpdateTask(tas); err != nil {
		return tas, err
	}
	return tas, nil
}

func (s *TaskService) DeleteTaskByID(id string) error {
	return s.repo.DeleteTaskByID(id)
}

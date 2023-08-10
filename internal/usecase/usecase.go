package usecase

type UseCase struct {
	TodoList CRUD
}

func New(TodoList CRUD) *UseCase {
	return &UseCase{
		TodoList: NewTodoList(TodoList),
	}
}
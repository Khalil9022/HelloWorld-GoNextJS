package todos

import (
	"github.com/khalil9022/HelloWorldGoNextJS/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]model.Todos, error)
	CreateTodos(task string) (model.Todos, error)
	UpdateTodos(id string) (model.Todos, error)
	DeleteTodos(id string) (model.Todos, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTodos() ([]model.Todos, error) {
	var todos []model.Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}

func (r *repository) CreateTodos(task string) (model.Todos, error) {
	todo := model.Todos{
		Task: task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return model.Todos{}, res.Error
	}

	return todo, nil
}

func (r *repository) UpdateTodos(id string) (model.Todos, error) {
	todos := model.Todos{
		Done: true,
	}

	err := r.db.Model(&todos).Where("id", id).Updates(&todos).Error
	if err != nil {
		return model.Todos{}, err
	}
	return todos, nil
}

func (r *repository) DeleteTodos(id string) (model.Todos, error) {
	todos := model.Todos{}

	err := r.db.Model(&todos).Where("id", id).Delete(&todos).Error
	if err != nil {
		return model.Todos{}, err
	}
	return model.Todos{}, nil
}

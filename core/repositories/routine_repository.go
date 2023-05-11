package repositories

import "vita/core/entities"

type RoutineRepository interface {
	GetRoutine(key string) (entities.Routine, error)
	GetAllRoutines() ([]entities.Routine, error)
	CreateRoutine(routine entities.Routine) error
	UpdateRoutine(routine entities.Routine) error
	DeleteRoutine(key string) error
}

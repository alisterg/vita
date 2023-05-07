package ports

import "vita/core/entities"

type RoutineRepository interface {
	GetRoutine(key string) (entities.Routine, error)
	GetAllRoutines() ([]entities.Routine, error)
	CreateRoutine(key string) error
	UpdateRoutine(key string, entryTypes []string) error
	DeleteRoutine(key string) error
}

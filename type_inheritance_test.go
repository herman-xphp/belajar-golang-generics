package belajar_golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employe interface {
	GetName() string
}

func GetName[T Employe](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Ucok", GetName[Manager](&MyManager{Name: "Ucok"}))
	assert.Equal(t, "Ucok", GetName[VicePresident](&MyVicePresident{Name: "Ucok"}))
}

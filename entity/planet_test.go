package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBook(t *testing.T) {
	b, err := NewPlanet("", "", "")
	assert.Nil(t, err)
	assert.Equal(t, b.Name, "")
	assert.NotNil(t, b.ID)
}

func TestPlanetValidate(t *testing.T) {

}

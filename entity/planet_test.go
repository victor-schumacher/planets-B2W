package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlanet(t *testing.T) {
	b, err := NewPlanet("Tatooine", "arid", "desert")
	assert.Nil(t, err)
	assert.Equal(t, b.Name, "Tatooine")
	assert.NotNil(t, b.ID)
}

func TestPlanetValidate(t *testing.T) {
	type table struct {
		name    string
		climate string
		terrain string
		want    error
	}

	tt := []table{
		{
			name:    "Tatooine",
			climate: "arid",
			terrain: "",
			want:    ErrInvalidEntity,
		},
		{
			name:    "",
			climate: "arid",
			terrain: "desert",
			want:    ErrInvalidEntity,
		},
	}
	for _, tc := range tt {
		_, err := NewPlanet(tc.name, tc.climate, tc.terrain)
		assert.Equal(t, err, tc.want)
	}
}

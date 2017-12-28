package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKotak(t *testing.T) {
	p := NewKotak(4, 8)

	if p.IsPersegi() == true {
		t.Error("Seharusnya false")
	}

	if p.Luas() != 32 {
		t.Error("Seharusnya 32")
	}

	if p.Keliling() != 24 {
		t.Error("seharusnya 24")
	}

	persegi := NewKotak(4, 4)

	if persegi.IsPersegi() == false {
		t.Error("seharusnya true")
	}
}

func TestKotakPakeAssert(t *testing.T) {
	assert := assert.New(t)
	p := NewKotak(4, 8)

	assert.Equal(p.IsPersegi(), false, "seharusnya false")
	assert.Equal(p.Luas(), 32, "seharusnya 32")
	assert.Equal(p.Keliling(), 24, "seharusnya 24")

	persegi := NewKotak(4, 4)
	assert.Equal(persegi.Luas(), 16, "seharusnya 16")
	assert.Equal(persegi.Keliling(), 16, "seharusnya 16")
	assert.Equal(persegi.IsPersegi(), true, "seharusnya true")
}

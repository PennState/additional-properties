package json

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArraysAreDereferenced(t *testing.T) {
	assert := assert.New(t)

	a := []string{"Test"}
	p := &a
	pp := &p
	ppp := &pp

	assert.Equal(reflect.String, dereferencedKind(a))
	assert.Equal(reflect.String, dereferencedKind(p))
	assert.Equal(reflect.String, dereferencedKind(pp))
	assert.Equal(reflect.String, dereferencedKind(ppp))
}

func TestPointersToPointersAreDereferenced(t *testing.T) {
	assert := assert.New(t)

	s := "Test"
	p := &s
	pp := &p
	ppp := &pp

	assert.Equal(reflect.String, dereferencedKind(p))
	assert.Equal(reflect.String, dereferencedKind(pp))
	assert.Equal(reflect.String, dereferencedKind(ppp))
}
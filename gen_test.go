package main

import (
	"fmt"
	"testing"

	v1 "github.com/daotl/protoc-gen-go-enum-extractor/gen/proto/v1"
	"github.com/magiconair/properties/assert"
)

func TestGen(t *testing.T) {
	enum := v1.PropertyChangeOp_PROPERTY_CHANGE_OP_ADD
	fmt.Print(enum.String())
	assert.Equal(t, enum.ExtractValue(), "ADD")

	enum.FromValue("ASSIGN")
	assert.Equal(t, enum, v1.PropertyChangeOp_PROPERTY_CHANGE_OP_ASSIGN)
}

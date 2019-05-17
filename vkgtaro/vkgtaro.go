package vkgtaro

import (
	"os"

	"github.com/vkg/taro/printer"
)

var _ Vkgtaro = (*vkgtaroImpl)(nil)

// Vkgtaro is vkg service
type Vkgtaro interface {
	Move()
}

type vkgtaroImpl struct {
	printer *printer.Printer
}

// New returns vkgtaro
func New() Vkgtaro {
	printer := printer.New(os.Stdout)
	return &vkgtaroImpl{printer: printer}
}

// Move starts moving
func (v *vkgtaroImpl) Move() {
	v.printer.PrintAndMove(printer.NewSource(photo))
}

package data

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

const (
	Size     = 5000
	MinSize  = 100
	StepSize = 100
	MaxSize  = 1000
)

func Initialize(collection collections.Collection[types.String]) {
	for i := MinSize; i <= MaxSize; i = i + StepSize {
		collection.Add(types.String(fmt.Sprint(i)))
	}
}

func InitializeReverse(collection collections.Collection[types.String]) {
	for i := MinSize; i <= MaxSize; i = i + StepSize {
		collection.Add(types.String(fmt.Sprint(MaxSize - i)))
	}
}

func InitializeWith(collection collections.Collection[types.String], minSize int, stepSize int, maxSize int) {
	for i := minSize; i <= maxSize; i = i + stepSize {
		collection.Add(types.String(fmt.Sprint(i)))
	}
}

func InitializeReverseWith(collection collections.Collection[types.String], minSize int, stepSize int, maxSize int) {
	for i := maxSize; i >= minSize; i = i - stepSize {
		collection.Add(types.String(fmt.Sprint(i)))
	}
}

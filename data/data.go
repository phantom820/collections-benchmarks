package data

import (
	"fmt"

	"github.com/phantom820/collections"
	"github.com/phantom820/collections/types"
)

func Initialize(collection collections.Collection[types.String], size int) {
	for i := 1; i <= size; i = i + 1 {
		collection.Add(types.String(fmt.Sprint(i)))
	}
}

func InitializeReverse(collection collections.Collection[types.String], size int) {
	for i := size; i >= 1; i = i - 1 {
		collection.Add(types.String(fmt.Sprint(i)))
	}
}

package storage_test

import (
	"testing"

	"github.com/andrei-pokhila/insights-dashboards/src/storage"
)

func TestClickhouse(t *testing.T) {
	storage.NewConnection()
}

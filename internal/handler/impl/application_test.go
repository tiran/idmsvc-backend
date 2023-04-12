package impl

import (
	"testing"

	"github.com/hmsidm/internal/config"
	"github.com/hmsidm/internal/metrics"
	"github.com/hmsidm/internal/test"
	"github.com/hmsidm/internal/test/mock/interface/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewHandler(t *testing.T) {
	sqlMock, gormDB, err := test.NewSqlMock(&gorm.Session{SkipHooks: true})
	inventoryMock := client.NewHostInventory(t)
	require.NoError(t, err)
	require.NotNil(t, sqlMock)
	require.NotNil(t, gormDB)
	assert.Panics(t, func() {
		NewHandler(nil, nil, nil, nil)
	})
	assert.Panics(t, func() {
		NewHandler(&config.Config{}, nil, nil, nil)
	})
	assert.NotPanics(t, func() {
		NewHandler(&config.Config{}, gormDB, &metrics.Metrics{}, inventoryMock)
	})
}
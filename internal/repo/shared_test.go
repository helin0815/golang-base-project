package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal/query"
)

func newMockQuery(t *testing.T) (q *query.Query, sqlMock sqlmock.Sqlmock) {
	sqlDB, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(t, err)

	q = query.Use(gormDB)
	return
}

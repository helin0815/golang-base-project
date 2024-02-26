package repo

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/entities"
)

func TestGet(t *testing.T) {
	q, mock := newMockQuery(t)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`id` = ? ORDER BY `books`.`id` LIMIT 1")).
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	bookRepo := NewBook(q)
	book, err := bookRepo.Get(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, &entities.Book{Id: 1}, book)
	assert.NoError(t, mock.ExpectationsWereMet())
}

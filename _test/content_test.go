package domain_test

import (
	"testing"
	"time"

	"github.com/mgarcia88/maycms/domain"
	"github.com/stretchr/testify/require"
)

func ShouldReturnErrorWhenUUIDIsNotValid(t *testing.T) {
	content := domain.NewContent()
	content.ID = "aa"
	content.Title = "Post de teste"
	content.ContentText = "Lorem ipsum"
	content.Status = "Ativo"
	content.CreatedAt = time.Now()
	content.UpdatedAt = time.Now()

	err := content.Validate()
	require.Error(t, err)
}

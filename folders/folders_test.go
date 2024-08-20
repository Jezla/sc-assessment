package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("emptyTest", func(t *testing.T) {
		// Test case 1: Empty request
		req := &folders.FetchFolderRequest{}
		resp, err := folders.GetAllFolders(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("validTest", func(t *testing.T) {
		// Test case 2: Request with valid OrgID
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}
		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Folders)
	})

	t.Run("nilTest", func(t *testing.T) {
		// Test case 3: Request with invalid OrgID
		req := &folders.FetchFolderRequest{OrgID: uuid.Nil}
		resp, err := folders.GetAllFolders(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

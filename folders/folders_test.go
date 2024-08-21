package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	// Use GenerateData to get sample folder data
	folderData := folders.GenerateData()

	// Test case: Successful retrieval with matching OrgID
	t.Run("successful retrieval", func(t *testing.T) {
		orgID := folderData[0].OrgId
		req := &folders.FetchFolderRequest{OrgID: orgID}

		expectedFolders := []*folders.Folder{}
		for _, folder := range folderData {
			if folder.OrgId == orgID {
				expectedFolders = append(expectedFolders, folder)
			}
		}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.ElementsMatch(t, expectedFolders, resp.Folders)
	})

	// Test case: No folders match the nil OrgID
	t.Run("no folders match", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: uuid.Nil}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Folders, 0)
	})
}

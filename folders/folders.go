package folders

import (
	"github.com/gofrs/uuid"
)

/* Improvements:
- There are three loops between the two functions which all essentially do the same thing.
They create a list of Folder struct pointers to be returned in the FetchFolderResponse.
Only the one loop in FetchAllFoldersByOrgID is required for this. As part of this, some
variables are unnecessarily repeated, such as fs, fp and f, which can be combined.
- Error handling is not implemented in GetAllFolders. */

// This function appears to fetch all folders given a FetchFolderRequest, which contains an orgID.
// It returns two variables: a response and a possible error for error handling.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	/* Initially, these variables are unused. Suggested improvements:
	Use err correctly for error handling,
	f1 is unnecessary now but may be useful for pagination,
	combine the logic for fs, fp and f into one. */
	var (
		err error
		fs  []*Folder
	)

	// Improvement: correct error handling when fetching folders.
	fs, err = FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// Improvement: declare and initialize in one step
	ffr := &FetchFolderResponse{Folders: fs}
	return ffr, nil
}

// This function appears to fetch all folders for a given organization ID.
// It returns two variables: a list of Folder struct pointers and a possible error for error handling.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}

package folders

import (
	"encoding/base64"
	"errors"
	"strconv"

	"github.com/gofrs/uuid"
)

/* Short Explanation: the Requests and Responses have been updated to hold a Token,
decoded from a base64 string to get the current index from which the data fetching should start.
If no token is provided, as in the example, fetching starts from 0.
If there are more items to fetch after the current request, a new token is generated.
Page size can be adjusted to change the number of items returned in each response.
*/

// Set a fixed page size for pagination.
const pageSize = 2

// FetchFolderRequest represents the request for fetching folders, including pagination.
type FetchFolderRequestPaginated struct {
	OrgID uuid.UUID
	Token string
}

// FetchFolderResponse represents the response with a list of folders and a pagination token.
type FetchFolderResponsePaginated struct {
	Folders []*Folder
	Token   string
}

// This function fetches all folders with a request, implemented with pagination.
// It returns two variables: a response and a possible error for error handling.
func GetAllFoldersPaginated(req *FetchFolderRequestPaginated) (*FetchFolderResponsePaginated, error) {
	// Fetch folders by OrgID
	folders, err := FetchAllFoldersByOrgIDPaginated(req.OrgID)
	if err != nil {
		return nil, err
	}

	// Parse the token to get the start index
	startIndex := 0
	if req.Token != "" {
		decodedToken, err := base64.StdEncoding.DecodeString(req.Token)
		if err != nil {
			return nil, errors.New("invalid token")
		}
		startIndex, err = strconv.Atoi(string(decodedToken))
		if err != nil {
			return nil, errors.New("invalid token format")
		}
	}

	// Calculate end index based on the fixed page size
	endIndex := startIndex + pageSize
	if endIndex > len(folders) {
		endIndex = len(folders)
	}

	// Get the paginated slice of folders
	paginatedFolders := folders[startIndex:endIndex]

	// Generate the next token if more data is available
	var nextToken string
	if endIndex < len(folders) {
		nextToken = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(endIndex)))
	}

	// Prepare the response
	response := &FetchFolderResponsePaginated{
		Folders: paginatedFolders,
		Token:   nextToken,
	}

	return response, nil
}

// This function fetches all folders for a given organization ID.
// It returns two variables: a list of Folder struct pointers and a possible error for error handling.
func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	var resFolder []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}

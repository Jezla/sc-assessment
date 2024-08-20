package folders

import (
	"errors"

	"github.com/gofrs/uuid"
)

//Suggested improvements:
//1. Remove unused variables
//2. Simplify the process of creating the slice of Folder pointers (unnecessarily complicated)
//3. Use more descriptive variable names
//4. Add error handling for FetchAllFoldersByOrgID in GetAllFolders

var ErrInvalidOrgID = errors.New("invalid OrgID")

// GetAllFolders takes a FetcHFolderRequest as input
// Calls FetchAllFoldersByOrgID with the provided OrgID
// Creates a new slice of Folder pointers
// Returns a FetchFolderResponse containing the Folder pointers.
func GetAllFolders(request *FetchFolderRequest) (*FetchFolderResponse, error) {
	if request.OrgID == uuid.Nil {
		return nil, ErrInvalidOrgID
	}
	folders, err := FetchAllFoldersByOrgID(request.OrgID)
	if err != nil {
		return nil, err
	}
	return &FetchFolderResponse{Folders: folders}, nil
}

// Takes an orgID as input
// Calls GetSampleData() to retrieve folders
// Filters folders based on the provided orgID
// Returns a slice of matching Folder pointers
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

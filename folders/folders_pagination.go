package folders

import (
	"encoding/base64"
	"strconv"

	"github.com/gofrs/uuid"
)

// 1. I have added a new Token field to the NewFetchFolderRequest struct and NewFetchFolderResponse struct.
//2. I set a PageSize to determine how many Folders are returned in each page.
//3. In PaginationGetAllFolders:
//   a. I fetch all the folders that match the provided OrgID.
//   b. I check if a token is provided, which would be used to get the starting index.
//   c. I slice the folders array based on the starting index and the PageSize.
//   d. I generate a new token (if there are more) and encode the next starting index.
//4. Token is base64 encoded.
//5. If there are no more folders to return, the token is empty again.

/**
//The main.go file can be updated to use the new PaginationGetAllFolders function as shown below:
	req := &NewFetchFolderRequest{OrgID: someOrgID, Token: ""}
	resp, err := PaginationGetAllFolders(req)
	if err != nil {
	// Handle error
	}

//Can add loop until no more Token is returned
	nextReq := &NewFetchFolderRequest{OrgID: someOrgID, Token: resp.Token}
*/

const PageSize = 2 // Adjust this value as needed

type NewFetchFolderRequest struct {
	OrgID uuid.UUID
	Token string
}

type NewFetchFolderResponse struct {
	Folders []*Folder
	Token   string
}

func PaginationGetAllFolders(req *NewFetchFolderRequest) (*NewFetchFolderResponse, error) {
	allFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	startIndex := 0
	if req.Token != "" {
		decodedToken, err := base64.StdEncoding.DecodeString(req.Token)
		if err != nil {
			return nil, err
		}
		startIndex, err = strconv.Atoi(string(decodedToken))
		if err != nil {
			return nil, err
		}
	}

	endIndex := startIndex + PageSize
	if endIndex > len(allFolders) {
		endIndex = len(allFolders)
	}

	paginatedFolders := allFolders[startIndex:endIndex]

	var nextToken string
	if endIndex < len(allFolders) {
		nextToken = base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(endIndex)))
	}

	return &NewFetchFolderResponse{
		Folders: paginatedFolders,
		Token:   nextToken,
	}, nil
}

func PaginationFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	allFolders := GetSampleData()

	var matchingFolders []*Folder
	for _, folder := range allFolders {
		if folder.OrgId == orgID {
			matchingFolders = append(matchingFolders, folder)
		}
	}

	return matchingFolders, nil
}

package models

import (
	"chi-domain-go/models/utils"
	"time"
)

type TodoRequestBody struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	IsDone      bool   `json:"is_done" bson:"is_done"`
}

type TodoResponse struct {
	ID          string    `json:"id" bson:"id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	IsDone      bool      `json:"is_done" bson:"is_done"`
	IsArchive   bool      `json:"is_archive" bson:"is_archive"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type TodoDBPayload struct {
	ID          string `json:"id" bson:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	IsDone      bool   `json:"is_done" bson:"is_done"`
	IsArchive   bool   `json:"is_archive" bson:"is_archive"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	UpdatedAt   int64  `json:"updated_at" bson:"updated_at"`
}

func (dbPayload TodoDBPayload) ToResponse() TodoResponse {
	responseData := new(TodoResponse)
	utils.MergeData(dbPayload, responseData)

	responseData.CreatedAt = utils.TimestampToTime(dbPayload.CreatedAt)
	responseData.UpdatedAt = utils.TimestampToTime(dbPayload.UpdatedAt)
	return *responseData
}

type TodoDBPayloadList []TodoDBPayload

func (dbList TodoDBPayloadList) Size() int {
	return len(dbList)
}

func (dbList TodoDBPayloadList) IsEmpty() bool {
	return dbList.Size() == 0
}

func (dbList TodoDBPayloadList) ToResponseList() []TodoResponse {
	output := []TodoResponse{}
	for _, dbPayload := range dbList {
		output = append(output, dbPayload.ToResponse())
	}
	return output
}

type TodoSelector struct {
	ID          *string `bson:"id,omitempty"`
	UserID      *string `bson:"user_id,omitempty"`
	Title       *string `bson:"title,omitempty"`
	Description *string `bson:"description,omitempty"`
	IsDone      *bool   `bson:"is_done,omitempty"`
	IsArchive   *bool   `bson:"is_archive,omitempty"`
	CreatedAt   *int64  `bson:"created_at,omitempty"`
	UpdatedAt   *int64  `bson:"updated_at,omitempty"`
}

func (ts *TodoSelector) SetID(id string) *TodoSelector {
	ts.ID = &id
	return ts
}

func (ts *TodoSelector) SetUserID(userID string) *TodoSelector {
	ts.UserID = &userID
	return ts
}

func (ts *TodoSelector) SetIsArchive(isArchive bool) *TodoSelector {
	ts.IsArchive = &isArchive
	return ts
}

package entities

type CreateFilesRequest struct {
	CreatorID int64     `json:"-"`
	ObjectID  int64     `json:"object_id,omitempty" binding:"required"`
	Urls      []*string `json:"urls,omitempty" binding:"required"`
}

type DeleteFilesRequest struct {
	ObjectID int64   `json:"object_id,omitempty" binding:"required"`
	IDs      []int64 `json:"ids,omitempty" binding:"required"`
}
type ListFileByObjectID struct {
	ID       int64  `json:"id,omitempty"`
	ObjectID int64  `json:"object_id,omitempty"`
	Url      string `json:"url,omitempty"`
}

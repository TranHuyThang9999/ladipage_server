package entities

type CreateFilesRequest struct {
	CreatorID int64     `json:"-"`
	ObjectID  int64     `json:"object_id,omitempty" binding:"required"`
	Url       []*string `json:"url,omitempty" binding:"required"`
}

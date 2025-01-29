package entities

type CreateChannelDescRequest struct {
	Name      string `json:"name,omitempty"`
	CreatorID int64  `json:"-"`
	ImageDesc string `json:"avatar,omitempty"`
}

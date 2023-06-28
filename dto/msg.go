package dto

import "github.com/gookit/validate"

type Message struct {
	SourceID string `json:"source_id,omitempty" validate:""`
	Content  string `json:"content,omitempty"`
	TargetID string `json:"target_id,omitempty"`
	CreateAt int64  `json:"create_at,omitempty"`
}

func (m *Message) Validate() error {
	v := validate.Struct(m)
	if !v.Validate() {
		return v.Errors.ErrOrNil()
	}
	return nil
}

package dto

// validate
type CreateUser struct {
	// ID       string `json:"id,omitempty" validate:"required,len>"`
	// Nickname string `json:"nickname,omitempty" binding:"required,min=3,max=10,trim"`
	// Name     string `json:"name,omitempty" binding:"required,min=3,max=10"`
	Nickname string `json:"nickname,omitempty" validate:"required|minLen:3|maxLen:10" filter:"trim"`
	Name     string `json:"name,omitempty" validate:"required|minLen:3|maxLen:10" filter:"trim"`
	Password string `json:"password,omitempty" validate:"required|minLen:3|maxLen:10" filter:"trim"`
}

type Login struct {
	Name       string `json:"name" filter:"trim" validate:"required|min:3|maxLen:10"`
	Password   string `json:"password" filter:"trim" validate:"required|minLen:3|maxLen:10"`
	DeviceType string `json:"deviceType filter:"trim" validate:"required"`
}

// func (u *User) Validate() error {
// 	v := validate.Struct(u)
// 	if !v.Validate() {
// 		return v.Errors.ErrOrNil()
// 	}
// 	log.Info(u)
// 	return nil
// }

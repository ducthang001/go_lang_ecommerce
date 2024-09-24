package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid
)

// message
var msg = map[int]string{
	ErrCodeSuccess: "Success",
	ErrCodeParamInvalid: "Email is invalid",
}
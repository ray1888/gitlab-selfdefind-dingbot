package response

type DingResponse interface {
	IsSuccess() bool
	Failed() error
}

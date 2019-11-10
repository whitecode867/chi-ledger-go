package standard

type Response interface {
	GetCode() int
	GetPayload() interface{}
	IsError() bool
}

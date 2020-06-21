package main

type Request struct {
	Method string
	Location string
	Version string
	Headers map[string]string
	Body string
}

func (req *Request) SetHeader(key string, value string) {
	req.Headers[key] = value
}

func (req *Request) GetHeader(key string) string {
	return req.Headers[key]
}

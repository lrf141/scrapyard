package main

import (
	"strings"
)

const CRLF = "\r\n"

func parse(raw string) *Request {

	rows := strings.Split(raw, CRLF)
	status := strings.Split(rows[0], " ")
	bodyPos := 0
	headers := map[string]string{}
	for index, row := range rows[1:] {
		if row == "" {
			bodyPos = index
			break
		}
		header := strings.Split(row, ": ")
		if len(header) > 1 {
			headers[header[0]] = header[1]
		}
	}
	body := ""
	for _, row := range rows[bodyPos+1:] {
		body += row
	}
	return &Request{
		Method: status[0],
		Location: status[1],
		Version: status[2],
		Headers: headers,
		Body: body,
	}
}

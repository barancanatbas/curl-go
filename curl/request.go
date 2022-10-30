package curl

import "encoding/json"

type Request struct {
	Method     string
	Url        string
	Header     map[string]string
	Body       string
	OutputName string
}

func (req *Request) FindParams(args []string) {
	if len(args) == 0 {
		return
	}
	for index, arg := range args {
		switch arg {
		case "-h":
			var header map[string]string
			json.Unmarshal([]byte(args[index+1]), &header)
			req.Header = header
		case "-b":
			req.Body = args[index+1]
		case "-get":
			req.Method = "GET"
			req.Url = args[index+1]
		case "-post":
			req.Method = "POST"
			req.Url = args[index+1]
		case "-put":
			req.Method = "PUT"
			req.Url = args[index+1]
		case "-o":
			req.OutputName = args[index+1]
		}
	}

	if req.Method == "" {
		req.Method = "GET"
	}
}

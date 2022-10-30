# curl-go

Curl package based on Golang. At a basic level you can create http requests.

The package currently supports get post and put requests.
It prints the response from the request directly to the screen, but you can get your output as a file.

Usage;
```bash
go build
```
```bash
./curl-go {request method} {request url}
```

- (-h) : We specify the headers of the request to be sent.
- (-b) : We specify the body of the request to be sent.
- (-get) : create a get request.
- (-post) : create a post request.
- (-put) : create a put request.
- (-o) : creates an output file named after the value.

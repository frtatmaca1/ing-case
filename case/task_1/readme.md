## Task - API calls
### Server specification
Inside this folder is a simple HTTP server written in Go.
The server handles one endpoint with two HTTP methods:

Endpoint: `person` for GET requests, with response:
```json
{
  "name": "William",
  "age": 42
}
```
The return code should be `200`

And the same endpoint `person` for POST requests. For example, with request:
```json
{
  "person": {
    "name": "William",
    "age": 42
  },
  "greeting": "Hello William (42)"
}
```
The return code should be `200`

### Task
Perform a `GET` request on endpoint `person` and parse the JSON information.
Use the JSON information to construct a greeting containing the name and age.

Post the greeting, as JSON, 2 times to the endpoint `person`. The server is able to 
process requests in parallel. The server should have two log entries of the greeting 
that was posted.

### Using the server
Navigate to the folder and build the Dockerfile. Alternatively, you can also build the go
file directly; there are no dependencies. E.g.
```bash
cd /path/to/folder
docker build -t example-task .
```

After that run it:
```bash
docker run -p 8080:8080 -it --rm example-task
```

The server is not an example of good quality code, it's only there to provide some endpoints.

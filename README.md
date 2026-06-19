# Pucora martian

The `velonetics-martian` package integrates the martian project into the Pucora framework.

## How to use it

Add your martian DSL definition under the "github.com/pucora/velonetics-martian" namespace of the backend section of the config file

```
"extra_config": {
  "github.com/pucora/velonetics-martian": {}
}
```

More details here: https://github.com/google/martian#modifiers-all-the-way-down

Check the [example](https://github.com/pucora/velonetics-martian/tree/master/example) folder for a complete demo.

## Example

* Build and run pucora edition with martian

```
$ cd example
$ go install ./...
$ GOPATH/bin/example -c example/pucora.json -d -p 8000
```

* Send a request to the configured endpoint

```
$ curl -i 127.0.0.1:8000/supu
  HTTP/1.1 200 OK
  Cache-Control: public, max-age=0
  Content-Type: application/json; charset=utf-8
  X-Pucora: Version undefined
  X-Pucora-Completed: true
  Date: Tue, 17 Jul 2018 11:56:33 GMT
  Content-Length: 19

  {"msg":"you rock!"}
 ```

And check the logs of the Pucora: the request modifiers have done their job!
See how `{"msg":"you rock!"}` was added to the payload.

```
[PUCORA] ▶ DEBUG config: {[0xc4201839a0] 3s 3.6µs [] 8000 2  map[github_com/pucora/velonetics-gologging:map[prefix:[PUCORA] stdout:true level:DEBUG syslog:false]] 0s 0s 0s 0s false false 0 250 0s 0s 0s 0s 0s 0s false <nil> true 1}
[PUCORA] ▶ DEBUG Debug enabled
[PUCORA] ▶ DEBUG Method: GET
[PUCORA] ▶ DEBUG URL: /__debug/supu
[PUCORA] ▶ DEBUG Query: map[]
[PUCORA] ▶ DEBUG Params: [{param /supu}]
[PUCORA] ▶ DEBUG Headers: map[User-Agent:[Pucora Version undefined] Content-Length:[19] Content-Type:[] X-Forwarded-For:[127.0.0.1] Accept-Encoding:[gzip]]
[PUCORA] ▶ DEBUG Body: {"msg":"you rock!"}
[GIN] | 200 |     313.798µs |       127.0.0.1 | GET      /__debug/supu
[GIN] | 200 |    1.556445ms |       127.0.0.1 | GET      /supu
```

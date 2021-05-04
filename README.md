# HTTP Interface Exercise in Golang

The code is for the Go-HTTP Interface project from the Udemy course Go: The Complete Developer's Guide. The focus of the exercise is to show how interfaces work in Golang.

Here, a receiver function `Write` is implemented to print out the body content of the received http request. This function is called when the `io.Copy(lw, resp.Body)` is executed, as it expects a `Writer` interface, as well as a `Reader` interface [(link)](https://golang.org/pkg/io/#Copy). The `Writer`interface is used as the `Write` function has the same expected input `(p []byte)` as well as same return values `(n int, err error)` as the `Writer` interface [(link)](https://golang.org/pkg/io/#Writer):

```bash
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

## Output

HTML Body content of the request webpage, as well as a string containing the amount of bytes processed.

### Example

```bash
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="sv"><head><meta content="text/html;
[...]
</script>        </body></html>
Just wrote these many bytes 13504
```

## How-To

### Build and Run the Application

```bash
go build main.go
./main
```

### Test

No tests were developed for this application

## Notes

N/A

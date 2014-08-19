# go-minimal-log [![wercker status](https://app.wercker.com/status/5dd4e07683beb9987a54af5e0fe653ce/s "wercker status")](https://app.wercker.com/project/bykey/5dd4e07683beb9987a54af5e0fe653ce)

Simple go logger similar to Perl's [Log::Minimal](http://search.cpan.org/dist/Log-Minimal/)

## Usage

```go
import "github.com/typester/go-minimal-log"

log.Debugf("%s", "foo") // 2014/08/19 09:16:23 example.go:12: [debug] foo

```

### Change log level

```go
log.LogLevel = log.INFO
```

Available levels are:

- MUTE
- DEBUG
- INFO
- WARN
- CRITICAL
- ERROR

# Author

Daisuke Murase (typester)

# License

MIT

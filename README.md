## Hub-Search

search repositories matched with the given words and options in github.com

### Installation

```
$ go get github.com/skang99/hub-search
```

### Example

To see help message of usage, 
```
$ hub-search -h

  Usage:
    hub-search <query>... [--lang=<type>] [--sort=<method>] [--order=<type>] [--text] [--down]
    hub-search -d | --down
    hub-search -t | --text
    hub-search -h | --help
    hub-search --version

  Options:
    --lang=<type>    implemenation language, default:ALL
    --sort=<method>  sort field, default: best match [stars|forks|updated]
    --order=<type>   the sort order, default: desc [asc|desc]
    -d, --down       download packages searched
    -t, --text       normal text without esc chars
    -h, --help       output help information
    -v, --version    output version
```

If you search repos related with "webrtc server" and written in Go,
```
$ hub-search --lang=go webrtc server
```

If you want to download repos searched, add -d option
```
$ hub-search --lang=go webrtc server -d
```


### Reference

- [tj/go-search](http://github.com/tj-go-search)
- Github Developer API Search](https://developer.github.com/v3/search/)
- [JSON-to-Go](http://mholt.github.io/json-to-go/)
- [ANSI escape code](http://en.wikipedia.org/wiki/ANSI_escape_code)

### LICENSE

MIT


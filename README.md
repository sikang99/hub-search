## Hub-Search

Search repositories matched with the given words and options in github.com


## Installation

```
$ go get github.com/sikang99/hub-search/...
```
optionally and
```
$ cd $GOPATH/sikang99/hub-search
$ go install
```


### Usage

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

To search repos matched with the given words, and written in go
```
$ hub-search --lang=go restful framework
```

To list up top number of items, use `--list=<number>` option
```
$ hub-search web --list=10
```

To get (download) the repos matched, use `--down` option
```
$ hub-search --lang=go restful framework -d
```

To display the list in plain text without color code, add `--text` option
```
$ hub-search --lang=go restful framework -d --text
```

For gophers, to use the utility simply, `alias gog="hub-search --lang=go"`


### Reference

- [tj/go-search](http://github.com/tj-go-search)
- [Github Developer API Search](https://developer.github.com/v3/search/)
- [JSON-to-Go](http://mholt.github.io/json-to-go/)
- [ANSI escape code](http://en.wikipedia.org/wiki/ANSI_escape_code)

### LICENSE

MIT


## Hub-Search

<<<<<<< HEAD
Search repositories matched with the given words and options in github.com


## Installation

```
$ go get github.com/skang99/hub-search

or
$ cd $GOPATH/sikang99/hub-search
$ go install
```


### Usage

First, just type the command with -h option
=======
search repositories matched with the given words and options in github.com

### Installation

```
$ go get github.com/skang99/hub-search
```

### Example

To see help message of usage, 
>>>>>>> 50e667568d9989c0b3f63ae4e3753011e117a13b
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

<<<<<<< HEAD
To search repos matched the given words, and written in go
```
$ hub-search --lang=go restful framework
```

To list up top number of items
```
$ hub-search web --list=10
```

To get (download) the repos matched
```
$ hub-search --lang=go restful framework -d
=======
If you search repos related with "webrtc server" and written in Go,
```
$ hub-search --lang=go webrtc server
```

If you want to download repos searched, add -d option
```
$ hub-search --lang=go webrtc server -d
>>>>>>> 50e667568d9989c0b3f63ae4e3753011e117a13b
```


### Reference

- [tj/go-search](http://github.com/tj-go-search)
- Github Developer API Search](https://developer.github.com/v3/search/)
- [JSON-to-Go](http://mholt.github.io/json-to-go/)
- [ANSI escape code](http://en.wikipedia.org/wiki/ANSI_escape_code)

<<<<<<< HEAD

=======
>>>>>>> 50e667568d9989c0b3f63ae4e3753011e117a13b
### LICENSE

MIT


### Modified by gdaisukesuzuki (3/29/2018)

newly added 4 "methods"

`SetEmpty()` ... modified from `Set()`

`SetEmptyPath()` ... modified from `SetPath()`

`SetEmptyIndex()` ... modified from `SetIndex()`

`AddEmptyIndex()` ... modified from `AddIndex()`

which replace/add JSON's emplty element `{}` instead of some `value`

### Uploaded by gdaisukesuzuki (3/23/2018)

SampleCode (`test01.go`) Handling Array of JSON/Struct Type.


### Modified by gdaisukesuzuki (3/21/2018)

This repository is folked from bitly, and changed by gdaisukesuzuki 

so that "ARRAY / SLICE" in JSON eaily handled using go-simplejson

BY newley added 5 "METHODS"

`SizeIndex()` ... returns a size of Array/Slice

`DelIndex()` ... deletes one element of Array/Slice

`ZeroIndex()` ... creats or resets Array/Slice consisting zero element.

`AddIndex()` ... adds one element of Array/Slice     ("Size" increased)

`SetIndex()` ... replaces one element of Array/Slice ("Size" not changed)


~~### Issues (as of 3/21/2018 18:00JST)~~

~~SetIndex can only set "Atom" type (ie int, string) but not set "Struct / Json" type.~~

~~eg. js.Get("Top").Get("hoge").SetIndex("poi",1,'{"a":1,"b":"B"}')~~


### go-simplejson

a Go package to interact with arbitrary JSON

~~[![Build Status]~~~~(https://secure.travis-ci.org/bitly/go-simplejson.png)~~~~](http://travis-ci.org/bitly/go-simplejson)~~

### Importing

1. do on Terminal

    `% go get github.com/gdaisukesuzuki/go-simplejson`

2. And in the source code, set as follows

`import (

   "github.com/gdaisukesuzuki/go-simplejson"

    ...
    
 )`
 

### Documentation

Visit the docs on [gopkgdoc](http://godoc.org/github.com/bitly/go-simplejson)

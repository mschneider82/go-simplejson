


### go-simplejson


this is a fork of github.com/gdaisukesuzuki/go-simplejson which is a fork of
bitlys simplejson.

a Go package to interact with arbitrary handling indexed/array-type data (bitly
lacks).

In this fork the json calls are replaced from encoding/json to the faster
json-iterator/go library.

json-iterator is expored to Jsoniter, so the properties like HTML encoding can be changed.


~~[![Build Status]~~~~(https://secure.travis-ci.org/bitly/go-simplejson.png)~~~~](http://travis-ci.org/bitly/go-simplejson)~~

### Importing

1. do on Terminal

    `% go get github.com/gdaisukesuzuki/go-simplejson`

2. And in the source code, set as follows

`import (`

`   "github.com/gdaisukesuzuki/go-simplejson"`

`    ...`

` )`


### Documentation

Visit the docs on [gopkgdoc](https://godoc.org/github.com/gdaisukesuzuki/go-simplejson)

## Modified by gdaisukesuzuki (10/3/2018)

Change Set/Add/Del..method returning value of error

`Set()`

`SetEmpty()`

`SetPath()` ... deprecated and will be deleted

`SetEmptyPath()` ... deprecated and will be deleted

`Del()`

`DelIndex()`

`ZeroIndex()`

`AddIndex()`

`AddEmptyIndex()`

`SetIndex()`

`SetEmptyIndex()`

## Modified by gdaisukesuzuki (4/8/2018)

trying adding 1 "method" (currently buggy)

`GetMatchesIndex()` ... return indexes and Json's from array whose keys(s) match value(s)

### Modified by gdaisukesuzuki (3/31/2018)

Fix the bugs.

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

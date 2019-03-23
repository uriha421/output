# output

## What is this?
This command converts a markdown file into a beautiful html file. It enables you easily to output what you input.

## Prerequisites
* Golang
* Chrome

## How to use?
Type those commands in your terminal.

```
$ go get github.com/uriha421/output
$ cd output
$ go build
$ ./output md/note.md html/note.html
$ open html/note.html
```
You have succeeded in creating a beautiful html output!

## Interface
### markdown

Markdown files must satisfy that interface.
You can use # and ## only once respectively.

```
# Title (you can use # once only.)
## Introduction (you can use ## once only.)
At introduction, you need write what you readers already know, understand or agree.
You need to make sure that the readers and you are on the same stage.
Then you need to write a question or problem on the situation. Finally write some points to support your argument.

* The first point.
* The second point.
* The third point.

### The first point
As you can see, you have item numbers automatically.

#### sample

#### sample

#### sample


### The second point

#### sample

#### sample

#### sample


### The third point

#### sample

#### sample

#### sample
```

### command

The output command must satisfy that interface.
You can omit -t flag. The default value of -t flag is "./templates/note.tpl".

```
$ ./output path/to/markdown path/to/html -t path/to/template
```

## Reference
* https://github.com/russross/blackfriday
* https://github.com/microcosm-cc/bluemonday
* https://github.com/google/code-prettify

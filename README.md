# link-select
Small application to get used to Go (http://golang.org/) - so don't blame me for mistakes as this is my first time using Go in practice.

# Description
Small command-line tool to select links randomly and store them.  
Links are stored in JSON files.  

## Link storage format
Link storage is kept simple: just an array of (title, link) pairs in which the title may be omitted (the link may be omitted as well but this makes no sense at all and thus is not checked by now).  

Sample JSON files (see also the *files* folder of this repository):  
*read.json*
```
[
	{
		"title": "Vim After 11 Years",
		"link": "https://statico.github.io/vim.html"
	},
	{
		"title": "The Kernel Boot Process",
		"link": "http://duartes.org/gustavo/blog/post/kernel-boot-process/"
	},
		{
		"title": "Code Injection Attacks on HTML5-based Mobile Apps",
		"link": "http://arxiv.org/ftp/arxiv/papers/1410/1410.7756.pdf"
	},
	{
		"title": "Cache is the new RAM - MemSQL",
		"link": "http://blog.memsql.com/cache-is-the-new-ram/"
	}
]
```
*watch.json*
```
[
 	{
		"title": "Donald Knuth's 20th Annual Christmas Tree Lecture: (3/2)-ary Trees",
		"link": "https://www.youtube.com/watch?v=P4AaGQIo0HY"
	},
	{
		"title": "Is it really Complex? Or did we just make it Complicated?",
		"link": "https://www.youtube.com/watch?v=ubaX1Smg6pY&="
	},
	{
		"title": "The F# Path to Relaxation",
		"link": "https://www.youtube.com/watch?v=W-D6W7EA8gw"
	}
]
```
*book.json*
```
[
	{
		"title": "How to Design Programs, Second Edition",
		"link": "http://www.ccs.neu.edu/home/matthias/HtDP2e/"
	},
	{
		"title": "The Rust Programming Language",
		"link": "https://doc.rust-lang.org/book/"
	}
]
```

## Configuration
The `config.json` file has two main purposes by now:  
1. Configure the paths to the files to select from.  
2. Configure the command to be used to open the link (most likely a browser, e.g. chromium-browser).  

This configuration file has to be present in the folder from which `link-select` is called. The file paths have to be absolute or relative from the configuration file.

Sample configuration JSON file:  
*config.json*
```
{
	"files": {
		"read": "files/read.json",
		"watch": "files/watch.json",
		"book": "files/book.json"
	},
	"system": {
		"browser": "chromium-browser"
	}
}
```

# Usage
Assuming a Go environment set up as described on golang.org.  
```
go install link-select
```
```
link-select OPTION

At least one mandatory argument.
 -a=read|watch|book, --add-link=read|watch|book
    Add a link to the list provided (read, watch or book).
 -s=read|watch|book, --sel-link=read|watch|book
    Select a link from the list provided (read, watch or book).
```


# Some useful references
* http://golang.org/doc/code.html
* https://golang.org/doc/effective_go.html
* https://golang.org/pkg/encoding/json
* http://spf13.com/post/is-go-object-oriented/
* http://blog.golang.org/godoc-documenting-go-code

# link-select
Small application to get used to Go (http://golang.org/) - so don't blame me for mistakes as this is my first time using Go in practice.

# Description
Small command-line tool to select links randomly and store them.  
Links are stored in JSON files.  

A more detailed description will follow.

# Usage
Assuming a Go environment set up as described on golang.org.  
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

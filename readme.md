http://www.golang-book.com/

### Notes

 * go run helloworld.go
 * short variable declaration (`:=`) serves as a declaration and initialization, staticly typed, short for `var name string; name = "John"`
 * `panic` and `recover` for exception handling
 * `*` pointer, followed by type, also used to dereference a pointer
 * `new` can be used to create a pointer
 * `&` find the address of a variable, `&x` returns a pointer to a type, e.g. `*int`
 * "anonymous fields"/"embedded types" can be used to achieve subclassing
 * compose functionality with Interfaces, Interface defines a method set (list of methods)
 * concurrency with goroutines and channels
 * channels `c <- "ping"` means send string "ping" to channel c
 * `msg := <- c` means receive message on channel c and assign it to msg
 * channels can be bi-directional, or only send, or only receive
 * `select` keyboard can be used like case statement with channels
 * normal channels are sychronous
 * buffered channel is asynchronous, sending or receiving does not wait unless the channel is already full
 * [packages information](http://www.golang-book.com/11), `src` and `pkg` directories are required, set `GOPATH`
 * importing a package, can create an alias
 * `container/list` package implements a doubly-linked list

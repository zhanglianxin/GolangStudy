## 结构体

Structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds.

Structures are used to represent a record.

### Definition

The `struct` statement defines a new data type, with multiple members for your program.

The `type` statement binds a name with the type which is struct in our case.

e.g.

```go
type Book struct {
    title string
    author string
    subject string
    book_id int
}
```

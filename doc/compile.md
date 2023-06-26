
## 编译优化

- 内联优化
- 逃逸分析


```golang
type User struct {
    name  string
    email string
}

func createUser(name string, email string) *User {
    u := User{name: name, email: email}
    return &u
}
```

在 createUser 函数中，创建了一个新的 User 并返回其地址。注意，由于返回了 User 值的地址，所以它被分配在栈上，因此不会逃逸到堆上。

如果我们在返回之前添加了一个获取 User 值地址的行：

```golang
func createUser(name string, email string) *User {
    u := User{name: name, email: email}
    up := &u
    return up
}
```

现在， User 值的地址被获取并存储在一个变量中，然后返回。这导致该值逃逸到堆上而不是分配在栈上。

逃逸分析很重要，因为堆分配比栈分配更昂贵，所以减少堆分配可以提高性能。

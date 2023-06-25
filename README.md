
## Run

```bash
# 运行单元测试
go test -v

# 运行压力测试
go test -bench=.

# 执行当前目录中的所有基准测试，并打印内存分配统计信息
go test -bench=. -benchmem

# 运行测试并生成性能分析数据：
go test -cpuprofile=prof.out

# 使用以下命令启动pprof工具的交互式 shell：
go tool pprof prof.out

# 我们可以使用 top 命令显示消耗CPU时间最多的函数：
(pprof) top
```

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


## 执行跟踪器

```golang
package main

import (
    "fmt"
    "os"
    "runtime/trace"
)

func main() {
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()

    fmt.Println("Hello, World!")
}
```

在这个示例中，我们创建了一个跟踪文件，开始跟踪，并停止跟踪。当程序运行时，跟踪数据将被写入到名为trace.out的文件中。然后，您可以分析这些跟踪数据，以更好地理解程序的运行情况。


#### 内存管理和垃圾回收调优

```golang
package main

import (
    "fmt"
    "runtime"
    "runtime/debug"
)

func main() {
    // Set the maximum number of CPUs to use
    runtime.GOMAXPROCS(2)

    // Set the minimum heap size to 1GB
    runtime.MemProfileRate = 1 << 30

    // Set the garbage collection percentage to 50%
    debug.SetGCPercent(50)

    fmt.Println("Hello, World!")
}
```

在这个示例中，我们设置了最大CPU使用数量、最小堆大小和垃圾回收百分比。这些设置可以根据程序的需求进行调整，以提高性能。





## Run

单元测试

```bash
go test ./... -v

# Run the fuzz test without fuzzing it to make sure the seed inputs pass.
# 运行了 fuzz test，但是没有 fuzzing，确保输入通过
```

Fuzz

```bash
# Only wish to run the fuzz test.
go test ./... -run=FuzzReverse

# Run fuzzing
go test -fuzz=Fuzz -fuzztime=10s ./test_fuzzing
# -fuzz=Fuzz 指定运行 fuzz
# -fuzztime=10s 运行时间限制
# ./test_fuzzing 指定运行模块
```

压力测试

```bash
go test -bench=.

# 执行当前目录中的所有基准测试，并打印内存分配统计信息
go test -bench=. -benchmem
```

性能分析

```bash
# 运行测试并生成性能分析数据：
go test -cpuprofile=prof.out

# 使用以下命令启动pprof工具的交互式 shell：
go tool pprof prof.out

# 我们可以使用 top 命令显示消耗CPU时间最多的函数：
(pprof) top
```

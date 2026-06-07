# 第 4 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 4 章（复合数据类型）

## 学习进度记录（持续更新）
- 记录日期：2026-06-07
- 当前进度：第 4 周 Day 2 已完成。
- 已完成：
  - [x] Day 1：完成 `printSliceDemo()`，定义 `[]string{"Tom", "Jack", "Lucy"}` 并用 `for` 循环逐个输出。
  - [x] Day 2：完成 `printArrayAndSliceDemo()`，同时定义数组 `[3]string{...}` 与切片 `[]string{...}` 并输出，并说明“切片更常用，因为长度更灵活”。
- 当前代码状态（`main.go`）：当前主流程调用 `printArrayAndSliceDemo()`；前几周的示例函数仍保留用于复习。
- 下次开始先做：进入 Day 3，练习用 `range` 同时遍历索引和值。

## Day 1：切片入门
### 任务清单
- [x] 定义一个切片并保存多条数据。
- [x] 用 `for` + `len(...)` 逐个打印切片元素。
- [x] 运行 `go run .` 与 `go test ./...` 验证。

## Day 2：数组 vs 切片
### 任务清单
- [x] 定义一个固定长度数组。
- [x] 再定义一个切片，对比两者写法和输出。
- [x] 用一句话说明：为什么切片在 Go 里更常用。

## Day 3：range 遍历
### 任务清单
- [ ] 把 `for i := 0; i < len(...)` 改写成 `for index, value := range ...`。
- [ ] 同时打印索引和值。

## Day 4：append 追加元素
### 任务清单
- [ ] 给切片追加 1~2 个新元素。
- [ ] 重新打印，观察长度变化。

## Day 5：周复盘 + 小作业
### 任务清单
- [ ] 用 5 句话复盘第 4 周。
- [ ] 小作业：写 `printNumberSlice()`，保存一组整数并打印它们的总和。
- [ ] 运行验证并提交输出。

## 固定命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```


# 第 4 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 4 章（复合数据类型）

## 学习进度记录（持续更新）
- 记录日期：2026-06-13
- 当前进度：第 4 周 Day 5 代码任务已完成（周复盘待补）。
- 已完成：
  - [x] Day 1：完成 `printSliceDemo()`，定义 `[]string{"Tom", "Jack", "Lucy"}` 并用 `for` 循环逐个输出。
  - [x] Day 2：完成 `printArrayAndSliceDemo()`，同时定义数组 `[3]string{...}` 与切片 `[]string{...}` 并输出，并说明“切片更常用，因为长度更灵活”。
  - [x] Day 3：完成 `printRangeDemo()`，使用 `range` 遍历字符串切片和整数切片，并同时打印索引和值。
  - [x] Day 4：完成 `printAppendDemo()`，使用 `append` 给切片追加多个新元素，并打印追加前后结果及最终长度。
  - [x] Day 5：完成 `printNumberSlice()`，定义整数切片并分别计算 `[]int{10,20,30,40}` 与 `[]int{1,2,3,4,5}` 的总和。
- 备注：本周代码任务已完成；若后续愿意，可再补 5 句话周复盘。
- 当前代码状态（`main.go`）：当前主流程直接调用 `sumSlice(numss)` 演示整数切片求和；`printNumberSlice()` 等第 4 周示例函数仍保留用于复习。
- 下次开始先做：进入下一周，继续做“函数与数据处理”的更系统练习。

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
- [x] 把 `for i := 0; i < len(...)` 改写成 `for index, value := range ...`。
- [x] 同时打印索引和值。

## Day 4：append 追加元素
### 任务清单
- [x] 给切片追加 1~2 个新元素。
- [x] 重新打印，观察长度变化。

## Day 5：周复盘 + 小作业
### 任务清单
- [ ] 用 5 句话复盘第 4 周。
- [x] 小作业：写 `printNumberSlice()`，保存一组整数并打印它们的总和。
- [x] 运行验证并提交输出。

## 固定命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```


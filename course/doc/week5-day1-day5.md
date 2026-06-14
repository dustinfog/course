# 第 5 周任务卡（Day 1 ~ Day 5)

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 5 章（函数）

## 学习进度记录（持续更新）
- 记录日期：2026-06-14
- 当前进度：第 5 周 Day 4 已完成；Day 5 本次跳过。
- 已完成：
  - [x] Day 1：完成 `sumSlice(nums []int) int` 与 `maxSlice(nums []int) int`，能把整数切片作为参数传入函数，并在 `main()` 中接收返回值后输出结果。
  - [x] Day 2：完成 `sumAndCount(nums []int) (int, int)`，能同时返回总和与元素个数，并在 `main()` 中接收两个返回值后输出。
  - [x] Day 3：完成 `minSlice(nums []int) int`，能返回整数切片中的最小值，并在 `main()` 中打印结果。
  - [x] Day 4：为 `maxSlice(nums []int) int` 与 `minSlice(nums []int) int` 增加空切片保护，并在 `main()` 中验证空切片返回 `0`。
- 当前代码状态（`main.go`）：主流程当前演示 `sumSlice(numss)`、`maxSlice(numse)`、`sumAndCount(numss)`、`minSlice(numse)` 以及空切片的最大/最小值保护；前 1~4 周的练习函数仍保留，便于回顾。
- 下次开始先做：进入第 6 周 Day 1。

## Day 1：函数参数与返回值
### 任务清单
- [x] 写一个接收切片参数并返回总和的函数 `sumSlice(nums []int) int`。
- [x] 写一个接收切片参数并返回最大值的函数 `maxSlice(nums []int) int`。
- [x] 在 `main()` 中调用并打印结果。
- [x] 运行 `go run .` 与 `go test ./...` 验证。

## Day 2：两个返回值（预留）
### 任务清单
- [x] 练习返回多个值：`sumAndCount(nums []int) (int, int)`。
- [x] 在 `main()` 中接收两个返回值并打印。

## Day 3：函数职责拆分（预留）
### 任务清单
- [x] 继续练习“计算”和“打印”分离。
- [x] 新增 `minSlice(nums []int) int`，让切片计算函数更完整。

## Day 4：边界输入（预留）
### 任务清单
- [x] 思考空切片等边界情况。
- [x] 为 `maxSlice` 与 `minSlice` 补最小保护逻辑。

## Day 5：周复盘 + 小作业（预留）
### 任务清单
- [ ] 用 5 句话复盘第 5 周。
- [ ] 完成一个基于函数返回值的小练习。
- 备注：本次选择跳过 Day 5，可后续回补，不影响继续进入下一周。

## 固定命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```


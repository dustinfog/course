# 第 3 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 3 章（基本数据类型）

## 学习进度记录（持续更新）
- 记录日期：2026-06-07
- 当前进度：第 3 周 Day 5 已完成（本周完成）。
- 已完成：
  - [x] Day 1：完成 `printTypeDemo()`，输出 `int/float64/string/bool`，并验证 `5/2` 与 `5.0/2.0` 差异。
  - [x] Day 2：完成 `fmt.Printf` 格式化输出练习，理解 `%d/%.2f/%q/%t` 的基本用途。
  - [x] Day 3：完成 `const` + `iota` 练习，使用 `printLevelDemo()` 输出 `A=0 B=1 C=2`。
  - [x] Day 4：完成 `printConvertDemo()`，演示 `float64(x)` 与 `int(y)` 的显式类型转换。
-  - [x] Day 5：完成 `printTypeSummary()`，集中输出基础类型、除法差异、`iota` 常量和类型转换结果。
- 当前代码状态（`main.go`）：当前主流程调用 `printTypeSummary()`；本周练习函数 `printLevelDemo()`、`printConvertDemo()`、`printTypeDemo()` 以及前两周示例函数仍保留用于复习。
- 下次开始先做：进入第 4 周 Day 1，学习数组、切片和一组数据的组织方式。

## Day 1：整数、浮点、字符串、布尔
### 任务清单
- [x] 定义并打印 `int`、`float64`、`string`、`bool`。
- [x] 对比整数除法与浮点除法结果。
- [x] 运行 `go run .` 与 `go test ./...` 验证。

## Day 2：字符串与格式化输出
### 任务清单
- [x] 用 `fmt.Printf` 格式化输出类型和值（如 `%d`、`%.2f`、`%q`、`%t`）。
- [x] 完成本节核心目标：理解并输出整数与浮点类型的格式化结果。
- [x] 保持当前 `divide` 演示可运行。

## Day 3：常量与 iota 入门
### 任务清单
- [x] 新增一组 `const`（至少 3 个），尝试 `iota` 生成递增值。
- [x] 打印这些常量并解释用途。

## Day 4：类型转换
### 任务清单
- [x] 演示 `int` 与 `float64` 的显式转换。
- [x] 理解“不同类型不能直接混算，必须显式转换”的规则。

## Day 5：周复盘 + 小作业
### 任务清单
- [ ] 用 5 句话复盘第 3 周。
- [x] 小作业：写 `printTypeSummary()`，集中输出本周所有类型演示结果。
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


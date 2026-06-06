# 第 3 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 3 章（基本数据类型）

## 学习进度记录（持续更新）
- 记录日期：2026-06-06
- 当前进度：第 3 周 Day 2 进行中（格式化输出已完成，待补说明文本）。
- 已完成：
  - [x] Day 1：完成 `printTypeDemo()`，输出 `int/float64/string/bool`，并验证 `5/2` 与 `5.0/2.0` 差异。
- 当前代码状态（`main.go`）：`printTypeDemo()` 已使用 `fmt.Printf` 输出 `%d/%.2f/%q/%t`，`divide` 示例可运行。
- 下次开始先做：继续 Day 2，补“整除与浮点除法差异”说明文本后进入 Day 3。

## Day 1：整数、浮点、字符串、布尔
### 任务清单
- [x] 定义并打印 `int`、`float64`、`string`、`bool`。
- [x] 对比整数除法与浮点除法结果。
- [x] 运行 `go run .` 与 `go test ./...` 验证。

## Day 2：字符串与格式化输出
### 任务清单
- [x] 用 `fmt.Printf` 格式化输出类型和值（如 `%d`、`%.2f`、`%q`、`%t`）。
- [ ] 增加一行说明文本，解释整除与浮点除法差异。
- [x] 保持当前 `divide` 演示可运行。

## Day 3：常量与 iota 入门
### 任务清单
- [ ] 新增一组 `const`（至少 3 个），尝试 `iota` 生成递增值。
- [ ] 打印这些常量并解释用途。

## Day 4：类型转换
### 任务清单
- [ ] 演示 `int` 与 `float64` 的显式转换。
- [ ] 加一处“未转换会报错”的注释说明（不用保留错误代码）。

## Day 5：周复盘 + 小作业
### 任务清单
- [ ] 用 5 句话复盘第 3 周。
- [ ] 小作业：写 `printTypeSummary()`，集中输出本周所有类型演示结果。
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


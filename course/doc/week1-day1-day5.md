# 第 1 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 1 章（入门与程序结构）

## 学习进度记录（持续更新）
- 记录日期：2026-05-17
- 当前进度：第 1 周 Day 3 已完成，准备进入 Day 4。
- 已完成：
  - [x] Day 1：修复欢迎语格式化输出（`Printf`）。
  - [x] Day 2：循环范围从 `1..5` 改为 `1..10`。
  - [x] Day 3：提取 `printSeries(limit int)`，并保持 `100/i` 输出逻辑。
- 当前代码状态（`main.go`）：`main()` 调用 `printSeries(10)`，输出 10 行 `100/i` 结果。
- 下次开始先做：Day 4「函数参数化」，把函数改为 `printSeries(base int, limit int)`。
- 下次继续前先运行：

```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```

## 使用方式
- 每天只完成当天任务，不跨天。
- 每次改动都运行命令验证，再发结果给 AI 复盘。
- 固定在项目根目录 `/Users/panzd/course/course` 执行命令。

---

## Day 1：认识 `main.go` + 修复欢迎语

### 当前目标
- 理解 `main.go` 的基本执行流程。
- 区分 `fmt.Println` 与 `fmt.Printf`。

### 任务清单
- [ ] 用自己的话解释 `main.go` 每一段代码。
- [ ] 修复 `%s` 被当普通文本输出的问题。
- [ ] 运行并确认输出。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```

### 预期结果
- 欢迎语正确显示变量值，不出现字面量 `%s`。
- `go test ./...` 不再报 `fmt.Println` 格式化告警。

### 常见报错排查
- 报错：`possible Printf formatting directive %s`
  - 含义：`Println` 不处理 `%s`。
  - 最短修复：改为 `fmt.Printf("Hello and welcome, %s!\n", s)`。

---

## Day 2：循环与整数除法

### 当前目标
- 理解 `for` 循环结构。
- 理解整型除法会截断小数。

### 任务清单
- [ ] 把循环范围从 `1..5` 改成 `1..10`。
- [ ] 观察 `100/i` 结果，并解释 `100/3` 为什么是 `33`。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

### 预期结果
- 循环输出变为 10 行。
- 能说清整数除法与浮点除法区别。

### 常见报错排查
- 报错：`syntax error`
  - 含义：常见于 `for` 语法（分号/花括号）写错。
  - 最短修复：检查 `for 初始化; 条件; 递增 {}`。

---

## Day 3：第一次提取函数

### 当前目标
- 把流程代码拆成函数，保持行为不变。

### 任务清单
- [ ] 新增函数（示例：`printSeries(limit int)`）。
- [ ] 在 `main()` 中调用该函数。
- [ ] 对比 Day 2 输出，确认一致。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```

### 预期结果
- `main()` 更短，更容易读。
- 功能与 Day 2 保持一致。

### 常见报错排查
- 报错：`undefined: printSeries`
  - 含义：函数名不一致或未正确定义。
  - 最短修复：统一名称并确保同属 `package main`。

---

## Day 4：函数参数化

### 当前目标
- 让同一函数支持不同输入。

### 任务清单
- [ ] 把函数改成 `printSeries(base int, limit int)`。
- [ ] 在 `main()` 中分别调用：`(100,10)` 和 `(50,10)`。
- [ ] 给两段输出加标题，方便区分。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

### 预期结果
- 同一函数可复用到两组场景。
- 输出结构更清晰。

### 常见报错排查
- 报错：`not enough arguments` / `too many arguments`
  - 含义：函数定义和调用参数个数不一致。
  - 最短修复：同步检查函数签名与调用处。

---

## Day 5：本周复盘 + 小作业

### 当前目标
- 巩固第 1 章核心概念（程序结构、循环、函数、输出）。

### 任务清单
- [ ] 用 5 句话复盘本周。
- [ ] 小作业：新增函数打印 `1..limit` 的平方值。
- [ ] 说明何时用 `Println`，何时用 `Printf`。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```

### 预期结果
- 能独立完成“新增函数 -> 调用 -> 验证”。
- 能准确说出本周 3 个核心概念。

### 常见报错排查
- 报错：`declared and not used`
  - 含义：变量声明了但没使用。
  - 最短修复：删除变量或接入实际逻辑。

---

## 每日提交模板（复制填写）

```text
【今天是 Day X】
1) 我完成了：
2) 我遇到的报错：
3) 我的理解（用自己的话）：
4) 我希望你重点检查：
```

## 给 AI 的批改偏好
- 先讲“你哪里做对了”。
- 每次最多指出 2 个最关键改进点。
- 只给最小改动方案，不直接大改。


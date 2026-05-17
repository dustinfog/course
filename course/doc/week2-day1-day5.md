# 第 2 周任务卡（Day 1 ~ Day 5）

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 2 章（程序结构）

## 使用方式
- 本周只引入 2 个新概念：作用域、函数返回值。
- 每次改动后都运行验证命令，再提交给 AI 复盘。
- 固定在项目根目录 `/Users/panzd/course/course` 执行命令。

---

## Day 1：变量声明与作用域

### 当前目标
- 理解“变量在哪声明，就在哪可用”（作用域）。
- 避免常见的“变量未定义/未使用”错误。

### 任务清单
- [ ] 在 `main.go` 中新增一个局部变量（如 `base := 100`）。
- [ ] 尝试在不同代码块中使用它，观察作用域行为。
- [ ] 用一句话解释“全局变量 vs 局部变量”（先直觉理解即可）。

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
- 你能判断一个变量在当前行是否可访问。

### 常见报错排查
- 报错：`undefined: xxx`
  - 含义：变量超出作用域或名字写错。
  - 最短修复：把声明移动到使用前，或统一命名。

---

## Day 2：函数参数与返回值

### 当前目标
- 理解函数是“输入 -> 处理 -> 输出”。

### 任务清单
- [ ] 新建函数：`func divide(base int, i int) int`。
- [ ] 在循环中调用 `divide`，打印返回值。
- [ ] 用自己的话解释“参数”和“返回值”的区别。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

### 预期结果
- 你能把计算逻辑放进函数，而不是都写在 `main()`。

### 常见报错排查
- 报错：`too many return values` / `not enough return values`
  - 含义：函数声明返回值与实际 `return` 不一致。
  - 最短修复：对齐函数签名和 `return` 表达式。

---

## Day 3：多返回值（结果 + 错误）

### 当前目标
- 初步理解 Go 常见模式：`value, err`。

### 任务清单
- [ ] 把 `divide` 改成：`func divide(base int, i int) (int, error)`。
- [ ] 当 `i == 0` 时返回错误；正常时返回结果和 `nil`。
- [ ] 调用处判断 `err != nil` 再决定是否打印结果。

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
- 程序不会因除零直接崩溃。
- 你能读懂 `if err != nil` 的基本意思。

### 常见报错排查
- 报错：`cannot use ... as type error`
  - 含义：返回值类型与 `error` 不匹配。
  - 最短修复：出错分支返回 `fmt.Errorf(...)`。

---

## Day 4：包级变量与常量

### 当前目标
- 认识 `const` 与 `var` 的区别。
- 学会用常量管理“固定配置值”。

### 任务清单
- [ ] 定义常量（如 `const defaultLimit = 10`）。
- [ ] 让主流程使用该常量，而不是硬编码数字。
- [ ] 尝试改一次常量值，观察输出变化。

### 可执行命令
```bash
cd /Users/panzd/course/course
go run .
```

### 预期结果
- 你知道“哪些值适合写成常量”。

### 常见报错排查
- 报错：`cannot assign to ...`
  - 含义：常量不能在运行中被重新赋值。
  - 最短修复：如果需要变化，改用变量 `var`。

---

## Day 5：第二周复盘 + 小作业

### 当前目标
- 把第 2 章核心知识串起来：声明、作用域、返回值、错误处理。

### 任务清单
- [ ] 用 5 句话总结本周学习。
- [ ] 小作业：新增 `safeModulo(base, i)`，返回 `(int, error)`。
- [ ] 在 `main()` 中同时演示 `divide` 和 `safeModulo`。

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
- 你可以独立写出“带错误返回值”的函数。
- 你能解释为什么 Go 里常用返回 `error`。

### 常见报错排查
- 报错：`imported and not used`
  - 含义：引入了包但代码没用到。
  - 最短修复：删除无用 import，或补上实际调用。

---

## 每日提交模板（复制填写）

```text
【第 2 周 Day X】
1) 我完成了：
2) 运行结果：
3) 遇到的报错：
4) 我对“作用域/返回值/error”的理解：
5) 我希望你重点检查：
```

## 给 AI 的批改偏好
- 先肯定正确点，再给改进建议。
- 一次最多指出 2 个最关键问题。
- 优先最短修复路径，不做大重构。


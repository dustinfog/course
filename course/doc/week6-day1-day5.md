# 第 6 周任务卡（Day 1 ~ Day 5)

> 教材对应：《Go 程序设计语言》中文站（`https://gopl-zh.github.io/`）第 6-7 章（方法与接口）

## 学习进度记录（持续更新）
- 记录日期：2026-06-14
- 当前进度：第 6 周 Day 1 已完成。
- 已完成：
  - [x] Day 1：定义 `Student` 结构体，包含 `Name string` 与 `Scores []int` 两个字段；在 `printStudentDemo()` 中创建 `Tom`、`Lucy` 两个学生并打印字段内容。
- 当前代码状态（`main.go`）：`main()` 当前调用 `printStudentDemo()`；`Student` 结构体已定义，程序输出两位学生的姓名与分数切片。
- 下次开始先做：进入第 6 周 Day 2（方法入门）。

## Day 1：结构体入门
### 任务清单
- [x] 定义 `Student` 结构体。
- [x] 创建两个 `Student` 实例。
- [x] 访问并打印 `Name`、`Scores` 字段。
- [x] 运行 `go run .` 与 `go test ./...` 验证。

## Day 2：方法入门（预留）
### 任务清单
- [ ] 给 `Student` 增加一个最小方法，例如计算总分或平均分。
- [ ] 在 `main()` 中调用该方法并打印结果。

## Day 3：多个学生数据处理（预留）
### 任务清单
- [ ] 把多个学生放入切片中。
- [ ] 遍历切片并打印每位学生的信息。

## Day 4：结构体与函数配合（预留）
### 任务清单
- [ ] 写一个接收 `Student` 的函数。
- [ ] 在函数中输出或计算学生数据。

## Day 5：周复盘 + 小作业（预留）
### 任务清单
- [ ] 用 5 句话复盘第 6 周。
- [ ] 做一个“学生信息 + 成绩”的小扩展练习。

## 固定命令
```bash
cd /Users/panzd/course/course
go run .
```

```bash
cd /Users/panzd/course/course
go test ./...
```


# [LeetCode](https://leetcode.com) 的 Go 解答 {{- /* 本文件是用来生成 README.md 的模板 */}}

[![LeetCode 排名](https://img.shields.io/badge/{{.Username}}-{{.Ranking}}-blue.svg)](https://leetcode.com/{{.Username}}/)
[![codecov](https://codecov.io/gh/zjbztianya/LeetCode/branch/master/graph/badge.svg)](https://codecov.io/gh/zjbztianya/LeetCode)
[![Build Status](https://www.travis-ci.org/zjbztianya/LeetCode.svg?branch=master)](https://www.travis-ci.org/zjbztianya/LeetCode)
 [![Go](https://img.shields.io/badge/Go-1.14-blue.svg)](https://golang.google.cn)

## 进度

> 统计规则：1.免费题，2.算法题，3.能提交 Go 解答

{{.ProgressTable}}

## 题解

{{.AvailableTable}}
以下免费的算法题，暂时不能提交 Go 解答

{{.UnavailableList}}

## helper

[gen_kit](./Helper) 会处理大部分琐碎的工作。


## 致谢

感谢[aQuaYi](https://github.com/aQuaYi)开源的Leetcode Helper工具

# go-zero
## 指令
- 格式化`core.api`: `goctl api format --dir 文件地址`
- 把`.api`文件内容转成代码：`goctl api go -api core.api -dir . -style go_zero`
  - 存在的文件它会忽略, 但是types.go的结构体会发生变化的

## Token
- 我们直接把token放在header头部时，`go-zero`阻止，我们就放在`Authorization`上就行了


## TODO
- id自增不好，容易被用户察觉搞事情

## 删除文件夹
- 如果递归的删除下去，那么删除的成本就太大了，我们直接把它删除然后保证目录下的其他东西用户不可见即可。
- 然后用户从回收站进行文件夹的恢复，也是直接改一处即可
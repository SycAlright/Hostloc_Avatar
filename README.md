# Hostloc头像一键上传，免cookie，免flash，非js脚本

## 使用

### 1. 获取 `Agent` + `Input` 参数
打开 https://hostloc.com/home.php?mod=spacecp&ac=avatar 查看源代码或者审查元素  
搜索`input=`和`agent=`并记录下对应参数，已知`agent`是用户唯一值。

### 2. 修改工具参数
修改代码：
```go
const (
	INPUT = "your_input"
	AGENT = "your_agent"
)
```

### 3. 准备你的头像
> L：200 x 200 px （最大支持 200 x 250 px）  
> M：120 x 120 px  
> S：45 x 45 px

修改代码：
```go
func main() {
	S := hexImage("45.jpg")     // 修改对应文件
	M := hexImage("120.jpg")    // 修改对应文件
	L := hexImage("200.jpg")    // 修改对应文件
	... 
}
```

### 4. 执行，判断结果

`go run main.go`

打印结果：

```xml
// 成功上传，刷新头像即可
<?xml version="1.0" ?><root><face success="1"/></root>

// 图像像素超出服务器后台限制 或 格式不支持（Discuz常见支持：JPG/PNG/GIF）
<?xml version="1.0" ?><root><face success="0"/></root>

// Agent 或 Input 错误
Access denied for agent changed

// 空文件 或 文件上传错误
<root><message type="error" value="-2" /></root>
```
## package 的引用

> 1. 在`GOPATH/src`建立`tepackage`包文件夾。
> 2. 在包文件夹新建 go 文件即可。
> 3. import 包名即可使用。

### 1. `other.go`

```go
package tepackage

import(
	"fmt"
)
var Index int = 100
func Other(){
	fmt.Println("other")
}
```

### 2. `main.go`

```go
package main

import (
	"tepackage"
	"fmt"
)
func main(){
	tepackage.Test()
	fmt.Println(tepackage.Index)
	tepackage.Other()
}
```

### 注意

> 1. 包文件的大写字母开头的都是`被导出的名称`，可以通过包+变量名、方法等进行引用。
> 2. 同一个包内，可以直接引用对应的方法、变量、常量等。

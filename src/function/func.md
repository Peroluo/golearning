## func 使用

### 1. 基本使用

```go
func Add(x int,y int) int{
	return x+y
}
```

### 2. 缩写参数类型

```go
func Add2(x,y int) int{
  return x+y
}
```

### 3. 多参返回

```go
func Swap(a,b string)(x,y string){
	x=b;
	y=a;
	return x,y
}
```

### 4.命名返回值

```go
func Swap(a,b string)(x,y string){
	x=b;
	y=a;
	return
}
```

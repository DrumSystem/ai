#基础知识
## 数组
如果数组中元素的个数小于或者等于 4 个，那么所有的变量会直接在栈上初始化，如果数组元素大于 4 个，变量就会在静态存储区初始化然后拷贝到栈上
```go
// 栈上初始化
var arr [3]int
arr[0] = 1
arr[1] = 2
arr[2] = 3

// 静态存储区初始化
var arr [5]int
statictmp_0[0] = 1
statictmp_0[1] = 2
statictmp_0[2] = 3
statictmp_0[3] = 4
statictmp_0[4] = 5
arr = statictmp_0
```

数组结构：指向数组开头的指针， 元素的数量， 元素类型的大小。
对数组的访问和赋值需要同时依赖编译器（使用整数和常量）和运行时（变量下标访问），它的大多数操作在编译期间都会转换成直接读写内存，在中间代码生成期间，编译器还会插入运行时方法 runtime.panicIndex 调用防止发生越界错误。

## 切片
切片引入了一个抽象层，提供了对数组中部分连续片段的引用，而作为数组的引用，我们可以在运行区间可以修改它的长度和范围。当切片底层的数组长度不足时就会触发扩容，切片指向的数组可能会发生变化，不过在上层看来切片是没有变化的，上层只需要与切片打交道不需要关心数组的变化
当切片发生逃逸或者非常大时，运行时需要 runtime.makeslice 在堆上初始化切片，
- Data 是指向数组的指针;
- Len 是当前切片的长度；
- Cap 是当前切片的容量，即 Data 数组的大小：
```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
// 初始化
arr[0:3] or slice[0:3]
slice := []int{1, 2, 3}
slice := make([]int, 10)
arr := [3]int{1, 2, 3}
slice := arr[0:2]
slice1 := arr[0:2]
slice1[0] = 5
fmt.Println(slice, slice1)
//输出都是5 2
```
切片扩容：还需要根据切片中的元素大小对齐内存，当数组中元素所占的字节大小为 1、8 或者 2 的倍数时，runtime.roundupsize 函数会将待申请的内存向上取整，取整时会使用 runtime.class_to_size 数组
- 如果期望容量大于当前容量的两倍就会使用期望容量；
- 如果当前切片的长度小于 1024 就会将容量翻倍；
- 如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量
```go
var arr []int64
arr = append(arr, 1, 2, 3, 4, 5)
// 执行上述代码时，会触发 runtime.growslice 函数扩容 arr 切片并传入期望的新容量 5，这时期望分配的内存大小为 40 字节；
//不过因为切片中的元素大小等于 sys.PtrSize，所以运行时会调用 runtime.roundupsize 向上取整内存的大小到 48 字节，所以新切片的容量为 48 / 8 = 6。
```

## 哈希map--一般编程语言都会采用拉链法实现
- count 表示当前哈希表中的元素数量；
- B 表示当前哈希表持有的 buckets 数量
- oldbuckets 是哈希在扩容时用于保存之前 buckets 的字段，它的大小是当前 buckets 的一半

runtime.hmap 的桶是 runtime.bmap。每一个 runtime.bmap 都能存储 8 个键值对，当哈希表中存储的数据过多，
单个桶已经装满时就会使用 extra.nextOverflow 中桶存储溢出的数据。
bmap包含 tophash 字段，tophash 存储了键的哈希的高 8 位，通过比较不同键的哈希的高 8 位可以减少访问键值对次数以提高性能, 桶的序号时低位控制的
```go
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}

type bmap struct {
    tophash [bucketCnt]uint8
}
```
哈希表扩容机制
- 等量扩容：如果扩容是溢出的桶太多导致的，那么这次扩容就是等量扩容 sameSizeGrow，sameSizeGrow 是一种特殊情况下发生的扩容，当我们持续向哈希中插入数据并将它们全部删除时，如果哈希表中的数据量没有超过阈值，就会不断积累溢出桶造成缓慢的内存泄漏
- 翻倍扩容：就是正常扩容，在扩容期间访问哈希表时会使用旧桶，向哈希表写入数据时会触发旧桶元素的分流（原本3号桶的数据会分流到3，7桶）

Go 语言使用拉链法来解决哈希碰撞的问题实现了哈希表，它的访问、写入和删除等操作都在编译期间转换成了运行时的函数或者方法。哈希在每一个桶中存储键对应哈希的前 8 位，当对哈希进行操作时，这些 tophash 就成为可以帮助哈希快速遍历桶中元素的缓存。

哈希表的每个桶都只能存储 8 个键值对，一旦当前哈希的某个桶超出 8 个，新的键值对就会存储到哈希的溢出桶中。随着键值对数量的增加，溢出桶的数量和哈希的装载因子也会逐渐升高，超过一定范围就会触发扩容，扩容会将桶的数量翻倍，元素再分配的过程也是在调用写操作时增量进行的，不会造成性能的瞬时巨大抖动。

## 字符串
与切片的结构体相比，字符串只少了一个表示容量的 Cap 字段， 字符串是一个只读的类型， 所有在字符串上的写入操作都是通过拷贝实现的

# 常用关键字
## for-range
对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 ha，在赋值的过程中就发生了拷贝，而我们又通过 len 关键字预先获取了切片的长度，所以在循环中追加新的元素也不会改变循环执行的次数
而遇到这种同时遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的 v2 变量存储切片中的元素，循环中使用的这个变量 v2 会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝
```go
func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
        //newArr = append(newArr, &arr[i])
		
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

$ go run main.go
3 3 3
```
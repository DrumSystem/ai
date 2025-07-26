# 高性能编程
## 字符串高效拼接
go语言中字符串是不变的， 拼接字符串实际上是创建了一个新的字符串对象
- 使用+号拼接
```go
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}
```
- 使用fmt.Sprintf
```go
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}
```

- 使用strings.Builder
```go
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
```
- 使用bytes.buffer
```go
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}
```
- 使用 []byte
```go
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
```
比较结果.使用 + 和 fmt.Sprintf 的效率是最低的，和其余的方式相比，性能相差约 1000 倍，而且消耗了超过 1000 倍的内存, strings.Builder、bytes.Buffer 和 []byte 的性能差距不大，而且消耗的内存也十分接近
```go
$ go test -bench="Concat$" -benchmem .
goos: darwin
goarch: amd64
pkg: example
BenchmarkPlusConcat-8         19      56 ms/op   530 MB/op   10026 allocs/op
BenchmarkSprintfConcat-8      10     112 ms/op   835 MB/op   37435 allocs/op
BenchmarkBuilderConcat-8    8901    0.13 ms/op   0.5 MB/op      23 allocs/op
BenchmarkBufferConcat-8     8130    0.14 ms/op   0.4 MB/op      13 allocs/op
BenchmarkByteConcat-8       8984    0.12 ms/op   0.6 MB/op      24 allocs/op
BenchmarkPreByteConcat-8   17379    0.07 ms/op   0.2 MB/op       2 allocs/op
PASS
ok      example 8.627s
```
### 性能比较
#### strings.Builder 和 +
使用 + 拼接 2 个字符串时，会生成一个新的字符串，那么就需要重新开辟一段内存空间，新空间的大小是两个字符串之和

使用string.Builder，bytes.Buffer，包括切片 []byte 的内存是以倍数申请的，2048 以前按倍数申请，2048 之后，以 640 递增。

#### strings.Builder 和 bytes.Buffer
strings.Builder 和 bytes.Buffer 底层都是 []byte 数组，但 strings.Builder 性能比 bytes.Buffer 略快约 10% 。
一个比较重要的区别在于，bytes.Buffer 转化为字符串时重新申请了一块空间，存放生成的字符串变量，而 strings.Builder 直接将底层的 []byte 转换成了字符串类型返回了回来

注释：unsafe.Pointer是一个万能的指针类型， Go语言是不允许两个指针类型进行转换的。即*float64 -> *int,
但是可以通过unsafe.Pointer进行转换，*float64 -> (*float64)unsafe.Pointer(*int)
```go
// String returns the accumulated string.//使用指针？
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}

// To build strings more efficiently, see the strings.Builder type.
func (b *Buffer) String() string {
    if b == nil {
    // Special case, useful in debugging.
        return "<nil>"
    }
    return string(b.buf[b.off:])
}
```

## 切片性能及陷阱
一个数组变量被赋值或者传递时，实际上会复制整个数组， 例如，将 a 赋值给 b，修改 a 中的元素并不会改变 b 中的元素：为了避免复制数组，一般会传递指向数组的指针
```go
a := [...]int{1, 2, 3} // ... 会自动计算数组长度
b := a
a[0] = 100
fmt.Println(a, b) // [100 2 3] [1 2 3]

func square(arr *[3]int) {
    for i, num := range *arr {
        (*arr)[i] = num * num
    }
}

func TestArrayPointer(t *testing.T) {
    a := [...]int{1, 2, 3}
    square(&a)
    fmt.Println(a) // [1 4 9]
    if a[1] != 4 && a[2] != 9 {
        t.Fatal("failed")
    }
}
```

切片实际上是对底层数组的引用（包含一个指向底层数组的unsafe.pointer）, 创建一个新的切片会复用原来切片的底层数组
```go
nums := make([]int, 0, 8)
nums = append(nums, 1, 2, 3, 4, 5)
nums2 := nums[2:4]
printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
printLenCap(nums2) // len: 2, cap: 6 [3 4]

nums2 = append(nums2, 50, 60)
printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]
// 因为 nums 和 nums2 指向的是同一个数组，因此 nums 被修改为 [1, 2, 3, 4, 50]
```
## 性能陷阱
在已有切片的基础上进行切片，不会创建新的底层数组。因为原来的底层数组没有发生变化，内存会一直占用，直到没有变量引用该数组。
因此很可能出现这么一种情况，原切片由大量的元素构成，但是我们在原切片的基础上切片，虽然只使用了很小一段，
但底层数组在内存中仍然占据了大量空间，得不到释放。比较推荐的做法，使用 copy 替代 re-slice
- 第一个函数直接在原切片基础上进行切片
- 第二个函数创建了一个新的切片，将 origin 的最后两个元素拷贝到新切片上，然后返回新切片
通过copy新切片会指向一个新的底层数组， 原来的大容量数组会因为没有对象引用而被释放
```go
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}
```

## for and range 性能比较
range 在迭代过程中返回的是迭代值的拷贝，如果每次迭代的元素的内存占用很低，那么 for 和 range 的性能几乎是一样。
如果想使用 range 同时迭代下标和值，则需要将切片/数组的元素改为指针，才能不影响性能。

## 反射性能
创建对象：通过反射创建对象的耗时约为 new 的 1.5 倍，相差不是特别大。

修改字段值：通过反射获取结构体的字段有两种方式，一种是 FieldByName，另一种是 Field(按照下标)
- 普通的赋值操作，每次耗时约为 0.3 ns，通过下标找到对应的字段再赋值，每次耗时约为 30 ns，通过名称找到对应字段再赋值，每次耗时约为 300 ns。

(t *structType) FieldByName 中使用 for 循环，逐个字段查找，字段名匹配时返回。也就是说，在反射的内部，字段是按顺序存储的，
因此按照下标访问查询效率为 O(1)，而按照 Name 访问，则需要遍历所有字段，查询效率为 O(N)。结构体所包含的字段(包括方法)越多，那么两者之间的效率差距则越大。

### 如何提高反射性能
使用反射赋值，效率非常低下，如果有替代方案，尽可能避免使用反射，特别是会被反复调用的热点代码。例如 RPC 协议中，需要对结构体进行序列化和反序列化，
这个时候避免使用 Go 语言自带的 json 的 Marshal 和 Unmarshal 方法，因为标准库中的 json 序列化和反序列化是利用反射实现的。
可选的替代方案有 easyjson，在大部分场景下，相比标准库，有 5 倍左右的性能提升。

使用缓存：可以利用字典将 Name 和 Index 的映射缓存起来。避免每次反复查找，耗费大量的时间。从10倍缩小为2倍
```go
func BenchmarkReflect_FieldByNameCacheSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	cache := make(map[string]int)
	//使用map缓存，方便查找
	for i := 0; i < typ.NumField(); i++ {
		cache[typ.Field(i).Name] = i
	}
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(cache["Name"]).SetString("name")
		ins.Field(cache["IP"]).SetString("ip")
		ins.Field(cache["URL"]).SetString("url")
		ins.Field(cache["Timeout"]).SetString("timeout")
	}
}
```

## 内存对齐
CPU 访问内存时，并不是逐个字节访问，而是以字长（word size）为单位访问。比如 32 位的 CPU ，字长为 4 字节，那么 CPU 访问内存的单位也是 4 字节。

这么设计的目的，是减少 CPU 访问内存的次数，加大 CPU 访问内存的吞吐量。比如同样读取 8 个字节的数据，一次读取 4 个字节那么只需要读取 2 次。
CPU 始终以字长访问内存，如果不进行内存对齐，很可能增加 CPU 访问内存的次数

每个字段按照自身的对齐倍数来确定在内存中的偏移量，字段排列顺序不同，上一个字段因偏移而浪费的大小也不同。
- a 是第一个字段，默认是已经对齐的，从第 0 个位置开始占据 1 字节。
- b 是第二个字段，对齐倍数为 2，因此，必须空出 1 个字节，偏移量才是 2 的倍数，从第 2 个位置开始占据 2 字节。
- c 是第三个字段，对齐倍数为 4，此时，内存已经是对齐的，从第 4 个位置开始占据 4 字节即可

- a 是第一个字段，默认是已经对齐的，从第 0 个位置开始占据 1 字节
- c 是第二个字段，对齐倍数为 4，因此，必须空出 3 个字节，偏移量才是 4 的倍数，从第 4 个位置开始占据 4 字节。
- b 是第三个字段，对齐倍数为 2，从第 8 个位置开始占据 2 字节。
```go
type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main() {
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
}
```


## 如何退出goroutine
- 尽量使用非阻塞 I/O（select语句包含default）（非阻塞 I/O 常用来实现高性能的网络库），阻塞 I/O 很可能导致 goroutine 在某个调用一直等待，而无法正确结束
- 业务逻辑总是考虑退出机制，避免死循环。
- 任务分段执行，超时后即时退出，避免 goroutine 无用的执行过多，浪费资源

## 其他场景
如果不判断channel是否关闭， go do(taskCh)会一直阻塞，不会被回收。

一个通道被其发送数据协程队列和接收数据协程队列中的所有协程引用着。因此，如果一个通道的这两个队列只要有一个不为空，则此通道肯定不会被垃圾回收。
另一方面，如果一个协程处于一个通道的某个协程队列之中，则此协程也肯定不会被垃圾回收，即使此通道仅被此协程所引用。事实上，一个协程只有在退出后才能被垃圾回收。
```go

func do(taskCh chan int) {
	for {
		select {
		case t, beforeClosed := <-taskCh:
			if !beforeClosed {
				fmt.PrintLn("taskChan has closed")
				return
			}   
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		//要么使用非阻塞
		default:
			return
		}
	}
}

func sendTasks() {
	taskCh := make(chan int, 10)
	go do(taskCh)
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	//执行完任务后需要关闭channel
	close(taskCh)
}

func TestDo(t *testing.T) {
    t.Log(runtime.NumGoroutine())
    sendTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}
```

### 通道关闭原则
粗鲁关闭：如果 channel 已经被关闭，再次关闭会产生 panic，这时通过 recover 使程序恢复正常。
```go
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果 ch 已关闭，将 panic
	return true // <=> justClosed = true; return
}
```

礼貌的方式: 使用 sync.Once 或互斥锁(sync.Mutex)确保 channel 只被关闭一次。
```go
ype MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}
```

优雅的方式：
情形一：M个接收者和一个发送者，发送者通过关闭用来传输数据的通道来传递发送结束信号。
情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要再发送数据了。
情形三：M个接收者和N个发送者，它们中的任何协程都可以让一个中间调解协程帮忙发出停止数据传送的信号。

## 控制协程的并发数量
创建带缓冲的channel来限制并发的协程数量, 或者使用第三方库，创建协程池，
Jeffail/tunny
panjf2000/ants
```go
// main_chan.go
func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}
```

## go sync.Pool（并发安全）
作用：保存和复用临时对象，减少内存分配，降低 GC 压力。

sync.Pool 是可伸缩的，同时也是并发安全的，其大小仅受限于内存的大小。sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值。这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能。
sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理

### 如何使用
只需要实现 New 函数即可。对象池中没有对象时，将会调用 New 函数创建。
- Get() 用于从对象池中获取对象，因为返回值是 interface{}，因此需要类型转换
- Put() 则是在对象使用完毕后，返回对象池
```go
var studentPool = sync.Pool{
    New: func() interface{} { 
        return new(Student) 
    },
}

stu := studentPool.Get().(*Student)
json.Unmarshal(buf, stu)
studentPool.Put(stu)
```

例子
```go
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
        //buf.Reset() 相当于时执行了 buf.Truncate(0)，把 slice 的大小设置为 0，底层的 byte 数组没有释放
		//这里 buf.Reset() 是必须执行的，如果不执行，下次从 pool 中取出来再用时，就是在原来的 buf 上追加写了，会导致错误的结果。
		buf.Reset() 
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
		//buf.Reset()这里使用reset没必要，因为内存每次都是新申请的
	}
}

BenchmarkBufferWithPool-8    8778160    133 ns/op       0 B/op   0 allocs/op
BenchmarkBuffer-8             906572   1299 ns/op   10240 B/op   1 allocs/op
```
标准库当中的应用：fmt.Printf 的调用是非常频繁的，利用 sync.Pool 复用 pp 对象能够极大地提升性能，减少内存占用，同时降低 GC 压力。
```go
var ppFree = sync.Pool{
	New: func() interface{} { return new(pp) },
}

// newPrinter allocates a new pp struct or grabs a cached one.
func newPrinter() *pp {
	p := ppFree.Get().(*pp)
	p.panicking = false
	p.erroring = false
	p.wrapErrs = false
	p.fmt.init(&p.buf)
	return p
}
```

## go sync.Once
sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等

使用场景：多个goroutine并发的读取全局的配置文件，为其初始化，实际上只要执行一次就可以了
- 在这个例子中，声明了 2 个全局变量，once 和 config；
- config 是需要在 ReadConfig 函数中初始化的(将环境变量转换为 Config 结构体)，ReadConfig 可能会被并发调用。

如果 ReadConfig 每次都构造出一个新的 Config 结构体，既浪费内存，又浪费初始化时间。如果 ReadConfig 中不加锁，初始化全局变量 config 就可能出现并发冲突。
这种情况下，使用 sync.Once 既能够保证全局变量初始化时是线程安全的，又能节省内存和初始化时间。
```go
type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		if err != nil {
			config.Port = 8080 // default port
        }
        log.Println("init config")
	})
	return config
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}
```


### 原理实现
保证变量仅被初始化一次，需要有个标志来判断变量是否已初始化过，若没有则需要初始化。第二：线程安全，支持并发，无疑需要互斥锁来实现
```go
package sync

import (
    "sync/atomic"
)

type Once struct {
    done uint32
    m    Mutex
}

func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 0 {
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func()) {
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

## go sync.Cond
sync.Cond 条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine

### 实现
每个 Cond 实例都会关联一个锁 L, 当修改条件或者调用 Wait 方法时，必须加锁。
```go
type Cond struct {
        noCopy noCopy

        // L is held while observing or changing the condition
        L Locker

        notify  notifyList
        checker copyChecker
}
```
调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine，因此当前协程会阻塞在 Wait 方法调用的地方。
如果其他协程调用了 Signal 或 Broadcast 唤醒了该协程，那么 Wait 方法在结束阻塞时， 会重新给 c.L 加锁，并且继续执行 Wait 后面的代码。
```go
c.L.Lock() //调用wait会释放锁L， 当唤醒协程后会自动加锁，继续执行后面的代码
for !condition() {
    c.Wait()
}
... make use of condition ...
c.L.Unlock()
```

```go
var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	//c.L.Lock()是为了Wait函数内部并发安全的将当前协程加入Cond的通知队列，之后会解锁并挂起等待通知
	//c.L.Unlock()并不是配对先前你所见的那个c.L.Lock()，而是配对的wait()内部的加锁。wait()获得通知后，
	//并不是立即就退出函数了，wait内部接着是去抢占c.L锁的， 以确保对共享资源的并发安全访问。
	//只有当前协程抢占到c.L锁后才会从wait()退出。之后处理你的业务逻辑，最后需要你显式的c.L.Unlock()去解锁wait()退出前的锁定
	c.L.Lock() 
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

$ go run main.go
2021/01/14 23:18:20 writer starts writing
2021/01/14 23:18:21 writer wakes all
2021/01/14 23:18:21 reader2 starts reading
2021/01/14 23:18:21 reader3 starts reading
2021/01/14 23:18:21 reader1 starts reading
```
package toolbox

import (
	"time"
	"os"
	"io"
	"runtime/pprof"
	"strconv"
	"fmt"
	"log"
	"runtime"
	"path"
	"runtime/debug"
)

//导入包的时间
var startTime = time.Now()
//返回调用者所在的进程ID
var pid int

func init() {
	pid = os.Getpid()
}

//ProcessInput parse input command string
func ProcessInput(input string,w io.Writer){
	switch input {
	case "lookup goroutine":
		//查询Go所有程的调用栈空间
		p := pprof.Lookup("goroutine")
		p.WriteTo(w,2)
	case "lookup heap":
		//所有堆分配的采样
		p := pprof.Lookup("heap")
		p.WriteTo(w,2)
	case "lookup threadcreate":
		//导致新的OS线程创建的调用栈总计
		p := pprof.Lookup("threadcreate")
		p.WriteTo(w,2)
	case "lookup block":
		//导致同步原语水平的阻塞的调用栈踪迹
		p:= pprof.Lookup("block")
		p.WriteTo(w,2)
	case "get cpuprof":
		//CPU prof
		GetCPUProfile(w)
	case "get memprof":
		//内存 prof
		MemProf(w)
	case "gc summary":
		//GCsummary计算
		PrintGCSummary(w)
	}
}

// MemProf record memory profile in pprof
func MemProf(w io.Writer) {
	filename := "mem-"+strconv.Itoa(pid) + ".memprof"
	if f,err := os.Create(filename);err != nil {
		fmt.Fprintf(w,"create file %s error %s \n",filename,err.Error())
		log.Fatal("record heap profile failed: ",err)
	}else{
		//进行GC
		runtime.GC()
		//相当于 pprof.Lookup("heap").WriteTo(w,0)
		pprof.WriteHeapProfile(f)
		//关闭文件句柄
		f.Close()
		fmt.Fprintf(w,"create heap profile %s \n",filename)
		_,fl := path.Split(os.Args[0])
		fmt.Fprintf(w,"Now you can use this to check it:go tool pprof %s %s \n",fl,filename)
	}
}
// GetCPUProfile start cpu profile monitor
func GetCPUProfile(w io.Writer){
	sec := 30
	filename := "cpu-"+strconv.Itoa(pid) + ".pprof"
	f,err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(w,"Could not enable CPU profiling: %s\n",err)
		log.Fatal("record cpu profile failed: ",err)
	}
	//为当前进程开启CPU profile
	pprof.StartCPUProfile(f)
	time.Sleep(time.Duration(sec)*time.Second)
	//停止当前的CPU profile
	pprof.StopCPUProfile()

	fmt.Fprintf(w,"create cpu profile %s \n",filename)
	_,fl := path.Split(os.Args[0])
	fmt.Fprintf(w,"Now you can use this to check it: go tool pprof %s %s\n",fl,filename)
}
// PrintGCSummary print gc information to io.Writer
func PrintGCSummary(w io.Writer){
	memStats := &runtime.MemStats{}
	//将内存申请和分配的信息的统计信息填入m
	runtime.ReadMemStats(memStats)
	//
	gcstats :=&debug.GCStats{PauseQuantiles:make([]time.Duration,100)}
	//ReadGCStats将所有的垃圾收集信息填入stats，
	debug.ReadGCStats(gcstats)
	printGC(memStats,gcstats,w)
}

//type GCStats struct {
//	LastGC         time.Time       // 最近一次垃圾收集的时间
//	NumGC          int64           // 垃圾收集的次数
//	PauseTotal     time.Duration   // 所有暂停收集垃圾消耗的总时间
//	Pause          []time.Duration // 每次暂停收集垃圾的消耗的时间
//	PauseQuantiles []time.Duration
//}

//type MemStats struct {
//	// 一般统计
//	Alloc      uint64 // 已申请且仍在使用的字节数
//	TotalAlloc uint64 // 已申请的总字节数（已释放的部分也算在内）
//	Sys        uint64 // 从系统中获取的字节数（下面XxxSys之和）
//	Lookups    uint64 // 指针查找的次数
//	Mallocs    uint64 // 申请内存的次数
//	Frees      uint64 // 释放内存的次数
//	// 主分配堆统计
//	HeapAlloc    uint64 // 已申请且仍在使用的字节数
//	HeapSys      uint64 // 从系统中获取的字节数
//	HeapIdle     uint64 // 闲置span中的字节数
//	HeapInuse    uint64 // 非闲置span中的字节数
//	HeapReleased uint64 // 释放到系统的字节数
//	HeapObjects  uint64 // 已分配对象的总个数
//	// L低层次、大小固定的结构体分配器统计，Inuse为正在使用的字节数，Sys为从系统获取的字节数
//	StackInuse  uint64 // 引导程序的堆栈
//	StackSys    uint64
//	MSpanInuse  uint64 // mspan结构体
//	MSpanSys    uint64
//	MCacheInuse uint64 // mcache结构体
//	MCacheSys   uint64
//	BuckHashSys uint64 // profile桶散列表
//	GCSys       uint64 // GC元数据
//	OtherSys    uint64 // 其他系统申请
//	// 垃圾收集器统计
//	NextGC       uint64 // 会在HeapAlloc字段到达该值（字节数）时运行下次GC
//	LastGC       uint64 // 上次运行的绝对时间（纳秒）
//	PauseTotalNs uint64
//	PauseNs      [256]uint64 // 近期GC暂停时间的循环缓冲，最近一次在[(NumGC+255)%256]
//	NumGC        uint32
//	EnableGC     bool
//	DebugGC      bool
//	// 每次申请的字节数的统计，61是C代码中的尺寸分级数
//	BySize [61]struct {
//		Size    uint32
//		Mallocs uint64
//		Frees   uint64
//	}
//}
func printGC(memStats *runtime.MemStats,gcstats *debug.GCStats,w io.Writer){
	if gcstats.NumGC > 0{
		//最后一次GC所耗时间
		lastPause := gcstats.Pause[0]
		//运行时间
		elapsed := time.Now().Sub(startTime)
		//PauseTotal所有GC总时间 /时间差 *100
		overhead := float64(gcstats.PauseTotal)/float64(elapsed) *100
		//TotalAlloc已申请的总字节数/ 时间差 s
		allocatedRate := float64(memStats.TotalAlloc) / elapsed.Seconds()
		//GC次数 最后一次GC所耗时间 总GC平均消耗时间 GC总时间/时间差*100
		//已申请但仍在使用的内存 从系统中获取的字节数
		fmt.Fprintf(w,"NumGC:%d Pause:%s Pause(Avg):%s Overhead:%3.2f%% Alloc:%s Sys:%s Alloc(Rate):%s/s Histogram:%s %s %s \n",
			gcstats.NumGC,
			toS(lastPause),
			toS(avg(gcstats.Pause)),
			overhead,
			toH(memStats.Alloc),
			toH(memStats.Sys),
			toH(uint64(allocatedRate)),
			toS(gcstats.PauseQuantiles[94]),
			toS(gcstats.PauseQuantiles[98]),
			toS(gcstats.PauseQuantiles[99]),
			)
	}else{
		// while GC has disabled
		elapsed := time.Now().Sub(startTime)
		allocateRate := float64(memStats.TotalAlloc) / elapsed.Seconds()

		fmt.Fprintf(w,"Alloc:%s Sys:%s Alloc(Rate):%s/s\n",
			toH(memStats.Alloc),
			toH(memStats.Sys),
			toH(uint64(allocateRate)),
		)
	}
}

//求平均值
func avg(items []time.Duration)time.Duration{
	var sum time.Duration
	for _, item := range items{
		sum += item
	}
	return time.Duration(int64(sum)/int64(len(items)))
}

//大小转换
// B ==> K ==> M ==> G
func toH(bytes uint64)string{
	switch  {
	case bytes <1024:
		return fmt.Sprintf("%dB",bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2fK",float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2fM", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2fG",float64(bytes)/1024/1024/1024)
	}
}
//时间片转换字符串 保留两位小数
func toS(d time.Duration) string {
	u := uint64(d)
	if u < uint64(time.Second) {
		switch {
		case u==0 :
			return "0"
		case u< uint64(time.Microsecond):
			return fmt.Sprintf("%.2fns",float64(u))
		case u<uint64(time.Millisecond):
			return fmt.Sprintf("%.2fus",float64(u)/1000)
		default:
			return fmt.Sprintf("%.2fms",float64(u)/1000/1000)
		}
	}else {
		switch {
		case u<uint64(time.Minute):
			return fmt.Sprintf("%.2fs",float64(u)/1000/1000/1000)
		case u<uint64(time.Hour):
			return fmt.Sprintf("%.2fm",float64(u)/1000/1000/1000/60)
		default:
			return fmt.Sprintf("%.2fh",float64(u)/1000/1000/1000/60/60)
		}
	}
}


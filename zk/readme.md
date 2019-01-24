# golang zookeeper client chinese version

```go
const (
    FlagEphemeral = 1 //暂时的节点
    FlagSequence  = 2 //序列的节点
)
```

- `func FLWRuok(servers []string, timeout time.Duration) []bool`

&emsp;&emsp;FLWRuok是FourLetterWord辅助函数。 特别是，此函数从每个服务器中提取ruok输出。
```go
var hosts = []string{"192.168.159.128"}
func TestZK(t *testing.T) {
	bools := zk.FLWRuok(hosts, 5*time.Second)
	fmt.Println(bools)
}
//output
//[true]
```
- `func FormatServers(servers []string) []string`

&emsp;&emsp;检查servers这个string类型的切片，确保他们都是addr:port这种格式的，如果只
有addr，`FormatServers`会为其补充一个默认的端口（默认端口是2181）
```go
var hosts = []string{"192.168.159.128"}
func TestZK(t *testing.T) {
	results := zk.FormatServers(hosts)
	fmt.Println(results)
}
//output
//[192.168.159.128:2181]
```
- `func WithDialer(dialer Dialer) connOption`

&emsp;&emsp;根据一个`Dialer`参数（无默认值），返回链接选项
```go
func DiaTestfunc(network, address string, timeout time.Duration) (net.Conn, error) {
	return nil, nil
}
func TestZK(t *testing.T) {
	results := zk.WithDialer(DiaTestfunc)
	fmt.Println(results)
}
//Output like an address
//0x5189b0
```
- `func WithEventCallback(cb EventCallback) connOption`

&emsp;&emsp;返回一个指定的回调时间的链接选项，但是这个回调函数不能阻塞，否则会影响到zk的goroutines
- `func WithHostProvider(hostProvider HostProvider) connOption`

&emsp;&emsp;根据`HostProvider`（无默认值），返回一个链接选项
- `func WithLogInfo(logInfo bool) connOption`

&emsp;&emsp;根据`logInfo`是否记录信息消息，返回一个链接选项
- `func WithLogger(logger Logger) connOption`

&emsp;&emsp;根据`Logger`（无默认值），返回一个链接选项
- `func WithMaxBufferSize(maxBufferSize int) connOption`

&emsp;&emsp;通过`maxBufferSize`来设置从zk服务端读取和解码的数据包的最大缓冲区大小，Java的
zk客户端的标准是限制在1mb，为了向后兼容，go的zk客户端并没有做限制，除非通过这个函数修改这个变量，
`maxBufferSize`为0和负数表示没有限制

&emsp;&emsp;为了防止zk中面临潜在恶意数据的资源耗尽。它通常应该与zk服务端设置一样（也默认为1mb）
，以便客户端和服务端在单个znode中的数据大小和事务的总大小等内容一致

&emsp;&emsp;对于生产系统，应将其设置为一个合理的值（理想情况下是与zk服务端的配置相匹配）。对于
ops工具，可以很方便地将其调整为更大的限制，以便在zk树中清除有问题的状态。例如，如果单个znode具
有大量子节点，则对该节点的操作的响应可能超过设置的缓冲区大小引起客户端报错。而在后面的清理zk树（通过
删除多余的子节点）的唯一方法是使用配置了较大缓冲区大小的客户端，该客户端可以成功查询所有子节点，
并得到数据，然后将其删除。（注意，还有其他工具可以在客户端中列出所有子节点却不需要增加缓冲区大小，
但它们是通过检查zk服务端的事务日志来枚举子节点而不是向服务端发送在线的请求来列出子节点。
- `func WithMaxConnBufferSize(maxBufferSize int) connOption`
    
&emsp;&emsp;`WithMaxConnBufferSize`设置用于向zk服务端发送、编码数据包的最大缓冲区大小。
java的标准Zookeepeer客户端默认限制为1mb。 此选项应用于非标准服务器设置，其中znode大于默认值1mb。

##### type ACL
```go
type ACL struct {
    Perms  int32
    Scheme string
    ID     string
}
```
权限常量

    PermRead = 1 << iota
    PermWrite
    PermCreate
    PermDelete
    PermAdmin
    PermAll = 0x1f
    
- `func AuthACL(perms int32) []ACL`

&emsp;&emsp;根据给的权限返回一个ACL切片，经过验证的ACL
```go
func TestZK(t *testing.T) {
	results := zk.AuthACL(zk.PermRead)
	fmt.Println(results)
}
//Output
//[{1 auth }]
```
- `func DigestACL(perms int32, user, password string) []ACL`
&emsp;&emsp;
```go
func TestZK(t *testing.T) {
	results := zk.DigestACL(zk.PermRead, "testuser", "123456")
	fmt.Println(results)
}
//Output (password be encrypt)
//[{1 digest testuser:o0epvaAv9Ypg33bJQFRZ+POkwLs=}]
```
- `func WorldACL(perms int32) []ACL`

&emsp;&emsp;根据给的权限返回一个ACL切片，代表所有用户
```go
func TestZK(t *testing.T) {
	results := zk.WorldACL(zk.PermRead)
	fmt.Println(results)
}
//Output
//[{1 world anyone}]
```
##### type CheckVersionRequest
```go
type PathVersionRequest struct {
    Path    string
    Version int32
}
```

##### type Conn
```go
type Conn struct {
    // contains filtered or unexported fields
    	lastZxid         int64
    	sessionID        int64
    	state            State // must be 32-bit aligned
    	xid              uint32
    	sessionTimeoutMs int32 // session timeout in milliseconds
    	passwd           []byte
    
    	dialer         Dialer
    	hostProvider   HostProvider
    	serverMu       sync.Mutex // protects server
    	server         string     // remember the address/port of the current server
    	conn           net.Conn
    	eventChan      chan Event
    	eventCallback  EventCallback // may be nil
    	shouldQuit     chan struct{}
    	pingInterval   time.Duration
    	recvTimeout    time.Duration
    	connectTimeout time.Duration
    	maxBufferSize  int
    
    	creds   []authCreds
    	credsMu sync.Mutex // protects server
    
    	sendChan     chan *request
    	requests     map[int32]*request // Xid -> pending request
    	requestsLock sync.Mutex
    	watchers     map[watchPathType][]chan Event
    	watchersLock sync.Mutex
    	closeChan    chan struct{} // channel to tell send loop stop
    
    	// Debug (used by unit tests)
    	reconnectLatch   chan struct{}
    	setWatchLimit    int
    	setWatchCallback func([]*setWatchesRequest)
    	// Debug (for recurring re-auth hang)
    	debugCloseRecvLoop bool
    	debugReauthDone    chan struct{}
    
    	logger  Logger
    	logInfo bool // true if information messages are logged; false if only errors are logged
    
    	buf []byte
}
```
- `func Connect(servers []string, sessionTimeout time.Duration, options ...connOption) (*Conn, <-chan Event, error)`

&emsp;&emsp;Connect建立与zookeeper服务端的连接池的新链接。 会话超时时间。 在设置的会话超时时间内，可以重新建立与其他服务端的
连接并保持相同的会话。 这意味着维护任何短暂的节点和监听
```go
var hosts = []string{"192.168.159.128"}

func TestZK(t *testing.T) {
	conn, evChan, err := zk.Connect(hosts, 5*time.Second)
	defer conn.Close()
	fmt.Printf("evChan [%v] \n", evChan)
	fmt.Printf("err [%v] \n", err)
}
//Output
/*
evChan [0xc0420800c0] 
err [<nil>] 
2019/01/24 13:58:32 Connected to 192.168.159.128:2181
2019/01/24 13:58:32 Authenticated: id=72058159611445269, timeout=5000
2019/01/24 13:58:32 Re-submitting `0` credentials after reconnect
2019/01/24 13:58:32 Recv loop terminated: err=EOF
2019/01/24 13:58:32 Send loop terminated: err=<nil>
*/
```
- `func ConnectWithDialer(servers []string, sessionTimeout time.Duration, dialer Dialer) (*Conn, <-chan Event, error)`

&emsp;&emsp;使用自定义的Dialer奖励一个新连接到一个zk服务端的连接池，有关会话超时的详细信息，请参阅connect，但是
 不推荐使用此方法，并为此提供兼容性：请改用WithDialer函数
- `func (c *Conn) AddAuth(scheme string, auth []byte) error`

&emsp;&emsp;添加用户
```go
err := conn.AddAuth("testuser", []byte("123456"))
fmt.Println(err)

//Output
//zk: client authentication failed
```
- `func (c *Conn) Children(path string) ([]string, *Stat, error)`

&emsp;&emsp;返回`path`下的子节点
```go
var hosts = []string{"192.168.159.128"}

func TestZK(t *testing.T) {
	conn, _, _ := zk.Connect(hosts, 5*time.Second)
	defer conn.Close()
	childs, stat, err := conn.Children("/")
	fmt.Printf("childrens [%v] \n", childs)
	fmt.Printf("stat [%v] \n", *stat)
	fmt.Printf("err [%v] \n", err)
}
//Output:
//childrens [[zookeeper test root]] 
//stat [{0 0 0 0 0 1 0 0 0 3 35}] 
//err [<nil>] 
```
- `func (c *Conn) ChildrenW(path string) ([]string, *Stat, <-chan Event, error)`

&emsp;&emsp;返回`path`下的子节点监听
- `func (c *Conn) Close()`

&emsp;&emsp;关闭连接
- `func (c *Conn) Create(path string, data []byte, flags int32, acl []ACL) (string, error)`

&emsp;&emsp;创建一个节点
```go
var hosts = []string{"192.168.159.128"}
func TestZK(t *testing.T) {
	conn, _, _ := zk.Connect(hosts, 5*time.Second)
	defer conn.Close()
	var data = []byte("hello zk")
	var flags = zk.FlagSequence //0永久 1暂时 2有序永久 3有序短暂
	var acls = zk.WorldACL(zk.PermAll) //控制访问权限模式
	result, err := conn.Create("/test/test01", data, int32(flags), acls)
	fmt.Printf("childrens [%v] \n", result)
	fmt.Printf("err [%v] \n", err)
}
//Output:
//childrens [/test/test010000000000] 
//err [<nil>] 
```
- `func (c *Conn) CreateProtectedEphemeralSequential(path string, data []byte, acl []ACL) (string, error)`

&emsp;&emsp;如果zk服务端在创建一个节点后崩溃了，`CreateProtectedEphemeralSequential`会修复这个竞争条件，
重新连接的会话可能依然有效，并且这个暂时的节点依然存在，所以在重新连接后我们需要在guid操作前检查这个节点是否存在
- `func (c *Conn) Delete(path string, version int32) error`

&emsp;&emsp;删除节点
- `func (c *Conn) Exists(path string) (bool, *Stat, error)`

&emsp;&emsp;节点是否存在
- `func (c *Conn) ExistsW(path string) (bool, *Stat, <-chan Event, error)`

&emsp;&emsp;节点是否存在监听
- `func (c *Conn) Get(path string) ([]byte, *Stat, error)`

&emsp;&emsp;获取节点数据
- `func (c *Conn) GetACL(path string) ([]ACL, *Stat, error)`

&emsp;&emsp;获取节点权限
- `func (c *Conn) GetW(path string) ([]byte, *Stat, <-chan Event, error)`

&emsp;&emsp;返回节点数据，并且对节点设置监听
- `func (c *Conn) Multi(ops ...interface{}) ([]MultiResponse, error)`

&emsp;&emsp;`Multi`执行多项操作，但是`ops`必须是`*CreateRequest`, `*DeleteRequest`,
 `*SetDataRequest`, or `*CheckVersionRequest`
- `func (c *Conn) Server() string`

&emsp;&emsp;返回conn当前的或者最近连接的zk server的name
- `func (c *Conn) SessionID() int64`

&emsp;&emsp;返回conn的sessionId
- `func (c *Conn) Set(path string, data []byte, version int32) (*Stat, error)`

&emsp;&emsp;设置节点数据
- `func (c *Conn) SetACL(path string, acl []ACL, version int32) (*Stat, error)`

&emsp;&emsp;设置节点权限
- `func (c *Conn) SetLogger(l Logger)`

&emsp;&emsp;设置节点打印日志，`Logger`是一个接口
- `func (c *Conn) State() State`

&emsp;&emsp;获取当前连接的状态
- `func (c *Conn) Sync(path string) (string, error)`

&emsp;&emsp;同步啥？？？
```go
var hosts = []string{"192.168.159.128"}
func TestZK(t *testing.T) {
	conn, _, _ := zk.Connect(hosts, 5*time.Second)
	defer conn.Close()
	result, err := conn.Sync("/test/test01")
	fmt.Printf("result [%v] \n", result)
	fmt.Printf("err [%v] \n", err)
}
//Output
//result [/test/test01] 
//err [<nil>]
```
##### type CreateRequest
```go
type CreateRequest struct {
    Path  string
    Data  []byte
    Acl   []ACL
    Flags int32
}
```
##### type DNSHostProvider
```go
type DNSHostProvider struct {
    // contains filtered or unexported fields
}
```
&emsp;&emsp;DNSHostProvider是默认的HostProvider。 它一般是与Java的StaticHostProvider（What 
the fuck？？？）匹配，在调用Init期间从DNS解析host。 它可以很容易地扩展到定期重新查询DNS或连接有问题。
- `func (hp *DNSHostProvider) Connected()`

&emsp;&emsp;`Connected`这个`DNSHostProvider`是不是成功连接
- `func (hp *DNSHostProvider) Init(servers []string) error`

&emsp;&emsp;第一次调用`Init`，并在连接字符串中指定服务器，使用DNS查找每个服务器的地址，然后将他们打乱
- `func (hp *DNSHostProvider) Len() int`

&emsp;&emsp;返回可用的服务器的数量
- `func (hp *DNSHostProvider) Next() (server string, retryStart bool)`

&emsp;&emsp;返回下一个连接的server，如果我们一直的服务器没有调用`Connected()`函数那
么`retryStart`为True
##### type DeleteRequest
```go
type DeleteRequest PathVersionRequest
```
##### type Dialer
```go
type Dialer func(network, address string, timeout time.Duration) (net.Conn, error)
```
##### type ErrCode
```go
type ErrCode int32
```
##### type ErrMissingServerConfigField
```go
type ErrMissingServerConfigField string
```
- `func (e ErrMissingServerConfigField) Error() string`
##### type Event
```go
type Event struct {
    Type   EventType
    State  State
    Path   string // For non-session events, the path of the watched node.
    Err    error
    Server string // For connection events
}
```
##### type EventCallback
```go
type EventCallback func(Event)
```
&emsp;&emsp;`EventCallback`是一个当对应事件发生的时候会被调用的函数
##### type EventType
```go
type EventType int32
```
```go
const (
    EventNodeCreated         EventType = 1
    EventNodeDeleted         EventType = 2
    EventNodeDataChanged     EventType = 3
    EventNodeChildrenChanged EventType = 4

    EventSession     EventType = -1
    EventNotWatching EventType = -2
)
```
- `func (t EventType) String() string`
##### type HostProvider
```go
type HostProvider interface {
    // Init is called first, with the servers specified in the connection string.
    Init(servers []string) error
    // Len returns the number of servers.
    Len() int
    // Next returns the next server to connect to. retryStart will be true if we've looped through
    // all known servers without Connected() being called.
    Next() (server string, retryStart bool)
    // Notify the HostProvider of a successful connection.
    Connected()
}
```
##### type Lock
```go
type Lock struct {
    // contains filtered or unexported fields
}
```
&emsp;&emsp;互斥锁
- `func NewLock(c *Conn, path string, acl []ACL) *Lock`

&emsp;&emsp;`NewLock`根据conn、path、权限生成一个互斥锁实例，path必须是一个节点，而且
只被这个锁使用，锁初始时unlocked除非调用Lock()
- `func (l *Lock) Lock() error`

&emsp;&emsp;`Lock()`获取一个新的锁，它会等到获得这个锁或者一个错误的返回，如果这个实例已经
获得了锁，那么就会返回一个错误`ErrDeadlock `
- `func (l *Lock) Unlock() error`

&emsp;&emsp;释放锁，如果没有锁或者已经释放锁就会返回一个错误`ErrNotLocked`
##### type Logger
```go
type Logger interface {
    Printf(string, ...interface{})
}
```
&emsp;&emsp;DefaultLogger uses the stdlib log package for logging.
##### type Mode
```go
type Mode uint8
```
```go
const (
    ModeUnknown    Mode = iota
    ModeLeader     Mode = iota
    ModeFollower   Mode = iota
    ModeStandalone Mode = iota
)
```
- `func (m Mode) String() string`
##### type MultiResponse
```go
type MultiResponse struct {
    Stat   *Stat
    String string
    Error  error
}
```
##### type PathVersionRequest
```go
type PathVersionRequest struct {
    Path    string
    Version int32
}
```
##### type Server
```go
type Server struct {
    JarPath        string
    ConfigPath     string
    Stdout, Stderr io.Writer
    // contains filtered or unexported fields
}
```
- `func (srv *Server) Start() error`
- `func (srv *Server) Stop() error`
##### type ServerClient
```go
type ServerClient struct {
    Queued        int64
    Received      int64
    Sent          int64
    SessionID     int64
    Lcxid         int64
    Lzxid         int64
    Timeout       int32
    LastLatency   int32
    MinLatency    int32
    AvgLatency    int32
    MaxLatency    int32
    Established   time.Time
    LastResponse  time.Time
    Addr          string
    LastOperation string // maybe?
    Error         error
}
```
##### type ServerClients
```go
type ServerClients struct {
    Clients []*ServerClient
    Error   error
}
```
- `func FLWCons(servers []string, timeout time.Duration) ([]*ServerClients, bool)`
##### type ServerConfig
```go
type ServerConfig struct {
    TickTime                 int    // Number of milliseconds of each tick
    InitLimit                int    // Number of ticks that the initial synchronization phase can take
    SyncLimit                int    // Number of ticks that can pass between sending a request and getting an acknowledgement
    DataDir                  string // Direcrory where the snapshot is stored
    ClientPort               int    // Port at which clients will connect
    AutoPurgeSnapRetainCount int    // Number of snapshots to retain in dataDir
    AutoPurgePurgeInterval   int    // Purge task internal in hours (0 to disable auto purge)
    Servers                  []ServerConfigServer
}
```
- `func (sc ServerConfig) Marshall(w io.Writer) error`
##### type ServerConfigServer
```go
type ServerConfigServer struct {
    ID                 int
    Host               string
    PeerPort           int
    LeaderElectionPort int
}
```
##### type ServerStats
```go
type ServerStats struct {
    Sent        int64
    Received    int64
    NodeCount   int64
    MinLatency  int64
    AvgLatency  int64
    MaxLatency  int64
    Connections int64
    Outstanding int64
    Epoch       int32
    Counter     int32
    BuildTime   time.Time
    Mode        Mode
    Version     string
    Error       error
}
```
- `func FLWSrvr(servers []string, timeout time.Duration) ([]*ServerStats, bool)`
##### type SetDataRequest
```go
type SetDataRequest struct {
    Path    string
    Data    []byte
    Version int32
}
```
##### type Stat
```go
type Stat struct {
    Czxid          int64 // The zxid of the change that caused this znode to be created.
    Mzxid          int64 // The zxid of the change that last modified this znode.
    Ctime          int64 // The time in milliseconds from epoch when this znode was created.
    Mtime          int64 // The time in milliseconds from epoch when this znode was last modified.
    Version        int32 // The number of changes to the data of this znode.
    Cversion       int32 // The number of changes to the children of this znode.
    Aversion       int32 // The number of changes to the ACL of this znode.
    EphemeralOwner int64 // The session id of the owner of this znode if the znode is an ephemeral node. If it is not an ephemeral node, it will be zero.
    DataLength     int32 // The length of the data field of this znode.
    NumChildren    int32 // The number of children of this znode.
    Pzxid          int64 // last modified children
}
```
##### type State
```go
type State int32
```
```go
const (
    StateUnknown           State = -1
    StateDisconnected      State = 0
    StateConnecting        State = 1
    StateAuthFailed        State = 4
    StateConnectedReadOnly State = 5
    StateSaslAuthenticated State = 6
    StateExpired           State = -112

    StateConnected  = State(100)
    StateHasSession = State(101)
)
```
- `func (s State) String() string`
##### type TestCluster
```go
type TestCluster struct {
    Path    string
    Servers []TestServer
}
```
- `func StartTestCluster(size int, stdout, stderr io.Writer) (*TestCluster, error)`
- `func (tc *TestCluster) Connect(idx int) (*Conn, error)`
- `func (tc *TestCluster) ConnectAll() (*Conn, <-chan Event, error)`
- `func (tc *TestCluster) ConnectAllTimeout(sessionTimeout time.Duration) (*Conn, <-chan Event, error)`
- `func (tc *TestCluster) ConnectWithOptions(sessionTimeout time.Duration, options ...connOption) (*Conn, <-chan Event, error)`
- `func (tc *TestCluster) StartAllServers() error`
- `func (tc *TestCluster) StartServer(server string)`
- `func (tc *TestCluster) Stop() error`
- `func (tc *TestCluster) StopAllServers() error`
- `func (tc *TestCluster) StopServer(server string)`
##### type TestServer
```go
type TestServer struct {
    Port int
    Path string
    Srv  *Server
}
```
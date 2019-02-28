package mdb

import (
	"container/heap"
	"fmt"
	"log"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	*mgo.Session
	ref   int
	index int
}

type SessionHeap []*Session

func (h SessionHeap) Len() int {
	return len(h)
}

func (h SessionHeap) Less(i, j int) bool {
	return h[i].ref < h[j].ref
}

func (h SessionHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *SessionHeap) Push(s interface{}) {
	s.(*Session).index = len(*h)
	*h = append(*h, s.(*Session))
}

func (h *SessionHeap) Pop() interface{} {
	l := len(*h)
	s := (*h)[l-1]
	s.index = -1
	*h = (*h)[:l-1]
	return s
}

type DialContext struct {
	sync.Mutex
	sessions SessionHeap
}

//goroutine safe
func Dial(url string, sessionNum int) (*DialContext, error) {
	c, err := DiaWithTimeout(url, sessionNum, 10*time.Second, 5*time.Second)
	return c, err
}

//goroutine safe
func DiaWithTimeout(url string, sessionNum int, dialTimeout, timeout time.Duration) (*DialContext, error) {
	if sessionNum <= 0 {
		sessionNum = 100
		log.Println("DiaWithTimeout invalid sessionNum,reset to ", sessionNum)
	}

	s, err := mgo.DialWithTimeout(url, dialTimeout)
	if err != nil {
		return nil, err
	}
	s.SetSyncTimeout(timeout)
	s.SetSocketTimeout(timeout)

	c := new(DialContext)

	c.sessions = make(SessionHeap, sessionNum)
	c.sessions[0] = &Session{s, 0, 0}

	for i := 1; i < sessionNum; i++ {
		c.sessions[i] = &Session{s.New(), 0, i}
	}
	heap.Init(&c.sessions)
	return c, nil

}

//goroutine safe
func (c *DialContext) Close() {
	c.Lock()
	for _, s := range c.sessions {
		s.Close()
		if s.ref != 0 {
			fmt.Println("error session ref = ", s.ref)
		}
	}
	c.Unlock()
}

//goroutine safe
func (c *DialContext) Ref() *Session {
	c.Lock()
	defer c.Unlock()

	s := c.sessions[0]
	if s.ref == 0 {
		s.Refresh()
	}
	s.ref++
	heap.Fix(&c.sessions, 0)
	return s
}

//goroutine safe
func (c *DialContext) UnRef(s *Session) {
	c.Lock()
	defer c.Unlock()

	s.ref--
	heap.Fix(&c.sessions, s.index)
}

//goroutine safe
func (c *DialContext) EnsureCounter(db, collection, id string) error {
	s := c.Ref()
	defer c.UnRef(s)

	err := s.DB(db).C(collection).Insert(bson.M{
		"_id": id,
		"seq": 0,
	})

	if mgo.IsDup(err) {
		return nil
	} else {
		return err
	}
}

//goroutine safe
func (c *DialContext) NextSeq(db, collection, id string) (int, error) {
	s := c.Ref()
	defer c.UnRef(s)

	var res struct {
		Seq int
	}

	_, err := s.DB(db).C(collection).FindId(id).Apply(mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 2}},
		ReturnNew: true,
	}, &res)

	return res.Seq, err
}

// goroutine safe
//创建唯一索引
func (c *DialContext) EnsureIndex(db, collection string, key []string) error {
	s := c.Ref()
	defer c.UnRef(s)

	return s.DB(db).C(collection).EnsureIndex(mgo.Index{
		Key:    key,
		Unique: false,
		Sparse: true,
	})
}

// goroutine safe
//添加唯一索引
func (c *DialContext) EnsureUniqueIndex(db, collection string, key []string) error {
	s := c.Ref()
	defer c.UnRef(s)

	return s.DB(db).C(collection).EnsureIndex(mgo.Index{
		Key:    key,
		Unique: true,
		Sparse: true,
	})
}

// goroutine safe
//得到索引
func (c *DialContext) GetIndexs(db, collection string) ([]mgo.Index, error) {
	s := c.Ref()
	defer c.UnRef(s)

	return s.DB(db).C(collection).Indexes()
}

//goroutine safe
//存库
func (c *DialContext) SendData(db, collection, id string, data interface{}) error {
	s := c.Ref()
	defer c.UnRef(s)

	tbl := s.DB(db).C(collection)

	_, err := tbl.UpsertId(id, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *DialContext) Find(db, collection string, data bson.M) (result interface{}, err error) {
	s := c.Ref()
	defer c.UnRef(s)

	err = s.DB(db).C(collection).Find(data).One(&result)
	return
}

func (c *DialContext) Remove(db, collection string, query bson.M) error {
	s := c.Ref()
	defer c.UnRef(s)

	err := s.DB(db).C(collection).Remove(query)
	return err
}

func getNewID() string {
	objid := bson.NewObjectId()
	id := fmt.Sprintf(`%x`, string(objid))
	return id
}



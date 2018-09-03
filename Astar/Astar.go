package main

import (
	"strconv"
	"fmt"
	"strings"
	"math"
	"container/heap"
	"time"
)
type Openll []*_AstarPoint

func (self Openll)Len() int            {return len(self)}
func (self Openll)Swap(i,j int)        {self[i],self[j] = self[j],self[i]}
func (self Openll)Less(i,j int)(bool)  {return self[i].fVal<self[j].fVal}
func (self *Openll)Push(x interface{}) {*self = append(*self,x.(*_AstarPoint))}
func (self *Openll)Pop()(interface{}){
	old := *self
	n := self.Len()
	x := old[n-1]
	*self  = old[0:n-1]
	return x
}

func point2string(p _Point)(key string){
	key = strconv.Itoa(p.X)+","+strconv.Itoa(p.Y)
	return
}
//节点信息
type _Point struct {
	X,Y int
	View string
}
//地图结构
type Mapconf struct {
	points [][]_Point //地图信息
	blocks map[string]*_Point //障碍物信息
	maxX int //长
	maxY int //宽
}

//点信息
type _AstarPoint struct {
	_Point //节点
	father *_AstarPoint //父节点信息
	gVal int //g量、与父节点的距离计算
	hVal int //h量、与终点的距离计算
	fVal int //g与h 的和
}


//初始化地图信息
func InitMap(inputs []string)(m Mapconf){
	m.points = make([][]_Point, len(inputs))
	m.blocks = make(map[string]*_Point)

	for i,row := range inputs {
		cols := strings.Split(row," ")
		m.points[i] = make([]_Point,len(cols))
		for k,v := range cols{
			m.points[i][k] = _Point{i,k,v}
			if v == "X" || v=="x"{
				m.blocks[point2string(m.points[i][k])] = &m.points[i][k]
			}
		}
	}
	m.maxX = len(m.points)
	m.maxY = len(m.points[0])
	return
}
//得到周围的点
func (m *Mapconf)GetOtherPoint(midpoint *_Point)(otherpoint []*_Point){
	if x,y := midpoint.X-1,midpoint.Y-1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X-1,midpoint.Y;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X-1,midpoint.Y+1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X,midpoint.Y-1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X,midpoint.Y+1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X+1,midpoint.Y-1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X+1,midpoint.Y;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}
	if x,y := midpoint.X+1,midpoint.Y+1;x>=0 && x<m.maxX && y>=0 && y<m.maxY{
		otherpoint = append(otherpoint,&m.points[x][y])
	}

	return
}


func NewAstarPoint(p *_Point,father *_AstarPoint,end *_AstarPoint)(ap *_AstarPoint){
	ap = &_AstarPoint{*p,father,0,0,0,}
	if end != nil{
		ap.CalcuFval(end._Point)
	}
	return ap
}

//计算G量
func (this *_AstarPoint)CalcuGVal()(g int){
	if this.father != nil {
		deltaX := math.Abs(float64(this.father.X-this.X))
		deltaY := math.Abs(float64(this.father.Y-this.Y))
		if deltaX == 1 && deltaY ==0{
			this.gVal = 10
		}else if deltaX == 0 && deltaY == 1{
			this.gVal = 10
		}else if deltaX == 1 && deltaY == 1{
			this.gVal = 14
		}
	}
	return this.gVal
}

//计算h量
//曼哈顿算法
func (this *_AstarPoint)CalcuHValM(end _Point)(h int){
	this.hVal = int(math.Abs(float64(end.X-this.X))+math.Abs(float64(end.Y-this.Y)))
	return this.hVal
}
//几何算法
func (this *_AstarPoint)CalcuHValE(end _Point)(h int){
	this.hVal = int(math.Sqrt(float64((end.X-this.X)*(end.X-this.X)+(end.Y-this.Y)*(end.Y-this.Y))))
	return this.hVal
}
//对角算法
func (this *_AstarPoint)CalcuHvalD(end _Point)(h int){
	disdanceX := int(math.Abs(float64(this.X-end.X)))
	distanceY := int(math.Abs(float64(this.Y-end.Y)))
	if disdanceX > distanceY {
		this.hVal = distanceY * 14 + 10*(disdanceX-distanceY)
	}else {
		this.hVal = disdanceX *14 + 10*(distanceY - disdanceX)
	}
	return this.hVal
}
//计算F量
func (this *_AstarPoint)CalcuFval(end _Point)(f int){
	this.fVal = this.CalcuGVal()+this.CalcuHValM(end)
	return f
}
//寻路结构
type SearchRoad struct {
	Map *Mapconf //地图信息
	Openlist Openll //最小堆
	start _AstarPoint //起点
	end _AstarPoint //终点
	openxlsx map[string]*_AstarPoint //open表
	closexlsx map[string]*_AstarPoint //close表
	road []*_AstarPoint //路径
}

//初始化寻找路径
func InitSearch(start,end _Point,m *Mapconf)(S SearchRoad){
	S.Map = m
	S.start =  *NewAstarPoint(&_Point{start.X,start.Y,"S"},nil,nil)
	S.end = *NewAstarPoint(&_Point{end.X,end.Y,"E"},nil,nil)
	heap.Init(&S.Openlist)
	heap.Push(&S.Openlist,&S.start)
	S.road = make([]*_AstarPoint,0)
	S.openxlsx = make(map[string]*_AstarPoint,S.Map.maxX+S.Map.maxY)
	S.closexlsx = make(map[string]*_AstarPoint,S.Map.maxY+S.Map.maxX)

	//将起点放入open表
	S.openxlsx[point2string(start)] = &S.start
	//将障碍物放入close表
	for k,v :=range S.Map.blocks {
		S.closexlsx[k] =  NewAstarPoint(v,nil,nil)
	}

	return S
}
func (S *SearchRoad)PrintTheRoad()(){
	if len(S.road) == 0 {
		fmt.Printf("PrintTheRoad is nil ,error length")
		return
	}
	r := S.road[0]

	for r.father != nil{
		S.Map.points[r.X][r.Y].View = "*"
		r = r.father
	}
	S.Map.points[S.start.X][S.start.Y].View = "S"
	S.Map.points[S.end.X][S.end.Y].View = "E"
	fmt.Println("输出路径：")
	for x := 0;x<S.Map.maxX;x++{
		for y:=0;y<S.Map.maxY;y++{
			if y == S.Map.maxY-1{
				fmt.Print(S.Map.points[x][y].View+"\n")
			}else{
				fmt.Print(S.Map.points[x][y].View)
			}
		}
	}
}

func (S *SearchRoad)SearchOneRoad()(bool){
	num := 0
	for len(S.Openlist)>0 {
		num++
		min := heap.Pop(&S.Openlist)
		now_Apoint := min.(*_AstarPoint)
		delete(S.openxlsx, point2string(now_Apoint._Point))
		//fmt.Println("本次弹出的点：",now_Apoint)
		around_point := S.Map.GetOtherPoint(&now_Apoint._Point)
		S.closexlsx[point2string(now_Apoint._Point)]=now_Apoint

		for _,v := range around_point{
			considerPoint := NewAstarPoint(v,now_Apoint,&S.end)
			//fmt.Println("考虑的点：",considerPoint)
			if S.closexlsx[point2string(considerPoint._Point)] != nil {
				continue
			}
			if point2string(*v) == point2string(S.end._Point){
				fmt.Println("已经找到路径：")
				S.road = append(S.road,considerPoint)
				S.PrintTheRoad()
				fmt.Println("弹出次数：",num)
				return true
			}

			openpoint,ok := S.openxlsx[point2string(*v)]
			//fmt.Println("ok:",ok)
			if !ok {
				heap.Push(&S.Openlist,considerPoint)
				S.openxlsx[point2string(*v)] = considerPoint
			}else{
				oldF,oldFather := openpoint.fVal,openpoint.father
				openpoint.father = now_Apoint
				openpoint.CalcuFval(S.end._Point)
				//fmt.Printf("openpoint.F[%v] oldF[%v].",openpoint.fVal,oldF)
				if openpoint.fVal > oldF{
					openpoint.father = oldFather
					openpoint.CalcuFval(S.end._Point)
				}
			}
		}
	}
	return false
}

func main(){
	stratime := time.Now()
	presetMap := []string{
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X . X X X X X X X X X X X X X X X X X X X X X X X X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		"X X X X X X X X X X X X X X X X X X X X X X X X . X X",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
		". . . . . . . . . . . . . . . . . . . . . . . . . . .",
	}
	m := InitMap(presetMap)
	start_p := _Point{0,0,"."}
	end_p := _Point{18,10,"."}
	ss := InitSearch(start_p,end_p,&m)
	_ = ss.SearchOneRoad()

	fmt.Println("花费时间：",time.Since(stratime)/1000)
}

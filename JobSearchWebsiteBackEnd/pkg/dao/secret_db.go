package dao

import (
	"math/rand"
	"strconv"
	"time"
)

func Secret(resume string)(score int){
	rs,_:=strconv.ParseInt(resume,0,1000)
	rand.Seed(rs)
	rand.Seed(time.Now().UnixNano())
	score=60+rand.Intn(40)
	return score
}

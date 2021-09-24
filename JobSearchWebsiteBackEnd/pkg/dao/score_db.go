package dao

import (
	"os"
)

func Score(resume string)(score []byte) {

	f, _ := os.Create("Resume.txt")
	_, _ = f.Write([]byte(resume))
	score,_=os.ReadFile("score.txt")
	return
}


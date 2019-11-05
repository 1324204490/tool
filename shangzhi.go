package main

import (
	"fmt"
	"math"
)

type Code struct {
	ID      int
	Pcount  float64
	paid    float64
	NPcount float64
	before  float64
	score   float64
}

func main() {
	var codes []Code
	var code1 Code
	code1.ID = 54
	code1.paid = 3800
	code1.Pcount = 3
	code1.NPcount = 5
	code1.before = 233
	var code2 Code
	code2.ID = 65
	code2.paid = 6001
	code2.Pcount = 4
	code2.NPcount = 13
	code2.before = 91
	var code3 Code
	code3.ID = 67
	code3.paid = 2900
	code3.Pcount = 4
	code3.NPcount = 6
	code3.before = 75
	var code4 Code
	code4.ID = 72
	code4.paid = 3300
	code4.Pcount = 5
	code4.NPcount = 5
	code4.before = 45
	var code5 Code
	code5.ID = 78
	code5.paid = 3400
	code5.Pcount = 7
	code5.NPcount = 5
	code5.before = 4
	var code6 Code
	code6.ID = 79
	code6.paid = 500
	code6.Pcount = 1
	code6.NPcount = 1
	code6.before = 12
	var code7 Code
	code7.ID = 80
	code7.paid = 5300
	code7.Pcount = 2
	code7.NPcount = 7
	code7.before = 218
	var code8 Code
	code8.ID = 84
	code8.paid = 2400
	code8.Pcount = 4
	code8.NPcount = 10
	code8.before = 19
	var code9 Code
	code9.ID = 86
	code9.paid = 5800
	code9.Pcount = 4
	code9.NPcount = 9
	code9.before = 49
	var code10 Code
	code10.ID = 89
	code10.paid = 3800
	code10.Pcount = 3
	code10.NPcount = 4
	code10.before = 189
	codes = append(codes, code1, code2, code3, code4, code5, code6, code7, code8, code9, code10)
	ID := Score(codes)
	fmt.Println(ID)
}

//求score
func Score(codes []Code) int {
	if len(codes) == 0 {
		return 0
	}
	var paidSlice []float64
	var pcountSlice []float64
	var npcountSlice []float64
	var beforeSlice []float64
	for _, v := range codes {
		paidSlice = append(paidSlice, v.paid)
		pcountSlice = append(pcountSlice, v.Pcount)
		npcountSlice = append(npcountSlice, v.NPcount)
		beforeSlice = append(beforeSlice, v.before)
	}
	paidMax := max(paidSlice)
	paidMin := min(paidSlice)
	pcountMax := max(pcountSlice)
	pcountMin := min(pcountSlice)
	npcountMax := max(npcountSlice)
	npcountMin := min(npcountSlice)
	beforeMax := max(beforeSlice)
	beforeMin := min(beforeSlice)
	var temp Code
	temp.paid = paidMax + 1
	temp.Pcount = pcountMax + 1
	temp.NPcount = npcountMax + 1
	temp.before = beforeMin - 1
	codes = append(codes, temp)
	for _, v := range codes {
		paidSlice = append(paidSlice, v.paid)
		pcountSlice = append(pcountSlice, v.Pcount)
		npcountSlice = append(npcountSlice, v.NPcount)
		beforeSlice = append(beforeSlice, v.before)
	}
	paidMax = max(paidSlice)
	paidMin = min(paidSlice)
	pcountMax = max(pcountSlice)
	pcountMin = min(pcountSlice)
	npcountMax = max(npcountSlice)
	npcountMin = min(npcountSlice)
	beforeMax = max(beforeSlice)
	beforeMin = min(beforeSlice)
	k := 0 - 1/math.Log(float64(len(codes)))
	var NormalizationCodes []Code
	for _, v := range codes {
		var code Code
		code.ID = v.ID
		if paidMax-paidMin != 0 {
			code.paid = (paidMax - v.paid) / (paidMax - paidMin)
		}
		if pcountMax-pcountMin != 0 {
			code.Pcount = (pcountMax - v.Pcount) / (pcountMax - pcountMin)
		}
		if npcountMax-npcountMin != 0 {
			code.NPcount = (npcountMax - v.NPcount) / (npcountMax - npcountMin)
		}
		if beforeMax-beforeMin != 0 {
			code.before = (v.before - beforeMin) / (beforeMax - beforeMin)
		}
		NormalizationCodes = append(NormalizationCodes, code)
	}
	normalizationPaidSum := 0.0
	normalizationpcountSum := 0.0
	normalizationnpcountSum := 0.0
	normalizationbeforeSum := 0.0
	for _, v := range NormalizationCodes {
		normalizationPaidSum += v.paid
		normalizationpcountSum += v.Pcount
		normalizationnpcountSum += v.NPcount
		normalizationbeforeSum += v.before
	}
	var NormalizationCodes2 []Code
	for _, v := range NormalizationCodes {
		var code Code
		code.ID = v.ID
		if normalizationPaidSum != 0 {
			code.paid = v.paid / normalizationPaidSum
		}
		if normalizationpcountSum != 0 {
			code.Pcount = v.Pcount / normalizationpcountSum
		}
		if normalizationnpcountSum != 0 {
			code.NPcount = v.NPcount / normalizationnpcountSum
		}
		if normalizationbeforeSum != 0 {
			code.before = v.before / normalizationbeforeSum
		}
		NormalizationCodes2 = append(NormalizationCodes2, code)
	}
	var NormalizationCodes3 []Code //E  自然对数
	for _, v := range NormalizationCodes2 {
		var e Code
		e.ID = v.ID
		if v.paid > 0 {
			e.paid = v.paid * math.Log(v.paid)
		}
		if v.Pcount > 0 {
			e.Pcount = v.Pcount * math.Log(v.Pcount)
		}
		if v.NPcount > 0 {
			e.NPcount = v.NPcount * math.Log(v.NPcount)
		}
		if v.before > 0 {
			e.before = v.before * math.Log(v.before)
		}
		NormalizationCodes3 = append(NormalizationCodes3, e)
	}
	ejpaidSum := 0.0
	ejpcountSum := 0.0
	ejnpcountSum := 0.0
	ejbeforeSum := 0.0
	for _, v := range NormalizationCodes3 {
		ejpaidSum += v.paid
		ejpcountSum += v.Pcount
		ejnpcountSum += v.NPcount
		ejbeforeSum += v.before
	}
	var ej Code
	ej.paid = k * ejpaidSum
	ej.Pcount = k * ejpcountSum
	ej.NPcount = k * ejnpcountSum
	ej.before = k * ejbeforeSum
	var dj Code
	dj.paid = 1 - ej.paid
	dj.Pcount = 1 - ej.Pcount
	dj.NPcount = 1 - ej.NPcount
	dj.before = 1 - ej.before
	var wj Code
	sum := dj.paid + dj.before + dj.NPcount + dj.Pcount
	if sum != 0 {
		wj.paid = dj.paid / sum
		wj.Pcount = dj.Pcount / sum
		wj.NPcount = dj.NPcount / sum
		wj.before = dj.before / sum
	}
	maxVal := codes[0].score
	for i, _ := range codes {
		sum1 := NormalizationCodes2[i].paid * wj.paid
		sum2 := NormalizationCodes2[i].Pcount * wj.Pcount
		sum3 := NormalizationCodes2[i].NPcount * wj.NPcount
		sum4 := NormalizationCodes2[i].before * wj.before
		codes[i].score = (sum1 + sum2 + sum3 + sum4) * 100
		if maxVal < codes[i].score {
			maxVal = codes[i].score
		}
	}
	fmt.Println(codes)
	for _, v := range codes {
		if v.score == maxVal {
			return v.ID
		}
	}
	return 0
}

func max(slice []float64) float64 {
	maxVal := slice[0]
	for i := 1; i < len(slice); i++ {
		if maxVal < slice[i] {
			maxVal = slice[i]
		}
	}
	return maxVal
}
func min(slice []float64) float64 {
	minVal := slice[0]
	for i := 1; i < len(slice); i++ {
		if minVal > slice[i] {
			minVal = slice[i]
		}
	}
	return minVal
}

/*
@Time : 2021/6/23 19:09
@Author : lai
@Description :
@File : main
*/
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Prize struct {
	PlayerId int
	Weight   int
}

func main() {
	//设置奖项名称、权重等数组
	var prizes = make([]*Prize, 0)
	p1 := &Prize{
		PlayerId: 10001,
		Weight:   79,
	}
	prizes = append(prizes, p1)
	p2 := &Prize{
		PlayerId: 10002,
		Weight:   10,
	}
	prizes = append(prizes, p2)
	p3 := &Prize{
		PlayerId: 10003,
		Weight:   10,
	}
	prizes = append(prizes, p3)
	p4 := &Prize{
		PlayerId: 10004,
		Weight:   10000,
	}
	prizes = append(prizes, p4)
	//for _, v := range prizes {
	//	fmt.Println(v.PlayerId, v.Weight)
	//}
	RandomDraw(prizes)

	//fmt.Println(prizes[0])
}

// 权重随机抽奖
func RandomDraw(prizes []*Prize) int {
	//权重累加求和
	var weightSum int
	for _, v := range prizes {
		weightSum += v.Weight
	}

	//生成一个权重随机数，介于0-weightSum之间
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(weightSum)

	//权重数组重组并排序
	randomNumTmp := &Prize{PlayerId: 1000, Weight: randomNum}
	concatWeightArr := make([]*Prize, 0)
	aa, _ := json.Marshal(prizes)
	_ = json.Unmarshal(aa, &concatWeightArr)
	//fmt.Println("concatWeightArr：", concatWeightArr)
	concatWeightArr = append(concatWeightArr, randomNumTmp) //将随机数加入权重数组

	//将包含随机数的新权重数组按从小到大（升序）排序
	sort.Slice(concatWeightArr, func(i, j int) bool {
		return concatWeightArr[i].Weight < concatWeightArr[j].Weight
	})
	sort.Slice(prizes, func(i, j int) bool {
		return prizes[i].Weight < prizes[j].Weight
	})

	for _, v := range prizes {
		fmt.Println(*v)
	}

	//索引权重随机数的数组下标
	var randomNumIndex = -1 //索引随机数在新权重数组中的位置
	for p, v := range concatWeightArr {
		if v.Weight == randomNum {
			randomNumIndex = p
		}
	}
	randomNumIndexTmp := math.Min(float64(randomNumIndex), float64(len(prizes)-1)) //权重随机数的下标不得超过奖项数组的长度-1，重新计算随机数在奖项数组中的索引位置
	randomNumIndex = int(randomNumIndexTmp)

	//fmt.Println(randomNumIndex)
	//取出对应奖项
	res := prizes[randomNumIndex] //从奖项数组中取出本次抽奖结果
	fmt.Println("本次抽奖结果：", res.PlayerId)
	return res.PlayerId
}

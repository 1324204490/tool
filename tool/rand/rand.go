/**
*@Author: Zzy
*@Date: 2019/4/22 16:22
*@Description:
 */
package rand

import (
	"math/rand"
	"time"
)

/**
生成一组不重复的随机数
start 从0开始
end   到end结束  包含start  不包含end
count 需要生成的随机数数量
条件
1.count 不能大于end-start
2.end   不能小于或者等于start
*/
func GenerateRandomNumber(start, end int, count int) []int {
	if count > (end - start) {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

/*
生成一组随机的数据
kind = 0.纯数字 1.小写字母2.大写字母3. 数字、大小写字母
*/
func Random(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

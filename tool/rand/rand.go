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

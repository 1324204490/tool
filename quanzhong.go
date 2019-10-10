package main

import "fmt"

type DownstreamUpstream struct {
	Name   string
	Weight int //权重
}
type Weighted struct {
	Server        *DownstreamUpstream //权重表
	Weight        int                 //权重
	CurrentWeight int                 //当前权重
}
type Load struct {
	servers  []*DownstreamUpstream
	weighted []*Weighted
}

func NewLoad(servers []*DownstreamUpstream) *Load {
	new := &Load{}
	new.NewServers(servers)
	return new
}

//初始化权重信息
func (l *Load) NewServers(servers []*DownstreamUpstream) {
	weighted := make([]*Weighted, 0)
	for _, v := range servers {
		w := &Weighted{
			Server:        v,
			Weight:        v.Weight,
			CurrentWeight: v.Weight,
		}
		weighted = append(weighted, w)
	}
	l.weighted = weighted
	l.servers = servers
}

//选择权重
func (l *Load) nextWeighted(servers []*Weighted) (best *Weighted) {
	total := 0
	for i := 0; i < len(servers); i++ {
		w := servers[i]
		if w == nil {
			continue
		}
		w.CurrentWeight += w.Weight
		total += w.Weight
		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}
	if best == nil {
		return best
	}
	best.CurrentWeight -= total
	return best
}
func (l *Load) Select() *DownstreamUpstream {
	if len(l.weighted) == 0 {
		return nil
	}
	w := l.nextWeighted(l.weighted)
	if w == nil {
		return nil
	}
	return w.Server
}

/*
a b c d  //权重名称
1 2 3 4  //权重值总为10  循环十次做测试
原理
遍历所有的权重  让其当前权重值+=权重值
把总权重值计算出来 这里为10
每次轮询选择最大权重值的 减去总权重值
轮询结果dcbdcdabcd
第一轮
2 4 6 8
选择最大的 减去10 选择d
2 4 6 -2
第二轮
3 6 9 2
选择最大的 减去10 选择c
3 6 -1 2
第三轮
4 8 2 6
选择最大的 减去10 选择b
4 -2 2 6
第四轮
5 0 5 10
选择最大的 减去10 选择d
5 0 5 0
第五轮
6 2 8 4
选择最大的 减去10 选择c
6 2 -2 4
第六轮
7 4 1 8
选择最大的 减去10 选择d
7 4 1 -2
第七轮
8 6 4 2
选择最大的 减去10 选择a
-2 6 4 2
第八轮
-1 8 7 6
选择最大的 减去10 选择b
-1 -2 7 6
第九轮
0 0 10 10
选择最大的 减去10 c和d都可以 假如选择c 第十轮必定为d 加入选择d 第十轮必定为c  此次选c
0 0 0 10
第十轮
1 2 3 14
选择最大的 减去10  d
1 2 3 4  回归最开始的值
*/
func main() {
	servers := make([]*DownstreamUpstream, 0)
	servers = append(servers, &DownstreamUpstream{Name: "a", Weight: 1})
	servers = append(servers, &DownstreamUpstream{Name: "b", Weight: 2})
	servers = append(servers, &DownstreamUpstream{Name: "c", Weight: 3})
	servers = append(servers, &DownstreamUpstream{Name: "d", Weight: 4})
	lb := NewLoad(servers)
	for i := 0; i < 10; i++ {
		s := lb.Select()
		fmt.Print(s.Name)
	}
}

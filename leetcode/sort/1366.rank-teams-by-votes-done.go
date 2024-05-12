/**
 * @Time : 2020-06-04 10:27
 * @Author : yz
 */

package sort

import (
	"sort"
	"strings"
)

type Vote struct {
	// 队伍名称
	team string
	// 在对应名次获得的投票数
	rank []int
}

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}

	// 将每个队伍的投票结果写入该结构体
	teamRanks := []Vote{}

	// 每个队伍对应的投票总数
	voteMap := make(map[string][26]int)

	// 构造voteMap
	for i := 0; i < len(votes); i++ {
		// 第i名投票者的投票结果
		vote := votes[i]
		voteResult := strings.Split(vote, "")
		for j := 0; j < len(voteResult); j++ {
			if _, exists := voteMap[voteResult[j]]; exists {
				temp := voteMap[voteResult[j]]
				temp[j]++
				voteMap[voteResult[j]] = temp
			} else {
				temp := [26]int{}
				temp[j] = 1
				voteMap[voteResult[j]] = temp
			}
		}
	}

	// 给key排序，从小到大
	keys := []string{}
	for key := range voteMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 写入结构体
	for _, key := range keys {
		temp := voteMap[key]
		vote := Vote{key, temp[:]}
		teamRanks = append(teamRanks, vote)
	}

	var result []string

	// 不能转成string进行整体比较
	// 冒泡排序
	flag := true
	for i := 0; i < len(teamRanks) && flag; i++ {
		flag = false
		for j := len(teamRanks) - 2; j >= i; j-- {
			rank1 := teamRanks[j].rank
			rank2 := teamRanks[j+1].rank
			if !compare(rank1, rank2) {
				swap2(teamRanks, j, j+1)
				flag = true
			}
		}
	}

	// 从排序后的结果获取队伍名写入返回结果
	for i := 0; i < len(teamRanks); i++ {
		result = append(result, teamRanks[i].team)
	}

	return strings.Join(result, "")
}

// 交换
func swap2(vote []Vote, i int, j int) {
	temp := vote[i]
	vote[i] = vote[j]
	vote[j] = temp
}

// 如果a比b大返回true
func compare(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] == b[i] {
			continue
		} else {
			return false
		}
	}
	return true
}

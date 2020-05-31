/**
 * @Time : 2020-05-30 21:48
 * @Author : yz
 */

package main

import (
	"fmt"
	"github.com/yinzhenzhen/learning-leetcode/types"
)

type TaskQueue struct {
	// 所有队列
	queues []*types.Queue
	// 最长队列的长度
	maxLength int
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	taskQueue := TaskQueue{}

	// 没有加入任务队列的元素
	nums := 0

	i := 0
	for ; i < len(tasks); i++ {
		// 如果所有队列个数没有达到n+1个
		if len(taskQueue.queues) <= n  {
			flag := false
			j := 0
			for ; j < len(taskQueue.queues); j++ {
				str := taskQueue.queues[j].Tail().(byte)
				// 如果当前元素属于这个队列，加入队列
				if str == tasks[i]  {
					taskQueue.queues[j].EnQueue(tasks[i])
					if taskQueue.queues[j].Size() >= taskQueue.maxLength {
						taskQueue.maxLength = taskQueue.queues[j].Size()
					}
					flag = true
				}
			}
			// 当前元素在之前的队列都没有出现过，构造一个新的队列，放入队列任务
			if !flag {
				q := types.ConstructQueue()
				q.EnQueue(tasks[i])
				if taskQueue.maxLength == 0 {
					taskQueue.maxLength = 1
				}
				taskQueue.queues = append(taskQueue.queues, q)
			}
		} else {
			// 如果已经达到最大的队列个数，剩下的元素，如果已经存在于队列中，则继续添加到原来的队列，如果没有，则记下这样的元素的个数
			j := 0
			flag := false
			for ; j < len(taskQueue.queues); j++ {
				str := taskQueue.queues[j].Tail().(byte)
				// 如果当前元素属于这个队列，加入队列
				if str == tasks[i]   {
					taskQueue.queues[j].EnQueue(tasks[i])
					if taskQueue.queues[j].Size() >= taskQueue.maxLength {
						taskQueue.maxLength = taskQueue.queues[j].Size()
					}
					flag = true
				}
			}
			if !flag {
				nums++
			}
		}
	}

	all := 0

	maxQueueNum := 0

	// 获取所有队列的长度
	for j := 0; j < len(taskQueue.queues); j++ {
		taskQueue.queues[j].PrintQueue()
		all += taskQueue.queues[j].Length
		if taskQueue.queues[j].Length == taskQueue.maxLength {
			maxQueueNum++
		}
	}

	// 没有加入任务队列的元素，在多余的空格位置都能填空
	result1 := (n+1) * taskQueue.maxLength
	result2 := (n+1) * (taskQueue.maxLength-1)

	// 如果所有任务队列都没有达到, 且没有入队列的元素足以填补空缺
	if len(tasks) < result1 {
		result2 += maxQueueNum
		if result2-all < nums {
			result2 += nums
		}
		return result2
	} else if len(tasks) >= result1 {
		more := len(tasks) - result1
		if more > nums {
			return len(tasks)
		}
		return more + result1
	}

	return result1
}

func main()  {
	//fmt.Println(leastInterval([]byte{97,97,97,97,97,97,98,99,100,101,102,103}, 2))
	//fmt.Println(leastInterval([]byte{97,98,99,100,97,98,103}, 3))
	fmt.Println(leastInterval([]byte{97,97,97,97,98,98,98,98,99,99,99,99,100,100,100,100,101,102}, 4))
}
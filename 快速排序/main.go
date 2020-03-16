
package main

import "fmt"

/*
	快速排序
	快速排序 复杂度 N * (logN)
    原理为有两个指针:low, high 分别指向列表的第一个元素,和列表最后一个元素,和一个中间值 mid(就是取列表的第一个元素)
    一开始 low 默认情况下就指向了mid, high 指向列表最后一个元素
    [8, 3, 15, 7, 6, 2]
     |               |
    low             high
    如果high 指向的值比mid大,则high往前移动,如果high比mid 小,就把high 指向的值和low交换
    lst[low], lst[high] = lst[high], lst[low]
    如果low比mid小low往前移动一位, 如果low比mid大,则low指向的值和high 指向的值交换
    lst[high], lst[low] = lst[low], lst[high]
    当low 和 high 指针移动到相同的位置时,则mid就应该放在这个位置
 */

func main()  {
	list := []int{22, 56, 38, 101, 1, 18, 20, 30}
	fmt.Println("排序前", list)
	quick_sorting(list, 0, len(list) - 1)
	//quick(list, 0, len(list) - 1)
	fmt.Println("排序后", list)

}

// 快速排序实现，由于要二分递归，所以要传入每次排序的起始和结束的索引
func quick_sorting(data []int, start, end int) {
	if (start < end) {
		// 获取基准值
		base := data[start]
		// 临时变量，左右索引
		left := start
		right := end
		// 进入循环
		for left < right {
			// 由于左边的(第0个)被取走当成基准值，所以在左边就留下一个洞，用于存储比基准小的值
			// 所以先从右边找，找到第一个比基准值小的
			for left < right && data[right] >= base {
				right--
			}
			// 找到了比基准值小的，那就把这个值填入左边的洞，做索引要++
			if left < right {
				data[left] = data[right]
				left++
			}
			// 因为上面的操作让右边留了一个洞，开始从左边找比基准值大的
			for left < right && data[left] <= base {
				left++
			}
			// 找到比基准值大的再次把右边洞填上，并且在左边又留了一个洞
			if left < right {
				data[right] = data[left]
				right--
			}
		}

		// 把基准值填入到洞中，这就是本应该他所在的位置
		data[left] = base
		// 继续分两组递归排序，注意此时基准值已经不用参与排序了
		quick_sorting(data, start, left - 1)
		quick_sorting(data, left + 1, end)
	}
}

func quick(lst []int, start, end int)  {
	if start >= end {
		return
	}
	low := start
	high := end
	mid := lst[start]

	for low < high {
		for low < high && lst[high] >= mid {
			high -= 1
		}
		//出了循环说明 high 小于等于 mid 则需要交换位置
		lst[high], lst[low] = lst[low], lst[high]

		for low < high && lst[low] < mid{
			low += 1
		}
		// 出了循环说明 low 大于等于mid 则需要交换位置
		lst[low], lst[high] = lst[high], lst[low]
	}
	// 大的循环退出以后就找到了 mid的位置,此时mid 左边的都小于mid, 右边都大于mid,以 mid 为界限再排序左右两边
	lst[low] = mid
	quick(lst, start, low - 1) // 排序左边的
	quick(lst, low + 1, end ) // 排序右边的

}



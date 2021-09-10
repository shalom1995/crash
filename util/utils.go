package util

import (
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Bool2Int(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

/**
 * @parameter:[arr 要排序的数组][left 排序区间的左下标][right 排序区间的右下标]
 * @return:
 * @Description: 快速排序,从小到大
 * @author: shalom
 * @date: 2020/12/21 15:20
 */
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	tem := arr[i]
	for i < j {
		for i < j && arr[j] >= tem {
			j--
		}
		if i < j && arr[j] < tem {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] < tem {
			i++
		}
		if i < j && arr[i] > tem {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tem
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

/**
 * @parameter:[arr 要排序的数组][left 排序区间的左下标][right 排序区间的右下标]
 * @return:
 * @Description: 快速排序,从小到大
 * @author: shalom
 * @date: 2020/12/21 15:20
 */
func QuickSortBigInt(arr []*big.Int, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	tem := arr[i]
	for i < j {
		for i < j && comBig(arr[j], tem) {
			j--
		}
		if i < j && !comBig(arr[j], tem) {
			arr[i] = arr[j]
			i++
		}

		for i < j && !comBig(arr[i], tem) {
			i++
		}
		if i < j && comBig(arr[i], tem) {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = tem
	QuickSortBigInt(arr, left, i-1)
	QuickSortBigInt(arr, i+1, right)
}

//	true >= ; false <
func comBig(a, b *big.Int) bool {
	cmp := a.Cmp(b)
	return cmp != -1
}

/**
 * @parameter:[arr][aim]
 * @return: 返回数组下标，-1代表没有找到
 * @Description: 二分查找
 * @author: shalom
 * @date: 2021/1/19 16:01
 */
func BinarySearch(arr []*big.Int, aim *big.Int) int {
	left := 0
	right := len(arr) - 1

	// 这里一定要加上等号
	for left <= right {
		mid := (left + right) / 2

		if !comBig(aim, arr[mid]) {
			right = mid - 1
		} else if comBig(aim, arr[mid]) {
			left = mid + 1
		} else {
			return mid
		}
	}
	// 返回-1表示没有找到
	return -1
}

func StrToLow(str string) (strLow string) {
	strLow = strings.ToLower(str)
	return
}

/**
 * @parameter:
 * @return:
 * @Description: 返回两个块高中包含多少个 unitSection ，若最后不足一整个 unitSection，则返回值加1
 * @author: shalom
 * @date: 2020/12/25 14:08
 */
func DivideTime(fromBlock, toBlock, unitSection int64) int64 {
	var section = toBlock - fromBlock + 1
	var time = section / unitSection
	if section%unitSection > 0 {
		time++
	}
	return time
}

/**
 * @parameter:
 * @return:
 * @Description: 获得可执行文件被执行时的目录路径
 * @author: shalom
 * @date: 2020/12/9 14:03
 */
func GetAppPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	index := strings.LastIndex(p, string(os.PathSeparator))
	return p[:index], nil
}

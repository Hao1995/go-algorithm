package timebasedkeyvaluestore

type Item struct {
	Val       string
	Timestamp int
}
type TimeMap struct {
	// []Item sorted by timestamp.
	Data map[string][]Item
}

func Constructor() TimeMap {
	return TimeMap{
		Data: make(map[string][]Item),
	}
}

// O(1) + O(nlogn) (sorting)
func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Data[key] = append(this.Data[key], Item{
		Val:       value,
		Timestamp: timestamp,
	})
}

// O(1) + O(logn) (binary search)
func (this *TimeMap) Get(key string, timestamp int) string {
	data, ok := this.Data[key]
	if !ok {
		return ""
	}

	var ans string
	var left, right int = 0, len(data)
	for left < right {
		var midd int = (left + right) / 2
		item := data[midd]

		if item.Timestamp <= timestamp {
			ans = item.Val
			left = midd + 1
		} else {
			right = midd
		}
	}

	return ans
}

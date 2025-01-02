package heap

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	queue := make(Queue, 10)
	for i := 0; i < 10; i++ {
		item := &Item{
			data:  i + 1,
			ref:   i + 1,
			index: i,
		}
		queue[i] = item
	}
	heap.Init(&queue)
	item := Item{
		data: 8,
		ref:  1,
	}
	heap.Push(&queue, &item)
	heap.Fix(&queue, 2)
	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		t.Log("index", item.index, "ref", item.ref, "val", item.data)
	}
}

// TestQueue_Len
//
//	@Description: 追加测试用例，上面的用例和写的heap包不一致
//	@param t
func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		m    Queue
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Len from non-empty queue",
			m:    Queue{{ref: 1, index: 0}, {ref: 2, index: 1}, {ref: 3, index: 2}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    Queue
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Less from non-empty queue",
			m:    Queue{{ref: 1, index: 0}, {ref: 2, index: 1}, {ref: 3, index: 2}},
			args: args{i: 0, j: 1},
			want: false,
		},
		{
			name: "Less from non-empty queue",
			m:    Queue{{ref: 1, index: 0}, {ref: 2, index: 1}, {ref: 3, index: 2}},
			args: args{i: 2, j: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		m    Queue
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "Pop from non-empty queue",
			m:    Queue{{ref: 1, index: 0}, {ref: 2, index: 1}, {ref: 3, index: 2}},
			want: &Item{ref: 3, index: -1}, // 最后一个元素被移除，index 设置为 -1
		},
		{
			name: "Pop from single-element queue",
			m:    Queue{{ref: 1, index: 0}},
			want: &Item{ref: 1, index: -1}, // 唯一元素被移除，index 设置为 -1
		},
		{
			name: "Pop from empty queue",
			m:    Queue{},
			want: nil, // 空队列，返回 nil
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args struct {
		d interface{}
	}
	tests := []struct {
		name string
		m    Queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Push(tt.args.d)
		})
	}
}

func TestQueue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		m    Queue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Swap(tt.args.i, tt.args.j)
		})
	}
}

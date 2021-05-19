package tstorage

import "testing"

func BenchmarkStorage_InsertRows(b *testing.B) {
	storage, err := NewStorage()
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		storage.InsertRows([]Row{
			{Metric: "metric1", DataPoint: DataPoint{Timestamp: int64(i), Value: 0.1}},
		})
	}
}

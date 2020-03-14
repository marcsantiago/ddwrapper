package ddwrapper

import (
	"reflect"
	"testing"
)

func TestTags_Pair(t *testing.T) {
	tests := []struct {
		name string
		t    Tags
		want []string
	}{
		{
			name: "should pair stat without issue",
			t:    []string{"network", "foobar"},
			want: []string{"network:foobar"},
		},
		{
			name: "should pair stat without issue, many stats",
			t:    []string{"network", "foobar", "platform", "android", "auction_type", "static", "publisher", "testpub"},
			want: []string{"network:foobar", "platform:android", "auction_type:static", "publisher:testpub"},
		},
		{
			name: "should return broken tags instead of panicking",
			t:    []string{"network", "foobar", "platform", "android", "auction_type", "static", "publisher"},
			want: []string{"network", "foobar", "platform", "android", "auction_type", "static", "publisher"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Pair(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pair() =\n%v want\n %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTags_Pair2(b *testing.B) {
	tags := Tags([]string{"network", "foobar"})
	var n []string
	_ = n
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = tags.Pair()
	}
}

func BenchmarkTags_Pair4(b *testing.B) {
	tags := Tags([]string{"network", "foobar", "platform", "android"})
	var n []string
	_ = n
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = tags.Pair()
	}
}

func BenchmarkTags_Pair6(b *testing.B) {
	tags := Tags([]string{"network", "foobar", "platform", "android", "auction_type", "static"})
	var n []string
	_ = n
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = tags.Pair()
	}
}

func BenchmarkTags_Pair8(b *testing.B) {
	tags := Tags([]string{"network", "foobar", "platform", "android", "auction_type", "static", "publisher", "testpub"})
	var n []string
	_ = n
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n = tags.Pair()
	}
}

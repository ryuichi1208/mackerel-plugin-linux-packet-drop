package packet

import (
	"reflect"
	"testing"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func Test_newNetDevLine(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    []NetDevLine
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := newNetDevLine(tt.args.file)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. newNetDevLine() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. newNetDevLine() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		rawLine string
	}
	tests := []struct {
		name    string
		args    args
		want    NetDevLine
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := parse(tt.args.rawLine)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. parse() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. parse() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_graphGen(t *testing.T) {
	type args struct {
		labelPrefix string
		device      string
	}
	tests := []struct {
		name string
		args args
		want map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := graphGen(tt.args.labelPrefix, tt.args.device); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. graphGen() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_mapMerge(t *testing.T) {
	type args struct {
		m []map[string]mp.Graphs
	}
	tests := []struct {
		name string
		args args
		want map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mapMerge(tt.args.m...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. mapMerge() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNetDevLine_FetchMetrics(t *testing.T) {
	tests := []struct {
		name    string
		n       NetDevLine
		want    map[string]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.n.FetchMetrics()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. NetDevLine.FetchMetrics() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NetDevLine.FetchMetrics() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNetDevLine_MetricKeyPrefix(t *testing.T) {
	tests := []struct {
		name string
		n    NetDevLine
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.MetricKeyPrefix(); got != tt.want {
			t.Errorf("%q. NetDevLine.MetricKeyPrefix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNetDevLine_GraphDefinition(t *testing.T) {
	tests := []struct {
		name string
		n    NetDevLine
		want map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.GraphDefinition(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NetDevLine.GraphDefinition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		Do()
	}
}

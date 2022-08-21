package packet

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

type NetDevLine struct {
	Prefix       string
	Name         string
	RxBytes      float64
	RxPackets    float64
	RxErrors     float64
	RxDropped    float64
	RxFIFO       float64
	RxFrame      float64
	RxCompressed float64
	RxMulticast  float64
	TxBytes      float64
	TxPackets    float64
	TxErrors     float64
	TxDropped    float64
	TxFIFO       float64
	TxCollisions float64
	TxCarrier    float64
	TxCompressed float64
}

type NetDev map[string]NetDevLine

func newNetDevLine(file string) ([]NetDevLine, error) {
	f, err := os.Open(file)
	netDevs := []NetDevLine{}
	if err != nil {
		return netDevs, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for n := 0; s.Scan(); n++ {
		// Skip the 2 header lines.
		if n < 2 {
			continue
		}
		t, err := parse(s.Text())
		if err != nil {
			return netDevs, err
		}
		netDevs = append(netDevs, t)
	}

	return netDevs, nil
}

func parse(rawLine string) (NetDevLine, error) {
	m := NetDevLine{}
	var err error
	idx := strings.LastIndex(rawLine, ":")
	if idx == -1 {
		return m, errors.New("invalid net/dev line, missing colon")
	}
	fields := strings.Fields(strings.TrimSpace(rawLine[idx+1:]))

	m.Name = strings.TrimSpace(rawLine[:idx])
	if m.Name == "" {
		return m, errors.New("invalid net/dev line, empty interface name")
	}

	m.RxBytes, err = strconv.ParseFloat(fields[0], 10)
	if err != nil {
		return m, err
	}
	m.RxPackets, err = strconv.ParseFloat(fields[1], 10)
	if err != nil {
		return m, err
	}
	m.RxErrors, err = strconv.ParseFloat(fields[2], 10)
	if err != nil {
		return m, err
	}
	m.RxDropped, err = strconv.ParseFloat(fields[3], 10)
	if err != nil {
		return m, err
	}
	m.RxFIFO, err = strconv.ParseFloat(fields[4], 10)
	if err != nil {
		return m, err
	}
	m.RxFrame, err = strconv.ParseFloat(fields[5], 10)
	if err != nil {
		return m, err
	}
	m.RxCompressed, err = strconv.ParseFloat(fields[6], 10)
	if err != nil {
		return m, err
	}
	m.RxMulticast, err = strconv.ParseFloat(fields[7], 10)
	if err != nil {
		return m, err
	}
	m.TxBytes, err = strconv.ParseFloat(fields[8], 10)
	if err != nil {
		return m, err
	}
	m.TxPackets, err = strconv.ParseFloat(fields[9], 10)
	if err != nil {
		return m, err
	}
	m.TxErrors, err = strconv.ParseFloat(fields[10], 10)
	if err != nil {
		return m, err
	}
	m.TxDropped, err = strconv.ParseFloat(fields[11], 10)
	if err != nil {
		return m, err
	}
	m.TxFIFO, err = strconv.ParseFloat(fields[12], 10)
	if err != nil {
		return m, err
	}
	m.TxCollisions, err = strconv.ParseFloat(fields[13], 10)
	if err != nil {
		return m, err
	}
	m.TxCarrier, err = strconv.ParseFloat(fields[14], 10)
	if err != nil {
		return m, err
	}
	m.TxCompressed, err = strconv.ParseFloat(fields[15], 10)
	if err != nil {
		return m, err
	}

	return m, nil

}

func graphGen(labelPrefix, device string) map[string]mp.Graphs {
	graphs := map[string]mp.Graphs{
		device: {
			Label: labelPrefix,
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: fmt.Sprintf("%s.RxDropped", device), Label: "", Diff: true},
				{Name: fmt.Sprintf("%s.TxDropped", device), Label: "", Diff: true},
				{Name: fmt.Sprintf("%s.RxErrors", device), Label: "", Diff: true},
				{Name: fmt.Sprintf("%s.TxErrors", device), Label: "", Diff: true},
			},
		},
	}

	return graphs
}

func mapMerge(m ...map[string]mp.Graphs) map[string]mp.Graphs {
	ans := make(map[string]mp.Graphs, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func (n NetDevLine) FetchMetrics() (map[string]float64, error) {
	m := make(map[string]float64)
	nd, err := newNetDevLine("/proc/net/dev")
	if err != nil {
		return m, err
	}

	for _, v := range nd {
		m[fmt.Sprintf("%s.RxDropped", v.Name)] = v.RxDropped
		m[fmt.Sprintf("%s.TxDropped", v.Name)] = v.TxDropped
		m[fmt.Sprintf("%s.RxErrors", v.Name)] = v.RxErrors
		m[fmt.Sprintf("%s.TxErrors", v.Name)] = v.TxErrors
	}

	return m, nil
}

func (n NetDevLine) MetricKeyPrefix() string {
	if n.Prefix == "" {
		n.Prefix = "packetdrop"
	}
	return n.Prefix
}

func (n NetDevLine) GraphDefinition() map[string]mp.Graphs {
	m := make(map[string]mp.Graphs)
	nd, _ := newNetDevLine("/proc/net/dev")
	for _, val := range nd {
		m = mapMerge(m, graphGen("packet", val.Name))
	}

	return m
}

func Do() {
	optPrefix := flag.String("metric-key-prefix", "PacketDrop", "Metric key prefix")
	flag.Parse()

	v := NetDevLine{
		Prefix: *optPrefix,
	}
	plugin := mp.NewMackerelPlugin(v)
	plugin.Run()
}

package protobench

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/schattian/protobench/gogo"
	"github.com/schattian/protobench/official"
	officialjsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

// github.com/golang/protobuf (aka V1)
// google.golang.org/protobuf (aka V2)

func generateGoV2(n int) []official.Official {
	a := make([]official.Official, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, official.Official{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     official.TypeOfficial(rand.Intn(4)),
			Values:   &official.Official_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_GoV2_Proto_Marshal(b *testing.B) {
	data := generateGoV2(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV2_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV2(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &official.Official{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_GoV2_JSON_Marshal(b *testing.B) {
	data := generateGoV2(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := officialjsonpb.Marshal(data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GoV2_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGoV2(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = officialjsonpb.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	randomI := randomI(b.N)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &official.Official{}
		err := officialjsonpb.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

// github.com/gogo/protobuf/proto (aka gogo)

func generateGogoV1(n int) []*gogoproto.Gogo {
	a := make([]*gogoproto.Gogo, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &gogoproto.Gogo{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     TypeGoGoV1(rand.Intn(4)),
			Values:   &gogoproto.Gogo_ValueS{ValueS: randString(5)},
		})
	}
	return a
}

func Benchmark_GogoV1_Proto_Marshal(b *testing.B) {
	data := generateGogoV1(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := data[rand.Intn(len(data))].Marshal()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GogoV1_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoV1(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = d.Marshal()
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(ser[i])
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &gogoproto.Gogo{}
		err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_GogoV1_JSON_Marshal(b *testing.B) {
	data := generateGogoV1(b.N)
	marshaler := gogojsonpb.Marshaler{}
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		err := marshaler.Marshal(&buf, data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += buf.Len()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_GogoV1_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoV1(b.N)
	marshaler := gogojsonpb.Marshaler{}
	ser := make([]bytes.Buffer, len(data))
	var serialSize int
	for i, d := range data {
		err := marshaler.Marshal(&ser[i], d)
		if err != nil {
			b.Fatal(err)
		}
		serialSize += ser[i].Len()
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	unmarshaler := gogojsonpb.Unmarshaler{}
	randomI := randomI(b.N)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := randomI[i]
		o := &gogoproto.Gogo{}
		err := unmarshaler.Unmarshal(&ser[n], o)
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n].String())
		}
	}
}

func randomI(n int) []int {
	randomI := make([]int, n)
	for i := 0; i < len(randomI); i++ {
		randomI[i] = i
	}
	rand.Shuffle(len(randomI), func(i, j int) {
		randomI[i], randomI[j] = randomI[j], randomI[i]
	})
	return randomI
}

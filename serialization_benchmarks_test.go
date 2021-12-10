package protobench

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	gogo "github.com/schattian/protobench/gogo"
	official "github.com/schattian/protobench/official"
	officialjsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func generateOfficial(n int) []official.Message {
	a := make([]official.Message, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, official.Message{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     official.Type(rand.Intn(4)),
			Values:   &official.Message_ValueS{ValueS: randString(5)},
			Book: &official.Book{
				RandomReader: randString(n),
				Readers:      []string{randString(10), randString(20)},
				Author: &official.Author{
					Name:      randString(15),
					LastName:  randString(5),
					Age:       int32(n),
					RandomNum: math.MaxInt64 / 4,
				},
			},
			Ts: timestamppb.Now(),
		})
	}
	return a
}

// func Benchmark_Official_Proto_Marshal(b *testing.B) {
// 	data := generateOfficial(b.N)
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	var serialSize int
// 	for i := 0; i < b.N; i++ {
// 		bytes, err := proto.Marshal(&data[rand.Intn(len(data))])
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		serialSize += len(bytes)
// 	}
// 	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
// }

func Benchmark_Official_Proto_Unmarshal(b *testing.B) {
	fmt.Println("9start")
	b.StopTimer()
	data := generateOfficial(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = proto.Marshal(&d)
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
		o := &official.Message{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_Official_JSON_Marshal(b *testing.B) {
	data := generateOfficial(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := officialjsonpb.Marshal(&data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_Official_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateOfficial(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = officialjsonpb.Marshal(&d)
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
		o := &official.Message{}
		err := officialjsonpb.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

// github.com/gogo/protobuf/proto (aka gogo)

func generateGogo(n int) []*gogo.Message {
	a := make([]*gogo.Message, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, &gogo.Message{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     gogo.Type(rand.Intn(4)),
			Values:   &gogo.Message_ValueS{ValueS: randString(5)},
			Book: &gogo.Book{
				RandomReader: randString(3),
				Readers:      []string{randString(10), randString(20)},
				Author: &gogo.Author{
					Name:      randString(15),
					LastName:  randString(5),
					Age:       int32(n),
					RandomNum: math.MaxInt64,
				},
			},
			Ts: time.Now(),
		})
	}
	return a
}

func Benchmark_Gogo_Proto_Marshal(b *testing.B) {
	data := generateGogo(b.N)
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

func Benchmark_Gogo_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogo(b.N)
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
		o := &gogo.Message{}
		err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("gogoprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
	}
}

func Benchmark_Gogo_JSON_Marshal(b *testing.B) {
	data := generateGogo(b.N)
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

func Benchmark_Gogo_JSON_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogo(b.N)
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
		o := &gogo.Message{}
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

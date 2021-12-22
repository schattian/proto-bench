package protobench

import (
	"encoding/json"
	"math"
	"math/rand"
	"testing"
	"time"

	official "github.com/schattian/protobench/official"
	officialjsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func genOfficial(n int) []official.Message {
	a := make([]official.Message, 0, n)
	for i := 0; i < n; i++ {
		msg := official.Message{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     official.Type(rand.Intn(4)),
			Values:   &official.Message_ValueS{ValueS: randString(5)},
			Ts:       timestamppb.Now(),
			RandMap: map[string]uint32{
				"4358i45ja":          1,
				"342342b":            1342,
				"cfe3":               4351,
				"43242d":             3281,
				"dsadsae":            438291,
				"sadsada":            321,
				"321sadsada":         2141,
				"sa342dsada":         213,
				"sads4396itgada":     43951,
				"sadsad423uitrjhera": 9345941,
				"sfrejiufadsada":     134298,
			},
		}
		for i := 0; i < testSize; i++ {
			msg.Books = append(msg.Books, genOfficialBook())
		}
		a = append(a, msg)
	}
	return a
}

func genOfficialBook() *official.Book {
	return &official.Book{
		RandomReader: randString(3),
		Readers:      []string{randString(10), randString(20)},
		Author: &official.Author{
			Name:      randString(15),
			LastName:  randString(5),
			Age:       rand.Int31(),
			RandomNum: math.MaxInt64 / 2,
		},
		Letter: &official.Book_AQt{AQt: math.MaxInt64},
	}
}

func Benchmark_Official_Proto_Marshal(b *testing.B) {
	data := genOfficial(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := proto.Marshal(&data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func Benchmark_Official_Proto_Unmarshal(b *testing.B) {
	b.StopTimer()
	data := genOfficial(b.N)
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
	data := genOfficial(b.N)
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
	data := genOfficial(b.N)
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		var err error
		ser[i], err = officialjsonpb.Marshal(&d)
		// fmt.Println(string(ser[i]))
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

func Benchmark_Official_Default_JSON_Marshal(b *testing.B) {
	data := genOfficial(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		bytes, err := json.Marshal(&data[rand.Intn(len(data))])
		if err != nil {
			b.Fatal(err)
		}
		serialSize += len(bytes)
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

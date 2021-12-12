package protobench

import (
	"bytes"
	"math"
	"math/rand"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	gogo "github.com/schattian/protobench/gogo"
)

func genGogo(n int) []*gogo.Message {
	a := make([]*gogo.Message, 0, n)
	for i := 0; i < n; i++ {
		msg := &gogo.Message{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
			Type:     gogo.Type(rand.Intn(4)),
			Values:   &gogo.Message_ValueS{ValueS: randString(5)},
			Ts:       time.Now(),
		}
		for i := 0; i < testSize; i++ {
			msg.Books = append(msg.Books, genGogoBook())
		}
		a = append(a, msg)
	}
	return a
}

func genGogoBook() *gogo.Book {
	return &gogo.Book{
		RandomReader: randString(3),
		Readers:      []string{randString(10), randString(20)},
		Author: &gogo.Author{
			Name:      randString(15),
			LastName:  randString(5),
			Age:       rand.Int31(),
			RandomNum: math.MaxInt64 / 2,
		},
		Letter: &gogo.Book_AQt{AQt: math.MaxInt64},
	}
}

func Benchmark_Gogo_Proto_Marshal(b *testing.B) {
	data := genGogo(b.N)
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
	data := genGogo(b.N)
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
	data := genGogo(b.N)
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
	data := genGogo(b.N)
	marshaler := gogojsonpb.Marshaler{}
	ser := make([]bytes.Buffer, len(data))
	var serialSize int
	for i, d := range data {
		err := marshaler.Marshal(&ser[i], d)
		// fmt.Println(ser[i].String())
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

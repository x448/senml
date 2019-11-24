package senml_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/x448/senml"
)

func ExampleEncode1() {
	v := 23.1
	s := senml.SenML{
		Records: []senml.SenMLRecord{
			senml.SenMLRecord{Value: &v, Unit: "Cel", Name: "urn:dev:ow:10e2073a01080063"},
		},
	}

	dataOut, err := senml.Encode(s, senml.JSON, senml.OutputOptions{})
	if err != nil {
		fmt.Println("Encode of SenML failed")
	} else {
		fmt.Println(string(dataOut))
	}
	// Output: [{"n":"urn:dev:ow:10e2073a01080063","u":"Cel","v":23.1}]
}

func ExampleEncode2() {
	v1 := 23.5
	v2 := 23.6
	s := senml.SenML{
		Records: []senml.SenMLRecord{
			senml.SenMLRecord{Value: &v1, Unit: "Cel", BaseName: "urn:dev:ow:10e2073a01080063", Time: 1.276020076305e+09},
			senml.SenMLRecord{Value: &v2, Unit: "Cel", Time: 1.276020091305e+09},
		},
	}

	dataOut, err := senml.Encode(s, senml.JSON, senml.OutputOptions{})
	if err != nil {
		fmt.Println("Encode of SenML failed")
	} else {
		fmt.Println(string(dataOut))
	}
	// Output: [{"bn":"urn:dev:ow:10e2073a01080063","u":"Cel","t":1276020076.305,"v":23.5},{"u":"Cel","t":1276020091.305,"v":23.6}]
}

type TestVector struct {
	testDecode bool
	format     senml.Format
	binary     bool
	value      string
}

var testVectors = []TestVector{
	{true, senml.JSON, false, "W3siYm4iOiJkZXYxMjMiLCJidCI6LTQ1LjY3LCJidSI6ImRlZ0MiLCJidmVyIjo1LCJuIjoidGVtcCIsInUiOiJkZWdDIiwidCI6LTEsInV0IjoxMCwidiI6MjIuMSwicyI6MH0seyJuIjoicm9vbSIsInQiOi0xLCJ2cyI6ImtpdGNoZW4ifSx7Im4iOiJkYXRhIiwidmQiOiJhYmMifSx7Im4iOiJvayIsInZiIjp0cnVlfV0="},
	{true, senml.CBOR, true, "hKohZmRldjEyMyL7wEbVwo9cKPYjZGRlZ0MgBQBkdGVtcAFkZGVnQwb7v/AAAAAAAAAH+0AkAAAAAAAAAvtANhmZmZmZmgX7AAAAAAAAAACjAGRyb29tBvu/8AAAAAAAAANna2l0Y2hlbqIAZGRhdGEIY2FiY6IAYm9rBPU="},
	{true, senml.XML, false, "PHNlbnNtbCB4bWxucz0idXJuOmlldGY6cGFyYW1zOnhtbDpuczpzZW5tbCI+PHNlbm1sIGJuPSJkZXYxMjMiIGJ0PSItNDUuNjciIGJ1PSJkZWdDIiBidmVyPSI1IiBuPSJ0ZW1wIiB1PSJkZWdDIiB0PSItMSIgdXQ9IjEwIiB2PSIyMi4xIiBzPSIwIj48L3Nlbm1sPjxzZW5tbCBuPSJyb29tIiB0PSItMSIgdnM9ImtpdGNoZW4iPjwvc2VubWw+PHNlbm1sIG49ImRhdGEiIHZkPSJhYmMiPjwvc2VubWw+PHNlbm1sIG49Im9rIiB2Yj0idHJ1ZSI+PC9zZW5tbD48L3NlbnNtbD4="},
	{false, senml.CSV, false, "dGVtcCwyNTU2OC45OTk5ODgsMjIuMTAwMDAwLGRlZ0MNCg=="},
	//{true, senml.MPACK, true, "lIyhX8CiYm6mZGV2MTIzomJ0y8BG1cKPXCj2omJ1pGRlZ0OkYnZlcgWhbqR0ZW1woXPLAAAAAAAAAAChdMu/8AAAAAAAAKF1pGRlZ0OidXTLQCQAAAAAAAChdstANhmZmZmZmqJ2YsCHoV/AoW6kcm9vbaFzwKF0y7/wAAAAAAAAoXbAonZiwKJ2c6draXRjaGVuhqFfwKFupGRhdGGhc8ChdsCidmLAonZko2FiY4WhX8ChbqJva6FzwKF2wKJ2YsM="},
	{false, senml.LINEP, false, "Zmx1ZmZ5U2VubWwsbj10ZW1wLHU9ZGVnQyB2PTIyLjEscz0wIC0xMDAwMDAwMDAwCg=="},
}

func TestEncode(t *testing.T) {
	value := 22.1
	sum := 0.0
	vb := true
	s := senml.SenML{
		Records: []senml.SenMLRecord{
			senml.SenMLRecord{BaseName: "dev123",
				BaseTime:    -45.67,
				BaseUnit:    "degC",
				BaseVersion: 5,
				Value:       &value, Unit: "degC", Name: "temp", Time: -1.0, UpdateTime: 10.0, Sum: &sum},
			senml.SenMLRecord{StringValue: "kitchen", Name: "room", Time: -1.0},
			senml.SenMLRecord{DataValue: "abc", Name: "data"},
			senml.SenMLRecord{BoolValue: &vb, Name: "ok"},
		},
	}
	options := senml.OutputOptions{Topic: "fluffySenml", PrettyPrint: false}
	for i, vector := range testVectors {

		dataOut, err := senml.Encode(s, vector.format, options)
		if err != nil {
			t.Fail()
		}
		if vector.binary {
			fmt.Print("Test Encode " + strconv.Itoa(i) + " got: ")
			fmt.Println(dataOut)
		} else {
			fmt.Println("Test Encode " + strconv.Itoa(i) + " got: " + string(dataOut))
		}

		if base64.StdEncoding.EncodeToString(dataOut) != vector.value {
			t.Error("Failed Encode for format " + strconv.Itoa(i) + " got: " + base64.StdEncoding.EncodeToString(dataOut))
		}
	}

}

func TestDecode(t *testing.T) {
	for i, vector := range testVectors {
		fmt.Println("Doing TestDecode for vector", i)

		if vector.testDecode {
			data, err := base64.StdEncoding.DecodeString(vector.value)
			if err != nil {
				t.Fail()
			}

			s, err := senml.Decode(data, vector.format)
			if err != nil {
				t.Fail()
			}

			dataOut, err := senml.Encode(s, senml.JSON, senml.OutputOptions{PrettyPrint: true})
			if err != nil {
				t.Fail()
			}

			fmt.Println("Test Decode " + strconv.Itoa(i) + " got: " + string(dataOut))
		}
	}
}

func TestNormalize(t *testing.T) {
	value := 22.1
	sum := 0.0
	vb := true
	s := senml.SenML{
		Records: []senml.SenMLRecord{
			senml.SenMLRecord{BaseName: "dev123/",
				BaseTime:    897845.67,
				BaseUnit:    "degC",
				BaseVersion: 5,
				Value:       &value, Unit: "degC", Name: "temp", Time: -1.0, UpdateTime: 10.0, Sum: &sum},
			senml.SenMLRecord{StringValue: "kitchen", Name: "room", Time: -1.0},
			senml.SenMLRecord{DataValue: "abc", Name: "data"},
			senml.SenMLRecord{BoolValue: &vb, Name: "ok"},
		},
	}

	n := senml.Normalize(s)

	dataOut, err := senml.Encode(n, senml.JSON, senml.OutputOptions{PrettyPrint: true})
	if err != nil {
		t.Fail()
	}
	fmt.Println("Test Normalize got: " + string(dataOut))

	if base64.StdEncoding.EncodeToString(dataOut) != "WwogIHsiYnZlciI6NSwibiI6ImRldjEyMy90ZW1wIiwidSI6ImRlZ0MiLCJ0Ijo4OTc4NDQuNjcsInV0IjoxMCwidiI6MjIuMSwicyI6MH0sCiAgeyJidmVyIjo1LCJuIjoiZGV2MTIzL3Jvb20iLCJ1IjoiZGVnQyIsInQiOjg5Nzg0NC42NywidnMiOiJraXRjaGVuIn0sCiAgeyJidmVyIjo1LCJuIjoiZGV2MTIzL2RhdGEiLCJ1IjoiZGVnQyIsInQiOjg5Nzg0NS42NywidmQiOiJhYmMifSwKICB7ImJ2ZXIiOjUsIm4iOiJkZXYxMjMvb2siLCJ1IjoiZGVnQyIsInQiOjg5Nzg0NS42NywidmIiOnRydWV9Cl0K" {
		t.Error("Failed Normalize got: " + base64.StdEncoding.EncodeToString(dataOut))
	}
}

func TestBadInput1(t *testing.T) {
	data := []byte(" foo ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

func TestBadInput2(t *testing.T) {
	data := []byte(" { \"n\":\"hi\" } ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

func TestBadInputNoValue(t *testing.T) {
	data := []byte("  [ { \"n\":\"hi\" } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

func TestInputNumericName(t *testing.T) {
	data := []byte("  [ { \"n\":\"3a\", \"v\":1.0 } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestBadInputNumericName(t *testing.T) {
	data := []byte("  [ { \"n\":\"-3b\", \"v\":1.0 } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

func TestInputWeirdName(t *testing.T) {
	data := []byte("  [ { \"n\":\"Az3-:./_\", \"v\":1.0 } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestBadInputWeirdName(t *testing.T) {
	data := []byte("  [ { \"n\":\"A;b\", \"v\":1.0 } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

func TestInputWeirdBaseName(t *testing.T) {
	data := []byte("[ { \"bn\": \"a\" , \"n\":\"/b\" , \"v\":1.0} ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestBadInputNumericBaseName(t *testing.T) {
	data := []byte("[ { \"bn\": \"/3h\" , \"n\":\"i\" , \"v\":1.0} ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err == nil {
		t.Fail()
	}
}

// TODO add
//func TestBadInputUnknownMtuField(t *testing.T) {
//	data := []byte("[ { \"n\":\"hi\", \"v\":1.0, \"mtu_\":1.0  } ] ")
//	_ , err := senml.Decode(data, senml.JSON)
//	if err == nil {
//		t.Fail()
//	}
//}

func TestInputSumOnly(t *testing.T) {
	data := []byte("[ { \"n\":\"a\", \"s\":1.0 } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestInputBoolean(t *testing.T) {
	data := []byte("[ { \"n\":\"a\", \"vd\": \"aGkgCg\" } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestInputData(t *testing.T) {
	data := []byte("  [ { \"n\":\"a\", \"vb\": true } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

func TestInputString(t *testing.T) {
	data := []byte("  [ { \"n\":\"a\", \"vs\": \"Hi\" } ] ")
	_, err := senml.Decode(data, senml.JSON)
	if err != nil {
		t.Fail()
	}
}

// CBOR example from https://tools.ietf.org/html/rfc8428#section-6
func TestCborDecode(t *testing.T) {
	data, _ := hex.DecodeString("87a721781b75726e3a6465763a6f773a3130653230373361303130383030363a22fb41d303a15b00106223614120050067766f6c7461676501615602fb405e066666666666a3006763757272656e74062402fb3ff3333333333333a3006763757272656e74062302fb3ff4cccccccccccda3006763757272656e74062202fb3ff6666666666666a3006763757272656e74062102f93e00a3006763757272656e74062002fb3ff999999999999aa3006763757272656e74060002fb3ffb333333333333")
	value := 120.1
	v1 := 1.2
	v2 := 1.3
	v3 := 1.4
	v4 := 1.5
	v5 := 1.6
	v6 := 1.7
	wantRecords := []senml.SenMLRecord{
		senml.SenMLRecord{
			BaseName:    "urn:dev:ow:10e2073a0108006:",
			BaseTime:    1276020076.001,
			BaseUnit:    "A",
			BaseVersion: 5,
			Value:       &value,
			Unit:        "V",
			Name:        "voltage"},
		senml.SenMLRecord{Name: "current", Time: -5, Value: &v1},
		senml.SenMLRecord{Name: "current", Time: -4, Value: &v2},
		senml.SenMLRecord{Name: "current", Time: -3, Value: &v3},
		senml.SenMLRecord{Name: "current", Time: -2, Value: &v4},
		senml.SenMLRecord{Name: "current", Time: -1, Value: &v5},
		senml.SenMLRecord{Name: "current", Time: 0, Value: &v6},
	}

	s, err := senml.Decode(data, senml.CBOR)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.Records, wantRecords) {
		t.Errorf("got: %+v, want %+v", s, wantRecords)
	}
}

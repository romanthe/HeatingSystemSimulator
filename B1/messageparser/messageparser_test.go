package messageparser

import (
	"testing"
)

type MessageParseTest struct {
	in  string
	out Message
}

var messageparsetests = []MessageParseTest{
	// data log register
	{
		in: "action=register id=name port=1234",
		out: Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		},
	},
	// data log register response
	{
		in: "action=registerResponse id=name status=registered",
		out: Message{
			"action": []string{"registerResponse"},
			"id":     []string{"name"},
			"status": []string{"registered"},
		},
	},
	// data log requestAll
	{
		in: "id=name action=requestAll",
		out: Message{
			"id":     []string{"name"},
			"action": []string{"requestAll"},
		},
	},
	// data log requestAllResponse
	{
		in: "id=name action=requestAllResponse fcob=75.40 th=45.55 tr=17.03 tpco=30.15",
		out: Message{
			"id":     []string{"name"},
			"action": []string{"requestAllResponse"},
			"fcob":   []string{"75.40"},
			"th":     []string{"45.55"},
			"tr":     []string{"17.03"},
			"tpco":   []string{"30.15"},
		},
	},
}

func TestParseMessage(t *testing.T) {
	for i, test := range messageparsetests {
		x := ParseMessage(test.in)
		if len(x) != len(test.out) {
			t.Errorf("test %d: len(x) = %d, want %d", i, len(x), len(test.out))
		}
		for k, evs := range test.out {
			vs, ok := x[k]
			if !ok {
				t.Errorf("test %d: Missing key %q", i, k)
				continue
			}
			if len(vs) != len(evs) {
				t.Errorf("test %d: len(x[%q]) = %d, want %d", i, k, len(vs), len(evs))
				continue
			}
			for j, ev := range evs {
				if v := vs[j]; v != ev {
					t.Errorf("test %d: x[%q][%d] = %q, want %q", i, k, j, v, ev)
				}
			}
		}
	}
}

var messageparsebenchmark = "id=name action=requestAllResponse fcob=75.40 th=45.55 tr=17.03 tpco=30.15"

func BenchmarkParseMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(messageparsebenchmark)
	}
}

type MessageGetTest struct {
	in  Message
	out string
	key string
}

var messagegettests = []MessageGetTest{
	// not empty map
	{
		in: Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		},
		out: "1234",
		key: "port",
	},
	// single element map
	{
		in: Message{
			"id": []string{"name"},
		},
		out: "name",
		key: "id",
	},
	// empty map
	{
		in:  Message{},
		out: "",
		key: "id",
	},
	// not existing key
	{
		in: Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		},
		out: "",
		key: "status",
	},
}

func TestGet(t *testing.T) {
	for i, test := range messagegettests {
		x := test.in.Get(test.key)
		if x != test.out {
			t.Errorf("test %d: Get(\"%w\") = %w, want %w", i, test.key, x, test.out)
		}
	}
}

var messagegetbenchmark = Message{
	"action": []string{"register"},
	"id":     []string{"name"},
	"port":   []string{"1234"},
}

var messagegetbenchmarkkey = "id"

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		messagegetbenchmark.Get(messagegetbenchmarkkey)
	}
}

type MessageSetTest struct {
	in    Message
	key   string
	value string
}

var messagesettests = []MessageSetTest{
	// set port
	{
		in: Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		},
		key:   "port",
		value: "4321",
	},
}

func TestSet(t *testing.T) {
	for i, test := range messagesettests {
		test.in.Set(test.key, test.value)
		x := test.in[test.key]
		for j := range x{ 
		if x[j] != []string{test.value}[j] {
			t.Errorf("test %d: Set(\"%w\", \"%w\") = %w, want %w", i,
				test.key, test.value, x, test.value)
		}
		}
	}
}

var messagesetbenchmark = Message{
	"action": []string{"register"},
	"id":     []string{"name"},
	"port":   []string{"1234"},
}

var messagesetbenchmarkkey = "id"
var messagesetbenchmarkvalue = "id"

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		messagegetbenchmark.Set(messagesetbenchmarkkey,
			messagesetbenchmarkvalue)
	}
}

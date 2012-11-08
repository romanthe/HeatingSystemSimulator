package messageparser

import "testing"

func TestParseMessage(t *testing.T) {
	var testCases = []struct {
		in  string
		out Message
	}{
		// normal message
		{"action=register id=name port=1234", Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		}},
	}
	for i, test := range testCases {
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

func BenchmarkParseMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage("id=name action=requestAllResponse fcob=75.40 th=45.55 tr=17.03 tpco=30.15")
	}
}

func TestGet(t *testing.T) {
	var testCases = []struct {
		in  Message
		out string
		key string
	}{
		{
			// not empty map
			in: Message{
				"action": []string{"register"},
				"id":     []string{"name"},
				"port":   []string{"1234"},
			},
			out: "1234",
			key: "port",
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
	for i, test := range testCases {
		x := test.in.Get(test.key)
		if x != test.out {
			t.Errorf("test %d: Get(\"%w\") = %w, want %w", i, test.key, x, test.out)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	m := Message{
		"action": []string{"register"},
		"id":     []string{"name"},
		"port":   []string{"1234"},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Get("id")
	}
}

func TestSet(t *testing.T) {
	var testCases = []struct {
		in    Message
		key   string
		value string
	}{
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
	for i, test := range testCases {
		test.in.Set(test.key, test.value)
		x := test.in[test.key]
		for j := range x {
			if x[j] != []string{test.value}[j] {
				t.Errorf("test %d: Set(\"%w\", \"%w\") set %w, but want set %w", i,
					test.key, test.value, x, test.value)
			}
		}
	}
}

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	m := Message{
		"action": []string{"register"},
		"id":     []string{"name"},
		"port":   []string{"1234"},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Set("id", "4321")
	}
}

func TestDel(t *testing.T) {
	var testCases = []struct {
		in  Message
		key string
	}{
		// del port
		{
			in: Message{
				"action": []string{"register"},
				"id":     []string{"name"},
				"port":   []string{"1234"},
			},
			key: "port",
		},
	}
	for i, test := range testCases {
		test.in.Del(test.key)
		x, ok := test.in[test.key]
		if ok {
			t.Errorf("test %d: Del(\"%w\") after delete result %w, but want delete it",
				i, test.key, x)
		}
	}
}

func BenchmarkDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := Message{
			"id":     []string{"name"},
		}
		b.StartTimer()
		m.Del("id")
	}
}

func TestEncode(t *testing.T) {
	var testCases = []struct {
		in  Message
		out string
	}{
		// normal message
		{Message{
			"action": []string{"register"},
			"id":     []string{"name"},
			"port":   []string{"1234"},
		}, "action=register id=name port=1234", 
		},
	}
	for i, test := range testCases {
		x := test.in.Encode()
		if len(x) != len(test.out) {
			t.Errorf("test %d: len(x) = %d, want %d", i, len(x), len(test.out))
		}
		// test is incomplete, need to be fix later
	}
}

func BenchmarkEncode(b *testing.B) {
	b.StopTimer()
	m := Message{
		"action": []string{"register"},
		"id":     []string{"name"},
		"port":   []string{"1234"},
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Encode()
	}
}

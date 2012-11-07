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

var messageparsbencgmark = "id=name action=requestAllResponse fcob=75.40 th=45.55 tr=17.03 tpco=30.15"

func BenchmarkParseMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(messageparsbencgmark)
	}
}
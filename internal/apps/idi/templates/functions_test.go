package templates

import "testing"

func TestTrimS(t *testing.T) {
	t.Parallel()

	cases := map[string]string{"Todos": "Todo", "Users": "User"}

	for k, v := range cases {
		got := trimS(k)
		if got != v {
			t.Errorf("want %s got %s", v, got)
		}
	}
}

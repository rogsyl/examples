package examples

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestONE(t *testing.T) {
	data := []byte(`{"name": "Roger", "something": "one"}`)

	result := &MyData{"", UNKNOWN}
	if err := json.Unmarshal(data, result); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)
	if result.Name != "Roger" {
		t.Error("Name is " + result.Name + ", but expected \"Roger\"")
	}
	if result.Something != ONE {
		t.Error("Something is " + fmt.Sprintf("%d", result.Something) + ", but expected \"ONE\"")
	}
}

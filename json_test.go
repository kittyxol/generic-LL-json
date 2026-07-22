package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLLToJson(t *testing.T) {
	t.Parallel()
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		list1 := &LinkedList[string]{}
		list1.Append("a")
		list1.Append("b")
		list1.Append("c")
		data, err := json.Marshal(list1)

		var eList LinkedList[string]
		data0, err0 := json.Marshal(&eList)
		require.NoError(t, err0)
		require.JSONEq(t, string(data0), `{"length": 0}`)
		require.NoError(t, err)
		require.JSONEq(t, string(data),
			`{
		"head": {
  		"value": "a",
  		"child": {
		"value": "b",
   		"child": {
		"value": "c"
			}
		}
		},
		"length": 3
		}`)

		var nList LinkedList[int]
		nList.Append(11)
		nList.Append(22)
		dataN, errn := json.Marshal(&nList)
		require.NoError(t, errn)
		require.JSONEq(t, string(dataN),
			`{
		"head": { 
		"value": 11, 
		"child": { 
		"value": 22
			}
		},
		"length": 2}`)
	})
}
func TestJsonToLL(t *testing.T) {
	t.Parallel()
	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()
		str2 := `{"head": { "value": "a", "child": { "value": "b", "child": {  "value": "c"}}},"length": 3}`
		var list2 LinkedList[string]
		err2 := json.Unmarshal([]byte(str2), &list2)
		require.NoError(t, err2)
		require.Equal(t, 3, list2.Len())
		require.NotNil(t, &list2.head)
		require.Equal(t, "a", list2.head.val)
		require.Equal(t, "b", list2.head.next.val)
		require.Equal(t, "c", list2.head.next.next.val)
		require.Nil(t, list2.head.next.next.next)

	})
}

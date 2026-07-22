package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

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

		var uList LinkedList[User]
		uList.Append(User{Name: "Bob", Id: 2121})
		uList.Append(User{Name: "Lily", Id: 23672})
		dataU, errU := json.Marshal(&uList)
		require.NoError(t, errU)
		require.JSONEq(t, string(dataU),
			`{
		"head": {
        "value": {
        "name": "Bob",
        "id": 2121
		},
        "child": {
        "value": {
        "name": "Lily",
        "id": 23672
            }
        }
    	},
    	"length": 2
		}
		`)
	})
}
func TestJsonToLL(t *testing.T) {
	t.Parallel()
	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()
		str := `{
		"head": {
		"value": "a",
		"child": 
		{ "value": "b",
		"child": {
		"value": "c"
		}
		}
		},
		"length": 3
		}`
		var list LinkedList[string]
		err := json.Unmarshal([]byte(str), &list)
		require.NoError(t, err)
		require.Equal(t, 3, list.Len())
		require.NotNil(t, &list.head)
		require.Equal(t, "a", list.head.val)
		require.Equal(t, "b", list.head.next.val)
		require.Equal(t, "c", list.head.next.next.val)
		require.Nil(t, list.head.next.next.next)

		str1 := `{
		"head": {
        "value": {
        "name": "Bob",
        "id": 2121
		},
        "child": {
        "value": {
        "name": "Lily",
        "id": 23672
            }
        }
    	},
    	"length": 2
		}
		`
		var list1 LinkedList[User]
		err2 := json.Unmarshal([]byte(str1), &list1)
		require.NoError(t, err2)
		require.Equal(t, 2, list1.Len())
		require.NotNil(t, &list1.head)
		require.Equal(t, User{Name: "Bob", Id: 2121}, list1.head.val)
		require.Equal(t, User{Name: "Lily", Id: 23672}, list1.head.next.val)
		require.Nil(t, list1.head.next.next)

	})
}

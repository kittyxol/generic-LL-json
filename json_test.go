package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func TestLinkedList_string_MarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	list := &LinkedList[string]{}
	list.Append("a")
	list.Append("b")
	list.Append("c")
	data, err := json.Marshal(list)
	require.NoError(t, err)
	require.JSONEq(t, string(data), `
		{
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
		}
			`)
}

func TestLinkedList_string_MarshalJSON_Empty(t *testing.T) {
	t.Parallel()
	var List LinkedList[string]
	data, err := json.Marshal(&List)
	require.NoError(t, err)
	require.JSONEq(t, string(data), `{"length": 0}`)
}

func TestLinkedList_int_MarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	var List LinkedList[int]
	List.Append(11)
	List.Append(22)
	data, err := json.Marshal(&List)
	require.NoError(t, err)
	require.JSONEq(t, string(data), `
		{
			"head": { 
				"value": 11, 
				"child": { 
					"value": 22
				}
			},
			"length": 2
		}
			`)
}
func TestLinkedList_structUser_MarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	var List LinkedList[User]
	List.Append(User{Name: "Bob", Id: 2121})
	List.Append(User{Name: "Lily", Id: 23672})
	data, err := json.Marshal(&List)
	require.NoError(t, err)
	require.JSONEq(t, string(data), `
		{
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
}

func TestLinkedList_string_UnmarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	str := `{
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
}

func TestLinkedList_structUser_UnmarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	str := `{
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
	var list LinkedList[User]
	err := json.Unmarshal([]byte(str), &list)
	require.NoError(t, err)
	require.Equal(t, 2, list.Len())
	require.NotNil(t, &list.head)
	require.Equal(t, User{Name: "Bob", Id: 2121}, list.head.val)
	require.Equal(t, User{Name: "Lily", Id: 23672}, list.head.next.val)
	require.Nil(t, list.head.next.next)

}

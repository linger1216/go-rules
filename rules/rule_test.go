package rules

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	m := map[string]interface{}{
		"bool":   false,
		"number": 4,
		"list_number": []int64{
			1, 2, 3,
		},
		"string": "value",
		"list_string": []string{
			"aaa", "bbb", "ccc",
		},
		"float": 3.1415,
		"list_float": []float64{
			3.1, 3.2, 3.3,
		},

		"m1": map[string]interface{}{
			"bool":   false,
			"number": 4,
			"list_number": []int64{
				1, 2, 3,
			},
			"string": "value",
			"list_string": []string{
				"aaa", "bbb", "ccc",
			},
			"float": 3.1415,
			"list_float": []float64{
				3.1, 3.2, 3.3,
			},
		},
		"m2": map[string]interface{}{
			"arr": []map[string]interface{}{
				{
					"bool2":   false,
					"number2": 4,
					"list_number2": []int64{
						1, 2, 3,
					},
					"string2": "value",
					"list_string2": []string{
						"aaa", "bbb", "ccc",
					},
					"float2": 3.1415,
					"list_float2": []float64{
						3.1, 3.2, 3.3,
					},
				},
				{
					"bool3":   false,
					"number3": 4,
					"list_number3": []int64{
						1, 2, 3,
					},
					"string3": "value",
					"list_string3": []string{
						"aaa", "bbb", "ccc",
					},
					"float3": 3.1415,
					"list_float3": []float64{
						3.1, 3.2, 3.3,
					},
				},
			},
		},
	}

	rule := `m2.arr[0].list_number2[0] == 1 And number == 4`
	Evaluate(rule, m)
}

package php

import "testing"

func TestArrayChangeKeyCase(t *testing.T) {
	ma := ArrMap{
		"hello": 1111,
		"World": 2222,
	}
	t.Log(ArrayChangeKeyCase(ma, CaseUpper))
}

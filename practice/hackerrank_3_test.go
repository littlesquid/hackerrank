package practice

import "testing"

func TestCompareTriplets_draw(t *testing.T) {
	aliceRate := "1,2,3"
	bobRate := "3,2,1"

	expected := []int32{1, 1}

	result := CompareTriplets(aliceRate, bobRate)

	AssertEqual(expected[0], result[0], t)
	AssertEqual(expected[1], result[1], t)
}

func TestCompareTriplets_aliceWin(t *testing.T) {
	aliceRate := "17,28,30"
	bobRate := "99,16,8"

	expected := []int32{2, 1}

	result := CompareTriplets(aliceRate, bobRate)

	AssertEqual(expected[0], result[0], t)
	AssertEqual(expected[1], result[1], t)
}

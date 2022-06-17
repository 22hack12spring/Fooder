package service

import (
	"math"
	"math/rand"
)

func SimilarityVec3(vecA [3]float64, vecB [3]float64) float64 {
	// ユークリッド距離
	var sum float64
	for i := 0; i < 3; i++ {
		sum += math.Pow(float64(vecA[i]-vecB[i]), 2)
	}
	return math.Sqrt(sum)
}

// 直交基底のランダムな回転
// 雑実装なので、ちょっと後で考える
func RandomRotate() [3][3]float64 {
	theta := math.Pi * 2 * rand.Float64()
	theta2 := math.Pi * 2 * rand.Float64()
	res := [3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	for i := 0; i < 3; i++ {
		res[i] = RotateZ(res[i], theta)
	}
	for i := 0; i < 3; i++ {
		res[i] = RotateX(res[i], theta2)
	}
	return res
}

// Z軸行列回転
func RotateZ(vec [3]float64, theta float64) [3]float64 {
	return [3]float64{
		vec[0]*math.Cos(theta) - vec[1]*math.Sin(theta),
		vec[0]*math.Sin(theta) + vec[1]*math.Cos(theta),
		vec[2],
	}
}

// X軸行列回転
func RotateX(vec [3]float64, theta float64) [3]float64 {
	return [3]float64{
		vec[0],
		vec[1]*math.Cos(theta) - vec[2]*math.Sin(theta),
		vec[1]*math.Sin(theta) + vec[2]*math.Cos(theta),
	}
}

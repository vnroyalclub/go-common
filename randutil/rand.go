/*
* 常用的随机方法
 */

package randutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandFloat64() float64  {
   return rand.Float64()
}

func RandFloat32() float32  {
   return rand.Float32()
}

// 生成最大范围内随机数
func RandIntn(value int) int {
   if value == 0 {
      return 0
   }
   return rand.Intn(value)
}

//生成一个区间范围的随机数 [min,max)
func RandInts(min, max int) int {
   if min > max {
      max = min
   }
   if max-min == 0 {
      return max
   }
   randNum := rand.Intn(max - min)
   randNum = randNum + min
   return randNum
}

//生成一个区间范围的随机数 [min,max]
func RandIntsc(min, max int) int {
   if min > max {
      max = min
   }
   if max-min == 0 {
      return max
   }
   randNum := rand.Intn(max - min + 1)
   randNum = randNum + min

   return randNum
}

func RandInt64() int64 {
   return rand.Int63()
}

// Rand 生成最大范围内随机数
func RandInt64n(value int64) int64 {
   if value == 0 {
      return 0
   }
   randNum := rand.Int63n(value)
   return randNum
}

// 生成一个区间范围的随机数 [min,max)
func RandInt64s(min, max int64) int64 {
   if min > max {
      max = min
   }
   if max-min == 0 {
      return max
   }
   randNum := rand.Int63n(max - min)
   randNum = randNum + min
   return randNum
}

// 生成一个区间范围的随机数 [min,max]
func RandInt64sc(min, max int64) int64 {
   if min > max {
      max = min
   }
   if max-min == 0 {
      return max
   }
   randNum := rand.Int63n(max - min+1)
   randNum = randNum + min
   return randNum
}
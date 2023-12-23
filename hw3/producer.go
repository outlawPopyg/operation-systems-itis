package main

import (
 "fmt"
 "math/rand"
 "strconv"
 "time"
)

func main() {
 rand.Seed(time.Now().UnixNano())

 numExpressions := rand.Intn(61) + 120 // генерируем количество выражений
 for i := 0; i < numExpressions; i++ {
  expression := expression()
  fmt.Println(expression)
  time.Sleep(time.Second) // пауза в 1 секунду между выражениями
 }

 fmt.Println()


func expression() string {
     digit1 := digit()
     sign := sign()
     digit2 := digit()

     return fmt.Sprintf("%v %v %v", digit1, sign, digit2)
}

func digit() string {
 digit := rand.Intn(9) + 1 // генерируем число от 1 до 9
 return strconv.Itoa(digit)
}

func sign() string {
    signs := []string{"+", "-", "*", "/"}
    randomIndex := rand.Intn(len(signs)) // генерируем случайный индекс в диапазоне от 0 до 3
    return signs[randomIndex]
 }
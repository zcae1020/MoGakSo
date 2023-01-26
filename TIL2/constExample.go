//비트 플래그

package main

import "fmt"

type Room uint8

const (
	MasterRoom Room = 1 << iota // 1<<0
	BathRoom                    // 1<<1
	Kitchen                     // 1<<2
	LivingRoom                  // 1<<3
)

const (
	pig int = iota
	dog
	cat
	cow
)

const PI = 3.14
const FloatPI float64 = 3.14

func SetLight(rooms, room Room) Room {
	return rooms | room
}

func ResetLight(rooms, room Room) Room {
	return rooms &^ room
}

func IsTurnOn(rooms, room Room) bool {
	return rooms&room == room
}

func TurnLights(rooms Room) {
	if IsTurnOn(rooms, MasterRoom) {
		fmt.Println("masterRoom")
	}
	if IsTurnOn(rooms, BathRoom) {
		fmt.Println("bathRoom")
	}
	if IsTurnOn(rooms, Kitchen) {
		fmt.Println("kitchen")
	}
	if IsTurnOn(rooms, LivingRoom) {
		fmt.Println("livingRoom")
	}
}

func main() {
	var rooms Room = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)
	TurnLights(rooms)

	// var a int = PI * 100 		오류 발생 x
	// var b int = FloatPI * 100	오류 발생 o
}

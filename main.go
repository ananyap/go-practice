package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	john := NewMan()

	john.SellingFish(1000)

	cook := NewCook(john)
	fish := cook.myFisherMan.CatchingFish()

	log.Println(fish)

	for _, f := range cook.myFisherMan.CatchingFish() {
		log.Println(f)
	}

	johnMoney := cook.myFisherMan.GetMoney()

	log.Println(johnMoney)

}

type Cook struct {
	myFisherMan Fisherman
}

func NewCook(fisherMan Fisherman) Cook {
	return Cook{myFisherMan: fisherMan}
}

type Man struct {
	Balance int
}

func (m *Man) GetMoney() int {
	return m.Balance
}

func (m *Man) SellingFish(money int) {
	m.Balance = m.Balance + money
}

func (m *Man) CatchingFish() []string {
	var fish []string
	for i := 1; i <= 3; i++ {
		println(i)
		time.Sleep(1 * time.Second)

		fish = append(fish, strings.Join([]string{"ปลาตัวที่", strconv.Itoa(i)}, " : "))

	}
	return fish
}

type Fisherman interface {
	SellingFish(money int)
	CatchingFish() []string
	GetMoney() int
}

func NewMan() Fisherman {
	return &Man{}
}

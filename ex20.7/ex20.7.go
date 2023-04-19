package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	name string
}

func (m *Monster) Attack() {

}

func DoAttack(att Attacker) Attacker {
	if att != nil {
		att.Attack()

		var monster *Monster
		monster = att.(*Monster) // 다른 구현타입이 들어오면 Panic 에러 발생
		return monster
	}
	return nil
}

func main() {
	// var att Attacker // interface의 기본값은 nil
	// att.Attack()     // Runtime Exception.
	// DoAttack(att)
	var att1 Attacker
	att1 = &Monster{"mons"}
	returnVal := DoAttack(att1)
	fmt.Println(returnVal)

}
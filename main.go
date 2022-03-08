package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Input struct {
	Ts         int64       `json:"ts"`
	Properties *Properties `json:"properties"`
}

type Properties struct {
	MegnetStatus bool `json:"megnet_status"`
	Ia           int  `json:"ia"`
	Ep           int  `json:"ep"`
	Num          int  `json:"num"`
	Status       int  `json:"status"`
	InstantEp    int  `json:"instant_ep"`
	Threadhold   int  `json:"threadhold"`
}

func main() {
	/*aaa := make([]*ruleData, 0)
		a := &ruleData{
			Name: "megnet_status",
			Type: "bool",
			Rule: "",
		}
		b := &ruleData{
			Name: "num",
			Type: "int",
			Rule: `if megnet_status_last == false && megnet_status == true {
		return num_last + 1 }
	else {
		return num_last}
	`,
		}
		c := &ruleData{
			Name: "status",
			Type: "int",
			Rule: `if ia == 0 {
		return 0
	} else if ia > 40 {
		return 2
	} else {
		return 1
	}`,
		}
		d := &ruleData{
			Name: "ia",
			Type: "int",
			Rule: "",
		}
		e := &ruleData{
			Name: "ep",
			Type: "int",
			Rule: "",
		}
		f := &ruleData{
			Name: "instant_ep",
			Type: "int",
			Rule: ` return ep - ep_last`,
		}
		g := &ruleData{
			Name: "threadhold",
			Type: "int",
			Rule: "40",
		}
		aaa = append(aaa, a, g, b, c, d, e, f)
		aaaByte , _ := json.Marshal(aaa)
		fmt.Println(string(aaaByte))*/

	InitMiddle("./rule.txt")
	for i := 0; i < 10000; i++ {
		tmp := false
		if i%2 == 0 {
			tmp = false
		} else {
			tmp = true
		}

		input := &Input{
			Ts: 123456,
			Properties: &Properties{
				MegnetStatus: tmp,
				Ia:           rand.Intn(60),
				Ep:           i,
				Threadhold:   40,
			},
		}
		buf, _ := json.Marshal(input)
		data, _ := Computing(buf, "./rule.txt")
		fmt.Println(string(data))

		time.Sleep(time.Second)
	}
}

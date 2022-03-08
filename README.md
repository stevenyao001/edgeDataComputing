# edgeDataComputing

/*
调用Computing API，
第一个参数为生成的input json串，
第二个参数为规则文件所放的位置和名称
*/
func main() {
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

/*
规则文件格式：
name:节点名称（key）
rule:节点的规则
*/
[{"name":"megnet_status","type":"bool","rule":""},{"name":"num","type":"int","rule":"if megnet_status_last == false \u0026\u0026 megnet_status == true {\n\treturn num_last + 1 }\nelse {\n\treturn num_last}\n"},{"name":"status","type":"int","rule":"if ia == 0 {\n\treturn 0\n} else if ia \u003e 40 {\n\treturn 2\n} else {\n\treturn 1\n}"},{"name":"ia","type":"int","rule":""},{"name":"ep","type":"int","rule":""},{"name":"instant_ep","type":"int","rule":" return ep - ep_last"},{"name":"threadhold","type":"int","rule":""}]
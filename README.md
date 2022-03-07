# edgeDataComputing

/*
调用Computing API，
第一个参数为生成的input json串，
第二个参数为规则文件所放的位置和名称
*/
func main() {
    edgeDataComputing.InitMiddle()
    input := &edgeDataComputing.Input{
        Ts:         123456,
        Properties: &edgeDataComputing.Properties{
            AddressNames1: 26.1,
            AddressNames2: 4,
        },
    }
    buf, _ := json.Marshal(input)
    data, _ := edgeDataComputing.Computing(buf, "./rule.txt")
    fmt.Println(string(data))
}

/*
规则文件格式：
name:节点名称（key）
rule:节点的规则
*/
[{"name":"MegnetStatus","rule":""},{"name":"Num","rule":"if MegnetStatus_last == false \u0026\u0026 MegnetStatus == true {\n\t Num = Num_last + 1\n} else {\n\t Num = Num_last\n}"},{"name":"Status","rule":"if Ia == 0 {\n\t Status = 0\n} else if Ia \u003e 40 {\n\t Status = 2\n} else {\n\t Status = 1\n}\n"},{"name":"Ia","rule":""},{"name":"Ep","rule":""},{"name":"InstantEp","rule":" InstantEp = Ep - Ep_last"}]
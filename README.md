# edgeDataComputing

/*
调用Computing API，
第一个参数为生成的input json串，
第二个参数为规则文件所放的位置和名称
*/
func main() {
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
[{"name":"AddressNames1","rule":"AddressNames1 = AddressNames1 / 10"},{"name":"AddressNames2","rule":""},{"name":"AddressNames3","rule":"AddressNames3 = AddressNames1 / AddressNames2"}]
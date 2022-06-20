package rpctester

// InitialBalance is used to get initialbalance based on network
func InitialBalance(url string, version string) string {
	switch url {
	case "http://localhost:9500":
		if version == "V2" {
			return "10000000000000000000000000000"
		} else if version == "V1" {
			return "0x204fce5e3e25026110000000"
		}
	case "https://api.s1.t.hmny.io":
		if version == "V2" {
			return "0"
		} else if version == "V1" {
			return "0x0"
		}
		return "0"
	}
	return "0"
}

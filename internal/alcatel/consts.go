package alcatel

const BaseURL = "http://192.168.1.1/jrd/webapi"

type NetworkType int

const (
	NoService NetworkType = 0
	Net2G NetworkType = 1
	Net2G2 NetworkType = 2
	Net3G NetworkType = 3
	Net3G2 NetworkType = 4
	Net3GPlus NetworkType = 5
	Net3GPlus2 NetworkType = 6
	Net4G NetworkType = 7
	Net4GPlus NetworkType = 8
)

func (n NetworkType) Value() int {
	return int(n)
}

var NetWorkLiteral = []string{"NO_SERVICE", "2G", "2G", "3G", "3G", "3G", "3G+", "3G+", "4G", "4G+" }

func GetNetworkStringLiteral( networkType NetworkType ) string {
	if networkType.Value() >= 0 && networkType.Value() < len(NetWorkLiteral) {
		return NetWorkLiteral[ networkType ]
	}
	return NetWorkLiteral[0]
}
package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"encoding/base64"
	bip21 "github.com/yassun/go-bip21"
	"fmt"
	"os"
	"strings"
)

var png []byte

func main() {
	bolt11 := "lnbc10n1pjpdymlsp5pxusduthtnx4w4yn5rvefvpyrl9fd2ptunnrqcm6l9dwt9x74g7qpp5t33002kvlpd78et7kpjye74k25tlqvtt4ark8dwdxvlzyhjl0n9qdquvf5hqgpjxys8getnwssxzempd9hqxqyjw5qcqpjrzjqwgvcyhhsgjegfaac3dzsdc97npqej7l7pjw6slkka87ljh7n6clwze27cqqd7qqqyqqqqqqqqqqqeqqfv9qx3qysgq30tyyxufvf3msac0sn590mr03vkzgljd4dppazj9sst277ncccnqfwv0ja5h5nz75ke7pzrth4s7pdl9cvkeddh3upx07nx4f2ugykspp8x92v"
	onchainAddr := "BC1QYLH3U67J673H6Y6ALV70M0PL2YZ53TZHVXGG7U"
	label := "for lunch"
	message := "for lunch"

	u := &bip21.URIResources {
		UrnScheme: "bitcoin",
		Address: strings.ToUpper(onchainAddr),
		Amount: 0.001,
		Label: label,
		Message: message,
		Params: make(map[string]string),
	}

	u.Params["lightning"] = strings.ToUpper(bolt11)

	bip21uri, err := u.BuildURI()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bolt11uri := fmt.Sprintf("lightning:%s", strings.ToUpper(bolt11))
	bolt11png, err := qrcode.Encode(bolt11uri, qrcode.Medium, 256)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bip21png, err := qrcode.Encode(bip21uri, qrcode.Medium, 256)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bolt11out := base64.StdEncoding.EncodeToString(bolt11png)
	bip21out := base64.StdEncoding.EncodeToString(bip21png)
	output := fmt.Sprintf("<html><body><image src=\"data:image/png;base64,%s\" /><image src=\"data:image/png;base64,%s\" /></body></html>\n", bip21out, bolt11out)

	err = os.WriteFile("index.html", []byte(output), 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

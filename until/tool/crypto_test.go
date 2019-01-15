package tool

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fixKey := "8aL3gmNw9bd77hRRc7sRgWSsPccxQGecybgyHFt7yfOj8LcEVcar4u2M75BebWpb"

	i1 := 1
	i2 := 2

	fmt.Printf("md5 from MD5Sum=[%x]\n",
		MD5Sum([]byte(fmt.Sprintf("%d%d%s", i1, i2, fixKey))))

	fmt.Printf("md5 from MD5Sumf=[%x]\n",
		MD5Sumf("%d%d%s", i1, i2, fixKey))
}

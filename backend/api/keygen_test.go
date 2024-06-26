package api

import (
	"fmt"
	"testing"
)

func TestKeyGen(t *testing.T) {
	// var license entity.License
	lispart := `{"licenseId":"7GC5726T07","licenseeName":"胡爸爸","assigneeName":"","assigneeEmail":"","licenseRestriction":"","checkConcurrentUse":false,"products":[{"code":"PCWMP","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true},{"code":"PMYBATISHELPER","fallbackDate":"2029-09-14","paidUpTo":"2029-09-14","extended":true},{"code":"PSI","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true}],"metadata":"0120230914PSAX000005","hash":"TRIAL:1805249793","gracePeriodDays":7,"autoProlongated":false,"isAutoProlongated":false, "aiAllowed":true}`
	res := KeyGenRequest(lispart)
	fmt.Println(res)
	// json,_ := json.Marshal(license)
	// fmt.Printf("%s",json)

}

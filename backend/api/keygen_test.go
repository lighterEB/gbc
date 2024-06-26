package api

import (
	"encoding/json"
	"gbc/backend/entity"
	"testing"
)

func TestKeyGen(t *testing.T){
	var license entity.License
	lispart := `{"licenseId":"7GC5726T07","licenseeName":"gurgles tumbles","assigneeName":"","assigneeEmail":"","licenseRestriction":"","checkConcurrentUse":false,"products":[{"code":"PCWMP","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true},{"code":"GO","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":false},{"code":"PSI","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true}],"metadata":"0120230914PSAX000005","hash":"TRIAL:1805249793","gracePeriodDays":7,"autoProlongated":false,"isAutoProlongated":false}`
	json.Unmarshal([]byte(lispart), &license)
	KeyGen(&license)
	// json,_ := json.Marshal(license)
	// fmt.Printf("%s",json)
	
}
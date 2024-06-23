package entity

type License struct {
	LicenseId          string         `json:"licenseId"`
	LicenseName        string         `json:"licenseName"`
	AssigneeName       string         `json:"assigneeName"`
	AssigneeEmail      string         `json:"assigneeEmail"`
	LicenseRestriction string         `json:"licenseRestriction"`
	CheckConcurrentUse bool           `json:"checkConcurrentUse"`
	Products           []*ProductList `json:"products"`
	Metadata           string         `json:"metadata"`
	Hash               string         `json:"hash"`
	GracePeriodDays    int            `json:"gracePeriodDays"`
	AutoProlongated    bool           `json:"autoProlongated"`
	IsAutoProlongated  bool           `json:"isAutoProlongated"`
}

type ProductList struct {
	Code         string `json:"code"`
	FallbackDate string `json:"fallbackDate"`
	PaidUpTo     string `json:"paidUpTo"`
	Extended     bool   `json:"extended"`
}

func (license *License) CreateLis() {
	license.Metadata = "0120230914PSAX000005"
	license.Hash = "TRIAL:1649058719"
	license.GracePeriodDays = 7
	license.AutoProlongated = false
	license.IsAutoProlongated = false
}

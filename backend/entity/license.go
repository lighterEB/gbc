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
	Trial              bool           `json:"trial"`
	AiAllowed          bool           `json:"aiAllowed"`
}

type ProductList struct {
	Code         string `json:"code"`
	FallbackDate string `json:"fallbackDate"`
	PaidUpTo     string `json:"paidUpTo"`
	Extended     bool   `json:"extended"`
}

func (license *License) CreateLis() {
	var productList = ProductList{
		FallbackDate: "2099-12-31",
		PaidUpTo:     "2099-12-31",
	}
	license.Products = append(license.Products, &productList)
	license.Metadata = "0120230914PSAX000005"
	license.Hash = ""
	license.GracePeriodDays = 7
	license.AutoProlongated = false
	license.IsAutoProlongated = false
	license.Trial = false
	license.AiAllowed = true
}

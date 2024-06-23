package entity

type Request struct {
	LicenseId     string `json:"licenseId"`
	LicenseName   string `json:"licenseeName"`
	AssigneeName  string `json:"assigneeName"`
	AssigneeEmail string `json:"assigneeEmail"`
	UpToPaid      string `json:"upToPaid"`
}

//func (r *Request) getRequest() string {}

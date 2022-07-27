package entity

type Rows struct {
	Userid string `json:"userid"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type ApplyForLeave struct {
	Email     string `json:"email"`
	Userid    string `json:"userid"`
	LeaveType string `json:"leavetype"`
	NoOfDays  string `json:"noofdays"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
	Status    string `json:"Status"`
	Reason    string `json:"Reason"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

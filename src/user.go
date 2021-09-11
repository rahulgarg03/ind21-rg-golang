package src

type User struct {
	ID                              uint   `json:"id" gorm:"primary_key"`
	Email                           string `json:"email"`
	Password                        string `json:"password"`
	Active                          bool   `json:"active"`
	Friends                         []User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
	FirstName                       string `json:"firstname"`
	LastName                        string `json:"lastname"`
	Age                             int    `json:"age"`
	Gender                          string `json:"gender"`
	MartialStatus                   string `json:"martialstatus"`
	ResidentialAddress              string `json:"residential_address"`
	ResidentialCity                 string `json:"residential_city"`
	ResidentialState                string `json:"residential_state"`
	ResidentialCountry              string `json:"residential_country"`
	ResidentialContactNo1           string `json:"residential_contact_no1"`
	ResidentialContactNo2           string `json:"residential_contact_no2"`
	OfficialDetailsAddress          string `json:"officialdetails_employeecode"`
	OfficialDetailsState            string `json:"officialdetails_address"`
	OfficialDetailsCity             string `json:"officialdetails_city"`
	OfficialDetailsCountry          string `json:"officialdetails_country"`
	OfficialDetailsCompanyContactNo string `json:"officialdetails_company_contact_no"`
	OfficialDetailsCompanyEmail     string `json:"officialdetails_company_email"`
	OfficialDetailsCompanyName      string `json:"officialdetails_company_name"`
}

// func ConnectDataBase() *gorm.DB {
// 	dsn := "host=localhost user=postgres password=postgres dbname=myDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("Successfully Connected!!")
// 	}
// 	db.AutoMigrate(&User{})
// 	return db
// }

package models

type User struct {
	Id 		      			string      `json:"id"`
	FirstName     			string      `json:"first_name"`
	SecondName    			string      `json:"second_name"`
	PassportSeria 			string		`json:"passport_seria"`			 
   	PassportNumber			int			`json:"passport_number"`
	MiddleName				string      `json:"middle_name"`
	Addres					string		`json:"addres"`		     		
}

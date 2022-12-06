package domain

import "github.com/Jay-Shah10/go-lang-microservices/banking/err"

type CustomerRepositoryDb struct {
	client *sql.db
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *err.AppErrors) {

	var row *sql.Rows
	var err error

	if status == "" {
		findAllSql := "query"
		rows, err := d.client.Query(findAllSql)
	}else {
	findAllSql := "query"
		rows, err := d.client.Query(findAllSql, status)
	}
		
	if err != nil {
		logger.Error("Error running find all query " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	customers := make([]Customers, 0)
	for rows.Next(){
		var c Customer
		err := rows.Scan(all &'c.attributes')
		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}
		customers = append(customers, c)
	}
	return customers


}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *err.AppErrors) {
	
		customersql := '<same as the one in findall, but it will have a where clause.>'

		row := d.client.QueryRow(customersql, id)
		var c Customer
		err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("Customer not found")
			}else {
				fmt.Println("Error while scanning customer... " + err.Error)
				return nil, errs.NewUnexpectedError("Unexpected Database Error")
			}
		}
		return &c, nil
	

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	/*
		client, err := sql.Open("mysql", "user:password@/dbname")
			if err != nil{
				panic(err)
			}

		// see "Important Settings" section.
		client.SetMaxLifttime(time.Minute * 3)
		client.SetMaxConns(10)
		client.SetMaxIdleConns(10)
		return CustomerRepositoryDb{client}
	*/

}

//helper function

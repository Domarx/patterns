package main

import "github.com/domarx/patterns/structural/composite/company"

func main() {
	c := company.NewCompany()

	ceo := company.NewCEO("alpha")
	c.Add(ceo)

	devs := company.NewDepartment("developers")
	devs.Add(company.NewDeveloper("bob"))
	devs.Add(company.NewDeveloper("alice"))
	c.Add(devs)

	accountants := company.NewDepartment("accountants")
	accountants.Add(company.NewAccountant("john rambo"))
	accountants.Add(company.NewAccountant("john matrix"))
	c.Add(accountants)

	c.Execute()
}

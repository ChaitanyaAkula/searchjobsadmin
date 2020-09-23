package searchjobsadmin

import (
	
	"log"
    "fmt"
	"github.com/ChaitanyaAkula/gittyjobsdb"
	"github.com/ChaitanyaAkula/newsearches"
)

var Idjobs string
var IDslice []string

func GetSearchJobs(keyword string,loc string)[]string{
	var Searchvalue string
	db:=dbconnection.Connection()
	defer db.Close()
	x := keyword
	location := loc
	result, err := db.Query("select searchvalue from searchresults where searchvalue=?", x)
	if err != nil {
		
		log.Fatal(err)
	}
	for result.Next() {
         		
		err1 := result.Scan(&Searchvalue)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	fmt.Println("search value",Searchvalue)
	if x=="" && location==""{

		result1, err1 := db.Query("select idjobs from jobs order by idjobs desc")
		if err1 != nil {

			log.Fatal(err1)
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Idjobs)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Idjobs)
		}


	}
	if x=="" && location!=""{
		fmt.Println("test location:",location)
		result1, err1 := db.Query("select idjobs from jobs where match(location,country) against(? IN boolean mode) order by idjobs desc ", location,)
		if err1 != nil {

			log.Fatal(err1)
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Idjobs)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Idjobs)
		}
	}
	if x!=""{
	if x == Searchvalue {
		if location==""{
			fmt.Println("test search",x)
		result1, err1 := db.Query("select idjobs from jobs where match(jobtitle,location,country,typeofemp,jobdescription,jobcategory,firstskill,secondskill,thirdskill,otherskill,firstsalary,lastsalary,salarytype) against(? IN boolean mode) order by idjobs desc", x)
		if err1 != nil {

			log.Fatal(err1)
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Idjobs)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Idjobs)
		}
		}
		if location!=""{
			fmt.Println("test message:",x,location)
			result1, err1 := db.Query("select idjobs from jobs where match(jobtitle,location,country,typeofemp,jobdescription,jobcategory,firstskill,secondskill,thirdskill,otherskill,firstsalary,lastsalary,salarytype) against(? IN boolean mode) and match(location,country) against(? IN boolean mode) order by idjobs desc", x,location)
		if err1 != nil {

			log.Fatal(err1)
		}
		
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Idjobs)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Idjobs)
		}
	}
	}
	
	if x != Searchvalue {
		newsearches.NewSearchKeyword(x)
		if location==""{
			result1, err1 := db.Query("select idjobs from jobs where match(jobtitle,location,country,typeofemp,jobdescription,jobcategory,firstskill,secondskill,thirdskill,otherskill,firstsalary,lastsalary,salarytype) against(? IN boolean mode) order by idjobs desc", x)
			if err1 != nil {
	
				log.Fatal(err1)
			}
			IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Idjobs)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Idjobs)
		}
			}
			if location!=""{
				result1, err1 := db.Query("select idjobs from jobs where match(jobtitle,location,country,typeofemp,jobdescription,jobcategory,firstskill,secondskill,thirdskill,otherskill,firstsalary,lastsalary,salarytype) against(? IN boolean mode) and match(location,country) against(? IN boolean mode) order by idjobs desc", x,location)
			if err1 != nil {
	
				log.Fatal(err1)
			}
			
			IDslice=nil
			for result1.Next() {
	
				err2 := result1.Scan(&Idjobs)
				if err2 != nil {
					log.Fatal(err2)
				}
				IDslice=append(IDslice,Idjobs)
			}	
			}
	}
}

	return IDslice
}
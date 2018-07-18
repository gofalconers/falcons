package main

import (
	"log"
	"gopkg.in/ldap.v2"
	"fmt"
)

func main() {

	// username := "cn=admin,dc=example,dc=org"

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.200.20", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("cn=admin,dc=example,dc=org", "admin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("login successfule!")
	
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=org",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		// fmt.Sprintf("&(objectClass=organizationalUnit)(uid=%s))", username),
		"(&(objectClass=organizationalUnit))",
		[]string{"dn"},
		nil,
	)
	
	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(sr.Entries))
	if len(sr.Entries) != 1 {
		log.Fatal("User dose not exsite or too many entries returned")
	}

	userdn := sr.Entries[0].DN

	fmt.Println(userdn)
// 	err = l.Bind(userdn, "admin")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
}
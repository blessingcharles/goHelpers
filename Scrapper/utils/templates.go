package utils

import "fmt"

func Banner() {

	fmt.Println(
		`
				BOUNTY CAT
                             ^_^                      
                            {' '}   
                        (   /  (           
                        (  /   )    
                        \(_)__))    
                    THOMAS THE CAT 
                                - th3h04x
		`)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

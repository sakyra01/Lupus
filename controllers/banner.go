package controllers

import "fmt"

type Color string

func Banner() {
	PreBanner := `
Framework for collecting logs about user actions on Yandex360 disk
1. Logging information for up to six months
2. Build CEF syslog format file potentially for SOC(Security Operation Center)
3. Make notifications`
	banner := `

                                                                          .:::::^~^^~^
  o                                                                      .7:.  ..     .!
 <|>                                                                    7^            J^.::.
 / \                                                                    ^?           ~~.....^~
 \o/           o       o   \o_ __o     o       o       __o__       :~^:::77.       ..        ~^ 
  |           <|>     <|>   |    v\   <|>     <|>     />  \  	  ^7       ..                ?.
 / \          < >     < >  / \    <\  < >     < >     \o          :?                          !.
 \o/           |       |   \o/     /   |       |       v\          .J.                        ^!
  |            o       o    |     o    o       o        <\          Y         .        :~...:^^
 / \ _\o__/_   <\__ __/>   / \ __/>    <\__ __/>   _\o__</          ^!.    :~.          ^?...
                            |                                        .:::::5.            ?
                            |                                              !^     ...  .~:
                           / \                                              :~^^^^^.::::
                                          Version 2.0
`

	OutBanner := `
----The goal is to find users who download documents locally or try to share them with someone---------
	Use flag (-r/--report) to activate notification about alerts on telegram
	Example: ~$ ./lupus --report  
`
	fmt.Printf(PreBanner)
	fmt.Printf("\x1b[94m%s\x1b[0m", banner)
	fmt.Printf(OutBanner)
}

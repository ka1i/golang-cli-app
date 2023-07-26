package info

import "fmt"

/*
** Slant GolangAppKit
** http://patorjk.com/software/taag
 */
func ShowBanner() {
	fmt.Printf(`>>> %s ---> %s
   ______      __                  ___                __ __ _ __ 
  / ____/___  / /___ _____  ____  /   |  ____  ____  / //_/(_) /_
 / / __/ __ \/ / __ \/ __ \/ __ )/ /| | / __ \/ __ \/ ,<  / / __/
/ /_/ / /_/ / / /_/ / / / / /_/ / ___ |/ /_/ / /_/ / /| |/ / /_  
\____/\____/_/\__,_/_/ /_/\__, /_/  |_/ .___/ .___/_/ |_/_/\__/  
                         /____/      /_/   /_/ %s
`, Services, MicroService, Version.ToString())
}

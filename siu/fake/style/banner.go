package style

import "fmt"




const AUTOR string = "@mikis2323"


var Banner2 string = fmt.Sprintf(`%s

 ____                       ____ ____  _   _ 
|  _ \ ___  ___ ___  _ __  / ___|  _ \| \ | |
| |_) / _ \/ __/ _ \| '_ \| |   | | | |  \| |
|  _ <  __/ (_| (_) | | | | |___| |_| | |\  |
|_| \_\___|\___\___/|_| |_|\____|____/|_| \_| 
	subDomain & CDN scanner%s
	%sby:telegram > %s  %s
`, GREEN, END, YELLOW, AUTOR, END )


var _ string = fmt.Sprintf(`%s
888b     d888 d8b 888      d8b           .d8888b.  8888888b.  888b    888 
8888b   d8888 Y8P 888      Y8P          d88P  Y88b 888  "Y88b 8888b   888 
88888b.d88888     888                   888    888 888    888 88888b  888 
888Y88888P888 888 888  888 888 .d8888b  888        888    888 888Y88b 888 
888 Y888P 888 888 888 .88P 888 88K      888        888    888 888 Y88b888 
888  Y8P  888 888 888888K  888 "Y8888b. 888    888 888    888 888  Y88888 
888   "   888 888 888 "88b 888      X88 Y88b  d88P 888  .d88P 888   Y8888 
888       888 888 888  888 888  88888P'  "Y8888P"  8888888P"  888    Y888 
by:%s telegram %s%s %s %s`,GREEN,END,WHITE, AUTOR, SUB, END  )










package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	VERSION = "3.5.0"
)

func putAsciiArt(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case " ":
			color.Set(color.BgBlue)
			d = " "
		case "@":
			color.Set(color.BgBlack)
			d = " "
		case "#":
			color.Set(color.BgHiBlue)
			d = " "
		case "W":
			color.Set(color.BgWhite)
			d = " "
		case "_":
			color.Unset()
			d = " "
		case "\n":
			color.Unset()
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printLogo(s string) {
	for _, c := range s {
		d := string(c)
		switch string(c) {
		case "_":
			color.Set(color.FgWhite)
		case "\n":
			color.Unset()
		default:
			color.Set(color.FgHiBlue)
		}
		fmt.Print(d)
	}
	color.Unset()
}

func printUpdateName() {
	nameClr := color.New(color.FgHiCyan)
	txt := nameClr.Sprintf("            -- Evil Ginx Itwizardo Edition --")
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner1() {
	handleClr := color.New(color.FgHiBlue)
	versionClr := color.New(color.FgGreen)
	textClr := color.New(color.FgHiBlack)
	spc := strings.Repeat(" ", 10-len(VERSION))
	txt := textClr.Sprintf("    by Team Evil Ginx @thebigonescore ") + handleClr.Sprintf("https://t.me/thebigonescore @evilginx") + spc + textClr.Sprintf("version ") + versionClr.Sprintf("%s", VERSION)
	fmt.Fprintf(color.Output, "%s", txt)
}

func printOneliner2() {
	textClr := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	txt := textClr.Sprintf("           no ") + red.Sprintf("nginx") + white.Sprintf(" - ") + textClr.Sprintf("pure ") + red.Sprintf("evil")
	fmt.Fprintf(color.Output, "%s", txt)
}

func Banner() {
	fmt.Println()

	putAsciiArt("__                                     __\n")
	putAsciiArt("_   @@     @@@@@@@@@@@@@@@@@@@     @@   _\n")
	printLogo(`    ___________      __ __           __               `)
	fmt.Println()
	putAsciiArt("  @@@@    @@@@@@@@@@@@@@@@@@@@@    @@@@  \n")
	printLogo(`    \_   _____/__  _|__|  |    ____ |__| ____ ___  ___`)
	fmt.Println()
	putAsciiArt("  @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  \n")
	printLogo(`     |    __)_\  \/ /  |  |   / __ \|  |/    \\  \/  /`)
	fmt.Println()
	putAsciiArt("    @@@@@@@@@@###@@@@@@@###@@@@@@@@@@    \n")
	printLogo(`     |        \\   /|  |  |__/ /_/  >  |   |  \>    < `)
	fmt.Println()
	putAsciiArt("      @@@@@@@#####@@@@@#####@@@@@@@      \n")
	printLogo(`    /_______  / \_/ |__|____/\___  /|__|___|  /__/\_ \`)
	fmt.Println()
	putAsciiArt("       @@@@@@@###@@@@@@@###@@@@@@@       \n")
	printLogo(`            \/              /_____/         \/      \/`)
	fmt.Println()
	putAsciiArt("      @@@@@@@@@@@@@@@@@@@@@@@@@@@@@      \n")
	putAsciiArt("     @@@@@WW@@@WW@@WWW@@WW@@@WW@@@@@     \n")
	printUpdateName()
	fmt.Println()
	putAsciiArt("    @@@@@@WW@@@WW@@WWW@@WW@@@WW@@@@@@    \n")
	putAsciiArt("_   @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   _\n")
	printOneliner1()
	fmt.Println()
	putAsciiArt("__                                     __\n")
	fmt.Println()
	printOneliner2()
	fmt.Println()
}

func main() {
	Banner()
}

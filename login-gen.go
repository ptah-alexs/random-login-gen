package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var syl0 = strings.Split("aeiouy", "")
var syl1 = strings.Split("bcdfghjklmnpqrstvwxz", "")
var chanceTable = make([]int, 100)
var chanceList = []int{12, 2, 0, 2, 2, 2, 2, 0, 0} // bov, bvo, obv, a, b, ae, bv, aa, bb
var numbers int = 25
var piece int = 3
var delim string = "-"
var piecesizei = 3
var pcf string = ""

func fillChanceTable(chsyl []int) { // bov, bvo, obv, a, b, ae, bv, aa, bb
	var capb = len(chanceTable)
	var reg = len(chanceTable)
	for i := range len(chsyl) {
		if chsyl[i] > capb {
			chsyl[i] = capb
		}
		if chsyl[i] == 0 {
			continue
		}
		var inc = chsyl[i]
		for {
			var gh = rand.Intn(reg)
			if chanceTable[gh] == 0 {
				chanceTable[gh] = i + 1
				inc = inc - 1
				capb = capb - 1
			}
			if inc < 1 {
				break
			}
			if capb == 0 {
				return
			}
		}
	}
}

func genSyl() string {
	var res = ""
	var tt = rand.Intn(100)
	switch chanceTable[tt] {
	case 1:
		res = syl1[rand.Intn(len(syl1))] + syl0[rand.Intn(len(syl0))] + syl1[rand.Intn(len(syl1))]
	case 2:
		res = syl1[rand.Intn(len(syl1))] + syl1[rand.Intn(len(syl1))] + syl0[rand.Intn(len(syl0))]
	case 3:
		res = syl0[rand.Intn(len(syl0))] + syl1[rand.Intn(len(syl1))] + syl1[rand.Intn(len(syl1))]
	case 4:
		res = syl0[rand.Intn(len(syl0))]
	case 5:
		res = syl1[rand.Intn(len(syl1))]
	case 6:
		res = syl0[rand.Intn(len(syl0))] + syl0[rand.Intn(len(syl0))]
	case 7:
		res = syl1[rand.Intn(len(syl1))] + syl1[rand.Intn(len(syl1))]
	case 8:
		var tb = syl0[rand.Intn(len(syl0))]
		res = tb + tb
	case 9:
		var tb = syl1[rand.Intn(len(syl1))]
		res = tb + tb
	default:
		res = syl1[rand.Intn(len(syl1))] + syl0[rand.Intn(len(syl0))]
	}
	return res
}

func parseColon(arg string) []int {
	var tempTable = strings.Split(arg, ":")
	var ttLen int = len(tempTable)
	var res = make([]int, ttLen)
	for i := range ttLen {
		res[i] = toInt(tempTable[i])
	}
	return res
}

func toInt(str string) int {
	var res int = 0
	if sd, err := strconv.Atoi(str); err == nil {
		res = sd
	}
	return res
}

func fillPiecesizes(piece int, piecesizei int, sz string) []int {
	var res = make([]int, piece)
	if sz != "" {
		var tempTable = parseColon(sz)
		if len(tempTable) >= piece {
			res = tempTable[:piece]
		} else {
			res = tempTable
		}
	} else {
		for i := range piece {
			res[i] = piecesizei
		}
	}
	return res
}

func genLogin(delim string, arr []int) string {
	var sb strings.Builder
	var zd int = len(arr)
	sb.Grow(zd)
	for i := range zd {
		for range arr[i] {
			//res += genSyl()
			sb.WriteString(genSyl())
		}
		if i < zd-1 {
			//res += delim
			sb.WriteString(delim)
		}
	}
	return sb.String()
}

func argparse(args []string) {
	for i := range args {
		if args[i] == "--help" || args[i] == "-h" {
			printHelp()
			os.Exit(0)
		}
		var tempTableArg = strings.Split(args[i], "=")
		if len(tempTableArg) != 2 {
			continue
		}
		switch tempTableArg[0] {
		case "-q", "--quantity":
			numbers = toInt(tempTableArg[1])
		case "-a", "--piece-amount":
			piece = toInt(tempTableArg[1])
		case "-s", "--piece-size":
			piecesizei = toInt(tempTableArg[1])
		case "-t", "--piece-tune":
			pcf = tempTableArg[1]
		case "-c", "--chance-tune":
			var tempTable = parseColon(tempTableArg[1])
			var clen = len(chanceList)
			if len(tempTable) >= clen {
				chanceList = tempTable[:clen]
			} else {
				chanceList = tempTable
			}
		}
	}
}

func printHelp() {
	fmt.Println(`login-gen arguments:
	--help, -h   — Print help
	--quantity, -q   — quantity of logins at same time
	--piece-amount, -a   — quantity of parts of login at same time
	--piece-size, -s   — quantity of syllabes in piece of login
	--piece-tune, -t   —  quantity of syllabes in piece of login, splited by colon, ex. "1:2:3" -> xi-neqa-zypoja
	--chance-tune, -c — tune chance of syllabes types, splited by colon, ex. "12:2:14:1:5:6:8:0:1.
					    Types of syllabes: bov, bvo, obv, a, b, ae, bv, aa, bb

	Default values:
	-q=25 -a=3 -s=3 -c="12:2:0:2:2:2:2:0:0"`)
}

func main() {
	argparse(os.Args[1:])
	fillChanceTable(chanceList)
	var fsz []int = fillPiecesizes(piece, piecesizei, pcf)
	for range numbers {
		fmt.Println(genLogin(delim, fsz))
	}
}

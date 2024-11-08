package sort

import (
	"fmt"
	"github.com/yinzhenzhen/learning-programming/types"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(mergeIntervals([][]int{{1, 9}, {2, 5}, {19, 20}, {10, 11}, {12, 20}, {0, 3}, {0, 1}, {0, 2}}))
	// [0, 1] [0, 2] [0, 3] [1, 9] [2, 5] [10, 11] [12, 20] [19, 20]
	// [0, 9] [10, 11] [12, 20]

	//0 1, 1 0
	//1 1, 2 0
	//2 1, 3 0
}

func Test_insertionSortList2(t *testing.T) {
	//head := types.ConstructDoubleLink([]int{6,5,4,3,2,1})
	//head.Print()
	//fmt.Println()
	//insertionSortList(head)
	//head.Print()

	head := types.ConstructLink([]int{6, 5, 4, 3, 2, 1})
	head.Print()
	fmt.Println()
	insertionSortList2(head)
	head.Print()
}

func Test_sortList(t *testing.T) {
	head := types.ConstructLink([]int{6, 5, 4, 3, 2, 1})
	head.Print()
	fmt.Println()
	sortList(head)
	head.Print()
}

func Test_rankTeams(t *testing.T) {
	//fmt.Println(rankTeams([]string{"ABC","ACB","ABC","ACB","ACB"}))
	//fmt.Println(rankTeams([]string{"WXYZ","XYZW"}))
	//fmt.Println(rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}))
	//fmt.Println(rankTeams([]string{"BCA","CAB","CBA","ABC","ACB","BAC"}))
	//fmt.Println(rankTeams([]string{"M","M","M","M"}))
	//fmt.Println(compare([]int{11,2}, []int{9, 3}))
	fmt.Println(rankTeams([]string{"XDELQOKZUGMCBTWAJNPFRVIS", "IQKGUDCWMNLSABZEPVXFJRTO", "SFPUJWZABNRLQEMXKIGDVCTO",
		"FISQWOZAGXJDNLTKPRECVUBM", "BMXIRQOWJSLPDNGAEFKZTUVC", "SCUFTIAQMWNXBGZEOJVDLKRP", "LMNXWKDBTAGPFUCJSROIQZEV",
		"NVDICPQFBZLRGWTMXJOAUESK", "PGZRUWNFMKTVJSLCDEXOABIQ", "BFKNDXRIGWTVSUJOCMPZAEQL", "MJLDKRNWPQXUFSTEGCOBVAIZ",
		"VCIMLFJSPZTDRKBAGXWUOEQN", "TERGMVDPILNQSFJCAKZWXBOU", "KOPFCIDQSVAMRBLNTWUZXJGE", "NKWPJGUAFOTZIERXMSLBDCQV",
		"QUKLSIFCDVZEAMXGJNTWROPB", "DSAGZRFWTMJEILKOPUNQVXCB", "ARQIVZNCMBJGWUPKEOSLFTDX", "ESMATLNJKPRFGXWVUCOQDIZB",
		"QNGTVDFICPSKMWUJBOZELAXR", "JAOMTSGFQPDWKRNLUXEVBCIZ", "FORAQBEDSTNPUZICVXGWMKLJ", "UWAFSXITRMPEGVQDLZBCNJOK",
		"IKFOVMNWJLCZBUAGERQPXTSD", "SLEOPWNDBUVIGQJZMRACXTFK", "GTDKMOJSILNAVXFPZQCUWRBE", "AZGRVFJDEUSCNLTBXWQIPMOK",
		"VQPNUMCJKSBXWIAFTGDOEZRL", "PVNODUXIMJSGACQTWRFBKLZE", "QGKDUAIMEVTXSZFBJLOWNPRC", "MDUOXBVQJPSAGNIKLRFZTCWE",
		"FMRTEDKWPVLGBICXSNZAQJOU", "WUEFMGXSIRNAOPCBDLQTKJZV", "UDGSPEQVZOBMNICJKXTWALRF", "SFNWLTARVKPJQZMGUDBOICEX",
		"FSPCMAJZEKNRLDWBGOXQIUVT", "VRUOAMJQGFSIZNXLEPWTDCBK", "VNCMBRLKUPDISJEWOAZGTXFQ", "VGJQZXMPKEIRLTNODCABWFSU",
		"IMALJVBCXSQPOUGEKNTRFDWZ", "UPJBVKQOAIDGWEZRLTNSXFMC", "IUAPCZFWOQTXBSNLKGVRDMEJ", "DOIRNGMKJLEUCPTWZQSVXFAB",
		"VJMWTPZFOGBNIDCQAULERKSX", "OAISKWCDLVRQFGNZPUTJEMXB", "NBTGWMEUROJDCKAZLFQXISPV", "DSOBAIJXFWNTUVERGQMCLKPZ",
		"SBLFZOGCNDWPXAEKTVRJUMIQ", "EKWMAQNURJBZPLFDSGOVIXCT", "BKXRSNVMWGAPCIFLZJTEUQOD", "VQNWJSZTCPEURAKBLIXOFGMD",
		"TQKFDXCWIUARVBELNOPJZGMS", "MSCPWBTKQUFVDGEXROANJLIZ", "XFQVITPRZMONWCLKBUAJGSED", "MVKLRFQWICOPDJBUSTXNEZGA",
		"DMKBRGLONJEIFTXZPUACQWSV", "NKSJBQZORVFTXWAUCDLGEMIP", "JQUIXNBRDMOLGSTFVAZCPEWK", "MEBXTJUNVWCQAZGRLIKSPDFO",
		"EKNUQILMTDPOSFJZWXGVCARB", "JOZNFECPVSXRLGDUKQWBITAM", "UXNZISEGTCVABJMKFLWDQPOR", "IJCFLPGZRXQUMANODKBVSETW",
		"KCJQRFEWNGAMBIXZVLPDSUTO", "TMSEDRAUWZGQXJPIFVLCOBKN", "ZPRXJNOEKCGFUSDQAILWMVTB", "WTBXANSJLRKFEPOIVDMQUZCG",
		"IRLVZNQPECOJATFWUSDMGKBX", "BGMPAZOIKWCSUDJNXLTVREQF", "BZXSTPRMLEKWCFIVONGADUQJ", "UFGZCWTARKSBDJXPMVQIENOL",
		"ZBROLDUFMNQAPJEGITXWCSKV", "FWLJXQDTRCGNVUZMAPBKISEO", "SJXLKAVZCPBIEWFQTNODGRMU", "VAIBDLOEXGUZQWMPNSJTCRKF",
		"XFLMPECKOQINGWSZJRTVBADU", "ZAJCPMSKUBFWEXQDTVGNLORI", "FUVECNMKZWJBGXTDLPASIOQR", "COFKSVMTNLBWXAREZDQIJUGP",
		"TKRVFXUIONEALPCDSJGBZWQM", "VJNQLTARZSGCPOMFBUWIXKDE", "GMUSFBPDZATKWEQONXCJILVR", "RFNCLTQZXMSVWPIUBGDEAKOJ",
		"NARSWLXMEIGQCFBOVDUJPZTK", "OBZVJGSKDUCXAQWPIFNEMTRL", "RIEQDWFSLCGTPKNBAJUVMZOX", "GBAXDSVUEOJNTLPIRKZCWQMF",
		"ZOETNALWVPGXRIMQKDBFUSCJ", "NAMBJKFZQTOSGWCDPIRXLEUV", "NWQJUZMKIXCGDSBREFPATLOV", "KWAECTUZODSPIJMNLVQRXFBG",
		"JXPWZNOCGRDKLQITUSEVAFBM", "FNAXSLPRWUBMQJICZTVKDOEG", "GMAWSUXVPZINTODEBRJFKQLC", "FACUKTLWNJSXRQMEDZVGBOIP",
		"FXEGPVLUJDRNQMTOWSIKAZCB", "ZNEKOJXDVQSUFARMLWPGICTB", "IRSZQGUKVJXCANFWLOBDEMPT", "IUEDPRJFVGQSAXNOWCBMZTLK",
		"XGKLOPJVQCWTBMSNUZARIEDF", "VSOJUPRFDCNQXTLMBKIGZWEA", "UDGIKPAFNTZWCJRMSLVOXEQB", "GQOIULRNMCJVPFEAZBSDWKTX",
		"ZJUTMKIFXRSCOPENVGLBDAQW", "ZUQOBCPWNJKVARLTDEGIXFSM", "ADBOKTGCWFMNVUEZLRPIQJSX", "ALDVEQZKCTOJRWBMISPNXFGU",
		"CEGUFDRXLAVTQNWOKMPZIJBS", "EMBXCVQDJLNWSGIZPRAFTUKO", "KJSQEPNOVGRWZLUIADFBCTXM", "UWVFQEDASJBGXCNPTZIORMLK",
		"XFLCZSMJAVEPIQDRTUBNKGWO", "URFBQMVATJWEPDCNGSILOZXK", "ZLDPUVIWOJEXGNTQFRBKSACM", "IEQBVDUPNLCGJASFRXTWZOMK",
		"TVWMROKSUPCXNFEALDQZIGBJ", "GFKEZOICNTXPASMBJVRUWLDQ", "NVWKLBMQJFPIRCSDOZGTXEAU", "QANRCGLPMDZKFOTXBVJEIWUS",
		"UPCLXNVQZTRKMEBGAOSFDJWI", "PANMJZIFCWGKQDXSRUETOLVB", "JGWVNUXLKAEODTBZIQRCMPFS", "TEOMUIZFSXPGWLADRBQKJVCN",
		"XPFDJLGZAWEKTRVCMBNUSOQI", "OCQSTRKAZPIMUBNDGEXJVWLF", "MILWTBRAVGZODFJUCSPXQENK", "FREBOSZDPMCNUJXAVKIGLWQT",
		"TQWPUCERNGIBMXDZKJVASFOL", "APLDRMWCQNFUTSKXVEGJIZBO", "SUWEMPJNFITABXVOCRZLDQKG", "FGRZBWMUTIACNEJKVQPXDOSL",
		"PBUWQEIDOLXCRNFVSZJTKAMG", "BRFZUGVLWEJKNIASPMTXDQOC", "XLIPMTZQSUFVANKOGBEJDRCW", "GSCRPUAKMEIDBVOLQXWZNFTJ",
		"KCQMJXRWSFGNBDLAUOTEIVZP", "GKDPJSRMALVEXUNWOICTFQZB"}))
}

func Test_getKth(t *testing.T) {
	fmt.Println(getKth(1, 1, 1))
}

func Test_findDiagonalOrder2(t *testing.T) {
	fmt.Println(findDiagonalOrder2([][]int{{6}, {8}, {6, 1, 6, 16}}))
	//fmt.Println(findDiagonalOrder2([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	fmt.Println(findDiagonalOrder2([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}))
}

func Test_arrangeWords(t *testing.T) {
	fmt.Println(arrangeWords("To be or not to be"))
}

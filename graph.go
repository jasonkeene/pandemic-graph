package pandemic

import "log"

var (
	Vertices = []string{
		// north america
		"atlanta",       // 0
		"washington",    // 1
		"san francisco", // 2
		"los angeles",   // 3
		"chicago",       // 4
		"new york",      // 5
		"montreal",      // 6
		"miami",         // 7
		"mexico city",   // 8

		// south america
		"bogata",       // 9
		"lima",         // 10
		"san paulo",    // 11
		"santiago",     // 12
		"buenos aires", // 13

		// europe
		"london",        // 14
		"madrid",        // 15
		"essen",         // 16
		"paris",         // 17
		"milan",         // 18
		"st petersburg", // 19
		"instanbul",     // 20
		"moscow",        // 21

		// africa
		"algiers",      // 22
		"cairo",        // 23
		"lagos",        // 24
		"kinsasha",     // 25
		"khartoum",     // 26
		"johannesburg", // 27

		// asia
		"tehran",    // 28
		"baghdad",   // 29
		"riyadh",    // 30
		"karachi",   // 31
		"delhi",     // 32
		"mumbai",    // 33
		"chennai",   // 34
		"kolkata",   // 35
		"bangkok",   // 36
		"hong kong", // 37
		"shanghai",  // 38
		"beijing",   // 39
		"seoul",     // 40

		// pacific rim
		"tokyo",            // 41
		"osaka",            // 42
		"taipei",           // 43
		"manila",           // 44
		"ho chi minh city", // 45
		"jakarta",          // 46
		"sydney",           // 47
	}
	AdjacencyList = makeAdjacencyList([][]int{
		// north america
		{1, 4, 7},         // atlanta (0)
		{0, 5, 6, 7},      // washington (1)
		{3, 4, 41, 44},    // san francisco (2)
		{2, 4, 8, 10, 47}, // los angeles (3)
		{0, 2, 3, 6, 8},   // chicago (4)
		{1, 6, 14, 15},    // new york (5)
		{1, 4, 5},         // montreal (6)
		{0, 1, 8, 9},      // miami (7)
		{3, 4, 7, 9, 10},  // mexico city (8)

		// south america
		{7, 8, 10, 11, 13}, // bogata (9)
		{3, 8, 9, 12},      // lima (10)
		{9, 13, 15, 24},    // san paulo (11)
		{10, 13},           // santiago (12)
		{9, 11, 12, 27},    // buenos aires (13)

		// europe
		{5, 15, 17, 16},          // london (14)
		{5, 11, 14, 17, 22},      // madrid (15)
		{14, 17, 18, 19},         // essen (16)
		{14, 15, 16, 18, 22},     // paris (17)
		{16, 17, 20},             // milan (18)
		{16, 20, 21},             // st petersburg (19)
		{18, 19, 21, 22, 23, 29}, // instanbul (20)
		{19, 20, 28},             // moscow (21)

		// africa
		{15, 17, 20, 23},     // algiers (22)
		{20, 22, 26, 29, 30}, // cairo (23)
		{11, 25, 26},         // lagos (24)
		{24, 26, 27},         // kinsasha (25)
		{23, 24, 25, 27},     // khartoum (26)
		{13, 25, 26},         // johannesburg (27)

		// asia
		{21, 29, 31, 32},         // tehran (28)
		{20, 23, 28, 30},         // baghdad (29)
		{23, 29, 31},             // riyadh (30)
		{28, 30, 32, 33},         // karachi (31)
		{28, 31, 33, 34, 35},     // delhi (32)
		{31, 32, 34},             // mumbai (33)
		{32, 33, 35, 46},         // chennai (34)
		{32, 34, 36, 37},         // kolkata (35)
		{35, 37, 45, 46},         // bangkok (36)
		{35, 36, 38, 43, 44, 45}, // hong kong (37)
		{37, 39, 40, 41, 43},     // shanghai (38)
		{38, 40},                 // beijing (39)
		{38, 39, 41},             // seoul (40)

		// pacific rim
		{2, 38, 40, 42},     // tokyo (41)
		{41, 43},            // osaka (42)
		{37, 38, 42, 44},    // taipei (43)
		{2, 37, 43, 45, 47}, // manila (44)
		{36, 37, 44, 46},    // ho chi minh city (45)
		{34, 36, 45, 47},    // jakarta (46)
		{3, 44, 46},         // sydney (47)
	})
)

type IntSet map[int]struct{}

func (s IntSet) Copy() IntSet {
	out := make(IntSet, len(s))
	for k, v := range s {
		out[k] = v
	}
	return out
}

func makeAdjacencyList(in [][]int) []IntSet {
	out := make([]IntSet, len(in))
	for vi, edges := range in {
		out[vi] = make(IntSet, len(edges))
		for _, edge := range edges {
			out[vi][edge] = struct{}{}
		}
	}
	return out
}

func init() {
	verify()
}

func verify() {
	if len(AdjacencyList) != len(Vertices) {
		log.Panic("AdjacencyList is not the same lenght as vertices")
	}

	for vi, edges := range AdjacencyList {
		for edge := range edges {
			if _, ok := AdjacencyList[edge][vi]; !ok {
				log.Panicf("No backwards edge for: %d vertex: %d", edge, vi)
			}
		}
	}
}

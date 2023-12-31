package internals

type Data struct {
	Blogs          []*BlogHeader
	DetailMap      map[int][]string
	SearchData     []string
	Accents        map[string]bool
	Diacritics_Map map[rune]string
}

type Dream struct {
	BlogHeader []*BlogHeader `json:"BlogHeader"`
	BlogDetail []*Detail     `json:"BlogDetail"`
}

type Detail struct {
	ID      int    `json:"BlogDetailId"`
	BlogID  int    `json:"BlogId"`
	Content string `json:"BlogContent"`
}

type BlogHeader struct {
	BlogId    int    `json:"BlogId"`
	BlogTitle string `json:"BlogTitle"`
}

// type IndexSearchCache struct {
// 	BlogDetailID int    `json:"DetailID"`
// 	BlogContent  string `json:"Content"`
// }

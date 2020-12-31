package bbs

import (
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUserecRaw   *ptttype.UserecRaw
	testUserec1     *Userec
	testUserec2     *Userec
	testUserec3     *Userec
	testUserec4     *Userec
	testUserec5     *Userec
	testUserecEmpty *Userec
	testUserec6     *Userec

	testOpenUserecFile1     []*Userec = nil
	TEST_N_OPEN_USER_FILE_1           = 50

	testBoardSummaryRaw6 *ptttype.BoardSummaryRaw

	testBoardSummary6  *BoardSummary
	testBoardSummary7  *BoardSummary
	testBoardSummary8  *BoardSummary
	testBoardSummary11 *BoardSummary

	testArticleSummary0 *ArticleSummary
	testArticleSummary1 *ArticleSummary

	testContent1 []byte
)

func initTestVars() {
	if testUserecRaw != nil {
		return
	}

	testUserecRaw = &ptttype.UserecRaw{
		Version: ptttype.PASSWD_VERSION,
		UserID:  ptttype.UserID_t{0x53, 0x59, 0x53, 0x4f, 0x50}, // SYSOP
		RealName: ptttype.RealName_t{ // CodingMan
			0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e,
		},
		Nickname:   ptttype.Nickname_t{0xaf, 0xab}, // 神
		PasswdHash: ptttype.Passwd_t{0x62, 0x68, 0x77, 0x76, 0x4f, 0x4a, 0x74, 0x66, 0x54, 0x31, 0x54, 0x41, 0x49, 0x00},

		UFlag:        33557088,
		UserLevel:    536871943,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost: ptttype.IPv4_t{ //59.124.167.226
			0x35, 0x39, 0x2e, 0x31, 0x32, 0x34, 0x2e, 0x31, 0x36, 0x37,
			0x2e, 0x32, 0x32, 0x36,
		},
		Address: ptttype.Address_t{ //新竹縣子虛鄉烏有村543號
			0xb7, 0x73, 0xa6, 0xcb, 0xbf, 0xa4, 0xa4, 0x6c, 0xb5, 0xea,
			0xb6, 0x6d, 0xaf, 0x51, 0xa6, 0xb3, 0xa7, 0xf8, 0x35, 0x34,
			0x33, 0xb8, 0xb9,
		},
		Over18:   true,
		Pager:    ptttype.PAGER_ON,
		Career:   ptttype.Career_t{0xa5, 0xfe, 0xb4, 0xba, 0xb3, 0x6e, 0xc5, 0xe9}, //全景軟體
		LastSeen: 1600681288,
	}

	testUserec1 = &Userec{
		Version:  4194,
		UUserID:  UUserID("SYSOP"),
		Username: "SYSOP",
		Realname: []byte{ // CodingMan
			0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e,
		},
		Nickname: []byte{0xaf, 0xab}, // 神

		Uflag:        33557088,
		Userlevel:    536871943,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		Lasthost:     "59.124.167.226",
		Address: []byte{ //新竹縣子虛鄉烏有村543號
			0xb7, 0x73, 0xa6, 0xcb, 0xbf, 0xa4, 0xa4, 0x6c, 0xb5, 0xea,
			0xb6, 0x6d, 0xaf, 0x51, 0xa6, 0xb3, 0xa7, 0xf8, 0x35, 0x34,
			0x33, 0xb8, 0xb9,
		},
		Over18: true,
		Pager:  ptttype.PAGER_ON,
		Career: []byte{0xa5, 0xfe, 0xb4, 0xba, 0xb3, 0x6e, 0xc5, 0xe9}, //全景軟體
	}

	testUserec6 = &Userec{
		Version:      ptttype.PASSWD_VERSION,
		UUserID:      UUserID("B1"),
		Username:     "B1",
		Lasthost:     "127.0.0.1",
		Uflag:        33557088,
		Userlevel:    7,
		Numlogindays: 1,
		Pager:        1,
		Over18:       true,
	}

	testUserecEmpty = &Userec{}

	testBoardSummaryRaw6 = &ptttype.BoardSummaryRaw{
		Bid:     6,
		BrdAttr: ptttype.BRD_POSTMASK,
		Brdname: &ptttype.BoardID_t{'A', 'L', 'L', 'P', 'O', 'S', 'T', 0x00, 0x2e, 0x2e, 0x2e, 0x2e},
		Title: &ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xf3, 0xaa,
			0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0xb7, 0x73,
			0xa4, 0xe5, 0xb3, 0xb9, 0x00, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
			0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49, 0x00, 0x6e,
		},
		BM:       []*ptttype.UserID_t{},
		StatAttr: ptttype.NBRD_FAV,
	}

	testBoardSummary6 = &BoardSummary{
		BBoardID: BBoardID("6_ALLPOST"),
		BrdAttr:  ptttype.BRD_POSTMASK,
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "ALLPOST",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb8, 0xf3, 0xaa, 0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41,
			0x4c, 0xb7, 0x73, 0xa4, 0xe5, 0xb3, 0xb9,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []UUserID{},
	}

	testBoardSummary7 = &BoardSummary{
		BBoardID: BBoardID("7_deleted"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "deleted",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb8, 0xea, 0xb7, 0xbd, 0xa6, 0x5e, 0xa6, 0xac, 0xb5, 0xa9,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []UUserID{},
	}

	testBoardSummary8 = &BoardSummary{
		BBoardID: BBoardID("8_Note"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "Note",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xb0, 0xca, 0xba, 0x41, 0xac, 0xdd, 0xaa, 0x4f, 0xa4, 0xce,
			0xba, 0x71, 0xa6, 0xb1, 0xa7, 0xeb, 0xbd, 0x5a,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []UUserID{},
	}

	testBoardSummary11 = &BoardSummary{
		BBoardID: BBoardID("11_EditExp"),
		StatAttr: ptttype.NBRD_FAV,
		Brdname:  "EditExp",
		BoardClass: []byte{
			0xbc, 0x54, 0xad, 0xf9,
		},
		RealTitle: []byte{
			0xbd, 0x64, 0xa5, 0xbb, 0xba, 0xeb, 0xc6, 0x46, 0xa7, 0xeb,
			0xbd, 0x5a, 0xb0, 0xcf,
		},
		BoardType: []byte{0xa1, 0xb7},
		BM:        []UUserID{},
	}

	testArticleSummary0 = &ArticleSummary{
		BBoardID:   BBoardID("10_WhoAmI"),
		ArticleID:  "1Vo_M_CDSYSOP",
		IsDeleted:  false,
		Filename:   "M.1607202239.A.30D",
		CreateTime: 1607202239,
		MTime:      1607202238,
		Owner:      "SYSOP",
		Title: []byte{
			0x5b, 0xb0, 0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7,
			0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x48, 0xa1,
			0xe3,
		},

		Class: []byte{0xb0, 0xdd, 0xc3, 0x44},
	}

	testArticleSummary1 = &ArticleSummary{
		BBoardID:   BBoardID("10_WhoAmI"),
		ArticleID:  "1Vo_f30DSYSOP",
		IsDeleted:  false,
		Filename:   "M.1607203395.A.00D",
		CreateTime: 1607203395,
		MTime:      1607203394,
		Owner:      "SYSOP",
		Title: []byte{
			0x5b, 0xa4, 0xdf, 0xb1, 0x6f, 0x5d, 0x20, 0xb5,
			0x4d, 0xab, 0xe1, 0xa9, 0x4f, 0xa1, 0x48, 0xa1,
			0xe3,
		},

		Filemode: ptttype.FILE_MARKED,

		Class: []byte{0xa4, 0xdf, 0xb1, 0x6f},
	}

	testContent1 = []byte{
		0xa7, 0x40, 0xaa, 0xcc, 0x3a, 0x20, 0x53, 0x59, 0x53,
		0x4f, 0x50, 0x20, 0x28, 0x29, 0x20, 0xac, 0xdd, 0xaa,
		0x4f, 0x3a, 0x20, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
		0x0a, 0xbc, 0xd0, 0xc3, 0x44, 0x3a, 0x20, 0x5b, 0xb0,
		0xdd, 0xc3, 0x44, 0x5d, 0x20, 0xa7, 0xda, 0xac, 0x4f,
		0xbd, 0xd6, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0xae, 0xc9,
		0xb6, 0xa1, 0x3a, 0x20, 0x53, 0x75, 0x6e, 0x20, 0x44,
		0x65, 0x63, 0x20, 0x20, 0x36, 0x20, 0x30, 0x35, 0x3a,
		0x30, 0x33, 0x3a, 0x35, 0x37, 0x20, 0x32, 0x30, 0x32,
		0x30, 0x0a, 0x0a, 0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6,
		0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0xa7, 0xda, 0xa6,
		0x62, 0xad, 0xfe, 0xb8, 0xcc, 0xa1, 0x48, 0xa1, 0xe3,
		0x0a, 0x0a, 0xa7, 0xda, 0xac, 0xb0, 0xa4, 0xb0, 0xbb,
		0xf2, 0xb7, 0x7c, 0xa6, 0x62, 0xb3, 0x6f, 0xb8, 0xcc,
		0xa9, 0x4f, 0xa1, 0x48, 0xa1, 0xe3, 0x0a, 0x0a, 0x2d,
		0x2d, 0x0a, 0xa1, 0xb0, 0x20, 0xb5, 0x6f, 0xab, 0x48,
		0xaf, 0xb8, 0x3a, 0x20, 0xa7, 0xe5, 0xbd, 0xf0, 0xbd,
		0xf0, 0x20, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x28,
		0x70, 0x74, 0x74, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
		0x2e, 0x74, 0x65, 0x73, 0x74, 0x29, 0x2c, 0x20, 0xa8,
		0xd3, 0xa6, 0xdb, 0x3a, 0x20, 0x31, 0x37, 0x32, 0x2e,
		0x31, 0x38, 0x2e, 0x30, 0x2e, 0x31, 0x0a,
	}

	testOpenUserecFile1 = make([]*Userec, TEST_N_OPEN_USER_FILE_1)
	for i := 0; i < TEST_N_OPEN_USER_FILE_1; i++ {
		testOpenUserecFile1[i] = testUserecEmpty
	}
	testOpenUserecFile1[0] = testUserec1
	testOpenUserecFile1[1] = testUserec2
	testOpenUserecFile1[2] = testUserec3
	testOpenUserecFile1[3] = testUserec4
	testOpenUserecFile1[4] = testUserec5
}

func freeTestVars() {
}
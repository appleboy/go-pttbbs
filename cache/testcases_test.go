package cache

import (
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	testUserInfo1 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'S', 'Y', 'S', 'O', 'P'},
		Uid:    1,
		From:   ptttype.From_t{'D'},
		Pid:    3,
	}

	testUserInfo2 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'A', '1'},
		Uid:    2,
		From:   ptttype.From_t{'B'},
		Pid:    2,
	}

	testUserInfo3 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'A', '0'},
		Uid:    3,
		From:   ptttype.From_t{'S'},
		Pid:    1,
	}
	testUserInfo4 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '2'},
		Uid:    5,
		From:   ptttype.From_t{'K'},
		Pid:    5,
	}
	testUserInfo5 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '1'},
		Uid:    4,
		From:   ptttype.From_t{'H'},
		Pid:    4,
	}

	testUserInfo6 = ptttype.UserInfoRaw{
		UserID: ptttype.UserID_t{'Z', '3'},
		Uid:    6,
		From:   ptttype.From_t{'K'},
		Pid:    6,
	}

	testBoardHeader0 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'b', '0'},
	}
	testBoardHeader1 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'b', '1'},
	}
	testBoardHeader2 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'b', '2'},
	}

	testBoardHeader3 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'S', 'Y', 'S', 'O', 'P'},
		Title: ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		BrdAttr: ptttype.BRD_POSTMASK,
		Gid:     2,
	}

	testBM4          ptttype.BM_t
	testBoardHeader4 ptttype.BoardHeaderRaw

	testBM13          ptttype.BM_t
	testBoardHeader13 ptttype.BoardHeaderRaw
)

var (
	testBCacheName  = make([]ptttype.BoardID_t, 12)
	testBCacheTitle = make([]ptttype.BoardTitle_t, 12)

	// strncasecmp:
	//   1........... (1)
	//   2........... (4)
	//   ALLHIDPOST (11)
	//   ALLPOST (5)
	//   deleted (6)
	//   EditExp (10)
	//   junk (2)
	//   Note (7)
	//   Record (8)
	//   Security (3)
	//   SYSOP (0)
	//   WhoAmI (9)
	testSortedByName = []ptttype.BidInStore{1, 4, 11, 5, 6, 10, 2, 7, 8, 3, 0, 9}

	// strncmp on title-type:
	//   .... 1........... (1)
	//   .... 2........... (4)
	//   發電 junk (2)
	//   發電 Security (3)
	//   嘰哩 ALLHIDPOST (11)
	//   嘰哩 ALLPOST (5)
	//   嘰哩 deleted (6)
	//   嘰哩 EditExp (10)
	//   嘰哩 Note (7)
	//   嘰哩 Record (8)
	//   嘰哩 SYSOP (0)
	//   嘰哩 WhoAmI (9)
	// strncmp title[:4]:
	testSortedByClass = []ptttype.BidInStore{1, 4, 2, 3, 11, 5, 6, 10, 7, 8, 0, 9}
)

var (
	isInitTestCases = false
)

func initTestCases() {
	if isInitTestCases {
		return
	}
	isInitTestCases = true
	copy(testBCacheName[0][:], []byte("SYSOP"))
	copy(testBCacheName[1][:], []byte("1..........."))
	copy(testBCacheName[2][:], []byte("junk"))
	copy(testBCacheName[3][:], []byte("Security"))
	copy(testBCacheName[4][:], []byte("2..........."))
	copy(testBCacheName[5][:], []byte("ALLPOST"))
	copy(testBCacheName[6][:], []byte("deleted"))
	copy(testBCacheName[7][:], []byte("Note"))
	copy(testBCacheName[8][:], []byte("Record"))
	copy(testBCacheName[9][:], []byte("WhoAmI"))
	copy(testBCacheName[10][:], []byte("EditExp"))
	copy(testBCacheName[11][:], []byte("ALLHIDPOST"))

	copy(testBCacheTitle[0][:], []byte{ //嘰哩 ◎站長好!
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
		0xf8, 0xa6, 0x6e, 0x21,
	})
	copy(testBCacheTitle[1][:], []byte{ //.... Σ中央政府  《高壓危險,非人可敵》
		0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0xa3, 0x55, 0xa4, 0xa4, 0xa5,
		0xa1, 0xac, 0x46, 0xa9, 0xb2, 0x20, 0x20, 0xa1, 0x6d, 0xb0,
		0xaa, 0xc0, 0xa3, 0xa6, 0x4d, 0xc0, 0x49, 0x2c, 0xab, 0x44,
		0xa4, 0x48, 0xa5, 0x69, 0xbc, 0xc4, 0xa1, 0x6e,
	})
	copy(testBCacheTitle[2][:], []byte{ //發電 ◎雜七雜八的垃圾
		0xb5, 0x6f, 0xb9, 0x71, 0x20, 0xa1, 0xb7, 0xc2, 0xf8, 0xa4,
		0x43, 0xc2, 0xf8, 0xa4, 0x4b, 0xaa, 0xba, 0xa9, 0x55, 0xa7,
		0xa3,
	})
	copy(testBCacheTitle[3][:], []byte{ //發電 ◎站內系統安全
		0xb5, 0x6f, 0xb9, 0x71, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xa4,
		0xba, 0xa8, 0x74, 0xb2, 0xce, 0xa6, 0x77, 0xa5, 0xfe,
	})
	copy(testBCacheTitle[4][:], []byte{ //.... Σ市民廣場     報告  站長  ㄜ！
		0x2e, 0x2e, 0x2e, 0x2e, 0x20, 0xa3, 0x55, 0xa5, 0xab, 0xa5,
		0xc1, 0xbc, 0x73, 0xb3, 0xf5, 0x20, 0x20, 0x20, 0x20, 0x20,
		0xb3, 0xf8, 0xa7, 0x69, 0x20, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
		0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49,
	})
	copy(testBCacheTitle[5][:], []byte{ //嘰哩 ◎跨板式LOCAL新文章
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xf3, 0xaa,
		0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0xb7, 0x73,
		0xa4, 0xe5, 0xb3, 0xb9,
	})
	copy(testBCacheTitle[6][:], []byte{ //嘰哩 ◎資源回收筒
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xea, 0xb7,
		0xbd, 0xa6, 0x5e, 0xa6, 0xac, 0xb5, 0xa9,
	})
	copy(testBCacheTitle[7][:], []byte{ //嘰哩 ◎動態看板及歌曲投稿
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb0, 0xca, 0xba,
		0x41, 0xac, 0xdd, 0xaa, 0x4f, 0xa4, 0xce, 0xba, 0x71, 0xa6,
		0xb1, 0xa7, 0xeb, 0xbd, 0x5a,
	})
	copy(testBCacheTitle[8][:], []byte{ //嘰哩 ◎我們的成果
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa7, 0xda, 0xad,
		0xcc, 0xaa, 0xba, 0xa6, 0xa8, 0xaa, 0x47,
	})
	copy(testBCacheTitle[9][:], []byte{ //嘰哩 ◎呵呵，猜猜我是誰！
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xa8, 0xfe, 0xa8,
		0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71, 0xa7, 0xda, 0xac,
		0x4f, 0xbd, 0xd6, 0xa1, 0x49,
	})
	copy(testBCacheTitle[10][:], []byte{ //嘰哩 ◎範本精靈投稿區
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xbd, 0x64, 0xa5,
		0xbb, 0xba, 0xeb, 0xc6, 0x46, 0xa7, 0xeb, 0xbd, 0x5a, 0xb0,
		0xcf,
	})
	copy(testBCacheTitle[11][:], []byte{ //嘰哩 ◎跨板式LOCAL新文章(隱板)
		0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xb8, 0xf3, 0xaa,
		0x4f, 0xa6, 0xa1, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0xb7, 0x73,
		0xa4, 0xe5, 0xb3, 0xb9, 0x28, 0xc1, 0xf4, 0xaa, 0x4f, 0x29,
	})

	copy(testBM4[:], []byte("SYSOP/CodingMan"))

	testBoardHeader4 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'S', 'Y', 'S', 'O', 'P'},
		Title: ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		BrdAttr: ptttype.BRD_POSTMASK,
		Gid:     2,
		BM:      testBM4,
	}

	copy(testBM13[:], []byte("CodingMan"))
	testBoardHeader13 = ptttype.BoardHeaderRaw{
		Brdname: ptttype.BoardID_t{'S', 'Y', 'S', 'O', 'P', '1', '3'},
		Title: ptttype.BoardTitle_t{
			0xbc, 0x54, 0xad, 0xf9, 0x20, 0xa1, 0xb7, 0xaf, 0xb8, 0xaa,
			0xf8, 0xa6, 0x6e, 0x21,
		},
		Gid: 5,
		BM:  testBM13,
	}

}

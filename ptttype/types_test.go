package ptttype

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/types"
	log "github.com/sirupsen/logrus"
)

func TestBoardTitle_t_RealTitle(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardTitle0 := &BoardTitle_t{ // CPBL ◎四海之內皆兄弟
		0x43, 0x50, 0x42, 0x4c, 0x20, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected0 := []byte{ // 四海之內皆兄弟
		0xa5, 0x7c, 0xae, 0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2,
		0xa5, 0x53, 0xa7, 0xcc,
	}

	boardTitle1 := &BoardTitle_t{ //*CPBL◎四海之內皆兄弟
		0x2a, 0x43, 0x50, 0x42, 0x4c, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected1 := []byte{ // 四海之內皆兄弟
		0xa5, 0x7c, 0xae, 0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2,
		0xa5, 0x53, 0xa7, 0xcc,
	}

	boardTitle2 := &BoardTitle_t{ // 里肌 ◎SYSOP
		0xa8, 0xbd, 0xa6, 0xd9, 0x20, 0xa1, 0xb7, 0x53, 0x59, 0x53,
		0x4f, 0x50,
	}
	expected2 := []byte("SYSOP")

	tests := []struct {
		name     string
		tr       *BoardTitle_t
		expected []byte
	}{
		// TODO: Add test cases.
		{
			tr:       boardTitle0,
			expected: expected0,
		},
		{
			tr:       boardTitle1,
			expected: expected1,
		},
		{
			tr:       boardTitle2,
			expected: expected2,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.tr.RealTitle(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BoardTitle_t.RealTitle() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestBoardTitle_t_BoardClass(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardTitle0 := &BoardTitle_t{ // CPBL ◎四海之內皆兄弟
		0x43, 0x50, 0x42, 0x4c, 0x20, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected0 := []byte("CPBL")

	boardTitle1 := &BoardTitle_t{ //*CPBL◎四海之內皆兄弟
		0x2a, 0x43, 0x50, 0x42, 0x4c, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected1 := []byte("*CPBL")

	boardTitle2 := &BoardTitle_t{ // 里肌 ◎SYSOP
		0xa8, 0xbd, 0xa6, 0xd9, 0x20, 0xa1, 0xb7, 0x53, 0x59, 0x53,
		0x4f, 0x50,
	}
	expected2 := []byte{0xa8, 0xbd, 0xa6, 0xd9}

	boardTitle3 := &BoardTitle_t{ // 里CP ◎SYSOP
		0xa8, 0xbd, 0x43, 0x50, 0x20, 0xa1, 0xb7, 0x53, 0x59, 0x53,
		0x4f, 0x50,
	}
	expected3 := []byte{0xa8, 0xbd, 0x43, 0x50}

	tests := []struct {
		name     string
		tr       *BoardTitle_t
		expected []byte
	}{
		// TODO: Add test cases.
		{
			tr:       boardTitle0,
			expected: expected0,
		},
		{
			tr:       boardTitle1,
			expected: expected1,
		},
		{
			tr:       boardTitle2,
			expected: expected2,
		},
		{
			tr:       boardTitle3,
			expected: expected3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.tr.BoardClass(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BoardTitle_t.BoardClass() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestBoardTitle_t_BoardType(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardTitle0 := &BoardTitle_t{ // CPBL ◎四海之內皆兄弟
		0x43, 0x50, 0x42, 0x4c, 0x20, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected0 := []byte{0xa1, 0xb7}

	boardTitle1 := &BoardTitle_t{ //*CPBL◎四海之內皆兄弟
		0x2a, 0x43, 0x50, 0x42, 0x4c, 0xa1, 0xb7, 0xa5, 0x7c, 0xae,
		0xfc, 0xa4, 0xa7, 0xa4, 0xba, 0xac, 0xd2, 0xa5, 0x53, 0xa7,
		0xcc,
	}
	expected1 := []byte{0xa1, 0xb7}

	boardTitle2 := &BoardTitle_t{ // 里肌 ◎SYSOP
		0xa8, 0xbd, 0xa6, 0xd9, 0x20, 0xa1, 0xb7, 0x53, 0x59, 0x53,
		0x4f, 0x50,
	}
	expected2 := []byte{0xa1, 0xb7}

	boardTitle3 := &BoardTitle_t{ // 里CP ◎SYSOP
		0xa8, 0xbd, 0x43, 0x50, 0x20, 0xa1, 0xb7, 0x53, 0x59, 0x53,
		0x4f, 0x50,
	}
	expected3 := []byte{0xa1, 0xb7}

	tests := []struct {
		name     string
		tr       *BoardTitle_t
		expected []byte
	}{
		// TODO: Add test cases.
		{
			tr:       boardTitle0,
			expected: expected0,
		},
		{
			tr:       boardTitle1,
			expected: expected1,
		},
		{
			tr:       boardTitle2,
			expected: expected2,
		},
		{
			tr:       boardTitle3,
			expected: expected3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.tr.BoardType(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BoardTitle_t.BoardType() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestBM_t_ToBMs(t *testing.T) {
	setupTest()
	defer teardownTest()

	bm0 := &BM_t{}
	copy(bm0[:], []byte("te1/te2/te3\x00te4"))
	expectedbm0 := &UserID_t{}
	copy(expectedbm0[:], []byte("te1"))
	expectedbm1 := &UserID_t{}
	copy(expectedbm1[:], []byte("te2"))
	expectedbm2 := &UserID_t{}
	copy(expectedbm2[:], []byte("te3"))
	expected0 := []*UserID_t{expectedbm0, expectedbm1, expectedbm2}
	tests := []struct {
		name     string
		bm       *BM_t
		expected []*UserID_t
	}{
		// TODO: Add test cases.
		{
			bm:       bm0,
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			got := tt.bm.ToBMs()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("BM_t.ToBMs() = %v, want %v", got, tt.expected)
			}
			for idx, each := range got {
				if idx >= len(tt.expected) {
					t.Errorf("BM_t: (%v/%v) %v", idx, len(got), each)
					continue
				}
				if !reflect.DeepEqual(each, tt.expected[idx]) {
					t.Errorf("BM_t: (%v/%v) %v want: %v", idx, len(got), each, tt.expected[idx])
				}
			}
		})
	}
	wg.Wait()
}

func TestFilename_t_CreateTime(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0 := &Filename_t{}
	copy(f0[:], []byte("M.1234567890.A.123"))

	expected0 := types.Time4(1234567890)

	tests := []struct {
		name     string
		f        *Filename_t
		expected types.Time4
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			f:        f0,
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, err := tt.f.CreateTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("Filename_t.CreateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Filename_t.CreateTime() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestFilename_t_Postfix(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0 := &Filename_t{}
	copy(f0[:], []byte("M.1234567890.A.123"))

	expected0 := []byte("123")

	tests := []struct {
		name     string
		f        *Filename_t
		expected []byte
	}{
		// TODO: Add test cases.
		{
			f:        f0,
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.f.Postfix(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Filename_t.Postfix() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestFilename_t_ToAidu(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0 := &Filename_t{}
	copy(f0[:], []byte("M.1234567890.A.123"))

	expected0 := Aidu(0x499602d2123)

	f1 := &Filename_t{}
	copy(f1[:], []byte("M.1607937174.A.081"))

	expected1 := Aidu(0x5fd72c96081)

	f2 := &Filename_t{}
	copy(f2[:], []byte("M.1607202239.A.30D"))

	expected2 := Aidu(0x5fcbf5bf30d)

	f3 := &Filename_t{}
	copy(f3[:], []byte("M.1607202240.A.30D"))

	expected3 := Aidu(0x5fcbf5c030d)

	tests := []struct {
		name     string
		f        *Filename_t
		expected Aidu
	}{
		// TODO: Add test cases.
		{
			f:        f0,
			expected: expected0,
		},
		{
			f:        f1,
			expected: expected1,
		},
		{
			f:        f2,
			expected: expected2,
		},
		{
			f:        f3,
			expected: expected3,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.f.ToAidu(); got != tt.expected {
				t.Errorf("Filename_t.ToAidu() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestFilename_t_Type(t *testing.T) {
	setupTest()
	defer teardownTest()

	f0 := &Filename_t{}
	copy(f0[:], []byte("M.1234567890.A.123"))

	f1 := &Filename_t{}
	copy(f1[:], []byte("G.1234567890.A.123"))

	tests := []struct {
		name     string
		f        *Filename_t
		expected RecordType
	}{
		// TODO: Add test cases.
		{
			f:        f0,
			expected: RECORD_TYPE_M,
		},
		{
			f:        f1,
			expected: RECORD_TYPE_G,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.f.Type(); got != tt.expected {
				t.Errorf("Filename_t.Type() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestAidu_Type(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := Aidu(0x0499602d2123)
	expected0 := RECORD_TYPE_M
	a1 := Aidu(0x1499602d2123)
	expected1 := RECORD_TYPE_G

	tests := []struct {
		name     string
		a        Aidu
		expected RecordType
	}{
		// TODO: Add test cases.
		{
			a:        a0,
			expected: expected0,
		},
		{
			a:        a1,
			expected: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.a.Type(); got != tt.expected {
				t.Errorf("Aidu.Type() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestAidu_Time(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := Aidu(0x0499602d2123)
	expected0 := types.Time4(1234567890)

	tests := []struct {
		name     string
		a        Aidu
		expected types.Time4
	}{
		// TODO: Add test cases.
		{
			a:        a0,
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.a.Time(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Aidu.Time() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestAidu_ToFN(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := Aidu(0x0499602d2123)
	expected0 := &Filename_t{}
	copy(expected0[:], []byte("M.1234567890.A.123"))

	a1 := Aidu(0x5fcbf5bf30d)
	expected1 := &Filename_t{}
	copy(expected1[:], []byte("M.1607202239.A.30D"))

	tests := []struct {
		name     string
		a        Aidu
		expected *Filename_t
	}{
		// TODO: Add test cases.
		{
			a:        a0,
			expected: expected0,
		},
		{
			a:        a1,
			expected: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.a.ToFN(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Aidu.ToFN() = %v, want %v", got, tt.expected)
				for idx, each := range got {
					if each != tt.expected[idx] {
						t.Errorf("Aidu: (%v/%v) %v want: %v", idx, len(got), each, tt.expected[idx])
					}
				}
			}
		})
	}
	wg.Wait()
}

func TestAidu_ToAidc(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := Aidu(0x0499602d2123)
	expected0 := &Aidc{}
	copy(expected0[:], []byte("19bWBI4Z"))

	f1 := &Filename_t{}
	copy(f1[:], []byte("M.1607937174.A.081"))
	a1 := f1.ToAidu()
	log.Infof("f1: %v a1: %x", f1, a1)
	expected1 := &Aidc{}
	copy(expected1[:], []byte("1VrooM21"))

	a2 := Aidu(0x05fcbf5bf30d)
	expected2 := &Aidc{}
	copy(expected2[:], []byte("1Vo_M_CD"))

	a3 := Aidu(0xffffffffffff)
	expected3 := &Aidc{}
	copy(expected3[:], []byte("________"))

	a4 := Aidu(0x010000000000) // 1978-07-04 21:24:16 UTC
	expected4 := &Aidc{}
	copy(expected4[:], []byte("0G000000"))

	a5 := Aidu(0x5fcbf5c030d)
	expected5 := &Aidc{}
	copy(expected5[:], []byte("1Vo_N0CD"))

	f6 := &Filename_t{}
	copy(f6[:], []byte("M.1608386280.A.BC9"))
	a6 := f6.ToAidu()
	expected6 := &Aidc{}
	copy(expected6[:], []byte("1VtWRel9"))

	f7 := &Filename_t{}
	copy(f7[:], []byte("M.1607937176.A.081"))
	a7 := f7.ToAidu()
	log.Infof("f7: %v a7: %x", f7, a7)
	expected7 := &Aidc{}
	copy(expected7[:], []byte("1VrooO21"))

	f8 := &Filename_t{}
	copy(f8[:], []byte("M.1234567892.A.123"))
	a8 := f8.ToAidu()
	log.Infof("f8: %v a8: %x", f8, a8)
	expected8 := &Aidc{}
	copy(expected8[:], []byte("19bWBK4Z"))

	tests := []struct {
		name     string
		a        Aidu
		expected *Aidc
	}{
		// TODO: Add test cases.
		{
			a:        a0,
			expected: expected0,
		},
		{
			a:        a1,
			expected: expected1,
		},
		{
			a:        a2,
			expected: expected2,
		},
		{
			a:        a3,
			expected: expected3,
		},
		{
			a:        a4,
			expected: expected4,
		},
		{
			a:        a5,
			expected: expected5,
		},
		{
			a:        a6,
			expected: expected6,
		},
		{
			a:        a7,
			expected: expected7,
		},
		{
			a:        a8,
			expected: expected8,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.a.ToAidc(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Aidu.ToAidc() = %v, want %v", string(got[:]), string(tt.expected[:]))
			}
		})
	}
	wg.Wait()
}

func TestAidc_ToAidu(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := &Aidc{}
	copy(a0[:], []byte("19bWBI4Z"))
	expected0 := Aidu(0x0499602d2123)

	a1 := &Aidc{}
	copy(a1[:], []byte("1VrooM21"))
	expected1 := Aidu(0x5fd72c96081)

	tests := []struct {
		name         string
		a            *Aidc
		expectedAidu Aidu
	}{
		// TODO: Add test cases.
		{
			a:            a0,
			expectedAidu: expected0,
		},
		{
			a:            a1,
			expectedAidu: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if gotAidu := tt.a.ToAidu(); gotAidu != tt.expectedAidu {
				t.Errorf("Aidc.ToAidu() = %x, want %x", gotAidu, tt.expectedAidu)
			}
		})
	}
	wg.Wait()
}

func TestOwner_t_ToUserID(t *testing.T) {
	setupTest()
	defer teardownTest()

	o := &Owner_t{}
	copy(o[:], []byte("test1."))

	u := &UserID_t{}
	copy(u[:], []byte("test1"))
	tests := []struct {
		name     string
		o        *Owner_t
		expected *UserID_t
	}{
		// TODO: Add test cases.
		{
			o:        o,
			expected: u,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.o.ToUserID(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Owner_t.ToUserID() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestUserID_t_IsValid(t *testing.T) {
	setupTest()
	defer teardownTest()

	userID0 := UserID_t{}

	userID1 := UserID_t{}
	copy(userID1[:], []byte("S"))

	userID2 := UserID_t{}
	copy(userID2[:], []byte("SYSOP"))

	userID3 := UserID_t{}
	copy(userID3[:], []byte("S1234567891234"))

	userID4 := UserID_t{}
	copy(userID4[:], []byte("SYSOP,-"))

	userID5 := UserID_t{}
	copy(userID5[:], []byte("SYSOP1"))

	userID6 := UserID_t{}
	copy(userID6[:], []byte("1SYSOP"))

	userID7 := UserID_t{}
	copy(userID7[:], []byte("S1"))

	userID8 := UserID_t{}
	copy(userID8[:], []byte("Ss"))

	tests := []struct {
		name     string
		u        *UserID_t
		expected bool
	}{
		// TODO: Add test cases.
		{
			name:     "nil",
			u:        nil,
			expected: false,
		},
		{
			name:     "",
			u:        &userID0,
			expected: false,
		},
		{
			name:     "S",
			u:        &userID1,
			expected: false,
		},
		{
			name:     "SYSOP",
			u:        &userID2,
			expected: true,
		},
		{
			name:     "too long",
			u:        &userID3,
			expected: false,
		},
		{
			name:     "not alnum",
			u:        &userID4,
			expected: false,
		},
		{
			name:     "SYSOP1",
			u:        &userID5,
			expected: true,
		},
		{
			name:     "1SYSOP (not alpha in 0st char)",
			u:        &userID6,
			expected: false,
		},
		{
			name:     "S1 (alnum)",
			u:        &userID7,
			expected: true,
		},
		{
			name:     "Ss (all alpha)",
			u:        &userID7,
			expected: true,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.u.IsValid(); got != tt.expected {
				t.Errorf("UserID_t.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestUserID_t_IsGuest(t *testing.T) {
	userID0Str := "guest"
	userID0 := &UserID_t{}
	copy(userID0[:], []byte(userID0Str))

	userID1Str := "guest\x00123"
	userID1 := &UserID_t{}
	copy(userID1[:], []byte(userID1Str))

	userID2Str := "st123123"
	userID2 := &UserID_t{}
	copy(userID2[:], []byte(userID2Str))

	tests := []struct {
		name     string
		u        *UserID_t
		expected bool
	}{
		// TODO: Add test cases.
		{
			u:        userID0,
			expected: true,
		},
		{
			u:        userID1,
			expected: true,
		},
		{
			u:        userID2,
			expected: false,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := tt.u.IsGuest(); got != tt.expected {
				t.Errorf("UserID_t.IsGuest() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}

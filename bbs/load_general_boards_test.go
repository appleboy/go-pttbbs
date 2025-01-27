package bbs

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestLoadGeneralBoards(t *testing.T) {
	setupTest()
	defer teardownTest()
	// setupTest moves in for-loop
	// teardownTest moves in for-loop
	type args struct {
		uuserID     UUserID
		startIdxStr string
		nBoards     int
		title       []byte
		keyword     []byte
		isAsc       bool
		sortBy      ptttype.BSortBy
	}
	tests := []struct {
		name               string
		args               args
		expectedSummaries  []*BoardSummary
		expectedNextIdxStr string
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{uuserID: "SYSOP3", startIdxStr: "", nBoards: 4, isAsc: true, sortBy: ptttype.BSORT_BY_NAME},
			expectedSummaries:  []*BoardSummary{testBoardSummary6, testBoardSummary11, testBoardSummary8, testBoardSummary9},
			expectedNextIdxStr: "SYSOP",
		},
		{
			args:               args{uuserID: "SYSOP3", startIdxStr: "", nBoards: 4, isAsc: true, sortBy: ptttype.BSORT_BY_CLASS},
			expectedSummaries:  []*BoardSummary{testBoardSummary6, testBoardSummary11, testBoardSummary8, testBoardSummary9},
			expectedNextIdxStr: "vFSt-Q@SYSOP",
		},
		{
			args:               args{uuserID: "SYSOP3", startIdxStr: "vFSt-Q@Record", nBoards: 4, isAsc: true, sortBy: ptttype.BSORT_BY_CLASS},
			expectedSummaries:  []*BoardSummary{testBoardSummary9, testBoardSummary1, testBoardSummary10},
			expectedNextIdxStr: "",
		},
		{
			args:               args{uuserID: "SYSOP3", startIdxStr: "Record", nBoards: 4, isAsc: true, sortBy: ptttype.BSORT_BY_NAME},
			expectedSummaries:  []*BoardSummary{testBoardSummary9, testBoardSummary1, testBoardSummary10},
			expectedNextIdxStr: "",
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			gotSummaries, gotNextIdx, err := LoadGeneralBoards(tt.args.uuserID, tt.args.startIdxStr, tt.args.nBoards, tt.args.title, tt.args.keyword, tt.args.isAsc, tt.args.sortBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "summaries", gotSummaries, tt.expectedSummaries)

			if gotNextIdx != tt.expectedNextIdxStr {
				t.Errorf("LoadGeneralBoards() gotNextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdxStr)
			}
		})
	}
	wg.Wait()
}

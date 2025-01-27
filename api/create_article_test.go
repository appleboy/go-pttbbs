package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptt"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestCreateArticle(t *testing.T) {
	setupTest(t.Name())
	defer teardownTest(t.Name())

	_ = ptt.SetupNewUser(testNewPostUserRaw1)

	class0 := []byte("test")
	title0 := []byte("this is a test")
	fullTitle0 := []byte("[test] this is a test")
	content0 := [][]byte{[]byte("test1"), []byte("test2")}
	ip0 := "127.0.0.1"

	params0 := &CreateArticleParams{
		PostType: class0,
		Title:    title0,
		Content:  content0,
	}

	path0 := &CreateArticlePath{
		BBoardID: "10_WhoAmI",
	}

	expected0 := CreateArticleResult(&bbs.ArticleSummary{
		BBoardID:  "10_WhoAmI",
		Owner:     "A1",
		FullTitle: fullTitle0,
		Class:     class0,
		Idx:       "",
		RealTitle: []byte("this is a test"),
	})

	type args struct {
		remoteAddr string
		uuser      bbs.UUserID
		params     interface{}
		path       interface{}
	}
	tests := []struct {
		name           string
		args           args
		expectedResult CreateArticleResult
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				remoteAddr: ip0,
				uuser:      "A1",
				params:     params0,
				path:       path0,
			},
			expectedResult: expected0,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := CreateArticle(tt.args.remoteAddr, tt.args.uuser, tt.args.params, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := gotResult.(CreateArticleResult)
			result.ArticleID = ""
			result.Filename = ""
			result.CreateTime = 0
			result.MTime = 0
			result.Idx = ""

			testutil.TDeepEqual(t, "result", result, tt.expectedResult)
		})
		wg.Wait()
	}
}

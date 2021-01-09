package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

const SET_ID_EMAIL_R = "/users/:uid/setidemail"

type SetIDEmailParams struct {
	IsSet bool   `json:"is_set"`
	Jwt   string `json:"email_token"`
}

type SetIDEmailPath struct {
	UserID bbs.UUserID `uri:"uid" binding:"required"`
}

type SetIDEmailResult struct {
	UserID     bbs.UUserID   `json:"user_id"`
	UserLevel2 ptttype.PERM2 `json:"user_level2"`
}

func SetIDEmailWrapper(c *gin.Context) {
	params := &SetIDEmailParams{}
	path := &SetIDEmailPath{}
	LoginRequiredPathJSON(SetIDEmail, params, path, c)
}

func SetIDEmail(remoteAddr string, uuserID bbs.UUserID, params interface{}, path interface{}) (result interface{}, err error) {

	theParams, ok := params.(*SetIDEmailParams)
	if !ok {
		return nil, ErrInvalidParams
	}

	thePath, ok := path.(*SetIDEmailPath)
	if !ok {
		return nil, ErrInvalidPath
	}

	isValid, _ := userIsValidEmailUser(uuserID, thePath.UserID, theParams.Jwt, true)
	if !isValid {
		return nil, ErrInvalidUser
	}

	userLevel2, err := bbs.SetIDEmail(thePath.UserID, theParams.IsSet)
	if err != nil {
		return nil, err
	}

	result = &SetIDEmailResult{
		UserID:     thePath.UserID,
		UserLevel2: userLevel2,
	}

	return result, nil
}

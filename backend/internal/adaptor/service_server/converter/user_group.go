package converter

import (
	"github.com/ritscc/Snack/internal/domain/model"
	pbUserGroup "github.com/ritscc/Snack/pb/user_group/v1"
)

func UserGroupTopbUserGroup(userGroup *model.UserGroup) *pbUserGroup.UserGroup {
	return &pbUserGroup.UserGroup{
		GroupId:    userGroup.GroupID,
		Name:        userGroup.Name,
		Description: userGroup.Description,
		Type:        userGroup.Type,
		Icon:        userGroup.Icon,
		Role:        userGroup.Role,
		Member:      converter.MemberTopbMember(userGroup.Member),
	}
}

func 
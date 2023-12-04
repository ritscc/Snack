package converter

import (
	"github.com/ritscc/Snack/internal/domain/model"
	pbUserGroup "github.com/ritscc/Snack/pb/user_group/v1"
)

func UserGroupTopbUserGroup(userGroup *model.UserGroup) *pbUserGroup.UserGroup {
	var members []*pbUserGroup.Member
	for _, member := range userGroup.Member {
		members = append(members, &pbUserGroup.Member{
			UserId: member.UserId,
			Grant:  member.Grant,
		})
	}

	return &pbUserGroup.UserGroup{
		GroupId:     userGroup.GroupID,
		Name:        userGroup.Name,
		Description: userGroup.Description,
		Type:        userGroup.Type,
		Icon:        userGroup.Icon,
		Role:        userGroup.Role,
		Member:      members,
	}
}

func PbUserGroupToUserGroup(userGroup *pbUserGroup.UserGroup) *model.UserGroup {
	var members []*model.Member
	for _, member := range userGroup.Member {
		members = append(members, &model.Member{
			UserId: member.UserId,
			Grant:  member.Grant,
		})
	}

	return &model.UserGroup{
		GroupID:     userGroup.GroupId,
		Name:        userGroup.Name,
		Description: userGroup.Description,
		Type:        userGroup.Type,
		Icon:        userGroup.Icon,
		Role:        userGroup.Role,
		Member:      members,
	}
}

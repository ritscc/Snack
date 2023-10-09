package dto

type CreateUserRequest struct {
	Username string
	Email    string
	Password string
}

type CreateUserGroupRequest struct {
	Group_id    int64  // UserGroupID
	Name        string // UserGroupName
	Description string
	Type        string // SystemAdmin or Community
	Icon        string // UserGroupIcon
	Role        string
	Member      []Member // Users who belong to UserGroup
}

type Member struct {
	group_id int64
	user_id  int64
	grant    string
}

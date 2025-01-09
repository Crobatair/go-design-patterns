package main

type Permission struct {
	Name string
}

type PermissionFlyweightFactory struct {
	permissions map[string]*Permission
}

func NewPermissionFlyweightFactory(systemPermissions []string) *PermissionFlyweightFactory {
	factory := PermissionFlyweightFactory{
		permissions: make(map[string]*Permission),
	}

	for _, permission := range systemPermissions {
		factory.permissions[permission] = &Permission{permission}
	}

	return &factory
}

// GetPermission returns a permission object if it exists in the factory, otherwise it returns nil
func (p *PermissionFlyweightFactory) GetPermission(permissionName string) *Permission {
	if _, ok := p.permissions[permissionName]; ok {
		return p.permissions[permissionName]
	}

	return nil
}

type User struct {
	Name        string
	Permissions []*Permission
}

type UserDataset struct {
	Name        string
	Permissions []string
}

func NewUser(userObject UserDataset, factory *PermissionFlyweightFactory) *User {
	user := User{Name: userObject.Name}
	for _, permissionName := range userObject.Permissions {
		user.Permissions = append(user.Permissions, factory.GetPermission(permissionName))
	}
	return &user
}

func main() {
	userDataset := []UserDataset{
		{"User1", []string{"can_claim_reward", "can_view_dashboard"}},
		{"User2", []string{"can_view_dashboard", "can_view_profile"}},
		{"User3", []string{"can_claim_reward", "can_view_profile"}},
	}
	systemPermissions := []string{"can_claim_reward", "can_view_dashboard", "can_view_profile"}
	permissionFactory := NewPermissionFlyweightFactory(systemPermissions)

	users := make([]*User, 0, len(userDataset))
	for _, user := range userDataset {
		users = append(users, NewUser(user, permissionFactory))
	}

	for _, user := range users {
		println(user.Name)
		for _, permission := range user.Permissions {
			println(permission.Name)
		}
	}
}

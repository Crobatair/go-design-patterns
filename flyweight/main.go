package main

import "fmt"

type Permission struct {
	Name string
}

type PermissionFlyweightFactory struct {
	permissions map[int8]*Permission
}

func NewPermissionFlyweightFactory(systemPermissions []string) *PermissionFlyweightFactory {
	factory := PermissionFlyweightFactory{
		permissions: make(map[int8]*Permission),
	}

	for i, permission := range systemPermissions {
		factory.permissions[int8(i)] = &Permission{permission}
	}

	return &factory
}

// GetPermission returns a permission object if it exists in the factory, otherwise it returns nil
func (p *PermissionFlyweightFactory) GetPermission(permissionName string) int8 {
	for idx, permission := range p.permissions {
		if permission.Name == permissionName {
			return idx
		}
	}

	return -1
}

type User struct {
	Name        string
	Permissions []int8
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

func (u *User) PermissionsList(p *PermissionFlyweightFactory) []string {
	permissions := make([]string, len(u.Permissions))
	for _, permission := range u.Permissions {
		permissions = append(permissions, p.permissions[permission].Name)
	}
	return permissions
}

type UserNotOptimized struct {
	Name        string
	Permissions []*Permission
}

func NewUserNotUsingFlyweight(userObject UserDataset) *UserNotOptimized {
	user := UserNotOptimized{Name: userObject.Name}
	for _, permissionName := range userObject.Permissions {
		user.Permissions = append(user.Permissions, &Permission{permissionName})
	}
	return &user
}

func main() {
	userDataset := []UserDataset{
		{"User1", []string{"can_claim_reward", "can_view_dashboard"}},
		{"User2", []string{"can_view_dashboard", "can_view_profile"}},
		{"User3", []string{"can_claim_reward", "can_view_profile"}},
		{"User4", []string{"can_claim_reward", "can_view_dashboard", "can_view_profile"}},
	}
	systemPermissions := []string{"can_claim_reward", "can_view_dashboard", "can_view_profile"}
	permissionFactory := NewPermissionFlyweightFactory(systemPermissions)

	users := make([]*User, 0, len(userDataset))
	for _, user := range userDataset {
		users = append(users, NewUser(user, permissionFactory))
	}

	fmt.Println("Done")
	fmt.Println("Total Memory usage, not using flyweight: ")
	memoryUsage := 0
	usersNotUsingFlyweight := make([]*UserNotOptimized, 0, len(userDataset))
	for idx, user := range userDataset {
		usersNotUsingFlyweight = append(usersNotUsingFlyweight, NewUserNotUsingFlyweight(user))
		for _, permission := range usersNotUsingFlyweight[idx].Permissions {
			memoryUsage += len([]byte(permission.Name))
		}
		memoryUsage += len(user.Name)
	}
	fmt.Printf("Memory usage no Flyweight: %d\n", memoryUsage)

	fmt.Println("Total Memory usage, using flyweight: ")
	memoryUsage = 0
	for _, permission := range permissionFactory.permissions {
		memoryUsage += len([]byte(permission.Name))
	}

	for _, user := range users {
		memoryUsage += len(user.Permissions)
		memoryUsage += len(user.Name)
	}
	fmt.Printf("Memory using Flyweight: %d\n", memoryUsage)
}

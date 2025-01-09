## Flyweight Pattern

The Flyweight Pattern is a structural design pattern that is used to reduce the memory footprint of objects. It is used to minimize the memory usage or computational expenses by sharing as much as possible with related objects. It is used to manage the state of objects by sharing as much as possible.

Example:
- We manage a large number of Permissions for a large number of users. We can use the Flyweight pattern to share the same Permission object for multiple users.
    We could have a Permission object, for each unique permission in the system, let's say 'can_claim_reward'
    A large set of users could have the same Permission object, `can_claim_reward` object.
    Instead of creating a new Permission object for each user, we can share the same Permission object for all users that have the 'can_claim_reward' permission.
  - This way we can reduce the memory footprint of the system.
  - We can also reduce the computational expenses of the system by sharing the same Permission object for multiple users.

### When to use the Flyweight Pattern
- When you need to create a large number of similar objects.
- When you need to reduce the memory footprint of the system.
- When you need to reduce the computational expenses of the system.

### Implementation
```go
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


```
package proxy

import "github.com/danilomarques1/design_patterns/proxy/service"

type UserServiceProxy struct {
	*service.UserService
	cachedUsers map[string]*service.User
}

func NewUserServiceProxy() *UserServiceProxy {
	p := &UserServiceProxy{UserService: service.NewUserService()}
	p.cachedUsers = make(map[string]*service.User)
	return p
}
func (p *UserServiceProxy) Add(user *service.User) {
	// cleaning cache
	p.cachedUsers = make(map[string]*service.User)
	p.UserService.Add(user)
}

func (p *UserServiceProxy) FindById(id string) (*service.User, error) {
	user, ok := p.cachedUsers[id]
	if !ok {
		var err error
		user, err = p.UserService.FindById(id)
		if err != nil {
			return nil, err
		}
		p.cachedUsers[id] = user
	}
	return user, nil
}

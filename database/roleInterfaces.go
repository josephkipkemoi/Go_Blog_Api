package database

func (r *Roles) Create() (*Roles, error) {
	err := DB.Create(r).Error
	if err != nil {
		return &Roles{}, err
	}
	return r, nil
}

func (r *Roles) Index() ([]Roles, error) {
	var roles []Roles
	res := DB.Order("id asc").Find(&roles)

	if res.Error != nil {
		return roles, res.Error
	}
	return roles, nil
}

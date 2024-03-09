package database

func (r *Roles) Create() (*Roles, error) {
	err := DB.Create(r).Error
	if err != nil {
		return &Roles{}, err
	}
	return r, nil
}

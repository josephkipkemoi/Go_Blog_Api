package database

import "fmt"

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

func (r *Roles) Show(id int) (*Roles, error) {
	res := DB.Find(&r, id)

	if res.Error != nil || res.RowsAffected == 0 {
		return &Roles{}, fmt.Errorf("%s", "record not found")
	}

	return r, nil
}

func (r *Roles) Delete(id int) error {
	res := DB.Delete(&r, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot delete record")
	}

	return nil
}

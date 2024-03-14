package database

import "fmt"

func (f *Favourite) Create() error {
	err := DB.Create(f).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Favourite) Index(id int) ([]Favourite, error) {
	var favourites []Favourite
	res := DB.Order("id asc").Where("user_id", id).Find(&favourites)

	if res.Error != nil {
		return favourites, res.Error
	}
	return favourites, nil
}

func (f *Favourite) Delete(user_id, id int) error {
	res := DB.Where("user_id", user_id).Delete(&f, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot delete record")
	}

	return nil
}

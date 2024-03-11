package database

import "fmt"

func (c *Contact) Create() error {
	err := DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Contact) Index() ([]Contact, error) {
	var contacts []Contact
	res := DB.Order("id desc").Take(20).Find(&contacts)

	if res.Error != nil {
		return contacts, res.Error
	}

	return contacts, nil
}

func (c *Contact) Show(id int) (*Contact, error) {
	res := DB.Find(&c, id)

	if res.Error != nil || res.RowsAffected == 0 {
		return &Contact{}, fmt.Errorf("%s", "record not found")
	}

	return c, nil
}

func (c *Contact) Delete(id int) error {
	res := DB.Delete(&c, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot delete record")
	}

	return nil
}

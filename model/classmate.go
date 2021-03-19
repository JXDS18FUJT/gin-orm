package model

import "orm.com/m/v2/dao"

type Classmate struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	Work      string `json:"work"`
	Age       int    `json:"age"`
	Education string `json:"education"`
	CreatedAt int64  `gorm:"autoUpdateTime:nano"`
	UpdatedAt int64  `gorm:"autoCreateTime"`
}

func CreateClassmate(classmate *Classmate) (err error) {
	err = dao.DB.Create(&classmate).Error
	return
}

func UpdateClassmate(classmate *Classmate) (err error) {
	err = dao.DB.Save(classmate).Error
	return err

}
func DeleteClassmate(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Classmate{}).Error
	return err
}
func GetClassmate(id string) (classmate *Classmate, err error) {
	classmate = new(Classmate)
	if err = dao.DB.Debug().Where("id=?", id).First(classmate).Error; err != nil {
		return nil, err
	}
	return
}
func GetAllClassmate() (classmateList []*Classmate, err error) {
	if err = dao.DB.Find(&classmateList).Error; err != nil {
		return nil, err
	}
	return
}

package repository

type gormRepository struct {
}

func NewRepository() Repository {
	return &gormRepository{}
}

type Repository interface {
	Add(uow *UnitOfWork, entity interface{}) error
	Get(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, entity interface{}) error
	ScopedCountRows(uow *UnitOfWork, entity interface{}) (int, error)
	UnScopedCountRows(uow *UnitOfWork, entity interface{}) (int, error)
	Delete(uow *UnitOfWork, entity interface{}) error
	HardDelete(uow *UnitOfWork, entity interface{}) error
}

func (*gormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	db := uow.DB

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) Get(uow *UnitOfWork, out interface{}) error {

	db := uow.DB

	if err := db.Debug().Preload("Orders").First(out).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) Update(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Model(entity).Update(entity).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) ScopedCountRows(uow *UnitOfWork, entity interface{}) (int, error) {

	db := uow.DB
	var count int

	if err := db.Debug().Model(entity).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func (*gormRepository) UnScopedCountRows(uow *UnitOfWork, entity interface{}) (int, error) {

	db := uow.DB
	var count int

	if err := db.Debug().Model(entity).Unscoped().Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func (*gormRepository) Delete(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	if err := db.Debug().Delete(entity).Error; err != nil {
		return err
	}

	return nil
}

func (*gormRepository) HardDelete(uow *UnitOfWork, entity interface{}) error {

	db := uow.DB
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}
	if err := db.Debug().Unscoped().Delete(entity).Error; err != nil {
		return err
	}

	return nil
}

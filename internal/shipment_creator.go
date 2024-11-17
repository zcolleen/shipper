package internal

import (
	"fmt"
	"log"
)

//go:generate mockgen -mock_names UserStore=MockUserStoreGomock -typed -package internal  -destination zzz_user_store_mock_gomock.go github.com/zcolleen/shipper/internal UserStore
//go:generate minimock -g -i .UserStore -o ./zzz_user_store_mock_minimock.go
type UserStore interface {
	GetUser(id int64) (User, error)
}

type User struct {
	ID       int64
	Name     string
	Age      int64
	PhoneNum string
	Country  string
	Address  string
}

type Shipment struct {
	ID              string `json:"ID"`
	DeliveryAddress string `json:"DeliveryAddress"`
}

type ShipmentCreator struct {
	userStore     UserStore
	withLogging   bool
	historyLogger historyLogger
}

func (u *ShipmentCreator) getUserAddress(userID int64) (string, error) {
	user, err := u.userStore.GetUser(userID)
	if err != nil {
		return "", err
	}

	return user.Address, nil
}

func (u *ShipmentCreator) validate(userID int64) error {
	user, err := u.userStore.GetUser(userID)
	if err != nil {
		return err
	}

	if user.Country != "Russia" {
		return fmt.Errorf("address is not available for delivery")
	}
	return nil
}

func (u *ShipmentCreator) CreateShipment(userID int64) error {
	if err := u.validate(userID); err != nil {
		return err
	}

	addr, err := u.getUserAddress(userID)
	if err != nil {
		return err
	}

	s := &Shipment{
		DeliveryAddress: addr,
	}
	if !u.withLogging {
		return nil
	}

	go func() {
		err = u.historyLogger.Log(s)
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

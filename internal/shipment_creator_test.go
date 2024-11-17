package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateShipment_Mockery_failIncorrectType(t *testing.T) {
	userStore := NewMockUserStore(t)
	userStore.EXPECT().GetUser(1).Return(User{
		Country: "Russia",
	}, nil)

	creator := &ShipmentCreator{
		userStore: userStore,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestCreateShipment_Mockery_success(t *testing.T) {
	userStore := NewMockUserStore(t)
	userStore.EXPECT().GetUser(int64(1)).Return(User{
		Country: "Russia",
	}, nil)

	creator := &ShipmentCreator{
		userStore: userStore,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestCreateShipment_Minimock_success(t *testing.T) {
	user := User{
		Country: "Russia",
	}
	userStore := NewUserStoreMock(t).GetUserMock.Expect(1).Return(user, nil)
	creator := &ShipmentCreator{
		userStore: userStore,
	}

	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestCreateShipment_Gomock_failIncorrectType(t *testing.T) {
	ctl := gomock.NewController(t)
	userStore := NewMockUserStoreGomock(ctl)
	userStore.EXPECT().GetUser(1).Return(User{
		Country: "Russia",
	}, nil).
		AnyTimes()

	creator := &ShipmentCreator{
		userStore: userStore,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestCreateShipment_Gomock_failTooManyCalls(t *testing.T) {
	ctl := gomock.NewController(t)
	userStore := NewMockUserStoreGomock(ctl)
	userStore.EXPECT().GetUser(int64(1)).Return(User{
		Country: "Russia",
	}, nil)

	creator := &ShipmentCreator{
		userStore: userStore,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestCreateShipment_Gomock_success(t *testing.T) {
	ctl := gomock.NewController(t)
	userStore := NewMockUserStoreGomock(ctl)
	userStore.EXPECT().GetUser(int64(1)).Return(User{
		Country: "Russia",
	}, nil).Times(2)

	creator := &ShipmentCreator{
		userStore: userStore,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

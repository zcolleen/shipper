package internal

import (
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestShipmentCreator_CreateShipment_AsyncMockery(t *testing.T) {
	resultCh := make(chan struct{})

	userStore := NewMockUserStore(t)
	userStore.EXPECT().GetUser(int64(1)).Return(User{
		Address: "some_street",
		Country: "Russia",
	}, nil)

	s := &Shipment{
		DeliveryAddress: "some_street",
	}

	logger := NewMockhistoryLogger(t)
	logger.EXPECT().Log(s).RunAndReturn(func(*Shipment) error {
		resultCh <- struct{}{}
		return nil
	})

	creator := &ShipmentCreator{
		userStore:     userStore,
		withLogging:   true,
		historyLogger: logger,
	}
	err := creator.CreateShipment(1)
	assert.NoError(t, err)

	select {
	case <-resultCh:
	case <-time.After(time.Second):
	}
}

func wait(ctl *gomock.Controller) {
	timer := time.NewTimer(time.Second)
	for !ctl.Satisfied() {
		select {
		case <-time.After(time.Millisecond):
		case <-timer.C:
			return
		}
	}
}

func TestShipmentCreator_CreateShipment_AsyncGomock(t *testing.T) {
	ctl := gomock.NewController(t)
	defer wait(ctl)

	s := &Shipment{
		DeliveryAddress: "some_street",
	}

	logger := NewMockHistoryLoggerGomock(ctl)
	logger.EXPECT().Log(s).Return(nil)

	userStore := NewMockUserStoreGomock(ctl)
	userStore.EXPECT().GetUser(int64(1)).Return(User{
		Country: "Russia",
		Address: "some_street",
	}, nil).Times(2)

	creator := &ShipmentCreator{
		userStore:     userStore,
		withLogging:   true,
		historyLogger: logger,
	}

	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

func TestShipmentCreator_CreateShipment_AsyncMinimock(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Wait(time.Second)

	userStore := NewUserStoreMock(t).GetUserMock.Expect(1).Return(User{
		Address: "some_street",
		Country: "Russia",
	}, nil)

	s := &Shipment{
		DeliveryAddress: "some_street",
	}

	logger := NewHistoryLoggerMock(mc).LogMock.Expect(s).Return(nil)

	creator := &ShipmentCreator{
		userStore:     userStore,
		withLogging:   true,
		historyLogger: logger,
	}

	err := creator.CreateShipment(1)
	assert.NoError(t, err)
}

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipmentCreator_CreateShipment_tableSuccess(t *testing.T) {
	testCases := []struct {
		name        string
		userStoreFn func(t *testing.T) UserStore
		userID      int64
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name:    "first case",
			userID:  1,
			wantErr: assert.NoError,
			userStoreFn: func(t *testing.T) UserStore {
				return NewUserStoreMock(t).
					GetUserMock.
					Expect(1).
					Return(User{ID: 1, Country: "Russia"}, nil)
			},
		},
		{
			name:    "second case",
			userID:  2,
			wantErr: assert.NoError,
			userStoreFn: func(t *testing.T) UserStore {
				return NewUserStoreMock(t).
					GetUserMock.
					Expect(2).
					Return(User{ID: 2, Country: "Russia"}, nil)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := &ShipmentCreator{
				userStore: tc.userStoreFn(t),
			}

			err := u.CreateShipment(tc.userID)
			tc.wantErr(t, err)
		})
	}
}

func TestShipmentCreator_CreateShipment_tableFail(t *testing.T) {
	testCases := []struct {
		name      string
		userStore UserStore
		userID    int64
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "first case",
			userID:  1,
			wantErr: assert.NoError,
			userStore: NewUserStoreMock(t).
				GetUserMock.
				Expect(1).
				Return(User{ID: 1, Country: "Russia"}, nil),
		},
		{
			name:    "second case",
			userID:  2,
			wantErr: assert.NoError,
			userStore: NewUserStoreMock(t).
				GetUserMock.
				Expect(2).
				Return(User{ID: 2, Country: "Russia"}, nil),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := &ShipmentCreator{
				userStore: tc.userStore,
			}

			err := u.CreateShipment(tc.userID)
			tc.wantErr(t, err)
		})
	}
}

package internal

import (
	"encoding/json"

	"github.com/zcolleen/redis"
)

//go:generate mockgen -typed -package internal  -destination zzz_client_mock_gomock.go github.com/zcolleen/redis Client
//go:generate minimock -g -i "github.com/zcolleen/redis.Client" -o ./zzz_client_mock_minimock.go

type redisClient interface {
	redis.Client
}

func setShipmentCache(s *Shipment, client redisClient) error {
	shipment, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return client.Set(s.ID, string(shipment))
}

func getShipmentCache(shipmentID string, client redisClient) (*Shipment, error) {
	s, err := client.Get(shipmentID)
	if err != nil {
		return nil, err
	}

	ship := &Shipment{}
	err = json.Unmarshal([]byte(s), ship)
	if err != nil {
		return nil, err
	}

	return ship, nil
}

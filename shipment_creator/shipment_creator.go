package shipment_creator

type Shipment struct {
}

//go:generate mockgen -typed -package shipment_creator -destination ./mocks/_shipment_creator_mock_test.go . ShipmentCreator
//go:generate minimock -i .ShipmentCreator -o ./mocks
type ShipmentCreator interface {
	CreateShipment() (*Shipment, error)
}

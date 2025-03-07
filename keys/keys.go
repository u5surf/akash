package keys

import (
	"bytes"
	"encoding/binary"

	"github.com/ovrclk/akash/types"
	"github.com/ovrclk/akash/types/base"
)

// XXX: interim hack (iteration!)

type Key interface {
	Bytes() []byte
}

type Address base.Bytes

type (
	Deployment = Address
	Provider   = Address
	Account    = Address
)

func (k Address) Bytes() []byte {
	return k
}

func (k Address) ID() base.Bytes {
	return base.Bytes(k)
}

func DeploymentID(id base.Bytes) Deployment {
	return Address(id)
}

type DeploymentGroup struct {
	types.DeploymentGroupID
}

func DeploymentGroupID(id types.DeploymentGroupID) DeploymentGroup {
	return DeploymentGroup{id}
}

func (k DeploymentGroup) ID() types.DeploymentGroupID {
	return k.DeploymentGroupID
}

func (k DeploymentGroup) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(k.Deployment)
	binary.Write(buf, binary.BigEndian, k.Seq)
	return buf.Bytes()
}

type Order struct {
	types.OrderID
}

func OrderID(id types.OrderID) Order {
	return Order{id}
}

func (k Order) ID() types.OrderID {
	return k.OrderID
}

func (k Order) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(k.Deployment)
	binary.Write(buf, binary.BigEndian, k.Group)
	binary.Write(buf, binary.BigEndian, k.Seq)
	return buf.Bytes()
}

func (k Order) GroupKey() DeploymentGroup {
	return DeploymentGroupID(types.DeploymentGroupID{
		Deployment: k.Deployment,
		Seq:        k.Group,
	})
}

type Fulfillment struct {
	types.FulfillmentID
}

func FulfillmentID(id types.FulfillmentID) Fulfillment {
	return Fulfillment{id}
}

func (k Fulfillment) ID() types.FulfillmentID {
	return k.FulfillmentID
}

func (k Fulfillment) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(k.Deployment)
	binary.Write(buf, binary.BigEndian, k.Group)
	binary.Write(buf, binary.BigEndian, k.Order)
	buf.Write(k.Provider)
	return buf.Bytes()
}

func (k Fulfillment) OrderKey() Order {
	return OrderID(types.OrderID{
		Deployment: k.Deployment,
		Group:      k.Group,
		Seq:        k.Order,
	})
}

func (k Fulfillment) GroupKey() DeploymentGroup {
	return k.OrderKey().GroupKey()
}

type Lease struct {
	types.LeaseID
}

func LeaseID(id types.LeaseID) Lease {
	return Lease{id}
}

func (k Lease) ID() types.LeaseID {
	return k.LeaseID
}

func (k Lease) Bytes() []byte {
	return k.FulfillmentKey().Bytes()
}

func (k Lease) FulfillmentKey() Fulfillment {
	return FulfillmentID(types.FulfillmentID{
		Deployment: k.Deployment,
		Group:      k.Group,
		Order:      k.Order,
		Provider:   k.Provider,
	})
}

func (k Lease) OrderKey() Order {
	return k.FulfillmentKey().OrderKey()
}

func (k Lease) GroupKey() DeploymentGroup {
	return k.FulfillmentKey().GroupKey()
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token_transfer_count_by_address.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type TokenTransferCountByAddressORM struct {
	Address         string `gorm:"primary_key"`
	Count           uint64
	LogIndex        uint64
	TransactionHash string
}

// TableName overrides the default tablename generated by GORM
func (TokenTransferCountByAddressORM) TableName() string {
	return "token_transfer_count_by_addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenTransferCountByAddress) ToORM(ctx context.Context) (TokenTransferCountByAddressORM, error) {
	to := TokenTransferCountByAddressORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByAddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TokenTransferCountByAddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenTransferCountByAddressORM) ToPB(ctx context.Context) (TokenTransferCountByAddress, error) {
	to := TokenTransferCountByAddress{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByAddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.Count = m.Count
	if posthook, ok := interface{}(m).(TokenTransferCountByAddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenTransferCountByAddress the arg will be the target, the caller the one being converted from

// TokenTransferCountByAddressBeforeToORM called before default ToORM code
type TokenTransferCountByAddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenTransferCountByAddressORM) error
}

// TokenTransferCountByAddressAfterToORM called after default ToORM code
type TokenTransferCountByAddressWithAfterToORM interface {
	AfterToORM(context.Context, *TokenTransferCountByAddressORM) error
}

// TokenTransferCountByAddressBeforeToPB called before default ToPB code
type TokenTransferCountByAddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenTransferCountByAddress) error
}

// TokenTransferCountByAddressAfterToPB called after default ToPB code
type TokenTransferCountByAddressWithAfterToPB interface {
	AfterToPB(context.Context, *TokenTransferCountByAddress) error
}

// DefaultCreateTokenTransferCountByAddress executes a basic gorm create call
func DefaultCreateTokenTransferCountByAddress(ctx context.Context, in *TokenTransferCountByAddress, db *gorm1.DB) (*TokenTransferCountByAddress, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferCountByAddressORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTokenTransferCountByAddress patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenTransferCountByAddress(ctx context.Context, patchee *TokenTransferCountByAddress, patcher *TokenTransferCountByAddress, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TokenTransferCountByAddress, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
		if f == prefix+"Address" {
			patchee.Address = patcher.Address
			continue
		}
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTokenTransferCountByAddress executes a gorm list call
func DefaultListTokenTransferCountByAddress(ctx context.Context, db *gorm1.DB) ([]*TokenTransferCountByAddress, error) {
	in := TokenTransferCountByAddress{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TokenTransferCountByAddressORM{}, &TokenTransferCountByAddress{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("address")
	ormResponse := []TokenTransferCountByAddressORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenTransferCountByAddress{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenTransferCountByAddressORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TokenTransferCountByAddressORM) error
}
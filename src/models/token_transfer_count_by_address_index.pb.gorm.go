// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token_transfer_count_by_address_index.proto

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

type TokenTransferCountByAddressIndexORM struct {
	Address         string `gorm:"primary_key"`
	BlockNumber     uint64 `gorm:"index:token_transfer_count_by_address_index_idx_block_number"`
	LogIndex        uint64 `gorm:"primary_key"`
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (TokenTransferCountByAddressIndexORM) TableName() string {
	return "token_transfer_count_by_address_indices"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TokenTransferCountByAddressIndex) ToORM(ctx context.Context) (TokenTransferCountByAddressIndexORM, error) {
	to := TokenTransferCountByAddressIndexORM{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByAddressIndexWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TokenTransferCountByAddressIndexWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TokenTransferCountByAddressIndexORM) ToPB(ctx context.Context) (TokenTransferCountByAddressIndex, error) {
	to := TokenTransferCountByAddressIndex{}
	var err error
	if prehook, ok := interface{}(m).(TokenTransferCountByAddressIndexWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	to.Address = m.Address
	to.BlockNumber = m.BlockNumber
	if posthook, ok := interface{}(m).(TokenTransferCountByAddressIndexWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TokenTransferCountByAddressIndex the arg will be the target, the caller the one being converted from

// TokenTransferCountByAddressIndexBeforeToORM called before default ToORM code
type TokenTransferCountByAddressIndexWithBeforeToORM interface {
	BeforeToORM(context.Context, *TokenTransferCountByAddressIndexORM) error
}

// TokenTransferCountByAddressIndexAfterToORM called after default ToORM code
type TokenTransferCountByAddressIndexWithAfterToORM interface {
	AfterToORM(context.Context, *TokenTransferCountByAddressIndexORM) error
}

// TokenTransferCountByAddressIndexBeforeToPB called before default ToPB code
type TokenTransferCountByAddressIndexWithBeforeToPB interface {
	BeforeToPB(context.Context, *TokenTransferCountByAddressIndex) error
}

// TokenTransferCountByAddressIndexAfterToPB called after default ToPB code
type TokenTransferCountByAddressIndexWithAfterToPB interface {
	AfterToPB(context.Context, *TokenTransferCountByAddressIndex) error
}

// DefaultCreateTokenTransferCountByAddressIndex executes a basic gorm create call
func DefaultCreateTokenTransferCountByAddressIndex(ctx context.Context, in *TokenTransferCountByAddressIndex, db *gorm1.DB) (*TokenTransferCountByAddressIndex, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressIndexORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressIndexORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TokenTransferCountByAddressIndexORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressIndexORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTokenTransferCountByAddressIndex patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTokenTransferCountByAddressIndex(ctx context.Context, patchee *TokenTransferCountByAddressIndex, patcher *TokenTransferCountByAddressIndex, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TokenTransferCountByAddressIndex, error) {
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
		if f == prefix+"BlockNumber" {
			patchee.BlockNumber = patcher.BlockNumber
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTokenTransferCountByAddressIndex executes a gorm list call
func DefaultListTokenTransferCountByAddressIndex(ctx context.Context, db *gorm1.DB) ([]*TokenTransferCountByAddressIndex, error) {
	in := TokenTransferCountByAddressIndex{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressIndexORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TokenTransferCountByAddressIndexORM{}, &TokenTransferCountByAddressIndex{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressIndexORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TokenTransferCountByAddressIndexORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TokenTransferCountByAddressIndexORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TokenTransferCountByAddressIndex{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TokenTransferCountByAddressIndexORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressIndexORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TokenTransferCountByAddressIndexORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TokenTransferCountByAddressIndexORM) error
}

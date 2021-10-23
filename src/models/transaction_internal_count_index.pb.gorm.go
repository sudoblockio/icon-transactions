// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction_internal_count_index.proto

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

type TransactionInternalCountIndexORM struct {
	LogIndex        int32  `gorm:"primary_key"`
	TransactionHash string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (TransactionInternalCountIndexORM) TableName() string {
	return "transaction_internal_count_indices"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *TransactionInternalCountIndex) ToORM(ctx context.Context) (TransactionInternalCountIndexORM, error) {
	to := TransactionInternalCountIndexORM{}
	var err error
	if prehook, ok := interface{}(m).(TransactionInternalCountIndexWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	if posthook, ok := interface{}(m).(TransactionInternalCountIndexWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TransactionInternalCountIndexORM) ToPB(ctx context.Context) (TransactionInternalCountIndex, error) {
	to := TransactionInternalCountIndex{}
	var err error
	if prehook, ok := interface{}(m).(TransactionInternalCountIndexWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	if posthook, ok := interface{}(m).(TransactionInternalCountIndexWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type TransactionInternalCountIndex the arg will be the target, the caller the one being converted from

// TransactionInternalCountIndexBeforeToORM called before default ToORM code
type TransactionInternalCountIndexWithBeforeToORM interface {
	BeforeToORM(context.Context, *TransactionInternalCountIndexORM) error
}

// TransactionInternalCountIndexAfterToORM called after default ToORM code
type TransactionInternalCountIndexWithAfterToORM interface {
	AfterToORM(context.Context, *TransactionInternalCountIndexORM) error
}

// TransactionInternalCountIndexBeforeToPB called before default ToPB code
type TransactionInternalCountIndexWithBeforeToPB interface {
	BeforeToPB(context.Context, *TransactionInternalCountIndex) error
}

// TransactionInternalCountIndexAfterToPB called after default ToPB code
type TransactionInternalCountIndexWithAfterToPB interface {
	AfterToPB(context.Context, *TransactionInternalCountIndex) error
}

// DefaultCreateTransactionInternalCountIndex executes a basic gorm create call
func DefaultCreateTransactionInternalCountIndex(ctx context.Context, in *TransactionInternalCountIndex, db *gorm1.DB) (*TransactionInternalCountIndex, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalCountIndexORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalCountIndexORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type TransactionInternalCountIndexORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalCountIndexORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskTransactionInternalCountIndex patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskTransactionInternalCountIndex(ctx context.Context, patchee *TransactionInternalCountIndex, patcher *TransactionInternalCountIndex, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*TransactionInternalCountIndex, error) {
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
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListTransactionInternalCountIndex executes a gorm list call
func DefaultListTransactionInternalCountIndex(ctx context.Context, db *gorm1.DB) ([]*TransactionInternalCountIndex, error) {
	in := TransactionInternalCountIndex{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalCountIndexORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &TransactionInternalCountIndexORM{}, &TransactionInternalCountIndex{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalCountIndexORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("transaction_hash")
	ormResponse := []TransactionInternalCountIndexORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(TransactionInternalCountIndexORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*TransactionInternalCountIndex{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type TransactionInternalCountIndexORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalCountIndexORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type TransactionInternalCountIndexORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]TransactionInternalCountIndexORM) error
}
package evm

import (
	"context"
	"math/big"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	relaytypes "github.com/smartcontractkit/chainlink-relay/pkg/types"
	"github.com/smartcontractkit/chainlink-relay/pkg/utils"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type encoder struct {
	Definitions map[string]*types.CodecEntry
}

var evmDecoderHook = mapstructure.ComposeDecodeHookFunc(utils.BigIntHook, utils.SliceToArrayVerifySizeHook, sizeVerifyBigIntHook)

var _ relaytypes.Encoder = &encoder{}

func (e *encoder) Encode(ctx context.Context, item any, itemType string) (ocrtypes.Report, error) {
	info, ok := e.Definitions[itemType]
	if !ok {
		return nil, relaytypes.InvalidTypeError{}
	}

	return encode(reflect.ValueOf(item), info)
}

func (e *encoder) GetMaxEncodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return 0, errors.New("TODO")
}

func (e *encoder) getEncodingType(itemType string, forceArray bool) (reflect.Type, error) {
	info, ok := e.Definitions[itemType]
	if !ok {
		return info.CheckedArrayType, nil
	}

	if forceArray {
		if info.CheckedArrayType == nil {
			return nil, relaytypes.InvalidTypeError{}
		}
		return info.CheckedArrayType, nil
	}

	return info.CheckedType, nil
}

func encode(item reflect.Value, info *types.CodecEntry) (ocrtypes.Report, error) {
	switch item.Kind() {
	case reflect.Pointer:
		return encode(item.Elem(), info)
	case reflect.Array, reflect.Slice:
		return encodeArray(item, info)
	case reflect.Struct, reflect.Map:
		return encodeItem(item, info)
	default:
		return nil, relaytypes.InvalidEncodingError{}
	}
}

func encodeArray(item reflect.Value, info *types.CodecEntry) (ocrtypes.Report, error) {
	var tmpMap []map[string]any
	if err := mapstructureDecode(item.Interface(), &tmpMap); err != nil {
		return nil, err
	}

	if info.ArraySize != 0 && info.ArraySize != len(tmpMap) {
		return nil, relaytypes.InvalidTypeError{}
	}

	singleMap, err := utils.MergeValueFields(tmpMap)
	if err != nil {
		return nil, relaytypes.InvalidTypeError{}
	}

	itemChecked := reflect.New(info.CheckedType)
	if err = mapstructureDecode(singleMap, itemChecked.Interface()); err != nil {
		return nil, err
	}

	return encodeItem(itemChecked, info)
}

func encodeItem(item reflect.Value, info *types.CodecEntry) (ocrtypes.Report, error) {
	if item.Type() == reflect.PointerTo(info.CheckedType) {
		item = reflect.NewAt(info.NativeType, item.UnsafePointer())
	} else if item.Type() != reflect.PointerTo(info.NativeType) {
		checked := reflect.New(info.CheckedType)
		if err := mapstructureDecode(item.Interface(), checked.Interface()); err != nil {
			return nil, err
		}
		item = reflect.NewAt(info.NativeType, checked.UnsafePointer())
	}

	item = item.Elem()
	length := item.NumField()
	values := make([]any, length)
	iType := item.Type()
	for i := 0; i < length; i++ {
		if iType.Field(i).IsExported() {
			values[i] = item.Field(i).Interface()
		}
	}

	if bytes, err := info.Args.Pack(values...); err == nil {
		withPrefix := make([]byte, 0, len(info.EncodingPrefix)+len(bytes))
		copy(withPrefix, info.EncodingPrefix)
		return append(withPrefix, bytes...), nil
	}

	return nil, relaytypes.InvalidEncodingError{}
}

func sizeVerifyBigIntHook(_, to reflect.Type, data any) (any, error) {
	bi, ok := data.(*big.Int)
	if !ok {
		return data, nil
	}

	if to.Implements(types.SizedBigIntType()) {
		converted := reflect.ValueOf(bi).Convert(to).Interface().(types.SizedBigInt)
		return converted, converted.Verify()
	}

	return data, nil
}

func mapstructureDecode(src, dest any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: evmDecoderHook,
		Result:     dest,
	})
	if err != nil || decoder.Decode(src) != nil {
		return relaytypes.InvalidTypeError{}
	}
	return nil
}

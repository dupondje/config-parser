// Code generated by go generate; DO NOT EDIT.
package simple

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *SimpleOption) GetParserName() string {
    return p.Name
}

func (p *SimpleOption) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.SimpleOption{}
			return p.data, nil
		}
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleOption) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.FetchError
	}
	if p.data == nil {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *SimpleOption) Delete(index int) error {
	p.Init()
	return nil
}

func (p *SimpleOption) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *SimpleOption) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.SimpleOption:
		p.data = newValue
	case types.SimpleOption:
		p.data = &newValue
	default:
		return errors.InvalidData
	}
	return nil
}
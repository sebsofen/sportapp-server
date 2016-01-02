// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package services

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type CitySvc interface {
	// Parameters:
	//  - Token
	//  - Title
	//  - Coords
	CreateCity(token string, title string, coords []*Coordinate) (err error)
	// Parameters:
	//  - Coordinate
	//  - Limit
	GetNearBy(coordinate *Coordinate, limit int32) (r []*City, err error)
	GetAllCities() (r []*City, err error)
}

type CitySvcClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewCitySvcClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *CitySvcClient {
	return &CitySvcClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewCitySvcClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *CitySvcClient {
	return &CitySvcClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Token
//  - Title
//  - Coords
func (p *CitySvcClient) CreateCity(token string, title string, coords []*Coordinate) (err error) {
	if err = p.sendCreateCity(token, title, coords); err != nil {
		return
	}
	return p.recvCreateCity()
}

func (p *CitySvcClient) sendCreateCity(token string, title string, coords []*Coordinate) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("createCity", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CitySvcCreateCityArgs{
		Token:  token,
		Title:  title,
		Coords: coords,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *CitySvcClient) recvCreateCity() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "createCity" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "createCity failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "createCity failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error147 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error148 error
		error148, err = error147.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error148
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "createCity failed: invalid message type")
		return
	}
	result := CitySvcCreateCityResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

// Parameters:
//  - Coordinate
//  - Limit
func (p *CitySvcClient) GetNearBy(coordinate *Coordinate, limit int32) (r []*City, err error) {
	if err = p.sendGetNearBy(coordinate, limit); err != nil {
		return
	}
	return p.recvGetNearBy()
}

func (p *CitySvcClient) sendGetNearBy(coordinate *Coordinate, limit int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getNearBy", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CitySvcGetNearByArgs{
		Coordinate: coordinate,
		Limit:      limit,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *CitySvcClient) recvGetNearBy() (value []*City, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getNearBy" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getNearBy failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getNearBy failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error149 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error150 error
		error150, err = error149.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error150
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getNearBy failed: invalid message type")
		return
	}
	result := CitySvcGetNearByResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

func (p *CitySvcClient) GetAllCities() (r []*City, err error) {
	if err = p.sendGetAllCities(); err != nil {
		return
	}
	return p.recvGetAllCities()
}

func (p *CitySvcClient) sendGetAllCities() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getAllCities", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CitySvcGetAllCitiesArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *CitySvcClient) recvGetAllCities() (value []*City, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getAllCities" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getAllCities failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getAllCities failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error151 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error152 error
		error152, err = error151.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error152
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getAllCities failed: invalid message type")
		return
	}
	result := CitySvcGetAllCitiesResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type CitySvcProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      CitySvc
}

func (p *CitySvcProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *CitySvcProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *CitySvcProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewCitySvcProcessor(handler CitySvc) *CitySvcProcessor {

	self153 := &CitySvcProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self153.processorMap["createCity"] = &citySvcProcessorCreateCity{handler: handler}
	self153.processorMap["getNearBy"] = &citySvcProcessorGetNearBy{handler: handler}
	self153.processorMap["getAllCities"] = &citySvcProcessorGetAllCities{handler: handler}
	return self153
}

func (p *CitySvcProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x154 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x154.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x154

}

type citySvcProcessorCreateCity struct {
	handler CitySvc
}

func (p *citySvcProcessorCreateCity) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CitySvcCreateCityArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("createCity", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CitySvcCreateCityResult{}
	var err2 error
	if err2 = p.handler.CreateCity(args.Token, args.Title, args.Coords); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing createCity: "+err2.Error())
		oprot.WriteMessageBegin("createCity", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("createCity", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type citySvcProcessorGetNearBy struct {
	handler CitySvc
}

func (p *citySvcProcessorGetNearBy) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CitySvcGetNearByArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getNearBy", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CitySvcGetNearByResult{}
	var retval []*City
	var err2 error
	if retval, err2 = p.handler.GetNearBy(args.Coordinate, args.Limit); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getNearBy: "+err2.Error())
		oprot.WriteMessageBegin("getNearBy", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getNearBy", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type citySvcProcessorGetAllCities struct {
	handler CitySvc
}

func (p *citySvcProcessorGetAllCities) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CitySvcGetAllCitiesArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getAllCities", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CitySvcGetAllCitiesResult{}
	var retval []*City
	var err2 error
	if retval, err2 = p.handler.GetAllCities(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getAllCities: "+err2.Error())
		oprot.WriteMessageBegin("getAllCities", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getAllCities", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Token
//  - Title
//  - Coords
type CitySvcCreateCityArgs struct {
	Token  string        `thrift:"token,1" db:"token" json:"token"`
	Title  string        `thrift:"title,2" db:"title" json:"title"`
	Coords []*Coordinate `thrift:"coords,3" db:"coords" json:"coords"`
}

func NewCitySvcCreateCityArgs() *CitySvcCreateCityArgs {
	return &CitySvcCreateCityArgs{}
}

func (p *CitySvcCreateCityArgs) GetToken() string {
	return p.Token
}

func (p *CitySvcCreateCityArgs) GetTitle() string {
	return p.Title
}

func (p *CitySvcCreateCityArgs) GetCoords() []*Coordinate {
	return p.Coords
}
func (p *CitySvcCreateCityArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcCreateCityArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *CitySvcCreateCityArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *CitySvcCreateCityArgs) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Coordinate, 0, size)
	p.Coords = tSlice
	for i := 0; i < size; i++ {
		_elem155 := &Coordinate{}
		if err := _elem155.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem155), err)
		}
		p.Coords = append(p.Coords, _elem155)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *CitySvcCreateCityArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createCity_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcCreateCityArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *CitySvcCreateCityArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("title", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:title: ", p), err)
	}
	if err := oprot.WriteString(string(p.Title)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.title (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:title: ", p), err)
	}
	return err
}

func (p *CitySvcCreateCityArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("coords", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:coords: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Coords)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Coords {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:coords: ", p), err)
	}
	return err
}

func (p *CitySvcCreateCityArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcCreateCityArgs(%+v)", *p)
}

type CitySvcCreateCityResult struct {
}

func NewCitySvcCreateCityResult() *CitySvcCreateCityResult {
	return &CitySvcCreateCityResult{}
}

func (p *CitySvcCreateCityResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcCreateCityResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createCity_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcCreateCityResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcCreateCityResult(%+v)", *p)
}

// Attributes:
//  - Coordinate
//  - Limit
type CitySvcGetNearByArgs struct {
	Coordinate *Coordinate `thrift:"coordinate,1" db:"coordinate" json:"coordinate"`
	Limit      int32       `thrift:"limit,2" db:"limit" json:"limit"`
}

func NewCitySvcGetNearByArgs() *CitySvcGetNearByArgs {
	return &CitySvcGetNearByArgs{}
}

var CitySvcGetNearByArgs_Coordinate_DEFAULT *Coordinate

func (p *CitySvcGetNearByArgs) GetCoordinate() *Coordinate {
	if !p.IsSetCoordinate() {
		return CitySvcGetNearByArgs_Coordinate_DEFAULT
	}
	return p.Coordinate
}

func (p *CitySvcGetNearByArgs) GetLimit() int32 {
	return p.Limit
}
func (p *CitySvcGetNearByArgs) IsSetCoordinate() bool {
	return p.Coordinate != nil
}

func (p *CitySvcGetNearByArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcGetNearByArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Coordinate = &Coordinate{}
	if err := p.Coordinate.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Coordinate), err)
	}
	return nil
}

func (p *CitySvcGetNearByArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Limit = v
	}
	return nil
}

func (p *CitySvcGetNearByArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getNearBy_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcGetNearByArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("coordinate", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:coordinate: ", p), err)
	}
	if err := p.Coordinate.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Coordinate), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:coordinate: ", p), err)
	}
	return err
}

func (p *CitySvcGetNearByArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("limit", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:limit: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Limit)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.limit (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:limit: ", p), err)
	}
	return err
}

func (p *CitySvcGetNearByArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcGetNearByArgs(%+v)", *p)
}

// Attributes:
//  - Success
type CitySvcGetNearByResult struct {
	Success []*City `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewCitySvcGetNearByResult() *CitySvcGetNearByResult {
	return &CitySvcGetNearByResult{}
}

var CitySvcGetNearByResult_Success_DEFAULT []*City

func (p *CitySvcGetNearByResult) GetSuccess() []*City {
	return p.Success
}
func (p *CitySvcGetNearByResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CitySvcGetNearByResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcGetNearByResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*City, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem156 := &City{}
		if err := _elem156.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem156), err)
		}
		p.Success = append(p.Success, _elem156)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *CitySvcGetNearByResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getNearBy_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcGetNearByResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *CitySvcGetNearByResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcGetNearByResult(%+v)", *p)
}

type CitySvcGetAllCitiesArgs struct {
}

func NewCitySvcGetAllCitiesArgs() *CitySvcGetAllCitiesArgs {
	return &CitySvcGetAllCitiesArgs{}
}

func (p *CitySvcGetAllCitiesArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcGetAllCitiesArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllCities_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcGetAllCitiesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcGetAllCitiesArgs(%+v)", *p)
}

// Attributes:
//  - Success
type CitySvcGetAllCitiesResult struct {
	Success []*City `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewCitySvcGetAllCitiesResult() *CitySvcGetAllCitiesResult {
	return &CitySvcGetAllCitiesResult{}
}

var CitySvcGetAllCitiesResult_Success_DEFAULT []*City

func (p *CitySvcGetAllCitiesResult) GetSuccess() []*City {
	return p.Success
}
func (p *CitySvcGetAllCitiesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CitySvcGetAllCitiesResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CitySvcGetAllCitiesResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*City, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem157 := &City{}
		if err := _elem157.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem157), err)
		}
		p.Success = append(p.Success, _elem157)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *CitySvcGetAllCitiesResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllCities_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CitySvcGetAllCitiesResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *CitySvcGetAllCitiesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CitySvcGetAllCitiesResult(%+v)", *p)
}

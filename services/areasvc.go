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

type AreaSvc interface {
	// Parameters:
	//  - Title
	//  - Description
	//  - Coords
	CreateArea(title string, description string, coords []*Coordinate) (err error)
	// Parameters:
	//  - Coordinate
	//  - Limit
	GetNearBy(coordinate *Coordinate, limit int32) (r []*Area, err error)
}

type AreaSvcClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewAreaSvcClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AreaSvcClient {
	return &AreaSvcClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewAreaSvcClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AreaSvcClient {
	return &AreaSvcClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Title
//  - Description
//  - Coords
func (p *AreaSvcClient) CreateArea(title string, description string, coords []*Coordinate) (err error) {
	if err = p.sendCreateArea(title, description, coords); err != nil {
		return
	}
	return p.recvCreateArea()
}

func (p *AreaSvcClient) sendCreateArea(title string, description string, coords []*Coordinate) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("createArea", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AreaSvcCreateAreaArgs{
		Title:       title,
		Description: description,
		Coords:      coords,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *AreaSvcClient) recvCreateArea() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "createArea" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "createArea failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "createArea failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error19 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error20 error
		error20, err = error19.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error20
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "createArea failed: invalid message type")
		return
	}
	result := AreaSvcCreateAreaResult{}
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
func (p *AreaSvcClient) GetNearBy(coordinate *Coordinate, limit int32) (r []*Area, err error) {
	if err = p.sendGetNearBy(coordinate, limit); err != nil {
		return
	}
	return p.recvGetNearBy()
}

func (p *AreaSvcClient) sendGetNearBy(coordinate *Coordinate, limit int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getNearBy", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AreaSvcGetNearByArgs{
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

func (p *AreaSvcClient) recvGetNearBy() (value []*Area, err error) {
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
		error21 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error22 error
		error22, err = error21.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error22
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getNearBy failed: invalid message type")
		return
	}
	result := AreaSvcGetNearByResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type AreaSvcProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AreaSvc
}

func (p *AreaSvcProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AreaSvcProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AreaSvcProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAreaSvcProcessor(handler AreaSvc) *AreaSvcProcessor {

	self23 := &AreaSvcProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self23.processorMap["createArea"] = &areaSvcProcessorCreateArea{handler: handler}
	self23.processorMap["getNearBy"] = &areaSvcProcessorGetNearBy{handler: handler}
	return self23
}

func (p *AreaSvcProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x24 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x24.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x24

}

type areaSvcProcessorCreateArea struct {
	handler AreaSvc
}

func (p *areaSvcProcessorCreateArea) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AreaSvcCreateAreaArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("createArea", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AreaSvcCreateAreaResult{}
	var err2 error
	if err2 = p.handler.CreateArea(args.Title, args.Description, args.Coords); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing createArea: "+err2.Error())
		oprot.WriteMessageBegin("createArea", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("createArea", thrift.REPLY, seqId); err2 != nil {
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

type areaSvcProcessorGetNearBy struct {
	handler AreaSvc
}

func (p *areaSvcProcessorGetNearBy) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AreaSvcGetNearByArgs{}
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
	result := AreaSvcGetNearByResult{}
	var retval []*Area
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

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Title
//  - Description
//  - Coords
type AreaSvcCreateAreaArgs struct {
	Title       string        `thrift:"title,1" json:"title"`
	Description string        `thrift:"description,2" json:"description"`
	Coords      []*Coordinate `thrift:"coords,3" json:"coords"`
}

func NewAreaSvcCreateAreaArgs() *AreaSvcCreateAreaArgs {
	return &AreaSvcCreateAreaArgs{}
}

func (p *AreaSvcCreateAreaArgs) GetTitle() string {
	return p.Title
}

func (p *AreaSvcCreateAreaArgs) GetDescription() string {
	return p.Description
}

func (p *AreaSvcCreateAreaArgs) GetCoords() []*Coordinate {
	return p.Coords
}
func (p *AreaSvcCreateAreaArgs) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
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

func (p *AreaSvcCreateAreaArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *AreaSvcCreateAreaArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Description = v
	}
	return nil
}

func (p *AreaSvcCreateAreaArgs) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Coordinate, 0, size)
	p.Coords = tSlice
	for i := 0; i < size; i++ {
		_elem25 := &Coordinate{}
		if err := _elem25.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem25), err)
		}
		p.Coords = append(p.Coords, _elem25)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AreaSvcCreateAreaArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createArea_args"); err != nil {
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

func (p *AreaSvcCreateAreaArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("title", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:title: ", p), err)
	}
	if err := oprot.WriteString(string(p.Title)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.title (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:title: ", p), err)
	}
	return err
}

func (p *AreaSvcCreateAreaArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("description", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:description: ", p), err)
	}
	if err := oprot.WriteString(string(p.Description)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.description (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:description: ", p), err)
	}
	return err
}

func (p *AreaSvcCreateAreaArgs) writeField3(oprot thrift.TProtocol) (err error) {
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

func (p *AreaSvcCreateAreaArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AreaSvcCreateAreaArgs(%+v)", *p)
}

type AreaSvcCreateAreaResult struct {
}

func NewAreaSvcCreateAreaResult() *AreaSvcCreateAreaResult {
	return &AreaSvcCreateAreaResult{}
}

func (p *AreaSvcCreateAreaResult) Read(iprot thrift.TProtocol) error {
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

func (p *AreaSvcCreateAreaResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createArea_result"); err != nil {
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

func (p *AreaSvcCreateAreaResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AreaSvcCreateAreaResult(%+v)", *p)
}

// Attributes:
//  - Coordinate
//  - Limit
type AreaSvcGetNearByArgs struct {
	Coordinate *Coordinate `thrift:"coordinate,1" json:"coordinate"`
	Limit      int32       `thrift:"limit,2" json:"limit"`
}

func NewAreaSvcGetNearByArgs() *AreaSvcGetNearByArgs {
	return &AreaSvcGetNearByArgs{}
}

var AreaSvcGetNearByArgs_Coordinate_DEFAULT *Coordinate

func (p *AreaSvcGetNearByArgs) GetCoordinate() *Coordinate {
	if !p.IsSetCoordinate() {
		return AreaSvcGetNearByArgs_Coordinate_DEFAULT
	}
	return p.Coordinate
}

func (p *AreaSvcGetNearByArgs) GetLimit() int32 {
	return p.Limit
}
func (p *AreaSvcGetNearByArgs) IsSetCoordinate() bool {
	return p.Coordinate != nil
}

func (p *AreaSvcGetNearByArgs) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
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

func (p *AreaSvcGetNearByArgs) readField1(iprot thrift.TProtocol) error {
	p.Coordinate = &Coordinate{}
	if err := p.Coordinate.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Coordinate), err)
	}
	return nil
}

func (p *AreaSvcGetNearByArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Limit = v
	}
	return nil
}

func (p *AreaSvcGetNearByArgs) Write(oprot thrift.TProtocol) error {
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

func (p *AreaSvcGetNearByArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *AreaSvcGetNearByArgs) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *AreaSvcGetNearByArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AreaSvcGetNearByArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AreaSvcGetNearByResult struct {
	Success []*Area `thrift:"success,0" json:"success,omitempty"`
}

func NewAreaSvcGetNearByResult() *AreaSvcGetNearByResult {
	return &AreaSvcGetNearByResult{}
}

var AreaSvcGetNearByResult_Success_DEFAULT []*Area

func (p *AreaSvcGetNearByResult) GetSuccess() []*Area {
	return p.Success
}
func (p *AreaSvcGetNearByResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AreaSvcGetNearByResult) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField0(iprot); err != nil {
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

func (p *AreaSvcGetNearByResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Area, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem26 := &Area{}
		if err := _elem26.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem26), err)
		}
		p.Success = append(p.Success, _elem26)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *AreaSvcGetNearByResult) Write(oprot thrift.TProtocol) error {
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

func (p *AreaSvcGetNearByResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *AreaSvcGetNearByResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AreaSvcGetNearByResult(%+v)", *p)
}

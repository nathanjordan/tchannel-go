// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hyperbahn

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Hyperbahn interface {
	// Parameters:
	//  - Query
	Discover(query *DiscoveryQuery) (r *DiscoveryResult_, err error)
}

type HyperbahnClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewHyperbahnClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HyperbahnClient {
	return &HyperbahnClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewHyperbahnClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HyperbahnClient {
	return &HyperbahnClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Query
func (p *HyperbahnClient) Discover(query *DiscoveryQuery) (r *DiscoveryResult_, err error) {
	if err = p.sendDiscover(query); err != nil {
		return
	}
	return p.recvDiscover()
}

func (p *HyperbahnClient) sendDiscover(query *DiscoveryQuery) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("discover", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := HyperbahnDiscoverArgs{
		Query: query,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HyperbahnClient) recvDiscover() (value *DiscoveryResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "discover" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "discover failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "discover failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error1 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error2 error
		error2, err = error1.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error2
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "discover failed: invalid message type")
		return
	}
	result := HyperbahnDiscoverResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.NoPeersAvailable != nil {
		err = result.NoPeersAvailable
		return
	} else if result.InvalidServiceName != nil {
		err = result.InvalidServiceName
		return
	}
	value = result.GetSuccess()
	return
}

type HyperbahnProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Hyperbahn
}

func (p *HyperbahnProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *HyperbahnProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *HyperbahnProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewHyperbahnProcessor(handler Hyperbahn) *HyperbahnProcessor {

	self3 := &HyperbahnProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["discover"] = &hyperbahnProcessorDiscover{handler: handler}
	return self3
}

func (p *HyperbahnProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x4

}

type hyperbahnProcessorDiscover struct {
	handler Hyperbahn
}

func (p *hyperbahnProcessorDiscover) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := HyperbahnDiscoverArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("discover", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := HyperbahnDiscoverResult{}
	var retval *DiscoveryResult_
	var err2 error
	if retval, err2 = p.handler.Discover(args.Query); err2 != nil {
		switch v := err2.(type) {
		case *NoPeersAvailable:
			result.NoPeersAvailable = v
		case *InvalidServiceName:
			result.InvalidServiceName = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing discover: "+err2.Error())
			oprot.WriteMessageBegin("discover", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("discover", thrift.REPLY, seqId); err2 != nil {
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
//  - Query
type HyperbahnDiscoverArgs struct {
	Query *DiscoveryQuery `thrift:"query,1,required" db:"query" json:"query"`
}

func NewHyperbahnDiscoverArgs() *HyperbahnDiscoverArgs {
	return &HyperbahnDiscoverArgs{}
}

var HyperbahnDiscoverArgs_Query_DEFAULT *DiscoveryQuery

func (p *HyperbahnDiscoverArgs) GetQuery() *DiscoveryQuery {
	if !p.IsSetQuery() {
		return HyperbahnDiscoverArgs_Query_DEFAULT
	}
	return p.Query
}
func (p *HyperbahnDiscoverArgs) IsSetQuery() bool {
	return p.Query != nil
}

func (p *HyperbahnDiscoverArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetQuery bool = false

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
			issetQuery = true
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
	if !issetQuery {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Query is not set"))
	}
	return nil
}

func (p *HyperbahnDiscoverArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Query = &DiscoveryQuery{}
	if err := p.Query.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Query), err)
	}
	return nil
}

func (p *HyperbahnDiscoverArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("discover_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *HyperbahnDiscoverArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("query", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:query: ", p), err)
	}
	if err := p.Query.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Query), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:query: ", p), err)
	}
	return err
}

func (p *HyperbahnDiscoverArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HyperbahnDiscoverArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - NoPeersAvailable
//  - InvalidServiceName
type HyperbahnDiscoverResult struct {
	Success            *DiscoveryResult_   `thrift:"success,0" db:"success" json:"success,omitempty"`
	NoPeersAvailable   *NoPeersAvailable   `thrift:"noPeersAvailable,1" db:"noPeersAvailable" json:"noPeersAvailable,omitempty"`
	InvalidServiceName *InvalidServiceName `thrift:"invalidServiceName,2" db:"invalidServiceName" json:"invalidServiceName,omitempty"`
}

func NewHyperbahnDiscoverResult() *HyperbahnDiscoverResult {
	return &HyperbahnDiscoverResult{}
}

var HyperbahnDiscoverResult_Success_DEFAULT *DiscoveryResult_

func (p *HyperbahnDiscoverResult) GetSuccess() *DiscoveryResult_ {
	if !p.IsSetSuccess() {
		return HyperbahnDiscoverResult_Success_DEFAULT
	}
	return p.Success
}

var HyperbahnDiscoverResult_NoPeersAvailable_DEFAULT *NoPeersAvailable

func (p *HyperbahnDiscoverResult) GetNoPeersAvailable() *NoPeersAvailable {
	if !p.IsSetNoPeersAvailable() {
		return HyperbahnDiscoverResult_NoPeersAvailable_DEFAULT
	}
	return p.NoPeersAvailable
}

var HyperbahnDiscoverResult_InvalidServiceName_DEFAULT *InvalidServiceName

func (p *HyperbahnDiscoverResult) GetInvalidServiceName() *InvalidServiceName {
	if !p.IsSetInvalidServiceName() {
		return HyperbahnDiscoverResult_InvalidServiceName_DEFAULT
	}
	return p.InvalidServiceName
}
func (p *HyperbahnDiscoverResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HyperbahnDiscoverResult) IsSetNoPeersAvailable() bool {
	return p.NoPeersAvailable != nil
}

func (p *HyperbahnDiscoverResult) IsSetInvalidServiceName() bool {
	return p.InvalidServiceName != nil
}

func (p *HyperbahnDiscoverResult) Read(iprot thrift.TProtocol) error {
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

func (p *HyperbahnDiscoverResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &DiscoveryResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *HyperbahnDiscoverResult) ReadField1(iprot thrift.TProtocol) error {
	p.NoPeersAvailable = &NoPeersAvailable{}
	if err := p.NoPeersAvailable.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.NoPeersAvailable), err)
	}
	return nil
}

func (p *HyperbahnDiscoverResult) ReadField2(iprot thrift.TProtocol) error {
	p.InvalidServiceName = &InvalidServiceName{}
	if err := p.InvalidServiceName.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.InvalidServiceName), err)
	}
	return nil
}

func (p *HyperbahnDiscoverResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("discover_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *HyperbahnDiscoverResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *HyperbahnDiscoverResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetNoPeersAvailable() {
		if err := oprot.WriteFieldBegin("noPeersAvailable", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:noPeersAvailable: ", p), err)
		}
		if err := p.NoPeersAvailable.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.NoPeersAvailable), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:noPeersAvailable: ", p), err)
		}
	}
	return err
}

func (p *HyperbahnDiscoverResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetInvalidServiceName() {
		if err := oprot.WriteFieldBegin("invalidServiceName", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:invalidServiceName: ", p), err)
		}
		if err := p.InvalidServiceName.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.InvalidServiceName), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:invalidServiceName: ", p), err)
		}
	}
	return err
}

func (p *HyperbahnDiscoverResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HyperbahnDiscoverResult(%+v)", *p)
}

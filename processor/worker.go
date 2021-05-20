/**
 * Copyright 2017 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ------------------------------------------------------------------------------
 */

package processor

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/sawtooth-sdk-go/messaging"
	"github.com/hyperledger/sawtooth-sdk-go/src/protobuf/smartcontract_pb2"
	"github.com/hyperledger/sawtooth-sdk-go/src/protobuf/transaction_pb2"
	zmq "github.com/pebbe/zmq4"
	"protobuf/processor_pb2"
	"protobuf/validator_pb2"
)

// The main worker thread finds an appropriate handler and processes the request
func worker(context *zmq.Context, uri string, queue <-chan *validator_pb2.Message, done chan<- bool, handlers []TransactionHandler) {
	// Connect to the main send/receive thread
	connection, err := messaging.NewConnection(context, zmq.DEALER, uri, false)
	if err != nil {
		logger.Errorf("Failed to connect to main thread: %v", err)
		done <- false
		return
	}
	defer connection.Close()

	// Receive work off of the queue until the queue is closed
	for msg := range queue {
		switch msg.MessageType {
		case validator_pb2.Message_TP_PROCESS_REQUEST:
			if err := execSmartContract(msg, handlers, connection); err != nil {
				break
			}
		case validator_pb2.Message_SMART_CONTRACT_VALIDATE_REQUEST:
			if err := execValidations(msg, handlers, connection); err != nil {
				break
			}
		}

	}

	done <- true
}

func execValidations(msg *validator_pb2.Message, handlers []TransactionHandler, connection *messaging.ZmqConnection) error {
	logger.Debug("Exec SmartContract Validations")
	id := connection.Identity()
	request := &smartcontract_pb2.SmartContractValidationRequest{}
	err := proto.Unmarshal(msg.GetContent(), request)
	if err != nil {
		logger.Errorf(
			"(%v) Failed to unmarshal TpProcessRequest: %v", id, err,
		)
		return err
	}

	handler, err := findHandlerForValidation(handlers, request.SmartContractAddress)
	if err != nil {
		logger.Errorf("(%v) Failed to find handler: %v", id, err)
		return err
	}

	handlerResponse, err := handler.Validate(request)

	response := smartcontract_pb2.SmartContractValidationResponse{}
	if err != nil {
		switch e := err.(type) {
		case *InvalidTransactionError:
			logger.Warnf("(%v) %v", id, e)
			response.Status = smartcontract_pb2.SmartContractValidationResponse_INVALID_TRANSACTION
		default:
			logger.Errorf("(%v) Internal error: %v", id, err)
			response.Status = smartcontract_pb2.SmartContractValidationResponse_INTERNAL_ERROR
		}
	}

	response.Status = smartcontract_pb2.SmartContractValidationResponse_OK
	response.T = int32(handlerResponse.T)
	response.N = int32(handlerResponse.N)
	response.SignatureScheme = handlerResponse.SignatureScheme

	responseData, err := proto.Marshal(&response)
	if err != nil {
		logger.Errorf("(%v) Failed to marshal TpProcessResponse: %v", id, err)
		return err
	}

	// Send back a response to the validator
	err = connection.SendMsg(
		validator_pb2.Message_SMART_CONTRACT_VALIDATE_RESPONSE,
		responseData, msg.GetCorrelationId(),
	)
	if err != nil {
		logger.Errorf("(%v) Error sending TpProcessResponse: %v", id, err)
		return err
	}

	return nil
}

func execSmartContract(msg *validator_pb2.Message, handlers []TransactionHandler, connection *messaging.ZmqConnection) error {
	id := connection.Identity()
	request := &processor_pb2.TpProcessRequest{}
	err := proto.Unmarshal(msg.GetContent(), request)
	if err != nil {
		logger.Errorf(
			"(%v) Failed to unmarshal TpProcessRequest: %v", id, err,
		)
		return err
	}

	header := request.GetHeader()

	// Try to find a handler
	handler, err := findHandler(handlers, header)
	if err != nil {
		logger.Errorf("(%v) Failed to find handler: %v", id, err)
		return err
	}

	// Construct a new Context instance for the handler
	contextId := request.GetContextId()
	context := NewContext(connection, contextId)

	// Run the handler
	err = handler.Apply(request, context)

	// Process the handler response
	response := &processor_pb2.TpProcessResponse{}
	if err != nil {
		switch e := err.(type) {
		case *InvalidTransactionError:
			logger.Warnf("(%v) %v", id, e)
			response.Status = processor_pb2.TpProcessResponse_INVALID_TRANSACTION
			response.Message = e.Msg
			response.ExtendedData = e.ExtendedData
		case *InternalError:
			logger.Warnf("(%v) %v", id, e)
			response.Status = processor_pb2.TpProcessResponse_INTERNAL_ERROR
			response.Message = e.Msg
			response.ExtendedData = e.ExtendedData
		case *AuthorizationException:
			logger.Warnf("(%v) %v", id, e)
			response.Status = processor_pb2.TpProcessResponse_INVALID_TRANSACTION
			response.Message = e.Msg
			response.ExtendedData = e.ExtendedData
		default:
			logger.Errorf("(%v) Unknown error: %v", id, err)
			response.Status = processor_pb2.TpProcessResponse_INTERNAL_ERROR
			response.Message = e.Error()
		}
	} else {
		response.Status = processor_pb2.TpProcessResponse_OK
	}

	responseData, err := proto.Marshal(response)
	if err != nil {
		logger.Errorf("(%v) Failed to marshal TpProcessResponse: %v", id, err)
		return err
	}

	// Send back a response to the validator
	err = connection.SendMsg(
		validator_pb2.Message_TP_PROCESS_RESPONSE,
		responseData, msg.GetCorrelationId(),
	)
	if err != nil {
		logger.Errorf("(%v) Error sending TpProcessResponse: %v", id, err)
		return err
	}
	return nil
}

// Searches for and returns a handler that matches the header. If a suitable
// handler is not found, returns an error.
func findHandler(handlers []TransactionHandler, header *transaction_pb2.TransactionHeader) (TransactionHandler, error) {
	for _, handler := range handlers {
		if header.GetFamilyName() != handler.FamilyName() {
			continue
		}

		HeaderVersion := header.GetFamilyVersion()
		HasVersion := false
		for _, version := range handler.FamilyVersions() {
			if version == HeaderVersion {
				HasVersion = true
				break
			}
		}

		if !HasVersion {
			continue
		}

		return handler, nil
	}
	return nil, fmt.Errorf(
		"Unknown handler: (%v, %v)", header.GetFamilyName(),
		header.GetFamilyVersion(),
	)
}

// Searches for and returns a handler that matches the header. If a suitable
// handler is not found, returns an error.
func findHandlerForValidation(handlers []TransactionHandler, smartcontractAddress string) (TransactionHandler, error) {
	for _, handler := range handlers {
		if smartcontractAddress != handler.FamilyName() {
			continue
		}

		return handler, nil
	}
	return nil, fmt.Errorf(
		"Unknown handler: (%v)", smartcontractAddress,
	)
}

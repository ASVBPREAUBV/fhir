// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import "encoding/json"

type ProcedureRequest struct {
	DomainResource          `bson:",inline"`
	Identifier              []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                 *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Code                    *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	BodySite                []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	ReasonCodeableConcept   *CodeableConcept  `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference         *Reference        `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	ScheduledDateTime       *FHIRDateTime     `bson:"scheduledDateTime,omitempty" json:"scheduledDateTime,omitempty"`
	ScheduledPeriod         *Period           `bson:"scheduledPeriod,omitempty" json:"scheduledPeriod,omitempty"`
	ScheduledTiming         *Timing           `bson:"scheduledTiming,omitempty" json:"scheduledTiming,omitempty"`
	Encounter               *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Performer               *Reference        `bson:"performer,omitempty" json:"performer,omitempty"`
	Status                  string            `bson:"status,omitempty" json:"status,omitempty"`
	Notes                   []Annotation      `bson:"notes,omitempty" json:"notes,omitempty"`
	AsNeededBoolean         *bool             `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	OrderedOn               *FHIRDateTime     `bson:"orderedOn,omitempty" json:"orderedOn,omitempty"`
	Orderer                 *Reference        `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Priority                string            `bson:"priority,omitempty" json:"priority,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ProcedureRequest) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		ProcedureRequest
	}{
		ResourceType:     "ProcedureRequest",
		ProcedureRequest: *resource,
	}
	return json.Marshal(x)
}

// The "procedureRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type procedureRequest ProcedureRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ProcedureRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := procedureRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ProcedureRequest(x2)
	}
	return
}

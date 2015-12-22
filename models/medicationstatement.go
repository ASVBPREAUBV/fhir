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

type MedicationStatement struct {
	DomainResource              `bson:",inline"`
	Identifier                  []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient                     *Reference                           `bson:"patient,omitempty" json:"patient,omitempty"`
	InformationSource           *Reference                           `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	DateAsserted                *FHIRDateTime                        `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	Status                      string                               `bson:"status,omitempty" json:"status,omitempty"`
	WasNotTaken                 *bool                                `bson:"wasNotTaken,omitempty" json:"wasNotTaken,omitempty"`
	ReasonNotTaken              []CodeableConcept                    `bson:"reasonNotTaken,omitempty" json:"reasonNotTaken,omitempty"`
	ReasonForUseCodeableConcept *CodeableConcept                     `bson:"reasonForUseCodeableConcept,omitempty" json:"reasonForUseCodeableConcept,omitempty"`
	ReasonForUseReference       *Reference                           `bson:"reasonForUseReference,omitempty" json:"reasonForUseReference,omitempty"`
	EffectiveDateTime           *FHIRDateTime                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod             *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Note                        string                               `bson:"note,omitempty" json:"note,omitempty"`
	SupportingInformation       []Reference                          `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	MedicationCodeableConcept   *CodeableConcept                     `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference         *Reference                           `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Dosage                      []MedicationStatementDosageComponent `bson:"dosage,omitempty" json:"dosage,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *MedicationStatement) MarshalJSON() ([]byte, error) {
	x := struct {
		ResourceType string `json:"resourceType"`
		MedicationStatement
	}{
		ResourceType:        "MedicationStatement",
		MedicationStatement: *resource,
	}
	return json.Marshal(x)
}

// The "medicationStatement" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type medicationStatement MedicationStatement

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *MedicationStatement) UnmarshalJSON(data []byte) (err error) {
	x2 := medicationStatement{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = MedicationStatement(x2)
	}
	return
}

type MedicationStatementDosageComponent struct {
	Text                    string           `bson:"text,omitempty" json:"text,omitempty"`
	Timing                  *Timing          `bson:"timing,omitempty" json:"timing,omitempty"`
	AsNeededBoolean         *bool            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	SiteCodeableConcept     *CodeableConcept `bson:"siteCodeableConcept,omitempty" json:"siteCodeableConcept,omitempty"`
	SiteReference           *Reference       `bson:"siteReference,omitempty" json:"siteReference,omitempty"`
	Route                   *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method                  *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	QuantitySimpleQuantity  *Quantity        `bson:"quantitySimpleQuantity,omitempty" json:"quantitySimpleQuantity,omitempty"`
	QuantityRange           *Range           `bson:"quantityRange,omitempty" json:"quantityRange,omitempty"`
	RateRatio               *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateRange               *Range           `bson:"rateRange,omitempty" json:"rateRange,omitempty"`
	MaxDosePerPeriod        *Ratio           `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
}

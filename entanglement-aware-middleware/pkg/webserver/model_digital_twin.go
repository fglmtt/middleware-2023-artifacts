/*
 * webserver
 *
 * rest api
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webserver

type DigitalTwin struct {
	Name string `json:"name"`

	// provides information about where to get the digital twin to deploy, e.g., link to the container image(s) to use
	Source string `json:"source"`

	// specifies if the digital twin is simple of composed. A composed digital twin is a digital twin built by composition of one ore more simple ones
	Type string `json:"type"`

	// lists the physical counterparts, i.e., the physical twins, of the digital twin
	TwinOf []string `json:"twinOf,omitempty"`

	Requirements Requirements `json:"requirements,omitempty"`

	Deployments []Deployment `json:"deployments"`
}

// AssertDigitalTwinRequired checks if the required fields are not zero-ed
func AssertDigitalTwinRequired(obj DigitalTwin) error {
	elements := map[string]interface{}{
		"name":        obj.Name,
		"source":      obj.Source,
		"type":        obj.Type,
		"deployments": obj.Deployments,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertRequirementsRequired(obj.Requirements); err != nil {
		return err
	}
	for _, el := range obj.Deployments {
		if err := AssertDeploymentRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseDigitalTwinRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DigitalTwin (e.g. [][]DigitalTwin), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDigitalTwinRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDigitalTwin, ok := obj.(DigitalTwin)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDigitalTwinRequired(aDigitalTwin)
	})
}

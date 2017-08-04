package issue483

import (
	"testing"
)

func TestPresence(t *testing.T) {
	issue483 := new(Issue483)

    if issue483.UnsetMetaAttribute(4) != "required" {
        t.Log("Unset attribute's presence should be 'required'")
        t.Fail()
    }

    if issue483.RequiredMetaAttribute(4) != "required" {
        t.Log("Required attribute's presence should be 'required'")
        t.Fail()
    }

    if issue483.ConstantMetaAttribute(4) != "constant" {
        t.Log("Constant attribute's presence should be 'constant'")
        t.Fail()
    }

    if issue483.OptionalMetaAttribute(4) != "optional" {
        t.Log("Optional attribute's presence should be 'optional'")
        t.Fail()
    }
}

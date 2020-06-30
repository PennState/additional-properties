# additional-properties

Generates ``MarshalJSON`` and ``UnmarshalJSON`` methods to allow arbitrary
fields when marshaling and unmarshaling structs.  See the official JSON
Schema definition of ``additional properties`` at
https://json-schema.org/understanding-json-schema/reference/object.html

This attempts to fix a problem with `encoding/json` where there's no
way to retrieve arbitrary fields when unmarshaling to a `struct` as
described in https://github.com/golang/go/issues/6213.

package testserver

import "github.com/fizx/gqlgen/codegen/testserver/otherpkg"

type WrappedScalar = otherpkg.Scalar
type WrappedStruct otherpkg.Struct
type WrappedMap otherpkg.Map
type WrappedSlice otherpkg.Slice

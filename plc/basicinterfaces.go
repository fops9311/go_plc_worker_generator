package plc

type inputbool func() bool
type inputint func() int
type inputfloat64 func() float64
type inputfloat32 func() float32

type outputbool func(bool)
type outputint func(int)
type outputfloat64 func(float64)
type outputfloat32 func(float32)

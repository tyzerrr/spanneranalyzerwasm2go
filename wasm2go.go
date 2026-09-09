package wasm2go

import (
	base "github.com/tyzerrr/spanneranalyzerwasm2go/base"
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
	_ "github.com/tyzerrr/spanneranalyzerwasm2go/p11"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	__memcap := reserveBytes
	if __memcap < 39518208 {
		__memcap = 39518208
	}
	m.Memory = make([]byte, 39518208, __memcap)
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	m.MemSize.Store(39518208)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 22053)
	m.G0 = int32(33554432)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_6_0(m)
	InitElemSeg_6_1(m)
	InitElemSeg_7_0(m)
	InitElemSeg_7_1(m)
	InitElemSeg_7_2(m)
	InitElemSeg_8_0(m)
	InitElemSeg_8_1(m)
	InitElemSeg_8_2(m)
	InitElemSeg_8_3(m)
	InitElemSeg_9_0(m)
	InitElemSeg_9_1(m)
	InitElemSeg_9_2(m)
	InitElemSeg_9_3(m)
	InitElemSeg_9_4(m)
	InitElemSeg_9_5(m)
	InitElemSeg_9_6(m)
	InitElemSeg_10_0(m)
	InitElemSeg_10_1(m)
	InitElemSeg_10_2(m)
	InitElemSeg_10_3(m)
	InitElemSeg_10_4(m)
	InitElemSeg_10_5(m)
	InitElemSeg_10_6(m)
	InitElemSeg_10_7(m)
	InitElemSeg_10_8(m)
	InitElemSeg_10_9(m)
	InitElemSeg_10_10(m)
	InitElemSeg_10_11(m)
	InitElemSeg_10_12(m)
	InitElemSeg_10_13(m)
	InitElemSeg_10_14(m)
	InitElemSeg_10_15(m)
	InitElemSeg_10_16(m)
	InitElemSeg_10_17(m)
	InitElemSeg_10_18(m)
	InitElemSeg_10_19(m)
	InitElemSeg_10_20(m)
	InitElemSeg_10_21(m)
	InitElemSeg_11_0(m)
	InitElemSeg_11_1(m)
	InitElemSeg_11_2(m)
	InitElemSeg_11_3(m)
	InitElemSeg_11_4(m)
	InitElemSeg_11_5(m)
	InitElemSeg_11_6(m)
	InitElemSeg_11_7(m)
	InitElemSeg_11_8(m)
	InitElemSeg_11_9(m)
	InitElemSeg_11_10(m)
	InitElemSeg_11_11(m)
	InitElemSeg_11_12(m)
	InitElemSeg_11_13(m)
	InitElemSeg_11_14(m)
	InitElemSeg_11_15(m)
	InitElemSeg_11_16(m)
	InitElemSeg_11_17(m)
	InitElemSeg_11_18(m)
	InitElemSeg_11_19(m)
	InitElemSeg_11_20(m)
	InitElemSeg_11_21(m)
	InitElemSeg_11_22(m)
	InitElemSeg_11_23(m)
	InitElemSeg_11_24(m)
	InitElemSeg_11_25(m)
	InitElemSeg_11_26(m)
	InitElemSeg_11_27(m)
	InitElemSeg_11_28(m)
	InitElemSeg_11_29(m)
	InitElemSeg_11_30(m)
	InitElemSeg_11_31(m)
	InitElemSeg_11_32(m)
	InitElemSeg_11_33(m)
	InitElemSeg_11_34(m)
	InitElemSeg_11_35(m)
	InitElemSeg_11_36(m)
	InitElemSeg_11_37(m)
	InitElemSeg_11_38(m)
	InitElemSeg_11_39(m)
	InitElemSeg_11_40(m)
	InitElemSeg_11_41(m)
	InitElemSeg_11_42(m)
	InitElemSeg_11_43(m)
	InitElemSeg_11_44(m)
	InitElemSeg_11_45(m)
	InitElemSeg_11_46(m)
	InitElemSeg_11_47(m)
	InitElemSeg_11_48(m)
	InitElemSeg_11_49(m)
	InitElemSeg_11_50(m)
	m.DataEnd = 39387139
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, env, 49397760)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New(env base.EnvImports) *base.Module {
	return NewWithWASI(base.DefaultWASI(), env)
}

const InitialMemoryBytes = 39518208

func NewWithMemory(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 22053)
	m.G0 = int32(33554432)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_6_0(m)
	InitElemSeg_6_1(m)
	InitElemSeg_7_0(m)
	InitElemSeg_7_1(m)
	InitElemSeg_7_2(m)
	InitElemSeg_8_0(m)
	InitElemSeg_8_1(m)
	InitElemSeg_8_2(m)
	InitElemSeg_8_3(m)
	InitElemSeg_9_0(m)
	InitElemSeg_9_1(m)
	InitElemSeg_9_2(m)
	InitElemSeg_9_3(m)
	InitElemSeg_9_4(m)
	InitElemSeg_9_5(m)
	InitElemSeg_9_6(m)
	InitElemSeg_10_0(m)
	InitElemSeg_10_1(m)
	InitElemSeg_10_2(m)
	InitElemSeg_10_3(m)
	InitElemSeg_10_4(m)
	InitElemSeg_10_5(m)
	InitElemSeg_10_6(m)
	InitElemSeg_10_7(m)
	InitElemSeg_10_8(m)
	InitElemSeg_10_9(m)
	InitElemSeg_10_10(m)
	InitElemSeg_10_11(m)
	InitElemSeg_10_12(m)
	InitElemSeg_10_13(m)
	InitElemSeg_10_14(m)
	InitElemSeg_10_15(m)
	InitElemSeg_10_16(m)
	InitElemSeg_10_17(m)
	InitElemSeg_10_18(m)
	InitElemSeg_10_19(m)
	InitElemSeg_10_20(m)
	InitElemSeg_10_21(m)
	InitElemSeg_11_0(m)
	InitElemSeg_11_1(m)
	InitElemSeg_11_2(m)
	InitElemSeg_11_3(m)
	InitElemSeg_11_4(m)
	InitElemSeg_11_5(m)
	InitElemSeg_11_6(m)
	InitElemSeg_11_7(m)
	InitElemSeg_11_8(m)
	InitElemSeg_11_9(m)
	InitElemSeg_11_10(m)
	InitElemSeg_11_11(m)
	InitElemSeg_11_12(m)
	InitElemSeg_11_13(m)
	InitElemSeg_11_14(m)
	InitElemSeg_11_15(m)
	InitElemSeg_11_16(m)
	InitElemSeg_11_17(m)
	InitElemSeg_11_18(m)
	InitElemSeg_11_19(m)
	InitElemSeg_11_20(m)
	InitElemSeg_11_21(m)
	InitElemSeg_11_22(m)
	InitElemSeg_11_23(m)
	InitElemSeg_11_24(m)
	InitElemSeg_11_25(m)
	InitElemSeg_11_26(m)
	InitElemSeg_11_27(m)
	InitElemSeg_11_28(m)
	InitElemSeg_11_29(m)
	InitElemSeg_11_30(m)
	InitElemSeg_11_31(m)
	InitElemSeg_11_32(m)
	InitElemSeg_11_33(m)
	InitElemSeg_11_34(m)
	InitElemSeg_11_35(m)
	InitElemSeg_11_36(m)
	InitElemSeg_11_37(m)
	InitElemSeg_11_38(m)
	InitElemSeg_11_39(m)
	InitElemSeg_11_40(m)
	InitElemSeg_11_41(m)
	InitElemSeg_11_42(m)
	InitElemSeg_11_43(m)
	InitElemSeg_11_44(m)
	InitElemSeg_11_45(m)
	InitElemSeg_11_46(m)
	InitElemSeg_11_47(m)
	InitElemSeg_11_48(m)
	InitElemSeg_11_49(m)
	InitElemSeg_11_50(m)
	m.DataEnd = 39387139
	return m
}
func NewFromSnapshot(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, env base.EnvImports, memory []byte, memSize uint64, globals []uint64) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1, Env: env}
	m.Memory = memory
	m.MemMu = &sync.Mutex{}
	m.MemSize = &atomic.Uint64{}
	m.Threads = &base.ThreadPool{}
	if memSize > 4294836224 {
		panic("wasm2go: memory size exceeds the implementation limit (4294836224 bytes)")
	}
	m.MemSize.Store(memSize)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = uint64(len(memory))
	m.T0 = make([]any, 22053)
	m.G0 = int32(33554432)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_4_0(m)
	InitElemSeg_5_0(m)
	InitElemSeg_6_0(m)
	InitElemSeg_6_1(m)
	InitElemSeg_7_0(m)
	InitElemSeg_7_1(m)
	InitElemSeg_7_2(m)
	InitElemSeg_8_0(m)
	InitElemSeg_8_1(m)
	InitElemSeg_8_2(m)
	InitElemSeg_8_3(m)
	InitElemSeg_9_0(m)
	InitElemSeg_9_1(m)
	InitElemSeg_9_2(m)
	InitElemSeg_9_3(m)
	InitElemSeg_9_4(m)
	InitElemSeg_9_5(m)
	InitElemSeg_9_6(m)
	InitElemSeg_10_0(m)
	InitElemSeg_10_1(m)
	InitElemSeg_10_2(m)
	InitElemSeg_10_3(m)
	InitElemSeg_10_4(m)
	InitElemSeg_10_5(m)
	InitElemSeg_10_6(m)
	InitElemSeg_10_7(m)
	InitElemSeg_10_8(m)
	InitElemSeg_10_9(m)
	InitElemSeg_10_10(m)
	InitElemSeg_10_11(m)
	InitElemSeg_10_12(m)
	InitElemSeg_10_13(m)
	InitElemSeg_10_14(m)
	InitElemSeg_10_15(m)
	InitElemSeg_10_16(m)
	InitElemSeg_10_17(m)
	InitElemSeg_10_18(m)
	InitElemSeg_10_19(m)
	InitElemSeg_10_20(m)
	InitElemSeg_10_21(m)
	InitElemSeg_11_0(m)
	InitElemSeg_11_1(m)
	InitElemSeg_11_2(m)
	InitElemSeg_11_3(m)
	InitElemSeg_11_4(m)
	InitElemSeg_11_5(m)
	InitElemSeg_11_6(m)
	InitElemSeg_11_7(m)
	InitElemSeg_11_8(m)
	InitElemSeg_11_9(m)
	InitElemSeg_11_10(m)
	InitElemSeg_11_11(m)
	InitElemSeg_11_12(m)
	InitElemSeg_11_13(m)
	InitElemSeg_11_14(m)
	InitElemSeg_11_15(m)
	InitElemSeg_11_16(m)
	InitElemSeg_11_17(m)
	InitElemSeg_11_18(m)
	InitElemSeg_11_19(m)
	InitElemSeg_11_20(m)
	InitElemSeg_11_21(m)
	InitElemSeg_11_22(m)
	InitElemSeg_11_23(m)
	InitElemSeg_11_24(m)
	InitElemSeg_11_25(m)
	InitElemSeg_11_26(m)
	InitElemSeg_11_27(m)
	InitElemSeg_11_28(m)
	InitElemSeg_11_29(m)
	InitElemSeg_11_30(m)
	InitElemSeg_11_31(m)
	InitElemSeg_11_32(m)
	InitElemSeg_11_33(m)
	InitElemSeg_11_34(m)
	InitElemSeg_11_35(m)
	InitElemSeg_11_36(m)
	InitElemSeg_11_37(m)
	InitElemSeg_11_38(m)
	InitElemSeg_11_39(m)
	InitElemSeg_11_40(m)
	InitElemSeg_11_41(m)
	InitElemSeg_11_42(m)
	InitElemSeg_11_43(m)
	InitElemSeg_11_44(m)
	InitElemSeg_11_45(m)
	InitElemSeg_11_46(m)
	InitElemSeg_11_47(m)
	InitElemSeg_11_48(m)
	InitElemSeg_11_49(m)
	InitElemSeg_11_50(m)
	m.DataEnd = 39387139
	base.RestoreGlobals(m, globals)
	return m
}
func initData_0(m *base.Module) {
	copy(m.Memory[33554432:], wasm2goData_data_bin[0:4617718])
	copy(m.Memory[38173888:], wasm2goData_data_bin[4617718:5830969])
}
func Initialize(m *base.Module) {
	Fn61(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn62(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn63(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn85(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn87(m)
}
func WasmShutdown(m *base.Module) {
	Fn88(m)
}
func Inv_0_0(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn64(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			if trapErr, trapIsErr := r.(error); trapIsErr {
				err = fmt.Errorf("wasm trap: %w", trapErr)
			} else {
				err = fmt.Errorf("wasm trap: %v", r)
			}
		}
	}()
	packed = Fn84(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte

package asm

import . "github.com/FMNSSun/bootstrap/defs"

import "strings"
import "strconv"
import "io"
import "bufio"

func EmitLDC(memory []uint8, op uint8, val uint32) []uint8 {
	memory[0] = op
	memory[1] = uint8((val >> 24) & 0xFF)
	memory[2] = uint8((val >> 16) & 0xFF)
	memory[3] = uint8((val >> 8) & 0xFF)
	memory[4] = uint8(val & 0xFF)
	return memory[5:]
}

func EmitOP(memory []uint8, op uint8, dst uint8, src uint8) []uint8 {
	memory[0] = op
	memory[1] = (dst << 4) | src
	return memory[2:]
}

func Str2Reg(i string) uint8 {
	switch i {
	case "ra":
		return REG_A
	case "rb":
		return REG_B
	case "rc":
		return REG_C
	}
	
	panic("no such reg")
}

func Str2Op(i string) uint8 {
	switch i{
	case "add":
		return OP_ADD
	case "sub":
		return OP_SUB
	case "div":
		return OP_DIV
	case "mul":
		return OP_MUL
	case "nop":
		return OP_NOP
	case "hlt":
		return OP_HLT
	case "ldcc":
		return OP_LDCC
	case "ldcb":
		return OP_LDCB
	case "ldca":
		return OP_LDCA
	case "dec":
		return OP_DEC
	case "inc":
		return OP_INC
	}
	
	panic("no such op")
}

func AsmReader(memory []uint8, rd io.Reader) []uint8 {
	sc := bufio.NewScanner(rd)
	
	for sc.Scan() {
		ln := sc.Text()
		memory = Asm(memory, ln)
	}
	
	return memory
}

func AsmLns(memory []uint8, i []string) []uint8 {
	for _,v := range(i) {
		memory = Asm(memory, v)
	}
	
	return memory
}

func Asm(memory []uint8, i string) []uint8 {
	flds := strings.Fields(i)
	
	if len(flds) == 0 {
		return nil
	}
	
	if len(flds) == 1 {
		switch flds[0] {
		case "nop":
			return EmitOP(memory, OP_NOP, 0, 0)
		case "hlt":
			return EmitOP(memory, OP_HLT, 0, 0)
		default:
			panic("can't handle this")
		}
	}
	
	if len(flds) == 2 {
		switch flds[0] {
		case "ldca","ldcb","ldcc":
			val, _ := strconv.ParseUint(flds[1], 0, 32)
			return EmitLDC(memory, Str2Op(flds[0]), uint32(val))
		case "inc","dec":
			return EmitOP(memory, Str2Op(flds[0]), Str2Reg(flds[1]), 0)
		default:
			panic("can't handle this")
		}
	}
	
	if len(flds) == 3 {
		return EmitOP(memory, Str2Op(flds[0]), Str2Reg(flds[1]), Str2Reg(flds[2]))
	}
	
	return nil
} 
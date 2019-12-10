package cpu

import . "github.com/FMNSSun/bootstrap/defs"
import "fmt"

func Execute(regs []uint32, memory []uint8) {
	
	for {
		if regs[REG_IP] >= uint32(len(memory)) {
			return
		}
		
		i := memory[regs[REG_IP]]
		regs[REG_IP]++
		
		if i <= OP_LDCC {
			// fetch uint32
			if regs[REG_IP]+3 >= uint32(len(memory)) {
				return
			}
			c0 := memory[regs[REG_IP]]
			regs[REG_IP]++
			
			c1 := memory[regs[REG_IP]]
			regs[REG_IP]++
			
			c2 := memory[regs[REG_IP]]
			regs[REG_IP]++
			
			c3 := memory[regs[REG_IP]]
			regs[REG_IP]++
			
			value := uint32(c0) << 24 | uint32(c1) << 16 | uint32(c2) << 8 | uint32(c3)
			
			switch i {
			case OP_LDCA:
				regs[REG_A] = value
			case OP_LDCB:
				regs[REG_B] = value
			case OP_LDCC:
				regs[REG_C] = value
			default:
				return
			}
		}else {
			if regs[REG_IP] >= uint32(len(memory)) {
				return
			}
			
			b1 := memory[regs[REG_IP]]
			regs[REG_IP]++
			
			src := b1 & 0x0F
			dst := (b1 >> 4) & 0x0F
		
			switch i {
			case OP_NOP:
				/* nop */
			case OP_XCHG:
				t := regs[dst]
				regs[dst] = regs[src]
				regs[src] = t
			case OP_ADD:
				regs[dst] += regs[src]
			case OP_SUB:
				regs[dst] -= regs[src]
			case OP_MUL:
				regs[dst] *= regs[src]
			case OP_DIV:
				regs[dst] /= regs[src]
			case OP_INC:
				fmt.Printf("inc %x\n", dst)
				regs[dst]++
			case OP_DEC:
				regs[dst]--
			case OP_HLT:
				return
			default:
				return
			}
		}
	}
}
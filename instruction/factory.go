package instruction

import (
	"fmt"
	"h-jvm/instruction/base"
	. "h-jvm/instruction/compare"
	. "h-jvm/instruction/constant"
	. "h-jvm/instruction/control"
	. "h-jvm/instruction/convert"
	. "h-jvm/instruction/extend"
	. "h-jvm/instruction/load"
	. "h-jvm/instruction/math"
	. "h-jvm/instruction/reference"
	. "h-jvm/instruction/stack"
	. "h-jvm/instruction/store"
)

var (
	nop         = &Nop{}
	aconst_null = &AConstNull{}
	iconst_m1   = &IConstM1{}
	iconst_0    = &IConst0{}
	iconst_1    = &IConst1{}
	iconst_2    = &IConst2{}
	iconst_3    = &IConst3{}
	iconst_4    = &IConst4{}
	iconst_5    = &IConst5{}
	lconst_0    = &LConst0{}
	lconst_1    = &LConst1{}
	fconst_0    = &FConst0{}
	fconst_1    = &FConst1{}
	fconst_2    = &FConst2{}
	dconst_0    = &DConst0{}
	dconst_1    = &DConst1{}
	iload_0     = &ILoad0{}
	iload_1     = &ILoad1{}
	iload_2     = &ILoad2{}
	iload_3     = &ILoad3{}
	lload_0     = &LLoad0{}
	lload_1     = &LLoad1{}
	lload_2     = &LLoad2{}
	lload_3     = &LLoad3{}
	fload_0     = &FLoad0{}
	fload_1     = &FLoad1{}
	fload_2     = &FLoad2{}
	fload_3     = &FLoad3{}
	dload_0     = &DLoad0{}
	dload_1     = &DLoad1{}
	dload_2     = &DLoad2{}
	dload_3     = &DLoad3{}
	aload_0     = &ALoad0{}
	aload_1     = &ALoad1{}
	aload_2     = &ALoad2{}
	aload_3     = &ALoad3{}
	// iaload      =
	// laload      =
	// faload      =
	// daload      =
	// aaload      =
	// baload      =
	// caload      =
	// saload      =
	istore_0 = &IStore0{}
	istore_1 = &IStore1{}
	istore_2 = &IStore2{}
	istore_3 = &IStore3{}
	lstore_0 = &LStore0{}
	lstore_1 = &LStore1{}
	lstore_2 = &LStore2{}
	lstore_3 = &LStore3{}
	fstore_0 = &FStore0{}
	fstore_1 = &FStore1{}
	fstore_2 = &FStore2{}
	fstore_3 = &FStore3{}
	dstore_0 = &DStore0{}
	dstore_1 = &DStore1{}
	dstore_2 = &DStore2{}
	dstore_3 = &DStore3{}
	astore_0 = &AStore0{}
	astore_1 = &AStore1{}
	astore_2 = &AStore2{}
	astore_3 = &AStore3{}
	// iastore  =
	// lastore  =
	// fastore  =
	// dastore  =
	// aastore  =
	// bastore  =
	// castore  =
	// sastore  =
	pop     = &Pop{}
	pop2    = &Pop2{}
	dup     = &Dup{}
	dup_x1  = &DupX1{}
	dup_x2  = &DupX2{}
	dup2    = &Dup2{}
	dup2_x1 = &Dup2X1{}
	dup2_x2 = &Dup2X2{}
	swap    = &Swap{}
	iadd    = &IAdd{}
	ladd    = &LAdd{}
	fadd    = &FAdd{}
	dadd    = &DAdd{}
	isub    = &ISub{}
	lsub    = &LSub{}
	fsub    = &FSub{}
	dsub    = &DSub{}
	imul    = &IMul{}
	lmul    = &LMul{}
	fmul    = &FMul{}
	dmul    = &DMul{}
	idiv    = &IDiv{}
	ldiv    = &LDiv{}
	fdiv    = &FDiv{}
	ddiv    = &DDiv{}
	irem    = &IRem{}
	lrem    = &LRem{}
	frem    = &FRem{}
	drem    = &DRem{}
	ineg    = &INeg{}
	lneg    = &LNeg{}
	fneg    = &FNeg{}
	dneg    = &DNeg{}
	ishl    = &IShl{}
	lshl    = &LShl{}
	ishr    = &IShr{}
	lshr    = &LShr{}
	iushr   = &IUShr{}
	lushr   = &LUShr{}
	iand    = &IAnd{}
	land    = &LAnd{}
	ior     = &IOr{}
	lor     = &LOr{}
	ixor    = &IXor{}
	lxor    = &LXor{}
	i2l     = &I2L{}
	i2f     = &I2F{}
	i2d     = &I2D{}
	l2i     = &L2I{}
	l2f     = &L2F{}
	l2d     = &L2D{}
	f2i     = &F2I{}
	f2l     = &F2L{}
	f2d     = &F2D{}
	d2i     = &D2I{}
	d2l     = &D2L{}
	d2f     = &D2F{}
	i2b     = &I2B{}
	i2c     = &I2C{}
	i2s     = &I2S{}
	lcmp    = &LCmp{}
	fcmpl   = &FCmpL{}
	fcmpg   = &FCmpG{}
	dcmpl   = &DCmpL{}
	dcmpg   = &DCmpG{}
	// ireturn =
	// lreturn =
	// freturn =
	// dreturn =
	// areturn =
	// _return =
	// arraylength   =
	// athrow        =
	// monitorenter  =
	// monitorexit   =
	// invoke_native =
	//

)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return &BIPush{}
	case 0x11:
		return &SIPush{}

	case 0x12:
		return &Ldc{}
	case 0x13:
		return &LdcW{}
	case 0x14:
		return &Ldc2W{}

	case 0x15:
		return &ILoad{}
	case 0x16:
		return &LLoad{}
	case 0x17:
		return &FLoad{}
	case 0x18:
		return &DLoad{}
	case 0x19:
		return &ALoad{}
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	// 数组aload
	case 0x2e:
		return &IALoad{}
	case 0x2f:
		return &LALoad{}
	case 0x30:
		return &FALoad{}
	case 0x31:
		return &DALoad{}
	case 0x32:
		return &AALoad{}
	case 0x33:
		return &BALoad{}
	case 0x34:
		return &CALoad{}
	case 0x35:
		return &SALoad{}


	case 0x36:
		return &IStore{}
	case 0x37:
		return &LStore{}
	case 0x38:
		return &FStore{}
	case 0x39:
		return &DStore{}
	case 0x3a:
		return &AStore{}
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3

	// 数组的astore
	case 0x4f:
		return &IAStore{}
	case 0x50:
		return &LAStore{}
	case 0x51:
		return &FAStore{}
	case 0x52:
		return &DAStore{}
	case 0x53:
		return &AAStore{}
	case 0x54:
		return &BAStore{}
	case 0x55:
		return &CAStore{}
	case 0x56:
		return &SAStore{}

	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &IInc{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return &IFeq{}
	case 0x9a:
		return &IFne{}
	case 0x9b:
		return &IFlt{}
	case 0x9c:
		return &IFge{}
	case 0x9d:
		return &IFgt{}
	case 0x9e:
		return &IFle{}
	case 0x9f:
		return &IFICmpEq{}
	case 0xa0:
		return &IFICmpNe{}
	case 0xa1:
		return &IFICmpLt{}
	case 0xa2:
		return &IFICmpGe{}
	case 0xa3:
		return &IFICmpGt{}
	case 0xa4:
		return &IFICmpLe{}
	case 0xa5:
		return &IFAcmEq{}
	case 0xa6:
		return &IFAcmNe{}
	case 0xa7:
		return &Goto{}

	case 0xaa:
		return &TableSwitch{}
	case 0xab:
		return &LookSwitch{}

	// return
	case 0xac:
		return &IReturn{}
	case 0xad:
		return &LReturn{}
	case 0xae:
		return &FReturn{}
	case 0xaf:
		return &DReturn{}
	case 0xb0:
		return &AReturn{}
	case 0xb1:
		return &Return{}

	case 0xb2:
		return &GetStatic{}
	case 0xb3:
		return &PutStatic{}
	case 0xb4:
		return &GetField{}
	case 0xb5:
		return &PutField{}


	case 0xb6:
		return &InvokeVirtual{}
	case 0xb7:
		return &InvokeSpecial{}
	case 0xb8:
		return &InvokeStatic{}
	case 0xb9:
		return &InvokeInterface{}


	case 0xbb:
		return &New{}
	// 创建数组指令
	case 0xbc:
		return &NewArray{}
	case 0xbd:
		return &ANewArray{}
	case 0xbe:
		return &ArrayLength{}

	case 0xc0:
		return &CheckCast{}
	case 0xc1:
		return &InstanceOf{}

	case 0xc6:
		return &IFNull{}
	case 0xc7:
		return &IFNonNull{}
	case 0xc8:
		return &GotoW{}

	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}

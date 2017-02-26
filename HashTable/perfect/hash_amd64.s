#include "textflag.h"

//func Hash(seed uint32, str string) uint32
TEXT ·MurmurHash(SB),NOSPLIT,$0
	MOVL		seed+0(FP),	AX
	MOVQ		str+8(FP),	SI
	MOVQ		str_len+16(FP), CX
	MOVQ		CX, R8
	
Lwordloop:
	SUBQ		$4, CX
	JB		Ltail
	MOVL		(SI), DX
	ADDQ		$4, SI
	IMULL	$0xcc9e2d51, DX
	ROLL		$15, DX
	IMULL	$0x1b873593, DX
	XORL		DX, AX
	ROLL		$13, AX
	LEAL		0xe6546b64(AX)(AX*4), AX
	JMP		Lwordloop
	
Ltail:
	CMPQ		CX, $-3
	JZ		L1byte
	CMPQ		CX, $-2
	JZ		L2byte
	CMPQ		CX, $-1
	JZ		L3byte
	
Lfinal:
	XORL		R8, AX
	MOVL		AX, DX
	SHRL		$16, DX
	XORL		DX, AX
	IMULL	$0x85ebca6b, AX
	MOVL		AX, DX
	SHRL		$13, DX
	XORL		DX, AX
	IMULL	$0xc2b2ae35, AX
	MOVL		AX, DX
	SHRL		$16, DX
	XORL		DX, AX
	
	MOVL		AX, 0x20(SP)
	RET
	
L1byte:
	XORL		DX, DX
	MOVB		(SI), DL
	JMP		Llast
L2byte:
	XORL		DX, DX
	MOVW		(SI), DX
	JMP		Llast
L3byte:
	XORL		DX, DX
	XORL		CX, CX
	MOVW		(SI), DX
	MOVB		2(SI), CL
	SHLL		$16, CX
	XORL		CX, DX
	
Llast:
	IMULL	$0xcc9e2d51, DX
	ROLL		$15, DX
	IMULL	$0x1b873593, DX
	XORL		DX, AX
	JMP		Lfinal

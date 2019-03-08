package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import(
	"unsafe"
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
)

// CUDA handle for shiftfudgey kernel
var shiftfudgey_code cu.Function

// Stores the arguments for shiftfudgey kernel invocation
type shiftfudgey_args_t struct{
	 arg_dst unsafe.Pointer
	 arg_src unsafe.Pointer
	 arg_Nx int
	 arg_Ny int
	 arg_Nz int
	 arg_shy int
	 argptr [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for shiftfudgey kernel invocation
var shiftfudgey_args shiftfudgey_args_t

func init(){
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	 shiftfudgey_args.argptr[0] = unsafe.Pointer(&shiftfudgey_args.arg_dst)
	 shiftfudgey_args.argptr[1] = unsafe.Pointer(&shiftfudgey_args.arg_src)
	 shiftfudgey_args.argptr[2] = unsafe.Pointer(&shiftfudgey_args.arg_Nx)
	 shiftfudgey_args.argptr[3] = unsafe.Pointer(&shiftfudgey_args.arg_Ny)
	 shiftfudgey_args.argptr[4] = unsafe.Pointer(&shiftfudgey_args.arg_Nz)
	 shiftfudgey_args.argptr[5] = unsafe.Pointer(&shiftfudgey_args.arg_shy)
	 }

// Wrapper for shiftfudgey CUDA kernel, asynchronous.
func k_shiftfudgey_async ( dst unsafe.Pointer, src unsafe.Pointer, Nx int, Ny int, Nz int, shy int,  cfg *config) {
	if Synchronous{ // debug
		Sync()
		timer.Start("shiftfudgey")
	}

	shiftfudgey_args.Lock()
	defer shiftfudgey_args.Unlock()

	if shiftfudgey_code == 0{
		shiftfudgey_code = fatbinLoad(shiftfudgey_map, "shiftfudgey")
	}

	 shiftfudgey_args.arg_dst = dst
	 shiftfudgey_args.arg_src = src
	 shiftfudgey_args.arg_Nx = Nx
	 shiftfudgey_args.arg_Ny = Ny
	 shiftfudgey_args.arg_Nz = Nz
	 shiftfudgey_args.arg_shy = shy
	

	args := shiftfudgey_args.argptr[:]
	cu.LaunchKernel(shiftfudgey_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous{ // debug
		Sync()
		timer.Stop("shiftfudgey")
	}
}

// maps compute capability on PTX code for shiftfudgey kernel.
var shiftfudgey_map = map[int]string{ 0: "" ,
30: shiftfudgey_ptx_30 ,
35: shiftfudgey_ptx_35 ,
37: shiftfudgey_ptx_37 ,
50: shiftfudgey_ptx_50 ,
52: shiftfudgey_ptx_52 ,
53: shiftfudgey_ptx_53 ,
60: shiftfudgey_ptx_60 ,
61: shiftfudgey_ptx_61 ,
70: shiftfudgey_ptx_70 ,
75: shiftfudgey_ptx_75  }

// shiftfudgey PTX code for various compute capabilities.
const(
  shiftfudgey_ptx_30 = `
.version 6.3
.target sm_30
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_35 = `
.version 6.3
.target sm_35
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_37 = `
.version 6.3
.target sm_37
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_50 = `
.version 6.3
.target sm_50
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_52 = `
.version 6.3
.target sm_52
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_53 = `
.version 6.3
.target sm_53
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_60 = `
.version 6.3
.target sm_60
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_61 = `
.version 6.3
.target sm_61
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_70 = `
.version 6.3
.target sm_70
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
   shiftfudgey_ptx_75 = `
.version 6.3
.target sm_75
.address_size 64

	// .globl	shiftfudgey

.visible .entry shiftfudgey(
	.param .u64 shiftfudgey_param_0,
	.param .u64 shiftfudgey_param_1,
	.param .u32 shiftfudgey_param_2,
	.param .u32 shiftfudgey_param_3,
	.param .u32 shiftfudgey_param_4,
	.param .u32 shiftfudgey_param_5
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<2>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<9>;


	ld.param.u64 	%rd1, [shiftfudgey_param_0];
	ld.param.u64 	%rd2, [shiftfudgey_param_1];
	ld.param.u32 	%r4, [shiftfudgey_param_2];
	ld.param.u32 	%r5, [shiftfudgey_param_3];
	ld.param.u32 	%r7, [shiftfudgey_param_4];
	ld.param.u32 	%r6, [shiftfudgey_param_5];
	mov.u32 	%r8, %ctaid.x;
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %tid.x;
	mad.lo.s32 	%r1, %r9, %r8, %r10;
	mov.u32 	%r11, %ntid.y;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %tid.y;
	mad.lo.s32 	%r2, %r11, %r12, %r13;
	mov.u32 	%r14, %ntid.z;
	mov.u32 	%r15, %ctaid.z;
	mov.u32 	%r16, %tid.z;
	mad.lo.s32 	%r3, %r14, %r15, %r16;
	setp.lt.s32	%p1, %r1, %r4;
	setp.lt.s32	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32	%p4, %r3, %r7;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	cvta.to.global.u64 	%rd3, %rd2;
	sub.s32 	%r17, %r2, %r6;
	setp.lt.s32	%p6, %r17, 0;
	setp.lt.s32	%p7, %r17, %r5;
	add.s32 	%r18, %r5, -1;
	selp.b32	%r19, %r17, %r18, %p7;
	selp.b32	%r20, 0, %r19, %p6;
	mul.lo.s32 	%r21, %r3, %r5;
	add.s32 	%r22, %r21, %r20;
	mad.lo.s32 	%r23, %r22, %r4, %r1;
	mul.wide.s32 	%rd4, %r23, 4;
	add.s64 	%rd5, %rd3, %rd4;
	ld.global.nc.f32 	%f1, [%rd5];
	add.s32 	%r24, %r21, %r2;
	mad.lo.s32 	%r25, %r24, %r4, %r1;
	cvta.to.global.u64 	%rd6, %rd1;
	mul.wide.s32 	%rd7, %r25, 4;
	add.s64 	%rd8, %rd6, %rd7;
	st.global.f32 	[%rd8], %f1;

BB0_2:
	ret;
}


`
 )

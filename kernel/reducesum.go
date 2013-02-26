package kernel

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var reducesum_code cu.Function

type reducesum_args struct {
	arg_src     unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [4]unsafe.Pointer
}

// Wrapper for reducesum CUDA kernel. Synchronizes before return.
func K_reducesum(src unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, gridDim, blockDim cu.Dim3) {
	if reducesum_code == 0 {
		reducesum_code = cu.ModuleLoadData(reducesum_ptx).GetFunction("reducesum")
	}

	var a reducesum_args

	a.arg_src = src
	a.argptr[0] = unsafe.Pointer(&a.arg_src)
	a.arg_dst = dst
	a.argptr[1] = unsafe.Pointer(&a.arg_dst)
	a.arg_initVal = initVal
	a.argptr[2] = unsafe.Pointer(&a.arg_initVal)
	a.arg_n = n
	a.argptr[3] = unsafe.Pointer(&a.arg_n)

	args := a.argptr[:]
	str := Stream()
	cu.LaunchKernel(reducesum_code, gridDim.X, gridDim.Y, gridDim.Z, blockDim.X, blockDim.Y, blockDim.Z, 0, str, args)
	SyncAndRecycle(str)
}

const reducesum_ptx = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry reducesum(
	.param .u64 reducesum_param_0,
	.param .u64 reducesum_param_1,
	.param .f32 reducesum_param_2,
	.param .u32 reducesum_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<38>;
	.reg .f32 	%f<30>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 __cuda_local_var_33845_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducesum_param_0];
	ld.param.u64 	%rd5, [reducesum_param_1];
	ld.param.f32 	%f29, [reducesum_param_2];
	ld.param.u32 	%r9, [reducesum_param_3];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 8 1
	mov.u32 	%r37, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r36, %r37, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r37, %r11;
	.loc 2 8 1
	setp.ge.s32 	%p1, %r36, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 2 8 1
	mul.wide.s32 	%rd6, %r36, 4;
	add.s64 	%rd7, %rd2, %rd6;
	ld.global.f32 	%f5, [%rd7];
	add.f32 	%f29, %f29, %f5;
	add.s32 	%r36, %r36, %r4;
	.loc 2 8 1
	setp.lt.s32 	%p2, %r36, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 2 8 1
	mul.wide.s32 	%rd8, %r2, 4;
	mov.u64 	%rd9, __cuda_local_var_33845_32_non_const_sdata;
	add.s64 	%rd3, %rd9, %rd8;
	st.shared.f32 	[%rd3], %f29;
	bar.sync 	0;
	.loc 2 8 1
	setp.lt.u32 	%p3, %r37, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 2 8 1
	mov.u32 	%r7, %r37;
	shr.u32 	%r37, %r7, 1;
	.loc 2 8 1
	setp.ge.u32 	%p4, %r2, %r37;
	@%p4 bra 	BB0_5;

	.loc 2 8 1
	ld.shared.f32 	%f6, [%rd3];
	add.s32 	%r15, %r37, %r2;
	mul.wide.u32 	%rd10, %r15, 4;
	add.s64 	%rd12, %rd9, %rd10;
	ld.shared.f32 	%f7, [%rd12];
	add.f32 	%f8, %f6, %f7;
	st.shared.f32 	[%rd3], %f8;

BB0_5:
	.loc 2 8 1
	bar.sync 	0;
	.loc 2 8 1
	setp.gt.u32 	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 2 8 1
	setp.gt.s32 	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 2 8 1
	ld.volatile.shared.f32 	%f9, [%rd3];
	ld.volatile.shared.f32 	%f10, [%rd3+128];
	add.f32 	%f11, %f9, %f10;
	st.volatile.shared.f32 	[%rd3], %f11;
	ld.volatile.shared.f32 	%f12, [%rd3+64];
	ld.volatile.shared.f32 	%f13, [%rd3];
	add.f32 	%f14, %f13, %f12;
	st.volatile.shared.f32 	[%rd3], %f14;
	ld.volatile.shared.f32 	%f15, [%rd3+32];
	ld.volatile.shared.f32 	%f16, [%rd3];
	add.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd3], %f17;
	ld.volatile.shared.f32 	%f18, [%rd3+16];
	ld.volatile.shared.f32 	%f19, [%rd3];
	add.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd3], %f20;
	ld.volatile.shared.f32 	%f21, [%rd3+8];
	ld.volatile.shared.f32 	%f22, [%rd3];
	add.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd3], %f23;
	ld.volatile.shared.f32 	%f24, [%rd3+4];
	ld.volatile.shared.f32 	%f25, [%rd3];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd3], %f26;

BB0_8:
	.loc 2 8 1
	setp.ne.s32 	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 2 8 1
	ld.shared.f32 	%f27, [__cuda_local_var_33845_32_non_const_sdata];
	.loc 3 1844 5
	atom.global.add.f32 	%f28, [%rd1], %f27;

BB0_10:
	.loc 2 9 2
	ret;
}


`

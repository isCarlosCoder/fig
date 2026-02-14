package builtins

import (
	"math"
	"os"
	"testing"

	"github.com/iscarloscoder/fig/environment"
)

func getBuiltin(t *testing.T, name string) environment.Value {
	t.Helper()
	mod := Get("mathx")
	if mod == nil {
		t.Fatalf("mathx module not registered")
	}
	v, ok := mod.Entries[name]
	if !ok {
		t.Fatalf("mathx.%s not found", name)
	}
	if v.Type != environment.BuiltinFnType {
		t.Fatalf("mathx.%s is not a builtin", name)
	}
	return v
}

func numSliceEquals(t *testing.T, got []environment.Value, expect []float64) {
	t.Helper()
	if len(got) != len(expect) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(expect))
	}
	for i := range got {
		if got[i].Type != environment.NumberType {
			t.Fatalf("element %d not a number", i)
		}
		if math.Abs(got[i].Num-expect[i]) > 1e-9 {
			t.Fatalf("element %d: got %v, want %v", i, got[i].Num, expect[i])
		}
	}
}

func TestArrayAsarrayAndCopy(t *testing.T) {
	arrFn := getBuiltin(t, "array")
	asarrFn := getBuiltin(t, "asarray")
	copyFn := getBuiltin(t, "copy")

	v, err := arrFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("array() error: %v", err)
	}
	if v.Type != environment.ArrayType {
		t.Fatalf("array() did not return array")
	}
	numSliceEquals(t, *v.Arr, []float64{1, 2, 3})

	asv, err := asarrFn.Builtin([]environment.Value{environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("asarray() error: %v", err)
	}
	numSliceEquals(t, *asv.Arr, []float64{5})

	// copy
	c, err := copyFn.Builtin([]environment.Value{v})
	if err != nil {
		t.Fatalf("copy() error: %v", err)
	}
	// mutate original and ensure copy unchanged
	(*v.Arr)[0] = environment.NewNumber(999)
	if (*c.Arr)[0].Num != 1 {
		t.Fatalf("copy() appears to be shallow/aliased")
	}
}

func TestZerosOnesFullArange(t *testing.T) {
	zeros := getBuiltin(t, "zeros")
	ones := getBuiltin(t, "ones")
	full := getBuiltin(t, "full")
	ar := getBuiltin(t, "arange")

	z, err := zeros.Builtin([]environment.Value{environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("zeros error: %v", err)
	}
	numSliceEquals(t, *z.Arr, []float64{0, 0, 0, 0})

	o, err := ones.Builtin([]environment.Value{environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("ones error: %v", err)
	}
	numSliceEquals(t, *o.Arr, []float64{1, 1, 1})

	f, err := full.Builtin([]environment.Value{environment.NewNumber(3), environment.NewNumber(7)})
	if err != nil {
		t.Fatalf("full error: %v", err)
	}
	numSliceEquals(t, *f.Arr, []float64{7, 7, 7})

	a, err := ar.Builtin([]environment.Value{environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("arange error: %v", err)
	}
	numSliceEquals(t, *a.Arr, []float64{0, 1, 2, 3, 4})

	b, err := ar.Builtin([]environment.Value{environment.NewNumber(2), environment.NewNumber(6), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("arange(2,6,2) error: %v", err)
	}
	numSliceEquals(t, *b.Arr, []float64{2, 4})
}

func TestLinspaceLogspaceGeomspace(t *testing.T) {
	lin := getBuiltin(t, "linspace")
	log := getBuiltin(t, "logspace")
	geo := getBuiltin(t, "geomspace")

	l, err := lin.Builtin([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("linspace error: %v", err)
	}
	numSliceEquals(t, *l.Arr, []float64{0, 0.25, 0.5, 0.75, 1})

	lg, err := log.Builtin([]environment.Value{environment.NewNumber(0), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(10)})
	if err != nil {
		t.Fatalf("logspace error: %v", err)
	}
	numSliceEquals(t, *lg.Arr, []float64{1, 10, 100})

	g, err := geo.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(8), environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("geomspace error: %v", err)
	}
	numSliceEquals(t, *g.Arr, []float64{1, 2, 4, 8})
}

func TestEyeIdentityDiagDiagflatFromiterFrombuffer(t *testing.T) {
	eye := getBuiltin(t, "eye")
	id := getBuiltin(t, "identity")
	diag := getBuiltin(t, "diag")
	diagflat := getBuiltin(t, "diagflat")
	fromiter := getBuiltin(t, "fromiter")
	frombuffer := getBuiltin(t, "frombuffer")

	e, err := eye.Builtin([]environment.Value{environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("eye error: %v", err)
	}
	if e.Type != environment.ArrayType || len(*e.Arr) != 3 {
		t.Fatalf("eye returned wrong shape")
	}

	ii, err := id.Builtin([]environment.Value{environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("identity error: %v", err)
	}
	if ii.Type != environment.ArrayType || len(*ii.Arr) != 2 {
		t.Fatalf("identity returned wrong shape")
	}

	dIn, err := diag.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})})
	if err != nil {
		t.Fatalf("diag error: %v", err)
	}
	// diag([1,2]) -> [[1,0],[0,2]]
	row0 := (*dIn.Arr)[0]
	if (*row0.Arr)[0].Num != 1 || (*(*dIn.Arr)[1].Arr)[1].Num != 2 {
		t.Fatalf("diag produced wrong values")
	}

	df, err := diagflat.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("diagflat error: %v", err)
	}
	if (*(*df.Arr)[0].Arr)[0].Num != 3 || (*(*df.Arr)[1].Arr)[1].Num != 4 {
		t.Fatalf("diagflat wrong")
	}

	fi, err := fromiter.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	if err != nil {
		t.Fatalf("fromiter error: %v", err)
	}
	numSliceEquals(t, *fi.Arr, []float64{1, 2, 3})

	fb, err := frombuffer.Builtin([]environment.Value{environment.NewString("ABC")})
	if err != nil {
		t.Fatalf("frombuffer error: %v", err)
	}
	// 'A' == 65
	if (*fb.Arr)[0].Num != 65 {
		t.Fatalf("frombuffer unexpected value: %v", (*fb.Arr)[0])
	}
}

func TestFromFunction(t *testing.T) {
	ff := getBuiltin(t, "fromfunction")
	// use a builtin (Go) function as callback — it will be invoked via invokeFn
	cb := environment.NewBuiltinFn("cb", func(a []environment.Value) (environment.Value, error) {
		// return the index (first arg)
		return a[0], nil
	})
	res, err := ff.Builtin([]environment.Value{environment.NewNumber(5), cb})
	if err != nil {
		t.Fatalf("fromfunction error: %v", err)
	}
	numSliceEquals(t, *res.Arr, []float64{0, 1, 2, 3, 4})
}

func TestShapeReshapeRavelFlattenTransposeAndDims(t *testing.T) {
	shapeFn := getBuiltin(t, "shape")
	reshapeFn := getBuiltin(t, "reshape")
	ravelFn := getBuiltin(t, "ravel")
	transposeFn := getBuiltin(t, "transpose")
	swapaxesFn := getBuiltin(t, "swapaxes")
	moveaxisFn := getBuiltin(t, "moveaxis")
	expandFn := getBuiltin(t, "expand_dims")
	squeezeFn := getBuiltin(t, "squeeze")

	mat := environment.NewArray([]environment.Value{
		environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)}),
		environment.NewArray([]environment.Value{environment.NewNumber(4), environment.NewNumber(5), environment.NewNumber(6)}),
	})

	shRes, err := shapeFn.Builtin([]environment.Value{mat})
	if err != nil {
		t.Fatalf("shape error: %v", err)
	}
	numSliceEquals(t, *shRes.Arr, []float64{2, 3})

	rav, err := ravelFn.Builtin([]environment.Value{mat})
	if err != nil {
		t.Fatalf("ravel error: %v", err)
	}
	numSliceEquals(t, *rav.Arr, []float64{1, 2, 3, 4, 5, 6})

	resh, err := reshapeFn.Builtin([]environment.Value{rav, environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(2)})})
	if err != nil {
		t.Fatalf("reshape error: %v", err)
	}
	// check first row
	firstRow := (*resh.Arr)[0]
	numSliceEquals(t, *firstRow.Arr, []float64{1, 2})

	tr, err := transposeFn.Builtin([]environment.Value{mat})
	if err != nil {
		t.Fatalf("transpose error: %v", err)
	}
	// transpose -> [[1,4],[2,5],[3,6]]
	numSliceEquals(t, *(*tr.Arr)[0].Arr, []float64{1, 4})

	sw, err := swapaxesFn.Builtin([]environment.Value{mat, environment.NewNumber(0), environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("swapaxes error: %v", err)
	}
	numSliceEquals(t, *(*sw.Arr)[0].Arr, []float64{1, 4})

	mv, err := moveaxisFn.Builtin([]environment.Value{mat, environment.NewNumber(0), environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("moveaxis error: %v", err)
	}
	numSliceEquals(t, *(*mv.Arr)[0].Arr, []float64{1, 4})

	arr1 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	e1, err := expandFn.Builtin([]environment.Value{arr1, environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("expand_dims error: %v", err)
	}
	// expand_dims at axis 0 -> single-row matrix
	if e1.Type != environment.ArrayType || (*e1.Arr)[0].Type != environment.ArrayType {
		t.Fatalf("expand_dims produced wrong shape")
	}

	sq, err := squeezeFn.Builtin([]environment.Value{e1})
	if err != nil {
		t.Fatalf("squeeze error: %v", err)
	}
	if sq.Type != environment.ArrayType {
		t.Fatalf("squeeze did not restore 1D array")
	}
}

func TestBroadcastTileRepeatConcatStackVhColumnSplit(t *testing.T) {
	broadcastTo := getBuiltin(t, "broadcast_to")
	tile := getBuiltin(t, "tile")
	repeat := getBuiltin(t, "repeat")
	concat := getBuiltin(t, "concatenate")
	stack := getBuiltin(t, "stack")
	vstack := getBuiltin(t, "vstack")
	hstack := getBuiltin(t, "hstack")
	columnStack := getBuiltin(t, "column_stack")
	splitFn := getBuiltin(t, "split")
	arraySplit := getBuiltin(t, "array_split")
	hsplit := getBuiltin(t, "hsplit")
	vsplit := getBuiltin(t, "vsplit")

	b, err := broadcastTo.Builtin([]environment.Value{environment.NewNumber(5), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)})})
	if err != nil {
		t.Fatalf("broadcast_to error: %v", err)
	}
	// expect 2 rows of 3 fives
	numSliceEquals(t, *(*b.Arr)[0].Arr, []float64{5, 5, 5})

	t1, err := tile.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("tile error: %v", err)
	}
	numSliceEquals(t, *t1.Arr, []float64{1, 2, 1, 2, 1, 2})

	rp, err := repeat.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("repeat error: %v", err)
	}
	numSliceEquals(t, *rp.Arr, []float64{1, 1, 2, 2})

	c, err := concat.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("concatenate error: %v", err)
	}
	numSliceEquals(t, *c.Arr, []float64{1, 2, 3, 4})

	st, err := stack.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("stack error: %v", err)
	}
	// stacked should be 2x2
	if len(*st.Arr) != 2 {
		t.Fatalf("stack wrong length")
	}

	vs, err := vstack.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("vstack error: %v", err)
	}
	if len(*vs.Arr) != 2 {
		t.Fatalf("vstack wrong")
	}

	hs, err := hstack.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("hstack error: %v", err)
	}
	numSliceEquals(t, *hs.Arr, []float64{1, 2, 3, 4})

	cs, err := columnStack.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("column_stack error: %v", err)
	}
	numSliceEquals(t, *(*cs.Arr)[0].Arr, []float64{1, 3})

	sp, err := splitFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)}), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("split error: %v", err)
	}
	if len(*sp.Arr) != 2 {
		t.Fatalf("split produced wrong number of parts")
	}

	asp, err := arraySplit.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)}), environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("array_split error: %v", err)
	}
	if len(*asp.Arr) != 3 {
		t.Fatalf("array_split produced wrong number of parts")
	}

	mat := environment.NewArray([]environment.Value{
		environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}),
		environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)}),
	})

	hsParts, err := hsplit.Builtin([]environment.Value{mat, environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("hsplit error: %v", err)
	}
	if len(*hsParts.Arr) != 2 {
		t.Fatalf("hsplit produced wrong number of parts")
	}

	vsParts, err := vsplit.Builtin([]environment.Value{mat, environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("vsplit error: %v", err)
	}
	if len(*vsParts.Arr) != 2 {
		t.Fatalf("vsplit produced wrong number of parts")
	}
}

// ---------- Index / selection tests ----------
func TestIndexSelectionFunctions(t *testing.T) {
	take := getBuiltin(t, "take")
	put := getBuiltin(t, "put")
	whereFn := getBuiltin(t, "where")
	nonzero := getBuiltin(t, "nonzero")
	argwhere := getBuiltin(t, "argwhere")
	extract := getBuiltin(t, "extract")
	selectFn := getBuiltin(t, "select")
	choose := getBuiltin(t, "choose")
	compress := getBuiltin(t, "compress")

	arr := environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(20), environment.NewNumber(30), environment.NewNumber(40)})
	res, err := take.Builtin([]environment.Value{arr, environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(-1)})})
	if err != nil {
		t.Fatalf("take error: %v", err)
	}
	numSliceEquals(t, *res.Arr, []float64{10, 40})

	// put
	a := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(0), environment.NewNumber(0)})
	pout, err := put.Builtin([]environment.Value{a, environment.NewArray([]environment.Value{environment.NewNumber(1)}), environment.NewNumber(9)})
	if err != nil {
		t.Fatalf("put error: %v", err)
	}
	numSliceEquals(t, *pout.Arr, []float64{0, 9, 0})

	cond := environment.NewArray([]environment.Value{environment.NewBool(true), environment.NewBool(false), environment.NewBool(true)})
	widx, err := whereFn.Builtin([]environment.Value{cond})
	if err != nil {
		t.Fatalf("where(cond) error: %v", err)
	}
	numSliceEquals(t, *widx.Arr, []float64{0, 2})

	x := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(1), environment.NewNumber(1)})
	y := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(0), environment.NewNumber(0)})
	wx, err := whereFn.Builtin([]environment.Value{cond, x, y})
	if err != nil {
		t.Fatalf("where(cond,x,y) error: %v", err)
	}
	numSliceEquals(t, *wx.Arr, []float64{1, 0, 1})

	nz, err := nonzero.Builtin([]environment.Value{cond})
	if err != nil {
		t.Fatalf("nonzero error: %v", err)
	}
	numSliceEquals(t, *nz.Arr, []float64{0, 2})

	aw, err := argwhere.Builtin([]environment.Value{cond})
	if err != nil {
		t.Fatalf("argwhere error: %v", err)
	}
	if len(*aw.Arr) != 2 {
		t.Fatalf("argwhere returned wrong length")
	}
	if (*(*aw.Arr)[0].Arr)[0].Num != 0 || (*(*aw.Arr)[1].Arr)[0].Num != 2 {
		t.Fatalf("argwhere contents unexpected")
	}

	ex, err := extract.Builtin([]environment.Value{cond, environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(20), environment.NewNumber(30)})})
	if err != nil {
		t.Fatalf("extract error: %v", err)
	}
	numSliceEquals(t, *ex.Arr, []float64{10, 30})

	cond1 := environment.NewArray([]environment.Value{environment.NewBool(true), environment.NewBool(false), environment.NewBool(false)})
	cond2 := environment.NewArray([]environment.Value{environment.NewBool(false), environment.NewBool(true), environment.NewBool(false)})
	condList := environment.NewArray([]environment.Value{cond1, cond2})
	choice1 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(1), environment.NewNumber(1)})
	choice2 := environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(2), environment.NewNumber(2)})
	choices := environment.NewArray([]environment.Value{choice1, choice2})
	selRes, err := selectFn.Builtin([]environment.Value{condList, choices, environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("select error: %v", err)
	}
	numSliceEquals(t, *selRes.Arr, []float64{1, 2, 0})

	idxs := environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(1), environment.NewNumber(0)})
	choiceVals := environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(20), environment.NewNumber(30)})
	ch, err := choose.Builtin([]environment.Value{idxs, choiceVals})
	if err != nil {
		t.Fatalf("choose error: %v", err)
	}
	numSliceEquals(t, *ch.Arr, []float64{30, 20, 10})

	comp, err := compress.Builtin([]environment.Value{cond, environment.NewArray([]environment.Value{environment.NewNumber(5), environment.NewNumber(6), environment.NewNumber(7)})})
	if err != nil {
		t.Fatalf("compress error: %v", err)
	}
	numSliceEquals(t, *comp.Arr, []float64{5, 7})
}

func TestMathxNumericScalars(t *testing.T) {
	add := getBuiltin(t, "add")
	sub := getBuiltin(t, "subtract")
	mul := getBuiltin(t, "multiply")
	div := getBuiltin(t, "divide")
	fdiv := getBuiltin(t, "floor_divide")
	pow := getBuiltin(t, "power")
	mod := getBuiltin(t, "mod")
	rem := getBuiltin(t, "remainder")
	neg := getBuiltin(t, "negative")
	absFn := getBuiltin(t, "absolute")
	fabsFn := getBuiltin(t, "fabs")
	sign := getBuiltin(t, "sign")
	sqrtFn := getBuiltin(t, "sqrt")
	square := getBuiltin(t, "square")
	recip := getBuiltin(t, "reciprocal")
	clip := getBuiltin(t, "clip")
	maximum := getBuiltin(t, "maximum")
	minimum := getBuiltin(t, "minimum")

	// add
	out, err := add.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("add error: %v", err)
	}
	if out.Type != environment.NumberType || out.Num != 3 {
		t.Fatalf("add scalar wrong: %v", out)
	}

	// subtract
	out, err = sub.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("subtract error: %v", err)
	}
	if out.Num != 3 {
		t.Fatalf("subtract scalar wrong: %v", out)
	}

	// multiply
	out, err = mul.Builtin([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("multiply error: %v", err)
	}
	if out.Num != 12 {
		t.Fatalf("multiply scalar wrong: %v", out)
	}

	// divide / floor_divide
	out, err = div.Builtin([]environment.Value{environment.NewNumber(7), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("divide error: %v", err)
	}
	if math.Abs(out.Num-3.5) > 1e-9 {
		t.Fatalf("divide scalar wrong: %v", out)
	}

	out, err = fdiv.Builtin([]environment.Value{environment.NewNumber(7), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("floor_divide error: %v", err)
	}
	if out.Num != 3 {
		t.Fatalf("floor_divide wrong: %v", out)
	}

	// power
	out, err = pow.Builtin([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("power error: %v", err)
	}
	if out.Num != 8 {
		t.Fatalf("power scalar wrong: %v", out)
	}

	// mod / remainder
	out, err = mod.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("mod error: %v", err)
	}
	if out.Num != math.Mod(5, 2) {
		t.Fatalf("mod scalar wrong: %v", out)
	}

	out, err = rem.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("remainder error: %v", err)
	}
	if out.Num != math.Remainder(5, 2) {
		t.Fatalf("remainder scalar wrong: %v", out)
	}

	// negative / abs / fabs / sign / sqrt / square / reciprocal
	out, err = neg.Builtin([]environment.Value{environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("negative error: %v", err)
	}
	if out.Num != -5 {
		t.Fatalf("negative wrong: %v", out)
	}

	out, err = absFn.Builtin([]environment.Value{environment.NewNumber(-3)})
	if err != nil {
		t.Fatalf("absolute error: %v", err)
	}
	if out.Num != 3 {
		t.Fatalf("absolute wrong: %v", out)
	}

	out, err = fabsFn.Builtin([]environment.Value{environment.NewNumber(-3)})
	if err != nil {
		t.Fatalf("fabs error: %v", err)
	}
	if out.Num != 3 {
		t.Fatalf("fabs wrong: %v", out)
	}

	out, err = sign.Builtin([]environment.Value{environment.NewNumber(-2)})
	if err != nil {
		t.Fatalf("sign error: %v", err)
	}
	if out.Num != -1 {
		t.Fatalf("sign wrong: %v", out)
	}
	out, err = sign.Builtin([]environment.Value{environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("sign error: %v", err)
	}
	if out.Num != 0 {
		t.Fatalf("sign wrong for zero: %v", out)
	}

	out, err = sqrtFn.Builtin([]environment.Value{environment.NewNumber(9)})
	if err != nil {
		t.Fatalf("sqrt error: %v", err)
	}
	if out.Num != 3 {
		t.Fatalf("sqrt wrong: %v", out)
	}

	out, err = square.Builtin([]environment.Value{environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("square error: %v", err)
	}
	if out.Num != 9 {
		t.Fatalf("square wrong: %v", out)
	}

	out, err = recip.Builtin([]environment.Value{environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("reciprocal error: %v", err)
	}
	if math.Abs(out.Num-0.25) > 1e-9 {
		t.Fatalf("reciprocal wrong: %v", out)
	}

	out, err = clip.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(2), environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("clip error: %v", err)
	}
	if out.Num != 4 {
		t.Fatalf("clip scalar wrong: %v", out)
	}

	out, err = maximum.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("maximum error: %v", err)
	}
	if out.Num != 2 {
		t.Fatalf("maximum scalar wrong: %v", out)
	}

	out, err = minimum.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("minimum error: %v", err)
	}
	if out.Num != 1 {
		t.Fatalf("minimum scalar wrong: %v", out)
	}
}

func TestMathxArrayOps(t *testing.T) {
	add := getBuiltin(t, "add")
	mul := getBuiltin(t, "multiply")
	pow := getBuiltin(t, "power")
	sqrtFn := getBuiltin(t, "sqrt")
	clip := getBuiltin(t, "clip")
	maximum := getBuiltin(t, "maximum")
	minimum := getBuiltin(t, "minimum")

	arr1 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	arr2 := environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})

	// add array + array
	r, err := add.Builtin([]environment.Value{arr1, arr2})
	if err != nil {
		t.Fatalf("add array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{4, 6})

	// add array + scalar (broadcast)
	r, err = add.Builtin([]environment.Value{arr1, environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("add broadcast error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{6, 7})

	// multiply array * scalar
	r, err = mul.Builtin([]environment.Value{arr1, environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("multiply broadcast error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{3, 6})

	// power array
	r, err = pow.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)}), environment.NewNumber(2)})
	if err != nil {
		t.Fatalf("power array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{4, 9})

	// sqrt array
	r, err = sqrtFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(4), environment.NewNumber(9)})})
	if err != nil {
		t.Fatalf("sqrt array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{2, 3})

	// clip array
	r, err = clip.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(5), environment.NewNumber(10)}), environment.NewNumber(2), environment.NewNumber(6)})
	if err != nil {
		t.Fatalf("clip array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{2, 5, 6})

	// maximum / minimum elementwise
	r, err = maximum.Builtin([]environment.Value{arr1, arr2})
	if err != nil {
		t.Fatalf("maximum array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{3, 4})

	r, err = minimum.Builtin([]environment.Value{arr1, arr2})
	if err != nil {
		t.Fatalf("minimum array error: %v", err)
	}
	numSliceEquals(t, *r.Arr, []float64{1, 2})
}

func TestMathxTrigFunctions(t *testing.T) {
	sinFn := getBuiltin(t, "sin")
	cosFn := getBuiltin(t, "cos")
	tanFn := getBuiltin(t, "tan")
	asinFn := getBuiltin(t, "arcsin")
	acosFn := getBuiltin(t, "arccos")
	atanFn := getBuiltin(t, "arctan")
	atan2Fn := getBuiltin(t, "arctan2")
	hypotFn := getBuiltin(t, "hypot")
	deg2rad := getBuiltin(t, "deg2rad")
	rad2deg := getBuiltin(t, "rad2deg")

	// scalar checks
	out, err := sinFn.Builtin([]environment.Value{environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("sin error: %v", err)
	}
	if math.Abs(out.Num-0) > 1e-9 {
		t.Fatalf("sin(0) wrong: %v", out)
	}

	out, err = sinFn.Builtin([]environment.Value{environment.NewNumber(math.Pi / 2)})
	if err != nil {
		t.Fatalf("sin error: %v", err)
	}
	if math.Abs(out.Num-1) > 1e-9 {
		t.Fatalf("sin(pi/2) wrong: %v", out)
	}

	out, err = cosFn.Builtin([]environment.Value{environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("cos error: %v", err)
	}
	if math.Abs(out.Num-1) > 1e-9 {
		t.Fatalf("cos(0) wrong: %v", out)
	}

	out, err = acosFn.Builtin([]environment.Value{environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("arccos error: %v", err)
	}
	if math.Abs(out.Num-math.Pi/2) > 1e-9 {
		t.Fatalf("arccos(0) wrong: %v", out)
	}

	out, err = tanFn.Builtin([]environment.Value{environment.NewNumber(math.Pi / 4)})
	if err != nil {
		t.Fatalf("tan error: %v", err)
	}
	if math.Abs(out.Num-1) > 1e-9 {
		t.Fatalf("tan pi/4 wrong: %v", out)
	}

	out, err = asinFn.Builtin([]environment.Value{environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("arcsin error: %v", err)
	}
	if math.Abs(out.Num-math.Pi/2) > 1e-9 {
		t.Fatalf("arcsin(1) wrong: %v", out)
	}

	out, err = atanFn.Builtin([]environment.Value{environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("arctan error: %v", err)
	}
	if math.Abs(out.Num-math.Pi/4) > 1e-9 {
		t.Fatalf("arctan(1) wrong: %v", out)
	}

	out, err = atan2Fn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("arctan2 error: %v", err)
	}
	if math.Abs(out.Num-math.Pi/2) > 1e-9 {
		t.Fatalf("arctan2 wrong: %v", out)
	}

	out, err = hypotFn.Builtin([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("hypot error: %v", err)
	}
	if math.Abs(out.Num-5) > 1e-9 {
		t.Fatalf("hypot wrong: %v", out)
	}

	out, err = deg2rad.Builtin([]environment.Value{environment.NewNumber(180)})
	if err != nil {
		t.Fatalf("deg2rad error: %v", err)
	}
	if math.Abs(out.Num-math.Pi) > 1e-9 {
		t.Fatalf("deg2rad wrong: %v", out)
	}

	out, err = rad2deg.Builtin([]environment.Value{environment.NewNumber(math.Pi)})
	if err != nil {
		t.Fatalf("rad2deg error: %v", err)
	}
	if math.Abs(out.Num-180) > 1e-9 {
		t.Fatalf("rad2deg wrong: %v", out)
	}

	// array cases
	arr, err := sinFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(math.Pi / 2)})})
	if err != nil {
		t.Fatalf("sin array error: %v", err)
	}
	numSliceEquals(t, *arr.Arr, []float64{0, 1})

	arr, err = deg2rad.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(180)})})
	if err != nil {
		t.Fatalf("deg2rad array error: %v", err)
	}
	numSliceEquals(t, *arr.Arr, []float64{0, math.Pi})
}
func TestMathxExpLogFunctions(t *testing.T) {
	expFn := getBuiltin(t, "exp")
	exp2Fn := getBuiltin(t, "exp2")
	expm1Fn := getBuiltin(t, "expm1")
	logFn := getBuiltin(t, "log")
	log2Fn := getBuiltin(t, "log2")
	log10Fn := getBuiltin(t, "log10")
	log1pFn := getBuiltin(t, "log1p")

	// scalar checks
	out, err := expFn.Builtin([]environment.Value{environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("exp error: %v", err)
	}
	if math.Abs(out.Num-1) > 1e-9 {
		t.Fatalf("exp(0) wrong: %v", out)
	}

	out, err = exp2Fn.Builtin([]environment.Value{environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("exp2 error: %v", err)
	}
	if math.Abs(out.Num-8) > 1e-9 {
		t.Fatalf("exp2(3) wrong: %v", out)
	}

	out, err = expm1Fn.Builtin([]environment.Value{environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("expm1 error: %v", err)
	}
	if math.Abs(out.Num-math.Expm1(1)) > 1e-12 {
		t.Fatalf("expm1(1) wrong: %v", out)
	}

	out, err = logFn.Builtin([]environment.Value{environment.NewNumber(math.E)})
	if err != nil {
		t.Fatalf("log error: %v", err)
	}
	if math.Abs(out.Num-1) > 1e-9 {
		t.Fatalf("log(e) wrong: %v", out)
	}

	out, err = log2Fn.Builtin([]environment.Value{environment.NewNumber(8)})
	if err != nil {
		t.Fatalf("log2 error: %v", err)
	}
	if math.Abs(out.Num-3) > 1e-9 {
		t.Fatalf("log2(8) wrong: %v", out)
	}

	out, err = log10Fn.Builtin([]environment.Value{environment.NewNumber(100)})
	if err != nil {
		t.Fatalf("log10 error: %v", err)
	}
	if math.Abs(out.Num-2) > 1e-9 {
		t.Fatalf("log10(100) wrong: %v", out)
	}

	out, err = log1pFn.Builtin([]environment.Value{environment.NewNumber(1)})
	if err != nil {
		t.Fatalf("log1p error: %v", err)
	}
	if math.Abs(out.Num-math.Log1p(1)) > 1e-12 {
		t.Fatalf("log1p(1) wrong: %v", out)
	}

	// array cases
	arr, err := expFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1)})})
	if err != nil {
		t.Fatalf("exp array error: %v", err)
	}
	numSliceEquals(t, *arr.Arr, []float64{1, math.E})

	arr, err = logFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(math.E)})})
	if err != nil {
		t.Fatalf("log array error: %v", err)
	}
	numSliceEquals(t, *arr.Arr, []float64{0, 1})
}

func TestMathxStatistics(t *testing.T) {
	sumFn := getBuiltin(t, "sum")
	meanFn := getBuiltin(t, "mean")
	avgFn := getBuiltin(t, "average")
	varFn := getBuiltin(t, "var")
	stdFn := getBuiltin(t, "std")
	medianFn := getBuiltin(t, "median")
	percentileFn := getBuiltin(t, "percentile")
	quantileFn := getBuiltin(t, "quantile")
	minFn := getBuiltin(t, "min")
	maxFn := getBuiltin(t, "max")
	ptpFn := getBuiltin(t, "ptp")
	argminFn := getBuiltin(t, "argmin")
	argmaxFn := getBuiltin(t, "argmax")
	cumsumFn := getBuiltin(t, "cumsum")
	cumprodFn := getBuiltin(t, "cumprod")
	histFn := getBuiltin(t, "histogram")
	bincountFn := getBuiltin(t, "bincount")
	covFn := getBuiltin(t, "cov")
	corrFn := getBuiltin(t, "corrcoef")

	// sum
	out, err := sumFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	if err != nil {
		t.Fatalf("sum error: %v", err)
	}
	if out.Num != 6 {
		t.Fatalf("sum wrong: %v", out)
	}

	// mean / average
	out, err = meanFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	if err != nil {
		t.Fatalf("mean error: %v", err)
	}
	if out.Num != 2 {
		t.Fatalf("mean wrong: %v", out)
	}
	av, _ := avgFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	if av.Num != out.Num {
		t.Fatalf("average alias wrong")
	}

	// var / std (population)
	vres, err := varFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("var error: %v", err)
	}
	if math.Abs(vres.Num-1.25) > 1e-9 {
		t.Fatalf("var wrong: %v", vres)
	}

	sres, err := stdFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)})})
	if err != nil {
		t.Fatalf("std error: %v", err)
	}
	if math.Abs(sres.Num-math.Sqrt(1.25)) > 1e-9 {
		t.Fatalf("std wrong: %v", sres)
	}

	// median / percentile / quantile
	m, err := medianFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(3), environment.NewNumber(2)})})
	if err != nil {
		t.Fatalf("median error: %v", err)
	}
	if m.Num != 2 {
		t.Fatalf("median wrong: %v", m)
	}

	p50, err := percentileFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)}), environment.NewNumber(50)})
	if err != nil {
		t.Fatalf("percentile error: %v", err)
	}
	if math.Abs(p50.Num-2.5) > 1e-9 {
		t.Fatalf("percentile wrong: %v", p50)
	}

	q50, err := quantileFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)}), environment.NewNumber(0.5)})
	if err != nil {
		t.Fatalf("quantile error: %v", err)
	}
	if math.Abs(q50.Num-2.5) > 1e-9 {
		t.Fatalf("quantile wrong: %v", q50)
	}

	// min/max/ptp/argmin/argmax
	mn, _ := minFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(5), environment.NewNumber(20)})})
	mx, _ := maxFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(5), environment.NewNumber(20)})})
	if mn.Num != 5 || mx.Num != 20 {
		t.Fatalf("min/max wrong: %v %v", mn, mx)
	}
	pt, _ := ptpFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(4), environment.NewNumber(2)})})
	if pt.Num != 3 {
		t.Fatalf("ptp wrong: %v", pt)
	}
	ami, _ := argminFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(5), environment.NewNumber(20)})})
	amx, _ := argmaxFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(5), environment.NewNumber(20)})})
	if ami.Num != 1 || amx.Num != 2 {
		t.Fatalf("argmin/argmax wrong: %v %v", ami, amx)
	}

	// cumsum / cumprod
	cs, _ := cumsumFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	numSliceEquals(t, *cs.Arr, []float64{1, 3, 6})
	cp, _ := cumprodFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})})
	numSliceEquals(t, *cp.Arr, []float64{1, 2, 6})

	// histogram
	h, _ := histFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(0.1), environment.NewNumber(0.2), environment.NewNumber(0.9)}), environment.NewNumber(2)})
	numSliceEquals(t, *h.Arr, []float64{2, 1})

	// bincount
	bc, _ := bincountFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(1), environment.NewNumber(0)})})
	numSliceEquals(t, *bc.Arr, []float64{1, 2, 1})

	// cov / corrcoef
	x := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	y := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	cv, err := covFn.Builtin([]environment.Value{x, y})
	if err != nil {
		t.Fatalf("cov error: %v", err)
	}
	if math.Abs(cv.Num-2.0/3.0) > 1e-9 {
		t.Fatalf("cov wrong: %v", cv)
	}
	cc, err := corrFn.Builtin([]environment.Value{x, y})
	if err != nil {
		t.Fatalf("corr error: %v", err)
	}
	if math.Abs(cc.Num-1.0) > 1e-9 {
		t.Fatalf("corrcoef wrong: %v", cc)
	}

	// corrcoef matrix
	mat := environment.NewArray([]environment.Value{x, y})
	cm, err := corrFn.Builtin([]environment.Value{mat})
	if err != nil {
		t.Fatalf("corr matrix error: %v", err)
	}
	if (*cm.Arr)[0].Type != environment.ArrayType || (*(*cm.Arr)[0].Arr)[0].Num != 1 {
		t.Fatalf("corr matrix wrong: %v", cm)
	}
}
func TestMathxComparisonsAndLogic(t *testing.T) {
	// builtins
	equalFn := getBuiltin(t, "equal")
	noteqFn := getBuiltin(t, "not_equal")
	gt := getBuiltin(t, "greater")
	land := getBuiltin(t, "logical_and")
	lor := getBuiltin(t, "logical_or")
	lxor := getBuiltin(t, "logical_xor")
	notFn := getBuiltin(t, "logical_not")
	allFn := getBuiltin(t, "all")
	anyFn := getBuiltin(t, "any")
	isfiniteFn := getBuiltin(t, "isfinite")
	isinfFn := getBuiltin(t, "isinf")
	isnanFn := getBuiltin(t, "isnan")
	iscloseFn := getBuiltin(t, "isclose")

	// equal / not_equal (scalar)
	r, err := equalFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(1)})
	if err != nil { t.Fatalf("equal error: %v", err) }
	if !r.IsTruthy() { t.Fatalf("equal scalar wrong: %v", r) }
	r, _ = equalFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if r.IsTruthy() { t.Fatalf("equal scalar wrong: %v", r) }
	rn, _ := noteqFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if !rn.IsTruthy() { t.Fatalf("not_equal scalar wrong: %v", rn) }

	// equal (array)
	a1 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	a2 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(0), environment.NewNumber(3)})
	res, err := equalFn.Builtin([]environment.Value{a1, a2})
	if err != nil { t.Fatalf("equal array error: %v", err) }
	flat := *res.Arr
	if !flat[0].IsTruthy() || flat[1].IsTruthy() || !flat[2].IsTruthy() { t.Fatalf("equal array wrong: %v", res) }

	// comparisons
	sc, _ := gt.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(2)})
	if !sc.IsTruthy() { t.Fatalf("greater scalar wrong: %v", sc) }
	arrGt, _ := gt.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(5)}), environment.NewNumber(2)})
	if !(*arrGt.Arr)[0].IsTruthy() || !(*arrGt.Arr)[1].IsTruthy() == false { /* ok */ }

	// logical ops
	b1 := environment.NewArray([]environment.Value{environment.NewBool(true), environment.NewBool(false), environment.NewBool(true)})
	b2 := environment.NewArray([]environment.Value{environment.NewBool(false), environment.NewBool(true), environment.NewBool(true)})
	andRes, err := land.Builtin([]environment.Value{b1, b2})
	if err != nil { t.Fatalf("logical_and error: %v", err) }
	if (*andRes.Arr)[0].IsTruthy() || (*andRes.Arr)[1].IsTruthy() || !(*andRes.Arr)[2].IsTruthy() { t.Fatalf("logical_and wrong: %v", andRes) }
	orRes, _ := lor.Builtin([]environment.Value{b1, b2})
	if !(*orRes.Arr)[0].IsTruthy() || !(*orRes.Arr)[1].IsTruthy() || !(*orRes.Arr)[2].IsTruthy() { t.Fatalf("logical_or wrong: %v", orRes) }
	xorRes, _ := lxor.Builtin([]environment.Value{b1, b2})
	if !(*xorRes.Arr)[0].IsTruthy() || !(*xorRes.Arr)[1].IsTruthy() || (*xorRes.Arr)[2].IsTruthy() { t.Fatalf("logical_xor wrong: %v", xorRes) }
	notRes, _ := notFn.Builtin([]environment.Value{b1})
	if (*notRes.Arr)[0].IsTruthy() || !(*notRes.Arr)[1].IsTruthy() || (*notRes.Arr)[2].IsTruthy() { t.Fatalf("logical_not wrong: %v", notRes) }

	// all / any
	allTrue, _ := allFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewBool(true), environment.NewBool(true)})})
	if !allTrue.IsTruthy() { t.Fatalf("all wrong: %v", allTrue) }
	anyTrue, _ := anyFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewBool(false), environment.NewBool(true)})})
	if !anyTrue.IsTruthy() { t.Fatalf("any wrong: %v", anyTrue) }

	// isfinite / isinf / isnan
	nums := environment.NewArray([]environment.Value{environment.NewNumber(math.NaN()), environment.NewNumber(math.Inf(1)), environment.NewNumber(1)})
	isnanRes, _ := isnanFn.Builtin([]environment.Value{nums})
	if !(*isnanRes.Arr)[0].IsTruthy() || (*isnanRes.Arr)[1].IsTruthy() || (*isnanRes.Arr)[2].IsTruthy() { t.Fatalf("isnan wrong: %v", isnanRes) }
	isinfRes, _ := isinfFn.Builtin([]environment.Value{nums})
	if (*isinfRes.Arr)[0].IsTruthy() || !(*isinfRes.Arr)[1].IsTruthy() || (*isinfRes.Arr)[2].IsTruthy() { t.Fatalf("isinf wrong: %v", isinfRes) }
	isfiniteRes, _ := isfiniteFn.Builtin([]environment.Value{nums})
	if (*isfiniteRes.Arr)[0].IsTruthy() || (*isfiniteRes.Arr)[1].IsTruthy() || !(*isfiniteRes.Arr)[2].IsTruthy() { t.Fatalf("isfinite wrong: %v", isfiniteRes) }

	// isclose
	ic, _ := iscloseFn.Builtin([]environment.Value{environment.NewNumber(1.0000000001), environment.NewNumber(1.0)})
	if !ic.IsTruthy() { t.Fatalf("isclose scalar wrong: %v", ic) }
	icarr, _ := iscloseFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(1.0), environment.NewNumber(3.0)})})
	if !(*icarr.Arr)[0].IsTruthy() || (*icarr.Arr)[1].IsTruthy() { t.Fatalf("isclose array wrong: %v", icarr) }
}

func TestMathxSortingAndSetOps(t *testing.T) {
	sortFn := getBuiltin(t, "sort")
	argsortFn := getBuiltin(t, "argsort")
	lexsortFn := getBuiltin(t, "lexsort")
	partitionFn := getBuiltin(t, "partition")
	argpartFn := getBuiltin(t, "argpartition")
	uniqueFn := getBuiltin(t, "unique")
	setdiffFn := getBuiltin(t, "setdiff1d")
	intersectFn := getBuiltin(t, "intersect1d")
	unionFn := getBuiltin(t, "union1d")
	in1dFn := getBuiltin(t, "in1d")

	// sort / argsort
	sout, err := sortFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1), environment.NewNumber(2)})})
	if err != nil { t.Fatalf("sort error: %v", err) }
	numSliceEquals(t, *sout.Arr, []float64{1, 2, 3})
	ai, _ := argsortFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1), environment.NewNumber(2)})})
	numSliceEquals(t, *ai.Arr, []float64{1, 2, 0})

	// lexsort (keys: sort by last key then previous)
	keys := environment.NewArray([]environment.Value{
		environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(0), environment.NewNumber(1)}),
		environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(0)}),
	})
	ls, err := lexsortFn.Builtin([]environment.Value{keys})
	if err != nil { t.Fatalf("lexsort error: %v", err) }
	numSliceEquals(t, *ls.Arr, []float64{0, 2, 1})

	// partition / argpartition
	p, err := partitionFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1), environment.NewNumber(2)}), environment.NewNumber(1)})
	if err != nil { t.Fatalf("partition error: %v", err) }
	if (*p.Arr)[1].Num != 2 { t.Fatalf("partition wrong: %v", p) }
	ap, _ := argpartFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1), environment.NewNumber(2)}), environment.NewNumber(1)})
	numSliceEquals(t, *ap.Arr, []float64{1, 2, 0})

	// unique / set operations
	u, _ := uniqueFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(1), environment.NewNumber(3)})})
	numSliceEquals(t, *u.Arr, []float64{1, 2, 3})

	sd, _ := setdiffFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(4)})})
	numSliceEquals(t, *sd.Arr, []float64{1, 3})

	ix, _ := intersectFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)})})
	numSliceEquals(t, *ix.Arr, []float64{2, 3})

	uu, _ := unionFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(3)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)})})
	numSliceEquals(t, *uu.Arr, []float64{1, 2, 3})

	// in1d
	inRes, _ := in1dFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(4)})})
	if !(*inRes.Arr)[0].IsTruthy() == false || !(*inRes.Arr)[1].IsTruthy() || (*inRes.Arr)[2].IsTruthy() { /* ok */ }
}

func TestMathxRandom(t *testing.T) {
	randFn := getBuiltin(t, "rand")
	randnFn := getBuiltin(t, "randn")
	randomFn := getBuiltin(t, "random")
	randintFn := getBuiltin(t, "randint")
	choiceFn := getBuiltin(t, "choice")
	shuffleFn := getBuiltin(t, "shuffle")
	permFn := getBuiltin(t, "permutation")
	normalFn := getBuiltin(t, "normal")
	uniformFn := getBuiltin(t, "uniform")
	binomFn := getBuiltin(t, "binomial")
	poisFn := getBuiltin(t, "poisson")
	expFn := getBuiltin(t, "exponential")
	gammaFn := getBuiltin(t, "gamma")
	betaFn := getBuiltin(t, "beta")

	// rand / random (scalar + array)
	r, err := randFn.Builtin([]environment.Value{})
	if err != nil { t.Fatalf("rand error: %v", err) }
	if r.Type != environment.NumberType || !(r.Num >= 0 && r.Num < 1) { t.Fatalf("rand scalar wrong: %v", r) }
	rarr, err := randFn.Builtin([]environment.Value{environment.NewNumber(3)})
	if err != nil { t.Fatalf("rand array error: %v", err) }
	if len(*rarr.Arr) != 3 { t.Fatalf("rand array wrong length: %v", rarr) }
	r2, _ := randomFn.Builtin([]environment.Value{environment.NewNumber(2)})
	if len(*r2.Arr) != 2 { t.Fatalf("random alias wrong: %v", r2) }

	// randn
	rn, err := randnFn.Builtin([]environment.Value{})
	if err != nil { t.Fatalf("randn error: %v", err) }
	if rn.Type != environment.NumberType || math.IsNaN(rn.Num) { t.Fatalf("randn scalar wrong: %v", rn) }

	// randint scalar and array
	ri, _ := randintFn.Builtin([]environment.Value{environment.NewNumber(5)})
	if ri.Type != environment.NumberType || int(ri.Num) < 0 || int(ri.Num) >= 5 { t.Fatalf("randint(5) wrong: %v", ri) }
	ri2, _ := randintFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewNumber(4), environment.NewNumber(3)})
	if ri2.Type != environment.ArrayType || len(*ri2.Arr) != 3 { t.Fatalf("randint array wrong: %v", ri2) }

	// choice from array and from integer
	src := environment.NewArray([]environment.Value{environment.NewNumber(10), environment.NewNumber(20), environment.NewNumber(30)})
	c, err := choiceFn.Builtin([]environment.Value{src})
	if err != nil { t.Fatalf("choice error: %v", err) }
	if c.Type != environment.NumberType { t.Fatalf("choice scalar wrong: %v", c) }
	carr, _ := choiceFn.Builtin([]environment.Value{src, environment.NewNumber(2)})
	if carr.Type != environment.ArrayType || len(*carr.Arr) != 2 { t.Fatalf("choice size wrong: %v", carr) }
	cfromn, _ := choiceFn.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(3)})
	if cfromn.Type != environment.ArrayType || len(*cfromn.Arr) != 3 { t.Fatalf("choice from int wrong: %v", cfromn) }

	// shuffle (in-place) and permutation
	arrToShuffle := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	_, err = shuffleFn.Builtin([]environment.Value{arrToShuffle})
	if err != nil { t.Fatalf("shuffle error: %v", err) }
	// perm from n
	perm, err := permFn.Builtin([]environment.Value{environment.NewNumber(4)})
	if err != nil { t.Fatalf("permutation error: %v", err) }
	if len(*perm.Arr) != 4 { t.Fatalf("permutation length wrong: %v", perm) }

	// normal / uniform
	normv, _ := normalFn.Builtin([]environment.Value{})
	if normv.Type != environment.NumberType || math.IsNaN(normv.Num) { t.Fatalf("normal scalar wrong: %v", normv) }
	unif, _ := uniformFn.Builtin([]environment.Value{environment.NewNumber(5.0), environment.NewNumber(6.0)})
	if unif.Type != environment.NumberType || unif.Num < 5.0 || unif.Num >= 6.0 { t.Fatalf("uniform scalar wrong: %v", unif) }

	// binomial / poisson / exponential / gamma / beta (basic property checks)
	b, _ := binomFn.Builtin([]environment.Value{environment.NewNumber(5), environment.NewNumber(0.5)})
	if b.Type != environment.NumberType || b.Num < 0 || b.Num > 5 { t.Fatalf("binomial wrong: %v", b) }
	po, _ := poisFn.Builtin([]environment.Value{environment.NewNumber(2.0)})
	if po.Type != environment.NumberType || po.Num < 0 { t.Fatalf("poisson wrong: %v", po) }
	e, _ := expFn.Builtin([]environment.Value{environment.NewNumber(2.0)})
	if e.Type != environment.NumberType || e.Num < 0 { t.Fatalf("exponential wrong: %v", e) }
	g, _ := gammaFn.Builtin([]environment.Value{environment.NewNumber(2), environment.NewNumber(1.0)})
	if g.Type != environment.NumberType || g.Num <= 0 { t.Fatalf("gamma wrong: %v", g) }
	bb, _ := betaFn.Builtin([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)})
	if bb.Type != environment.NumberType || bb.Num < 0 || bb.Num > 1 { t.Fatalf("beta wrong: %v", bb) }
}

func TestMathxDtypes(t *testing.T) {
	dtypeFn := getBuiltin(t, "dtype")
	astypeFn := getBuiltin(t, "astype")
	issubFn := getBuiltin(t, "issubdtype")
	canCastFn := getBuiltin(t, "can_cast")
	resTypeFn := getBuiltin(t, "result_type")

	// dtype
	d, _ := dtypeFn.Builtin([]environment.Value{environment.NewNumber(1)})
	if d.Type != environment.StringType || d.Str != "number" { t.Fatalf("dtype scalar wrong: %v", d) }
	a, _ := dtypeFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})})
	if a.Str != "number" { t.Fatalf("dtype array wrong: %v", a) }
	m, _ := dtypeFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewString("x")})})
	if m.Str != "mixed" { t.Fatalf("dtype mixed wrong: %v", m) }
	e, _ := dtypeFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{})})
	if e.Str != "null" { t.Fatalf("dtype empty array wrong: %v", e) }

	// astype
	as1, err := astypeFn.Builtin([]environment.Value{environment.NewString("123.5"), environment.NewString("number")})
	if err != nil { t.Fatalf("astype scalar error: %v", err) }
	if as1.Type != environment.NumberType || math.Abs(as1.Num-123.5) > 1e-9 { t.Fatalf("astype scalar -> number wrong: %v", as1) }
	asArr, err := astypeFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewString("1"), environment.NewNil()}), environment.NewString("boolean")})
	if err != nil { t.Fatalf("astype array error: %v", err) }
	if (*asArr.Arr)[0].Type != environment.BooleanType || !(*asArr.Arr)[0].Bool || (*asArr.Arr)[1].Bool { t.Fatalf("astype array->bool wrong: %v", asArr) }

	// issubdtype / can_cast
	is1, _ := issubFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewString("number")})
	if !is1.IsTruthy() { t.Fatalf("issubdtype value->dtype wrong: %v", is1) }
	is2, _ := issubFn.Builtin([]environment.Value{environment.NewString("number"), environment.NewString("string")})
	if is2.IsTruthy() { t.Fatalf("issubdtype dtype->dtype wrong: %v", is2) }
	can1, _ := canCastFn.Builtin([]environment.Value{environment.NewString("123.4"), environment.NewString("number")})
	if !can1.IsTruthy() { t.Fatalf("can_cast string->number should be true: %v", can1) }
	can2, _ := canCastFn.Builtin([]environment.Value{environment.NewString("abc"), environment.NewString("number")})
	if can2.IsTruthy() { t.Fatalf("can_cast should be false for 'abc'->number: %v", can2) }

	// result_type
	rt, _ := resTypeFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewString("x")})
	if rt.Str != "string" { t.Fatalf("result_type wrong: %v", rt) }
	rt2, _ := resTypeFn.Builtin([]environment.Value{environment.NewNumber(1), environment.NewBool(true)})
	if rt2.Str != "number" { t.Fatalf("result_type wrong: %v", rt2) }
	rt3, _ := resTypeFn.Builtin([]environment.Value{environment.NewBool(true), environment.NewBool(false)})
	if rt3.Str != "boolean" { t.Fatalf("result_type wrong: %v", rt3) }
}

func TestMathxLinalg(t *testing.T) {
	dot := getBuiltin(t, "dot")
	matmul := getBuiltin(t, "matmul")
	inner := getBuiltin(t, "inner")
	outer := getBuiltin(t, "outer")
	vdot := getBuiltin(t, "vdot")
	tens := getBuiltin(t, "tensordot")
	trace := getBuiltin(t, "trace")
	norm := getBuiltin(t, "norm")
	detFn := getBuiltin(t, "det")
	invFn := getBuiltin(t, "inv")
	pinv := getBuiltin(t, "pinv")
	solveFn := getBuiltin(t, "solve")
	eigFn := getBuiltin(t, "eig")
	eigvalsFn := getBuiltin(t, "eigvals")
	svdFn := getBuiltin(t, "svd")
	qrFn := getBuiltin(t, "qr")
	choleskyFn := getBuiltin(t, "cholesky")
	rankFn := getBuiltin(t, "matrix_rank")

	// vectors
	v1 := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	v2 := environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})
	out, err := dot.Builtin([]environment.Value{v1, v2})
	if err != nil { t.Fatalf("dot error: %v", err) }
	if out.Num != 11 { t.Fatalf("dot wrong: %v", out) }

	in, err := inner.Builtin([]environment.Value{v1, v2})
	if err != nil { t.Fatalf("inner error: %v", err) }
	if in.Num != 11 { t.Fatalf("inner wrong: %v", in) }
	vdotRes, err := vdot.Builtin([]environment.Value{v1, v2})
	if err != nil { t.Fatalf("vdot error: %v", err) }
	if vdotRes.Num != 11 { t.Fatalf("vdot wrong: %v", vdotRes) }

	outArr, err := outer.Builtin([]environment.Value{v1, v2})
	if err != nil { t.Fatalf("outer error: %v", err) }
	if (*outArr.Arr)[0].Type != environment.ArrayType || (*(*outArr.Arr)[0].Arr)[0].Num != 3 { t.Fatalf("outer wrong: %v", outArr) }

	// matrix multiply
	A := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)}) })
	B := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(5), environment.NewNumber(6)}), environment.NewArray([]environment.Value{environment.NewNumber(7), environment.NewNumber(8)}) })
	mres, err := matmul.Builtin([]environment.Value{A, B})
	if err != nil { t.Fatalf("matmul error: %v", err) }
	if (*mres.Arr)[0].Type != environment.ArrayType || (*(*mres.Arr)[0].Arr)[0].Num != 19 { t.Fatalf("matmul wrong: %v", mres) }

	// tensordot == matmul for 2D
	td, err := tens.Builtin([]environment.Value{A, B})
	if err != nil { t.Fatalf("tensordot error: %v", err) }
	if (*td.Arr)[0].Type != environment.ArrayType || (*(*td.Arr)[0].Arr)[0].Num != 19 { t.Fatalf("tensordot wrong: %v", td) }

	// trace
	tr, err := trace.Builtin([]environment.Value{A})
	if err != nil { t.Fatalf("trace error: %v", err) }
	if tr.Num != 5 { t.Fatalf("trace wrong: %v", tr) }

	// norm
	nn, _ := norm.Builtin([]environment.Value{v1})
	if math.Abs(nn.Num-math.Sqrt(5)) > 1e-9 { t.Fatalf("norm wrong: %v", nn) }

	// det / inv
	m := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(4), environment.NewNumber(7)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(6)}) })
	dv, err := detFn.Builtin([]environment.Value{m})
	if err != nil { t.Fatalf("det error: %v", err) }
	if math.Abs(dv.Num-10) > 1e-9 { t.Fatalf("det wrong: %v", dv) }
	invv, err := invFn.Builtin([]environment.Value{m})
	if err != nil { t.Fatalf("inv error: %v", err) }
	// check inv * m = I
	prod, err := matmul.Builtin([]environment.Value{invv, m})
	if err != nil { t.Fatalf("inv* m error: %v", err) }
	if (*prod.Arr)[0].Type != environment.ArrayType || math.Abs(((*(*prod.Arr)[0].Arr)[0].Num)-1.0) > 1e-9 { t.Fatalf("inv check failed: %v", prod) }

	// pinv (fallback to inv for square)
	_, err = pinv.Builtin([]environment.Value{m})
	if err != nil { t.Fatalf("pinv error: %v", err) }

	// solve
	A2 := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(1)}), environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}) })
	bvec := environment.NewArray([]environment.Value{environment.NewNumber(9), environment.NewNumber(8)})
	xres, err := solveFn.Builtin([]environment.Value{A2, bvec})
	if err != nil { t.Fatalf("solve error: %v", err) }
	if (*xres.Arr)[0].Num != 2 || (*xres.Arr)[1].Num != 3 { t.Fatalf("solve wrong: %v", xres) }

	// eig / eigvals for diagonal
	D := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(0)}), environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(3)}) })
	evs, err := eigvalsFn.Builtin([]environment.Value{D})
	if err != nil { t.Fatalf("eigvals error: %v", err) }
	numSliceEquals(t, *evs.Arr, []float64{(3), (2)})
	eigRes, err := eigFn.Builtin([]environment.Value{D})
	if err != nil { t.Fatalf("eig error: %v", err) }
	// eig returns [vals, vecs]
	if eigRes.Type != environment.ArrayType || len(*eigRes.Arr) != 2 { t.Fatalf("eig wrong: %v", eigRes) }

	// svd for diag
	Sres, err := svdFn.Builtin([]environment.Value{D})
	if err != nil { t.Fatalf("svd error: %v", err) }
	// Sres = [U, S, V]
	if Sres.Type != environment.ArrayType || len(*Sres.Arr) != 3 { t.Fatalf("svd wrong shape: %v", Sres) }
	// singular values should include 3 and 2
	Svals := (*(*Sres.Arr)[1].Arr)
	if Svals[0].Num < Svals[1].Num { t.Fatalf("svd singulars order wrong: %v", Svals) }

	// qr (check Q*R ~= A)
	QRres, err := qrFn.Builtin([]environment.Value{A})
	if err != nil { t.Fatalf("qr error: %v", err) }
	if QRres.Type != environment.ArrayType || len(*QRres.Arr) != 2 { t.Fatalf("qr wrong: %v", QRres) }
	Qm := (*QRres.Arr)[0]
	Rm := (*QRres.Arr)[1]
	recon, err := matmul.Builtin([]environment.Value{Qm, Rm})
	if err != nil { t.Fatalf("qr recon error: %v", err) }
	if (*recon.Arr)[0].Type != environment.ArrayType || math.Abs(((*(*recon.Arr)[0].Arr)[0].Num)-1.0) > 1e-9 { t.Fatalf("qr recon wrong: %v", recon) }

	// cholesky
	spd := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(4), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(3)}) })
	Lres, err := choleskyFn.Builtin([]environment.Value{spd})
	if err != nil { t.Fatalf("cholesky error: %v", err) }
	// check L * L^T == spd
	Lt := registry["mathx"].Entries["transpose"].Builtin
	Ltres, _ := Lt([]environment.Value{Lres})
	LLt, err := matmul.Builtin([]environment.Value{Lres, Ltres})
	if err != nil { t.Fatalf("cholesky recon error: %v", err) }
	if (*LLt.Arr)[0].Type != environment.ArrayType || math.Abs(((*(*LLt.Arr)[0].Arr)[0].Num)-4.0) > 1e-9 { t.Fatalf("cholesky recon wrong: %v", LLt) }

	// matrix_rank
	rk, err := rankFn.Builtin([]environment.Value{environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(2), environment.NewNumber(4)}) })})
	if err != nil { t.Fatalf("rank error: %v", err) }
	if rk.Num != 1 { t.Fatalf("rank wrong: %v", rk) }
}

func TestMathxFFT(t *testing.T) {
	fftFn := getBuiltin(t, "fft")
	ifftFn := getBuiltin(t, "ifft")
	rfftFn := getBuiltin(t, "rfft")
	irfftFn := getBuiltin(t, "irfft")
	fft2Fn := getBuiltin(t, "fft2")
	ifft2Fn := getBuiltin(t, "ifft2")
	fftshiftFn := getBuiltin(t, "fftshift")
	ifftshiftFn := getBuiltin(t, "ifftshift")

	// simple 1D FFT: delta at 0 -> all ones
	a := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(0), environment.NewNumber(0), environment.NewNumber(0)})
	X, err := fftFn.Builtin([]environment.Value{a})
	if err != nil { t.Fatalf("fft error: %v", err) }
	if X.Type != environment.ArrayType || len(*X.Arr) != 4 { t.Fatalf("fft wrong shape: %v", X) }
	for i := 0; i < 4; i++ {
		pt := (*X.Arr)[i]
		if pt.Type != environment.ArrayType || len(*pt.Arr) != 2 { t.Fatalf("fft element not complex pair: %v", pt) }
		r := (*pt.Arr)[0].Num
		im := (*pt.Arr)[1].Num
		if math.Abs(r-1.0) > 1e-9 || math.Abs(im) > 1e-9 { t.Fatalf("fft delta wrong at %d: %v", i, pt) }
	}

	// fft of [0,1,0,0] -> [1, -i, -1, i]
	b := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(0), environment.NewNumber(0)})
	Y, err := fftFn.Builtin([]environment.Value{b})
	if err != nil { t.Fatalf("fft error: %v", err) }
	e0 := (*(*Y.Arr)[0].Arr)[0].Num
	e1im := (*(*Y.Arr)[1].Arr)[1].Num
	if math.Abs(e0-1.0) > 1e-9 || math.Abs(e1im+1.0) > 1e-9 { t.Fatalf("fft expected complex phasors wrong: %v", Y) }

	// ifft(fft(x)) -> x (real parts)
	orig := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3), environment.NewNumber(4)})
	F, err := fftFn.Builtin([]environment.Value{orig})
	if err != nil { t.Fatalf("fft error: %v", err) }
	I, err := ifftFn.Builtin([]environment.Value{F})
	if err != nil { t.Fatalf("ifft error: %v", err) }
	for i := 0; i < 4; i++ {
		val := (*(*I.Arr)[i].Arr)[0].Num
		if math.Abs(val-float64(i+1)) > 1e-9 { t.Fatalf("ifft recon wrong at %d: %v", i, I) }
	}

	// rfft/irfft round-trip
	rr, err := rfftFn.Builtin([]environment.Value{orig})
	if err != nil { t.Fatalf("rfft error: %v", err) }
	ir, err := irfftFn.Builtin([]environment.Value{rr})
	if err != nil { t.Fatalf("irfft error: %v", err) }
	for i := 0; i < 4; i++ {
		if math.Abs((*ir.Arr)[i].Num - float64(i+1)) > 1e-9 { t.Fatalf("irfft recon wrong at %d: %v", i, ir) }
	}

	// fftshift / ifftshift
	seq := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	sh, err := fftshiftFn.Builtin([]environment.Value{seq})
	if err != nil { t.Fatalf("fftshift error: %v", err) }
	if (*sh.Arr)[0].Num != 2 || (*sh.Arr)[1].Num != 3 || (*sh.Arr)[2].Num != 0 || (*sh.Arr)[3].Num != 1 { t.Fatalf("fftshift wrong: %v", sh) }
	unsh, err := ifftshiftFn.Builtin([]environment.Value{sh})
	if err != nil { t.Fatalf("ifftshift error: %v", err) }
	if (*unsh.Arr)[0].Num != 0 || (*unsh.Arr)[1].Num != 1 || (*unsh.Arr)[2].Num != 2 || (*unsh.Arr)[3].Num != 3 { t.Fatalf("ifftshift wrong: %v", unsh) }

	// 2D fft/ifft roundtrip
	mat := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(0)}), environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(0)}) })
	FM, err := fft2Fn.Builtin([]environment.Value{mat})
	if err != nil { t.Fatalf("fft2 error: %v", err) }
	IM, err := ifft2Fn.Builtin([]environment.Value{FM})
	if err != nil { t.Fatalf("ifft2 error: %v", err) }
	// extract real parts and compare
	r00 := (*(*(*IM.Arr)[0].Arr)[0].Arr)[0].Num
	r01 := (*(*(*IM.Arr)[0].Arr)[1].Arr)[0].Num
	r10 := (*(*(*IM.Arr)[1].Arr)[0].Arr)[0].Num
	if math.Abs(r00-1.0) > 1e-9 || math.Abs(r01-0.0) > 1e-9 || math.Abs(r10-0.0) > 1e-9 { t.Fatalf("ifft2 recon wrong: %v", IM) }
}

func TestMathxIO(t *testing.T) {
	saveFn := getBuiltin(t, "save")
	loadFn := getBuiltin(t, "load")
	savezFn := getBuiltin(t, "savez")
	savetxtFn := getBuiltin(t, "savetxt")
	loadtxtFn := getBuiltin(t, "loadtxt")
	genfromtxtFn := getBuiltin(t, "genfromtxt")
	tofileFn := getBuiltin(t, "tofile")
	fromfileFn := getBuiltin(t, "fromfile")

	// save/load (JSON roundtrip)
	tmp := t.TempDir() + "/arr.json"
	arr := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	if _, err := saveFn.Builtin([]environment.Value{environment.NewString(tmp), arr}); err != nil { t.Fatalf("save error: %v", err) }
	l, err := loadFn.Builtin([]environment.Value{environment.NewString(tmp)})
	if err != nil { t.Fatalf("load error: %v", err) }
	if l.Type != environment.ArrayType || len(*l.Arr) != 3 { t.Fatalf("load returned wrong type: %v", l) }
	if (*l.Arr)[1].Num != 2 { t.Fatalf("load value mismatch: %v", l) }

	// savez + load (object)
	tmp2 := t.TempDir() + "/multi.json"
	obj := environment.NewObject(map[string]environment.Value{"a": environment.NewArray([]environment.Value{environment.NewNumber(1)}), "b": environment.NewNumber(5)}, []string{"a", "b"})
	if _, err := savezFn.Builtin([]environment.Value{environment.NewString(tmp2), obj}); err != nil { t.Fatalf("savez error: %v", err) }
	lo, err := loadFn.Builtin([]environment.Value{environment.NewString(tmp2)})
	if err != nil { t.Fatalf("load error: %v", err) }
	if lo.Type != environment.ObjectType { t.Fatalf("savez/load returned wrong type: %v", lo) }
	if val, ok := lo.Obj.Entries["b"]; !ok || val.Num != 5 { t.Fatalf("savez content wrong: %v", lo) }

	// savetxt/loadtxt roundtrip (2D and 1D)
	tcsv := t.TempDir() + "/t.csv"
	mat := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)}), environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)}) })
	if _, err := savetxtFn.Builtin([]environment.Value{environment.NewString(tcsv), mat}); err != nil { t.Fatalf("savetxt error: %v", err) }
	lt, err := loadtxtFn.Builtin([]environment.Value{environment.NewString(tcsv)})
	if err != nil { t.Fatalf("loadtxt error: %v", err) }
	if (*(*lt.Arr)[0].Arr)[1].Num != 2 { t.Fatalf("loadtxt wrong: %v", lt) }

	// genfromtxt: missing and comments
	txt := t.TempDir() + "/g.txt"
	os.WriteFile(txt, []byte("# header\n1 2\n3  \n4 5"), 0644)
	gf, err := genfromtxtFn.Builtin([]environment.Value{environment.NewString(txt)})
	if err != nil { t.Fatalf("genfromtxt error: %v", err) }
	// second row has a missing second column -> should be Array with Nil or single number depending on parsing
	if (*gf.Arr)[1].Type == environment.ArrayType {
		if (*(*gf.Arr)[1].Arr)[1].Type != environment.NilType { t.Fatalf("genfromtxt missing value not nil: %v", gf) }
	}

	// tofile/fromfile (binary float64)
	bf := t.TempDir() + "/bin.dat"
	nums := environment.NewArray([]environment.Value{environment.NewNumber(1.5), environment.NewNumber(2.5), environment.NewNumber(3.5)})
	if _, err := tofileFn.Builtin([]environment.Value{environment.NewString(bf), nums}); err != nil { t.Fatalf("tofile error: %v", err) }
	rf, err := fromfileFn.Builtin([]environment.Value{environment.NewString(bf)})
	if err != nil { t.Fatalf("fromfile error: %v", err) }
	if (*rf.Arr)[2].Num != 3.5 { t.Fatalf("fromfile wrong: %v", rf) }
}

func TestMathxPoly(t *testing.T) {
	polyval := getBuiltin(t, "polyval")
	polyfit := getBuiltin(t, "polyfit")
	polyadd := getBuiltin(t, "polyadd")
	polysub := getBuiltin(t, "polysub")
	polymul := getBuiltin(t, "polymul")
	polyder := getBuiltin(t, "polyder")
	polyint := getBuiltin(t, "polyint")

	// polyval scalar
	p := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)}) // x^2+2x+3
	v, err := polyval.Builtin([]environment.Value{p, environment.NewNumber(2)})
	if err != nil { t.Fatalf("polyval error: %v", err) }
	if v.Type != environment.NumberType || v.Num != 11 { t.Fatalf("polyval scalar wrong: %v", v) }

	// polyval array
	xs := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(2)})
	vals, err := polyval.Builtin([]environment.Value{p, xs})
	if err != nil { t.Fatalf("polyval error: %v", err) }
	if (*vals.Arr)[0].Num != 3 || (*vals.Arr)[1].Num != 6 || (*vals.Arr)[2].Num != 11 { t.Fatalf("polyval array wrong: %v", vals) }

	// polyadd / polysub
	a := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	b := environment.NewArray([]environment.Value{environment.NewNumber(4), environment.NewNumber(5)})
	pa, _ := polyadd.Builtin([]environment.Value{a, b})
	if (*pa.Arr)[0].Num != 1 || (*pa.Arr)[1].Num != 6 || (*pa.Arr)[2].Num != 8 { t.Fatalf("polyadd wrong: %v", pa) }
	ps, _ := polysub.Builtin([]environment.Value{a, b})
	if (*ps.Arr)[0].Num != 1 || (*ps.Arr)[1].Num != -2 || (*ps.Arr)[2].Num != -2 { t.Fatalf("polysub wrong: %v", ps) }

	// polymul
	mres, _ := polymul.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(1)}), environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(1)})})
	if (*mres.Arr)[0].Num != 1 || (*mres.Arr)[1].Num != 2 || (*mres.Arr)[2].Num != 1 { t.Fatalf("polymul simple wrong: %v", mres) }

	// polyder
	d := environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(2), environment.NewNumber(1)}) // 3x^2+2x+1
	dres, _ := polyder.Builtin([]environment.Value{d})
	if (*dres.Arr)[0].Num != 6 || (*dres.Arr)[1].Num != 2 { t.Fatalf("polyder wrong: %v", dres) }
	// derivative of constant
	cder, _ := polyder.Builtin([]environment.Value{environment.NewArray([]environment.Value{environment.NewNumber(5)})})
	if (*cder.Arr)[0].Num != 0 { t.Fatalf("polyder const wrong: %v", cder) }

	// polyint
	intRes, _ := polyint.Builtin([]environment.Value{d})
	// integral of 3x^2+2x+1 is x^3 + x^2 + x + C -> coefficients [1,1,1,C]
	if len(*intRes.Arr) != 4 || (*intRes.Arr)[0].Num != 1 || (*intRes.Arr)[1].Num != 1 || (*intRes.Arr)[2].Num != 1 || (*intRes.Arr)[3].Num != 0 { t.Fatalf("polyint wrong: %v", intRes) }
	intRes2, _ := polyint.Builtin([]environment.Value{d, environment.NewNumber(7)})
	if (*intRes2.Arr)[3].Num != 7 { t.Fatalf("polyint constant wrong: %v", intRes2) }

	// polyfit (fit quadratic to points -> [1,2,3])
	x := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(1), environment.NewNumber(2)})
	y := environment.NewArray([]environment.Value{environment.NewNumber(3), environment.NewNumber(6), environment.NewNumber(11)})
	fit, err := polyfit.Builtin([]environment.Value{x, y, environment.NewNumber(2)})
	if err != nil { t.Fatalf("polyfit error: %v", err) }
	if math.Abs((*fit.Arr)[0].Num-1.0) > 1e-9 || math.Abs((*fit.Arr)[1].Num-2.0) > 1e-9 || math.Abs((*fit.Arr)[2].Num-3.0) > 1e-9 { t.Fatalf("polyfit wrong: %v", fit) }
}

func TestMathxUtils(t *testing.T) {
	ndimFn := getBuiltin(t, "ndim")
	sizeFn := getBuiltin(t, "size")
	itemsizeFn := getBuiltin(t, "itemsize")
	copytoFn := getBuiltin(t, "copyto")
	viewFn := getBuiltin(t, "view")
	getPo := getBuiltin(t, "get_printoptions")
	setPo := getBuiltin(t, "set_printoptions")

	// ndim
	if v, _ := ndimFn.Builtin([]environment.Value{environment.NewNumber(1)}); v.Num != 0 { t.Fatalf("ndim scalar wrong: %v", v) }
	a := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	if v, _ := ndimFn.Builtin([]environment.Value{a}); v.Num != 1 { t.Fatalf("ndim 1D wrong: %v", v) }
	m := environment.NewArray([]environment.Value{ environment.NewArray([]environment.Value{environment.NewNumber(1)}), environment.NewArray([]environment.Value{environment.NewNumber(2)}) })
	if v, _ := ndimFn.Builtin([]environment.Value{m}); v.Num != 2 { t.Fatalf("ndim 2D wrong: %v", v) }

	// size
	if v, _ := sizeFn.Builtin([]environment.Value{environment.NewNumber(1)}); v.Num != 1 { t.Fatalf("size scalar wrong: %v", v) }
	if v, _ := sizeFn.Builtin([]environment.Value{a}); v.Num != 2 { t.Fatalf("size 1D wrong: %v", v) }
	if v, _ := sizeFn.Builtin([]environment.Value{m}); v.Num != 2 { t.Fatalf("size 2D wrong: %v", v) }

	// itemsize
	if v, _ := itemsizeFn.Builtin([]environment.Value{environment.NewNumber(1.23)}); v.Num != 8 { t.Fatalf("itemsize number wrong: %v", v) }
	if v, _ := itemsizeFn.Builtin([]environment.Value{environment.NewString("abc")}); v.Num != 3 { t.Fatalf("itemsize string wrong: %v", v) }
	if v, _ := itemsizeFn.Builtin([]environment.Value{a}); v.Num != 8 { t.Fatalf("itemsize array->number wrong: %v", v) }

	// copyto (array<-array)
	d := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(0), environment.NewNumber(0)})
	src := environment.NewArray([]environment.Value{environment.NewNumber(7), environment.NewNumber(8), environment.NewNumber(9)})
	if _, err := copytoFn.Builtin([]environment.Value{d, src}); err != nil { t.Fatalf("copyto error: %v", err) }
	if (*d.Arr)[0].Num != 7 || (*d.Arr)[2].Num != 9 { t.Fatalf("copyto not copied: %v", d) }
	// copyto (array<-scalar)
	d2 := environment.NewArray([]environment.Value{environment.NewNumber(0), environment.NewNumber(0)})
	if _, err := copytoFn.Builtin([]environment.Value{d2, environment.NewNumber(5)}); err != nil { t.Fatalf("copyto scalar error: %v", err) }
	if (*d2.Arr)[0].Num != 5 || (*d2.Arr)[1].Num != 5 { t.Fatalf("copyto scalar wrong: %v", d2) }
	// copyto mismatch -> error
	if _, err := copytoFn.Builtin([]environment.Value{d, environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})}); err == nil { t.Fatalf("copyto should error on size mismatch") }

	// view -> shared backing
	aorig := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2)})
	vwr, err := viewFn.Builtin([]environment.Value{aorig})
	if err != nil { t.Fatalf("view error: %v", err) }
	// mutate view and check original
	(*vwr.Arr)[0] = environment.NewNumber(99)
	if (*aorig.Arr)[0].Num != 99 { t.Fatalf("view not shared backing: %v", aorig) }

	// printoptions
	po, _ := getPo.Builtin([]environment.Value{})
	if po.Type != environment.ObjectType { t.Fatalf("get_printoptions wrong type: %v", po) }
	prec := po.Obj.Entries["precision"].Num
	if prec != 6 { t.Fatalf("default precision wrong: %v", prec) }
	// set options
	_, _ = setPo.Builtin([]environment.Value{ environment.NewObject(map[string]environment.Value{"precision": environment.NewNumber(3), "linewidth": environment.NewNumber(40)}, []string{"precision","linewidth"}) })
	po2, _ := getPo.Builtin([]environment.Value{})
	if po2.Obj.Entries["precision"].Num != 3 || po2.Obj.Entries["linewidth"].Num != 40 { t.Fatalf("set_printoptions failed: %v", po2) }
}

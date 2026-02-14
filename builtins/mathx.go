package builtins

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/iscarloscoder/fig/environment"
)

// --- helpers for mathx operations (operate on nested Fig arrays) ---

func product(shape []int) int {
	p := 1
	for _, d := range shape {
		p *= d
	}
	return p
}

// global print options for mathx
var mathxPrintOptions = struct{
	Precision int
	LineWidth int
	Threshold int
}{
	Precision: 6,
	LineWidth: 75,
	Threshold: 1000,
}

// shapeOf returns the shape of a (possibly-nested) array. For non-arrays returns empty slice.
func shapeOf(v environment.Value) []int {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return []int{}
	}
	out := []int{len(*v.Arr)}
	if len(*v.Arr) == 0 {
		return out
	}
	first := (*v.Arr)[0]
	if first.Type == environment.ArrayType {
		sub := shapeOf(first)
		out = append(out, sub...)
	}
	return out
}

// flattenValues flattens a nested array to 1D (pre-order).
func flattenValues(v environment.Value) []environment.Value {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return []environment.Value{v}
	}
	var out []environment.Value
	for _, e := range *v.Arr {
		if e.Type == environment.ArrayType {
			out = append(out, flattenValues(e)...)
		} else {
			out = append(out, e)
		}
	}
	return out
}

// buildFromFlat constructs nested arrays from flat values using provided shape.
func buildFromFlat(flat []environment.Value, shape []int) (environment.Value, error) {
	if len(shape) == 0 {
		if len(flat) == 0 {
			return environment.NewNil(), nil
		}
		return flat[0], nil
	}
	if product(shape) != len(flat) {
		return environment.NewNil(), fmt.Errorf("cannot reshape: size mismatch")
	}
	var idx int
	var build func(dim int) (environment.Value, error)
	build = func(dim int) (environment.Value, error) {
		n := shape[dim]
		arr := make([]environment.Value, n)
		if dim == len(shape)-1 {
			for i := 0; i < n; i++ {
				arr[i] = flat[idx]
				idx++
			}
			return environment.NewArray(arr), nil
		}
		for i := 0; i < n; i++ {
			v, err := build(dim + 1)
			if err != nil {
				return environment.NewNil(), err
			}
			arr[i] = v
		}
		return environment.NewArray(arr), nil
	}
	return build(0)
}

func reshapeValue(v environment.Value, shape []int) (environment.Value, error) {
	flat := flattenValues(v)
	if product(shape) != len(flat) {
		return environment.NewNil(), fmt.Errorf("cannot reshape: requested size %d does not match source size %d", product(shape), len(flat))
	}
	return buildFromFlat(flat, shape)
}

// transpose2D transposes a 2D array (matrix). If input is 1D returns it unchanged.
func transpose2D(v environment.Value) (environment.Value, error) {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return v, nil
	}
	// check if it's 2D
	rows := *v.Arr
	if len(rows) == 0 {
		return environment.NewArray([]environment.Value{}), nil
	}
	first := rows[0]
	if first.Type != environment.ArrayType || first.Arr == nil {
		// 1D — return unchanged
		return v, nil
	}
	cols := len(*first.Arr)
	// ensure regular
	for i := 1; i < len(rows); i++ {
		r := rows[i]
		if r.Type != environment.ArrayType || r.Arr == nil || len(*r.Arr) != cols {
			return environment.NewNil(), fmt.Errorf("transpose: irregular matrix")
		}
	}
	out := make([]environment.Value, cols)
	for c := 0; c < cols; c++ {
		row := make([]environment.Value, len(rows))
		for r := 0; r < len(rows); r++ {
			row[r] = (*rows[r].Arr)[c]
		}
		out[c] = environment.NewArray(row)
	}
	return environment.NewArray(out), nil
}

// concatAlongAxis concatenates 1D or 2D arrays along axis 0 (default) or 1 (for 2D).
func concatAlongAxis(arrs []environment.Value, axis int) (environment.Value, error) {
	if len(arrs) == 0 {
		return environment.NewArray([]environment.Value{}), nil
	}
	// handle 1D case
	all1D := true
	for _, a := range arrs {
		if a.Type != environment.ArrayType || a.Arr == nil {
			all1D = false
			break
		}
		// check elements non-array
		if len(*a.Arr) > 0 && (*a.Arr)[0].Type == environment.ArrayType {
			all1D = false
			break
		}
	}
	if all1D {
		// concatenate flat
		var out []environment.Value
		for _, a := range arrs {
			out = append(out, *a.Arr...)
		}
		return environment.NewArray(out), nil
	}
	// assume 2D
	if axis == 0 {
		// vertical concat (append rows)
		var out []environment.Value
		for _, a := range arrs {
			if a.Type != environment.ArrayType || a.Arr == nil {
				return environment.NewNil(), fmt.Errorf("concatenate: expected 2D arrays for axis=0")
			}
			for _, r := range *a.Arr {
				if r.Type != environment.ArrayType || r.Arr == nil {
					return environment.NewNil(), fmt.Errorf("concatenate: expected 2D arrays for axis=0")
				}
				out = append(out, r)
			}
		}
		return environment.NewArray(out), nil
	}
	// axis == 1: horizontal concat — rows must align
	// determine number of rows from first
	first := arrs[0]
	if first.Type != environment.ArrayType || first.Arr == nil {
		return environment.NewNil(), fmt.Errorf("concatenate: expected 2D arrays for axis=1")
	}
	nrows := len(*first.Arr)
	// build output rows
	outRows := make([][]environment.Value, nrows)
	for i := 0; i < nrows; i++ {
		outRows[i] = []environment.Value{}
	}
	for _, a := range arrs {
		if a.Type != environment.ArrayType || a.Arr == nil || len(*a.Arr) != nrows {
			return environment.NewNil(), fmt.Errorf("concatenate: all inputs must have same number of rows for axis=1")
		}
		for i := 0; i < nrows; i++ {
			row := (*a.Arr)[i]
			if row.Type != environment.ArrayType || row.Arr == nil {
				return environment.NewNil(), fmt.Errorf("concatenate: expected 2D arrays for axis=1")
			}
			outRows[i] = append(outRows[i], *row.Arr...)
		}
	}
	res := make([]environment.Value, nrows)
	for i := 0; i < nrows; i++ {
		res[i] = environment.NewArray(outRows[i])
	}
	return environment.NewArray(res), nil
}

// stackArrays stacks arrays along new axis 0
func stackArrays(arrs []environment.Value) (environment.Value, error) {
	// ensure same shape
	if len(arrs) == 0 {
		return environment.NewArray([]environment.Value{}), nil
	}
	// simply return array of the inputs (shallow)
	out := make([]environment.Value, len(arrs))
	for i := range arrs {
		out[i] = arrs[i]
	}
	return environment.NewArray(out), nil
}

// simple 1D split into equal parts
func split1D(arr environment.Value, parts int) ([]environment.Value, error) {
	if arr.Type != environment.ArrayType || arr.Arr == nil {
		return nil, fmt.Errorf("split: expects 1D array")
	}
	n := len(*arr.Arr)
	if parts <= 0 || n%parts != 0 {
		return nil, fmt.Errorf("split: array of length %d cannot be split into %d equal parts", n, parts)
	}
	sz := n / parts
	out := make([]environment.Value, parts)
	for i := 0; i < parts; i++ {
		chunk := make([]environment.Value, sz)
		copy(chunk, (*arr.Arr)[i*sz:(i+1)*sz])
		out[i] = environment.NewArray(chunk)
	}
	return out, nil
}

func arraySplit1D(arr environment.Value, parts int) ([]environment.Value, error) {
	if arr.Type != environment.ArrayType || arr.Arr == nil {
		return nil, fmt.Errorf("array_split: expects 1D array")
	}
	n := len(*arr.Arr)
	if parts <= 0 {
		return nil, fmt.Errorf("array_split: parts must be > 0")
	}
	base := n / parts
	rem := n % parts
	out := make([]environment.Value, parts)
	idx := 0
	for i := 0; i < parts; i++ {
		sz := base
		if i < rem {
			sz++
		}
		chunk := make([]environment.Value, sz)
		copy(chunk, (*arr.Arr)[idx:idx+sz])
		out[i] = environment.NewArray(chunk)
		idx += sz
	}
	return out, nil
}

// tile1D: repeat full array 'reps' times
func tile1D(arr environment.Value, reps int) (environment.Value, error) {
	if arr.Type != environment.ArrayType || arr.Arr == nil {
		return environment.NewNil(), fmt.Errorf("tile: expects 1D array")
	}
	if reps < 0 {
		return environment.NewNil(), fmt.Errorf("tile: reps must be non-negative")
	}
	n := len(*arr.Arr)
	out := make([]environment.Value, n*reps)
	for i := 0; i < reps; i++ {
		copy(out[i*n:(i+1)*n], *arr.Arr)
	}
	return environment.NewArray(out), nil
}

// repeat1D: repeat each element 'repeats' times
func repeat1D(arr environment.Value, repeats int) (environment.Value, error) {
	if arr.Type != environment.ArrayType || arr.Arr == nil {
		return environment.NewNil(), fmt.Errorf("repeat: expects 1D array")
	}
	if repeats < 0 {
		return environment.NewNil(), fmt.Errorf("repeat: repeats must be non-negative")
	}
	n := len(*arr.Arr)
	out := make([]environment.Value, n*repeats)
	k := 0
	for i := 0; i < n; i++ {
		for r := 0; r < repeats; r++ {
			out[k] = (*arr.Arr)[i]
			k++
		}
	}
	return environment.NewArray(out), nil
}

// broadcast scalar or 1D arrays to simple 2D/1D targets (limited support)
func broadcastToSimple(v environment.Value, target []int) (environment.Value, error) {
	// scalar
	if v.Type != environment.ArrayType || v.Arr == nil {
		// treat as scalar
		if len(target) == 0 {
			return v, nil
		}
		// create nested arrays of the correct shape filled with v
		size := product(target)
		flat := make([]environment.Value, size)
		for i := range flat {
			flat[i] = v
		}
		return buildFromFlat(flat, target)
	}
	// array -> support broadcasting 1D->2D or scalar dims
	srcShape := shapeOf(v)
	if len(srcShape) == 0 {
		return broadcastToSimple(environment.Value{}, target)
	}
	// if shapes equal just return copy
	if len(srcShape) == len(target) {
		eq := true
		for i := range srcShape {
			if srcShape[i] != target[i] {
				eq = false
				break
			}
		}
		if eq {
			// return copy
			flat := flattenValues(v)
			return buildFromFlat(flat, target)
		}
	}
	// support broadcasting 1D to (n,m) when len(srcShape)==1
	if len(srcShape) == 1 && len(target) == 2 {
		rowLen := srcShape[0]
		if target[1] != rowLen {
			return environment.NewNil(), fmt.Errorf("cannot broadcast: incompatible dimensions")
		}
		nrows := target[0]
		out := make([]environment.Value, nrows)
		for i := 0; i < nrows; i++ {
			// copy row
			row := make([]environment.Value, rowLen)
			copy(row, *v.Arr)
			out[i] = environment.NewArray(row)
		}
		return environment.NewArray(out), nil
	}
	// scalar to 1D
	if len(srcShape) == 0 && len(target) == 1 {
		n := target[0]
		arr := make([]environment.Value, n)
		for i := range arr {
			arr[i] = v
		}
		return environment.NewArray(arr), nil
	}
	return environment.NewNil(), fmt.Errorf("broadcast_to: unsupported broadcast case")
}

// broadcast_arrays: broadcast multiple arrays to a common shape (limited support)
func broadcastArraysSimple(arrs []environment.Value) ([]environment.Value, error) {
	// compute a target shape (support scalar, 1D -> 2D)
	maxDims := 0
	for _, a := range arrs {
		sh := shapeOf(a)
		if len(sh) > maxDims {
			maxDims = len(sh)
		}
	}
	// only support up to 2D
	if maxDims > 2 {
		return nil, fmt.Errorf("broadcast_arrays: only up to 2D supported")
	}
	// compute target shape by right-alignment
	var target []int
	for _, a := range arrs {
		sh := shapeOf(a)
		if len(sh) == 0 {
			// scalar skip
			continue
		}
		if len(sh) > len(target) {
			target = make([]int, len(sh))
			copy(target[len(target)-len(sh):], sh)
			copy(target, sh)
		}
	}
	if len(target) == 0 && len(arrs) > 0 {
		// all scalars -> target is scalar
		res := make([]environment.Value, len(arrs))
		for i, a := range arrs {
			res[i] = a
		}
		return res, nil
	}
	// fallback: if target still empty set to first non-empty shape
	if len(target) == 0 {
		if len(arrs) > 0 && len(shapeOf(arrs[0])) > 0 {
			target = shapeOf(arrs[0])
		}
	}
	results := make([]environment.Value, len(arrs))
	for i, a := range arrs {
		b, err := broadcastToSimple(a, target)
		if err != nil {
			return nil, err
		}
		results[i] = b
	}
	return results, nil
}

// --- elementwise / unary helpers for numeric mathx ops ---

func flattenNumbers(v environment.Value) ([]float64, error) {
	if v.Type == environment.NumberType {
		return []float64{v.Num}, nil
	}
	if v.Type != environment.ArrayType || v.Arr == nil {
		return nil, fmt.Errorf("expects number or array")
	}
	vals := flattenValues(v)
	out := make([]float64, 0, len(vals))
	for _, e := range vals {
		if e.Type != environment.NumberType {
			return nil, fmt.Errorf("array elements must be numbers")
		}
		out = append(out, e.Num)
	}
	return out, nil
}

// helper: build array of given shape using generator fn
func buildRandomArray(shape []int, gen func() float64) (environment.Value, error) {
	if len(shape) == 0 {
		return environment.NewNumber(gen()), nil
	}
	sz := product(shape)
	flat := make([]environment.Value, sz)
	for i := 0; i < sz; i++ {
		flat[i] = environment.NewNumber(gen())
	}
	return buildFromFlat(flat, shape)
}

func applyUnaryNumeric(name string, v environment.Value, op func(float64) float64) (environment.Value, error) {
	if v.Type != environment.ArrayType || v.Arr == nil {
		if v.Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s() expects number or array", name)
		}
		return environment.NewNumber(op(v.Num)), nil
	}
	flat := flattenValues(v)
	out := make([]environment.Value, len(flat))
	for i, e := range flat {
		if e.Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): array elements must be numbers", name)
		}
		out[i] = environment.NewNumber(op(e.Num))
	}
	return buildFromFlat(out, shapeOf(v))
}

func applyBinaryNumeric(name string, args []environment.Value, op func(a, b float64) float64) (environment.Value, error) {
	if len(args) < 2 {
		return environment.NewNil(), fmt.Errorf("%s() expects at least 2 arguments", name)
	}
	allNumbers := true
	for _, a := range args {
		if a.Type != environment.NumberType {
			allNumbers = false
			break
		}
	}
	if allNumbers {
		res := args[0].Num
		for i := 1; i < len(args); i++ {
			res = op(res, args[i].Num)
		}
		return environment.NewNumber(res), nil
	}
	// broadcast arrays/scalars to common shape
	bcast, err := broadcastArraysSimple(args)
	if err != nil {
		return environment.NewNil(), fmt.Errorf("%s(): %v", name, err)
	}
	if len(bcast) == 0 {
		return environment.NewNil(), fmt.Errorf("%s(): no inputs", name)
	}
	targetShape := shapeOf(bcast[0])
	if len(targetShape) == 0 {
		// scalar result
		if bcast[0].Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): expects numeric arguments", name)
		}
		res := bcast[0].Num
		for i := 1; i < len(bcast); i++ {
			if bcast[i].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("%s(): expects numeric arguments", name)
			}
			res = op(res, bcast[i].Num)
		}
		return environment.NewNumber(res), nil
	}
	// elementwise across flattened broadcasted arrays
	flats := make([][]environment.Value, len(bcast))
	for i, v := range bcast {
		flats[i] = flattenValues(v)
	}
	n := len(flats[0])
	outFlat := make([]environment.Value, n)
	for idx := 0; idx < n; idx++ {
		if flats[0][idx].Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): array elements must be numbers", name)
		}
		val := flats[0][idx].Num
		for j := 1; j < len(flats); j++ {
			if flats[j][idx].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("%s(): array elements must be numbers", name)
			}
			val = op(val, flats[j][idx].Num)
		}
		outFlat[idx] = environment.NewNumber(val)
	}
	return buildFromFlat(outFlat, targetShape)
}

// --- helpers for boolean / comparison operations ---

func applyUnaryPredicate(name string, v environment.Value, pred func(float64) bool) (environment.Value, error) {
	// scalar
	if v.Type != environment.ArrayType || v.Arr == nil {
		if v.Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s() expects number or array", name)
		}
		return environment.NewBool(pred(v.Num)), nil
	}
	flat := flattenValues(v)
	out := make([]environment.Value, len(flat))
	for i, e := range flat {
		if e.Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): array elements must be numbers", name)
		}
		out[i] = environment.NewBool(pred(e.Num))
	}
	return buildFromFlat(out, shapeOf(v))
}

func applyBinaryCompare(name string, a, b environment.Value, cmp func(x, y float64) bool) (environment.Value, error) {
	// both scalars
	if a.Type != environment.ArrayType && b.Type != environment.ArrayType {
		if a.Type != environment.NumberType || b.Type != environment.NumberType {
			// fallback to generic equality only for 'equal' handled elsewhere
			return environment.NewNil(), fmt.Errorf("%s(): expects numeric arguments", name)
		}
		return environment.NewBool(cmp(a.Num, b.Num)), nil
	}
	// broadcast
	bcast, err := broadcastArraysSimple([]environment.Value{a, b})
	if err != nil {
		return environment.NewNil(), fmt.Errorf("%s(): %v", name, err)
	}
	if len(bcast) == 0 {
		return environment.NewNil(), fmt.Errorf("%s(): no inputs", name)
	}
	targetShape := shapeOf(bcast[0])
	if len(targetShape) == 0 {
		// scalar
		if bcast[0].Type != environment.NumberType || bcast[1].Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): expects numeric arguments", name)
		}
		return environment.NewBool(cmp(bcast[0].Num, bcast[1].Num)), nil
	}
	fl0 := flattenValues(bcast[0])
	fl1 := flattenValues(bcast[1])
	n := len(fl0)
	outFlat := make([]environment.Value, n)
	for i := 0; i < n; i++ {
		if fl0[i].Type != environment.NumberType || fl1[i].Type != environment.NumberType {
			return environment.NewNil(), fmt.Errorf("%s(): array elements must be numbers", name)
		}
		outFlat[i] = environment.NewBool(cmp(fl0[i].Num, fl1[i].Num))
	}
	return buildFromFlat(outFlat, targetShape)
}

func applyBinaryTruth(name string, a, b environment.Value, op func(x, y bool) bool) (environment.Value, error) {
	// both scalars
	if a.Type != environment.ArrayType && b.Type != environment.ArrayType {
		return environment.NewBool(op(a.IsTruthy(), b.IsTruthy())), nil
	}
	bcast, err := broadcastArraysSimple([]environment.Value{a, b})
	if err != nil {
		return environment.NewNil(), fmt.Errorf("%s(): %v", name, err)
	}
	if len(bcast) == 0 {
		return environment.NewNil(), fmt.Errorf("%s(): no inputs", name)
	}
	targetShape := shapeOf(bcast[0])
	if len(targetShape) == 0 {
		return environment.NewBool(op(bcast[0].IsTruthy(), bcast[1].IsTruthy())), nil
	}
	fl0 := flattenValues(bcast[0])
	fl1 := flattenValues(bcast[1])
	n := len(fl0)
	outFlat := make([]environment.Value, n)
	for i := 0; i < n; i++ {
		outFlat[i] = environment.NewBool(op(fl0[i].IsTruthy(), fl1[i].IsTruthy()))
	}
	return buildFromFlat(outFlat, targetShape)
}

func applyUnaryTruth(name string, v environment.Value, op func(x bool) bool) (environment.Value, error) {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return environment.NewBool(op(v.IsTruthy())), nil
	}
	flat := flattenValues(v)
	out := make([]environment.Value, len(flat))
	for i, e := range flat {
		out[i] = environment.NewBool(op(e.IsTruthy()))
	}
	return buildFromFlat(out, shapeOf(v))
}

// castToType converts a single Value to the requested dtype string ("number","string","boolean").
func castToType(v environment.Value, dtype string) (environment.Value, error) {
	switch dtype {
	case "number":
		if v.Type == environment.NumberType {
			return v, nil
		}
		if v.Type == environment.BooleanType {
			if v.Bool {
				return environment.NewNumber(1), nil
			} else {
				return environment.NewNumber(0), nil
			}
		}
		if v.Type == environment.StringType {
			if s := v.Str; s != "" {
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("cannot cast string to number: %v", err)
				}
				return environment.NewNumber(f), nil
			}
			return environment.NewNumber(0), nil
		}
		return environment.NewNil(), fmt.Errorf("cannot cast %s to number", v.TypeName())
	case "string":
		return environment.NewString(v.String()), nil
	case "boolean":
		return environment.NewBool(v.IsTruthy()), nil
	default:
		return environment.NewNil(), fmt.Errorf("unsupported dtype: %s", dtype)
	}
}

// typeNameForValue returns canonical dtype string for a Value (number/string/boolean/array/object/null/unknown)
func typeNameForValue(v environment.Value) string {
	return v.TypeName()
}

// isSimpleTypeName checks canonical simple type names we support
func isSimpleTypeName(s string) bool {
	switch s {
	case "number", "string", "boolean", "array", "object", "null":
		return true
	}
	return false
}

// valueToInterface converts environment.Value -> plain Go types for JSON/file I/O
func valueToInterface(v environment.Value) (interface{}, error) {
	switch v.Type {
	case environment.NumberType:
		return v.Num, nil
	case environment.StringType:
		return v.Str, nil
	case environment.BooleanType:
		return v.Bool, nil
	case environment.NilType:
		return nil, nil
	case environment.ArrayType:
		if v.Arr == nil {
			return []interface{}{}, nil
		}
		out := make([]interface{}, len(*v.Arr))
		for i, e := range *v.Arr {
			iv, err := valueToInterface(e)
			if err != nil {
				return nil, err
			}
			out[i] = iv
		}
		return out, nil
	case environment.ObjectType:
		m := make(map[string]interface{})
		for _, k := range v.Obj.Keys {
			iv, err := valueToInterface(v.Obj.Entries[k])
			if err != nil {
				return nil, err
			}
			m[k] = iv
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupported value type for I/O: %s", v.TypeName())
	}
}

// interfaceToValue converts decoded JSON/Go types -> environment.Value
func interfaceToValue(i interface{}) environment.Value {
	switch t := i.(type) {
	case nil:
		return environment.NewNil()
	case float64:
		return environment.NewNumber(t)
	case string:
		return environment.NewString(t)
	case bool:
		return environment.NewBool(t)
	case []interface{}:
		arr := make([]environment.Value, len(t))
		for i := range t {
			arr[i] = interfaceToValue(t[i])
		}
		return environment.NewArray(arr)
	case map[string]interface{}:
		entries := make(map[string]environment.Value)
		keys := make([]string, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			entries[k] = interfaceToValue(t[k])
		}
		return environment.NewObject(entries, keys)
	default:
		// JSON decoder may produce numbers as other numeric types in some paths; handle fallback
		if num, ok := t.(int); ok {
			return environment.NewNumber(float64(num))
		}
		if num64, ok := t.(int64); ok {
			return environment.NewNumber(float64(num64))
		}
		return environment.NewString(fmt.Sprintf("%v", t))
	}
}

// ----------------- linalg helpers -----------------

func to1DFloatSlice(v environment.Value) ([]float64, error) {
	if v.Type == environment.NumberType {
		return []float64{v.Num}, nil
	}
	if v.Type != environment.ArrayType || v.Arr == nil {
		return nil, fmt.Errorf("expected number or 1D array")
	}
	out := make([]float64, len(*v.Arr))
	for i, e := range *v.Arr {
		if e.Type != environment.NumberType {
			return nil, fmt.Errorf("array elements must be numbers")
		}
		out[i] = e.Num
	}
	return out, nil
}

func to2DFloatSlice(v environment.Value) ([][]float64, error) {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return nil, fmt.Errorf("expected 2D array")
	}
	nrows := len(*v.Arr)
	if nrows == 0 {
		return [][]float64{}, nil
	}
	first := (*v.Arr)[0]
	if first.Type != environment.ArrayType || first.Arr == nil {
		return nil, fmt.Errorf("expected 2D array")
	}
	ncols := len(*first.Arr)
	out := make([][]float64, nrows)
	for i, rowVal := range *v.Arr {
		if rowVal.Type != environment.ArrayType || rowVal.Arr == nil {
			return nil, fmt.Errorf("to2D: rows must be arrays")
		}
		if len(*rowVal.Arr) != ncols {
			return nil, fmt.Errorf("to2D: irregular matrix")
		}
		row := make([]float64, ncols)
		for j, cell := range *rowVal.Arr {
			if cell.Type != environment.NumberType {
				return nil, fmt.Errorf("to2D: non-number element")
			}
			row[j] = cell.Num
		}
		out[i] = row
	}
	return out, nil
}

func from1DSlice(a []float64) environment.Value {
	out := make([]environment.Value, len(a))
	for i := range a {
		out[i] = environment.NewNumber(a[i])
	}
	return environment.NewArray(out)
}

func from2DFloatSlice(m [][]float64) environment.Value {
	rows := make([]environment.Value, len(m))
	for i := range m {
		r := make([]environment.Value, len(m[i]))
		for j := range m[i] {
			r[j] = environment.NewNumber(m[i][j])
		}
		rows[i] = environment.NewArray(r)
	}
	return environment.NewArray(rows)
}

// --- FFT / spectral helpers (1D/2D, simple implementations) ---

// convert an environment.Value (number or 1D array) to a slice of complex128
func envArrayToComplexSlice(v environment.Value) ([]complex128, error) {
	if v.Type == environment.NumberType {
		return []complex128{complex(v.Num, 0)}, nil
	}
	if v.Type != environment.ArrayType || v.Arr == nil {
		return nil, fmt.Errorf("expected number or 1D array")
	}
	out := make([]complex128, len(*v.Arr))
	for i, e := range *v.Arr {
		if e.Type == environment.NumberType {
			out[i] = complex(e.Num, 0)
			continue
		}
		// allow [real, imag] pair to represent complex
		if e.Type == environment.ArrayType && e.Arr != nil && len(*e.Arr) == 2 {
			r := (*e.Arr)[0]
			im := (*e.Arr)[1]
			if r.Type == environment.NumberType && im.Type == environment.NumberType {
				out[i] = complex(r.Num, im.Num)
				continue
			}
		}
		return nil, fmt.Errorf("array elements must be numbers or [real,imag] pairs")
	}
	return out, nil
}

func complexSliceToEnvArray(cs []complex128) environment.Value {
	out := make([]environment.Value, len(cs))
	for i, c := range cs {
		out[i] = environment.NewArray([]environment.Value{environment.NewNumber(real(c)), environment.NewNumber(imag(c))})
	}
	return environment.NewArray(out)
}

// naive DFT (fallback for non-power-of-two)
func dft(x []complex128) []complex128 {
	n := len(x)
	out := make([]complex128, n)
	for k := 0; k < n; k++ {
		var sum complex128
		for t := 0; t < n; t++ {
			angle := -2 * math.Pi * float64(k*t) / float64(n)
			sum += x[t] * complex(math.Cos(angle), math.Sin(angle))
		}
		out[k] = sum
	}
	return out
}

// recursive radix-2 FFT; falls back to DFT for odd lengths
func fftComplex(x []complex128) []complex128 {
	n := len(x)
	if n == 0 {
		return []complex128{}
	}
	if n == 1 {
		return []complex128{x[0]}
	}
	if n%2 != 0 {
		return dft(x)
	}
	// split even/odd
	e := make([]complex128, n/2)
	o := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		e[i] = x[2*i]
		o[i] = x[2*i+1]
	}
	Fe := fftComplex(e)
	Fo := fftComplex(o)
	out := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		t := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * Fo[k]
		out[k] = Fe[k] + t
		out[k+n/2] = Fe[k] - t
	}
	return out
}

func ifftComplex(x []complex128) []complex128 {
	n := len(x)
	if n == 0 {
		return []complex128{}
	}
	// conj(fft(conj(x))) / n
	cx := make([]complex128, n)
	for i := range x {
		cx[i] = cmplx.Conj(x[i])
	}
	y := fftComplex(cx)
	for i := range y {
		y[i] = cmplx.Conj(y[i]) / complex(float64(n), 0)
	}
	return y
}

func to2DComplexSlice(v environment.Value) ([][]complex128, error) {
	if v.Type != environment.ArrayType || v.Arr == nil {
		return nil, fmt.Errorf("expected 2D array")
	}
	nrows := len(*v.Arr)
	if nrows == 0 {
		return [][]complex128{}, nil
	}
	first := (*v.Arr)[0]
	if first.Type != environment.ArrayType || first.Arr == nil {
		return nil, fmt.Errorf("expected 2D array")
	}
	ncols := len(*first.Arr)
	out := make([][]complex128, nrows)
	for i, rowVal := range *v.Arr {
		if rowVal.Type != environment.ArrayType || rowVal.Arr == nil || len(*rowVal.Arr) != ncols {
			return nil, fmt.Errorf("to2DComplex: irregular matrix")
		}
		row := make([]complex128, ncols)
		for j, cell := range *rowVal.Arr {
			if cell.Type == environment.NumberType {
				row[j] = complex(cell.Num, 0)
				continue
			}
			if cell.Type == environment.ArrayType && cell.Arr != nil && len(*cell.Arr) == 2 {
				r := (*cell.Arr)[0]
				im := (*cell.Arr)[1]
				if r.Type == environment.NumberType && im.Type == environment.NumberType {
					row[j] = complex(r.Num, im.Num)
					continue
				}
			}
			return nil, fmt.Errorf("to2DComplex: expected number or [real,imag]")
		}
		out[i] = row
	}
	return out, nil
}

// roll a 1D slice by shift (positive = right)
func roll1D(vals []environment.Value, shift int) []environment.Value {
	n := len(vals)
	if n == 0 {
		return vals
	}
	shift = ((shift % n) + n) % n
	out := make([]environment.Value, n)
	for i := 0; i < n; i++ {
		out[i] = vals[(i+shift)%n]
	}
	return out
}

func fftshift1D(vals []environment.Value) []environment.Value {
	n := len(vals)
	if n == 0 {
		return vals
	}
	shift := n / 2
	return roll1D(vals, shift)
}

func ifftshift1D(vals []environment.Value) []environment.Value {
	n := len(vals)
	if n == 0 {
		return vals
	}
	// inverse shift (negative floor(n/2)) -> positive equivalent = ceil(n/2)
	shift := (n + 1) / 2
	return roll1D(vals, shift)
}

func matMulFloat(A, B [][]float64) ([][]float64, error) {
	if len(A) == 0 || len(B) == 0 {
		return nil, fmt.Errorf("matmul: empty input")
	}
	m := len(A)
	k := len(A[0])
	if k != len(B) {
		return nil, fmt.Errorf("matmul: shape mismatch")
	}
	n := len(B[0])
	C := make([][]float64, m)
	for i := 0; i < m; i++ {
		C[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for t := 0; t < k; t++ {
				s += A[i][t] * B[t][j]
			}
			C[i][j] = s
		}
	}
	return C, nil
}

func matVecMulFloat(A [][]float64, v []float64) ([]float64, error) {
	if len(A) == 0 {
		return nil, fmt.Errorf("matvec: empty matrix")
	}
	cols := len(A[0])
	if cols != len(v) {
		return nil, fmt.Errorf("matvec: shape mismatch")
	}
	out := make([]float64, len(A))
	for i := 0; i < len(A); i++ {
		s := 0.0
		for j := 0; j < cols; j++ {
			s += A[i][j] * v[j]
		}
		out[i] = s
	}
	return out, nil
}

func vecMatMulFloat(v []float64, B [][]float64) ([]float64, error) {
	if len(B) == 0 {
		return nil, fmt.Errorf("vecmat: empty matrix")
	}
	if len(v) != len(B) {
		return nil, fmt.Errorf("vecmat: shape mismatch")
	}
	n := len(B[0])
	out := make([]float64, n)
	for j := 0; j < n; j++ {
		s := 0.0
		for i := 0; i < len(v); i++ {
			s += v[i] * B[i][j]
		}
		out[j] = s
	}
	return out, nil
}

func detFloat(A [][]float64) (float64, error) {
	n := len(A)
	if n == 0 {
		return 1, nil
	}
	for i := 0; i < n; i++ {
		if len(A[i]) != n {
			return 0, fmt.Errorf("det: not square")
		}
	}
	// copy
	M := make([][]float64, n)
	for i := range A {
		M[i] = make([]float64, n)
		copy(M[i], A[i])
	}
	sign := 1.0
	for k := 0; k < n; k++ {
		// pivot
		p := k
		max := math.Abs(M[k][k])
		for i := k + 1; i < n; i++ {
			if math.Abs(M[i][k]) > max {
				p = i
				max = math.Abs(M[i][k])
			}
		}
		if max < 1e-12 {
			return 0, nil
		}
		if p != k {
			M[k], M[p] = M[p], M[k]
			sign = -sign
		}
		for i := k + 1; i < n; i++ {
			f := M[i][k] / M[k][k]
			for j := k + 1; j < n; j++ {
				M[i][j] -= f * M[k][j]
			}
		}
	}
	prod := sign
	for i := 0; i < n; i++ {
		prod *= M[i][i]
	}
	return prod, nil
}

func invFloat(A [][]float64) ([][]float64, error) {
	n := len(A)
	if n == 0 {
		return [][]float64{}, nil
	}
	for i := 0; i < n; i++ {
		if len(A[i]) != n {
			return nil, fmt.Errorf("inv: not square")
		}
	}
	// build augmented
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, 2*n)
		for j := 0; j < n; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][n+i] = 1.0
	}
	// gauss-jordan
	for i := 0; i < n; i++ {
		// pivot
		p := i
		max := math.Abs(aug[i][i])
		for r := i + 1; r < n; r++ {
			if math.Abs(aug[r][i]) > max {
				p = r
				max = math.Abs(aug[r][i])
			}
		}
		if max < 1e-12 {
			return nil, fmt.Errorf("matrix singular")
		}
		if p != i {
			aug[i], aug[p] = aug[p], aug[i]
		}
		// normalize
		pv := aug[i][i]
		for c := 0; c < 2*n; c++ {
			aug[i][c] /= pv
		}
		// eliminate
		for r := 0; r < n; r++ {
			if r == i {
				continue
			}
			f := aug[r][i]
			if f == 0 {
				continue
			}
			for c := 0; c < 2*n; c++ {
				aug[r][c] -= f * aug[i][c]
			}
		}
	}
	inv := make([][]float64, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			inv[i][j] = aug[i][n+j]
		}
	}
	return inv, nil
}

func solveLinear(A [][]float64, b []float64) ([]float64, error) {
	n := len(A)
	if n == 0 {
		return nil, fmt.Errorf("solve: empty matrix")
	}
	if len(b) != n {
		return nil, fmt.Errorf("solve: dimension mismatch")
	}
	for i := 0; i < n; i++ {
		if len(A[i]) != n {
			return nil, fmt.Errorf("solve: not square")
		}
	}
	// augment
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			aug[i][j] = A[i][j]
		}
		aug[i][n] = b[i]
	}
	// elimination
	for i := 0; i < n; i++ {
		p := i
		max := math.Abs(aug[i][i])
		for r := i + 1; r < n; r++ {
			if math.Abs(aug[r][i]) > max {
				p = r
				max = math.Abs(aug[r][i])
			}
		}
		if max < 1e-12 {
			return nil, fmt.Errorf("singular matrix")
		}
		if p != i {
			aug[i], aug[p] = aug[p], aug[i]
		}
		for r := i + 1; r < n; r++ {
			f := aug[r][i] / aug[i][i]
			for c := i; c <= n; c++ {
				aug[r][c] -= f * aug[i][c]
			}
		}
	}
	// back substitution
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		s := aug[i][n]
		for j := i + 1; j < n; j++ {
			s -= aug[i][j] * x[j]
		}
		x[i] = s / aug[i][i]
	}
	return x, nil
}

func qrGramSchmidt(A [][]float64) ([][]float64, [][]float64, error) {
	m := len(A)
	if m == 0 {
		return nil, nil, fmt.Errorf("qr: empty matrix")
	}
	n := len(A[0])
	// copy and validate rectangular
	for i := 0; i < m; i++ {
		if len(A[i]) != n {
			return nil, nil, fmt.Errorf("qr: irregular matrix")
		}
	}
	Q := make([][]float64, m)
	for i := 0; i < m; i++ {
		Q[i] = make([]float64, n)
	}
	R := make([][]float64, n)
	for i := 0; i < n; i++ {
		R[i] = make([]float64, n)
	}
	// process columns
	for j := 0; j < n; j++ {
		// v = column j of A
		v := make([]float64, m)
		for i := 0; i < m; i++ {
			v[i] = A[i][j]
		}
		for i := 0; i < j; i++ {
			// R[i][j] = dot(Q[:,i], A[:,j])
			s := 0.0
			for k := 0; k < m; k++ {
				s += Q[k][i] * A[k][j]
			}
			R[i][j] = s
			for k := 0; k < m; k++ {
				v[k] -= s * Q[k][i]
			}
		}
		// R[j][j] = norm(v)
		rjj := 0.0
		for k := 0; k < m; k++ {
			rjj += v[k] * v[k]
		}
		rjj = math.Sqrt(rjj)
		if rjj < 1e-12 {
			return nil, nil, fmt.Errorf("qr: rank deficient")
		}
		R[j][j] = rjj
		for k := 0; k < m; k++ {
			Q[k][j] = v[k] / rjj
		}
	}
	return Q, R, nil
}

func choleskyDecomp(A [][]float64) ([][]float64, error) {
	n := len(A)
	if n == 0 {
		return [][]float64{}, nil
	}
	for i := 0; i < n; i++ {
		if len(A[i]) != n {
			return nil, fmt.Errorf("cholesky: not square")
		}
	}
	L := make([][]float64, n)
	for i := 0; i < n; i++ {
		L[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			s := A[i][j]
			for k := 0; k < j; k++ {
				s -= L[i][k] * L[j][k]
			}
			if i == j {
				if s <= 0 {
					return nil, fmt.Errorf("cholesky: not positive definite")
				}
				L[i][j] = math.Sqrt(s)
			} else {
				L[i][j] = s / L[j][j]
			}
		}
	}
	return L, nil
}

func matrixRankFloat(A [][]float64) (int, error) {
	m := len(A)
	if m == 0 {
		return 0, nil
	}
	n := len(A[0])
	// copy
	M := make([][]float64, m)
	for i := range A {
		M[i] = make([]float64, n)
		copy(M[i], A[i])
	}
	rank := 0
	eps := 1e-9
	for col := 0; col < n && rank < m; col++ {
		// find pivot
		pivot := -1
		maxv := 0.0
		for r := rank; r < m; r++ {
			if math.Abs(M[r][col]) > maxv {
				maxv = math.Abs(M[r][col])
				pivot = r
			}
		}
		if pivot == -1 || maxv <= eps {
			continue
		}
		M[rank], M[pivot] = M[pivot], M[rank]
		// normalize
		pv := M[rank][col]
		for c := col; c < n; c++ {
			M[rank][c] /= pv
		}
		// eliminate
		for r := 0; r < m; r++ {
			if r == rank {
				continue
			}
			fac := M[r][col]
			for c := col; c < n; c++ {
				M[r][c] -= fac * M[rank][c]
			}
		}
		rank++
	}
	return rank, nil
}

func eig2x2(A [][]float64) ([]float64, [][]float64, error) {
	if len(A) == 1 && len(A[0]) == 1 {
		return []float64{A[0][0]}, [][]float64{{1}}, nil
	}
	if len(A) != 2 || len(A[0]) != 2 {
		return nil, nil, fmt.Errorf("eig: only 1x1 or 2x2 supported")
	}
	a := A[0][0]
	b := A[0][1]
	c := A[1][0]
	d := A[1][1]
	tr := (a + d)
	det := a*d - b*c
	term := math.Sqrt((tr*tr)/4.0 - det)
	lambda1 := tr/2.0 + term
	lambda2 := tr/2.0 - term
	// eigenvectors
	v1 := []float64{1, 0}
	v2 := []float64{0, 1}
	if math.Abs(b) > 1e-12 || math.Abs(c) > 1e-12 {
		// solve (A - lambda I) v = 0: pick vector
		if math.Abs(b) > 1e-12 {
			v1 = []float64{lambda1 - d, b}
			v2 = []float64{lambda2 - d, b}
		} else {
			v1 = []float64{c, lambda1 - a}
			v2 = []float64{c, lambda2 - a}
		}
	}
	return []float64{lambda1, lambda2}, [][]float64{v1, v2}, nil
}

// ---------------------------------------------------------------

func init() {
	register(newModule("mathx",
		// array(...vals) -> wraps arguments into an array or returns argument if already array
		fn("array", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 1 && args[0].Type == environment.ArrayType {
				return args[0], nil
			}
			return environment.NewArray(args), nil
		}),

		fn("asarray", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("asarray() expects 1 argument, got %d", len(args))
			}
			v := args[0]
			if v.Type == environment.ArrayType {
				return v, nil
			}
			return environment.NewArray([]environment.Value{v}), nil
		}),

		// alias for asarray
		fn("asanyarray", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("asanyarray() expects 1 argument, got %d", len(args))
			}
			v := args[0]
			if v.Type == environment.ArrayType {
				return v, nil
			}
			return environment.NewArray([]environment.Value{v}), nil
		}),

	// copy(value) -> returns a shallow copy of arrays/objects (immutable scalars returned as-is)
	fn("copy", func(args []environment.Value) (environment.Value, error) {
		if len(args) != 1 { return environment.NewNil(), fmt.Errorf("copy() expects 1 argument") }
		v := args[0]
		switch v.Type {
		case environment.ArrayType:
			if v.Arr == nil { return environment.NewArray([]environment.Value{}), nil }
			cp := make([]environment.Value, len(*v.Arr))
			copy(cp, *v.Arr)
			return environment.NewArray(cp), nil
		case environment.ObjectType:
			if v.Obj == nil { return environment.NewObject(map[string]environment.Value{}, []string{}), nil }
			entries := make(map[string]environment.Value)
			keys := make([]string, len(v.Obj.Keys))
			copy(keys, v.Obj.Keys)
			for _, k := range keys { entries[k] = v.Obj.Entries[k] }
			return environment.NewObject(entries, keys), nil
		default:
			// scalars/others returned as-is
			return v, nil
		}
	}),

	fn("zeros", func(args []environment.Value) (environment.Value, error) {
		if len(args) != 1 {
			return environment.NewNil(), fmt.Errorf("zeros() expects 1 argument (length), got %d", len(args))
		}
		n, err := args[0].AsNumber()
		if err != nil {
			return environment.NewNil(), fmt.Errorf("zeros() argument must be a number")
		}
		if n < 0 {
			return environment.NewNil(), fmt.Errorf("zeros() length must be non-negative")
		}
		arr := make([]environment.Value, int(n))
		for i := range arr {
			arr[i] = environment.NewNumber(0)
		}
		return environment.NewArray(arr), nil
	}),

		fn("zeros_like", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("zeros_like() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("zeros_like() argument must be an array")
			}
			n := len(*args[0].Arr)
			arr := make([]environment.Value, n)
			for i := range arr {
				arr[i] = environment.NewNumber(0)
			}
			return environment.NewArray(arr), nil
		}),

		fn("ones", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ones() expects 1 argument (length), got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("ones() argument must be a number")
			}
			arr := make([]environment.Value, int(n))
			for i := range arr {
				arr[i] = environment.NewNumber(1)
			}
			return environment.NewArray(arr), nil
		}),

		fn("ones_like", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("ones_like() expects 1 array argument")
			}
			n := len(*args[0].Arr)
			arr := make([]environment.Value, n)
			for i := range arr {
				arr[i] = environment.NewNumber(1)
			}
			return environment.NewArray(arr), nil
		}),

		fn("empty", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("empty() expects 1 argument (length), got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("empty() argument must be a number")
			}
			arr := make([]environment.Value, int(n))
			for i := range arr {
				arr[i] = environment.NewNil()
			}
			return environment.NewArray(arr), nil
		}),

		fn("empty_like", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 || args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("empty_like() expects 1 array argument")
			}
			n := len(*args[0].Arr)
			arr := make([]environment.Value, n)
			for i := range arr {
				arr[i] = environment.NewNil()
			}
			return environment.NewArray(arr), nil
		}),

		fn("full", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("full() expects 2 arguments (length, fill_value), got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("full() first argument must be a number")
			}
			val := args[1]
			arr := make([]environment.Value, int(n))
			for i := range arr {
				arr[i] = val
			}
			return environment.NewArray(arr), nil
		}),

		fn("full_like", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 || args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("full_like() expects (array, fill_value)")
			}
			n := len(*args[0].Arr)
			val := args[1]
			arr := make([]environment.Value, n)
			for i := range arr {
				arr[i] = val
			}
			return environment.NewArray(arr), nil
		}),

		fn("arange", func(args []environment.Value) (environment.Value, error) {
			// arange(stop) or arange(start, stop[, step])
			if len(args) == 0 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("arange() expects 1..3 numeric arguments")
			}
			if len(args) == 1 {
				stop, err := args[0].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("arange() arguments must be numbers")
				}
				start := 0.0
				step := 1.0
				var out []environment.Value
				for v := start; v < stop; v += step {
					out = append(out, environment.NewNumber(v))
				}
				return environment.NewArray(out), nil
			}
			start, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("arange() arguments must be numbers")
			}
			stop, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("arange() arguments must be numbers")
			}
			step := 1.0
			if len(args) == 3 {
				s, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("arange() arguments must be numbers")
				}
				step = s
			}
			if step == 0 {
				return environment.NewNil(), fmt.Errorf("arange() step cannot be zero")
			}
			var out []environment.Value
			if step > 0 {
				for v := start; v < stop; v += step {
					out = append(out, environment.NewNumber(v))
				}
			} else {
				for v := start; v > stop; v += step {
					out = append(out, environment.NewNumber(v))
				}
			}
			return environment.NewArray(out), nil
		}),

		fn("linspace", func(args []environment.Value) (environment.Value, error) {
			// linspace(start, stop, num)
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("linspace() expects 2..3 arguments")
			}
			start, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("linspace() start must be a number")
			}
			stop, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("linspace() stop must be a number")
			}
			num := 50
			if len(args) == 3 {
				nv, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("linspace() num must be a number")
				}
				num = int(nv)
			}
			if num <= 0 {
				return environment.NewNil(), fmt.Errorf("linspace() num must be > 0")
			}
			out := make([]environment.Value, num)
			if num == 1 {
				out[0] = environment.NewNumber(stop)
				return environment.NewArray(out), nil
			}
			step := (stop - start) / float64(num-1)
			for i := 0; i < num; i++ {
				out[i] = environment.NewNumber(start + float64(i)*step)
			}
			return environment.NewArray(out), nil
		}),

		fn("logspace", func(args []environment.Value) (environment.Value, error) {
			// logspace(start, stop, num, base=10) -> base**linspace(start,stop,num)
			if len(args) < 2 || len(args) > 4 {
				return environment.NewNil(), fmt.Errorf("logspace() expects 2..4 arguments")
			}
			start, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("logspace() start must be a number")
			}
			stop, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("logspace() stop must be a number")
			}
			num := 50
			if len(args) >= 3 {
				nv, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("logspace() num must be a number")
				}
				num = int(nv)
			}
			base := 10.0
			if len(args) == 4 {
				bv, err := args[3].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("logspace() base must be a number")
				}
				base = bv
			}
			// generate linear space of exponents
			if num <= 0 {
				return environment.NewNil(), fmt.Errorf("logspace() num must be > 0")
			}
			out := make([]environment.Value, num)
			if num == 1 {
				out[0] = environment.NewNumber(math.Pow(base, stop))
				return environment.NewArray(out), nil
			}
			step := (stop - start) / float64(num-1)
			for i := 0; i < num; i++ {
				exp := start + float64(i)*step
				out[i] = environment.NewNumber(math.Pow(base, exp))
			}
			return environment.NewArray(out), nil
		}),

		fn("geomspace", func(args []environment.Value) (environment.Value, error) {
			// geomspace(start, stop, num)
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("geomspace() expects 2..3 arguments")
			}
			start, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("geomspace() start must be a number")
			}
			stop, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("geomspace() stop must be a number")
			}
			num := 50
			if len(args) == 3 {
				nv, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("geomspace() num must be a number")
				}
				num = int(nv)
			}
			if num <= 0 {
				return environment.NewNil(), fmt.Errorf("geomspace() num must be > 0")
			}
			logStart := math.Log(start)
			logStop := math.Log(stop)
			out := make([]environment.Value, num)
			if num == 1 {
				out[0] = environment.NewNumber(stop)
				return environment.NewArray(out), nil
			}
			step := (logStop - logStart) / float64(num-1)
			for i := 0; i < num; i++ {
				out[i] = environment.NewNumber(math.Exp(logStart + float64(i)*step))
			}
			return environment.NewArray(out), nil
		}),

		fn("eye", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("eye() expects 1 argument (n), got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("eye() argument must be a number")
			}
			N := int(n)
			out := make([]environment.Value, N)
			for i := 0; i < N; i++ {
				row := make([]environment.Value, N)
				for j := 0; j < N; j++ {
					if i == j {
						row[j] = environment.NewNumber(1)
					} else {
						row[j] = environment.NewNumber(0)
					}
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("identity", func(args []environment.Value) (environment.Value, error) {
			// same as eye(n)
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("identity() expects 1 argument (n), got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("identity() argument must be a number")
			}
			N := int(n)
			out := make([]environment.Value, N)
			for i := 0; i < N; i++ {
				row := make([]environment.Value, N)
				for j := 0; j < N; j++ {
					if i == j {
						row[j] = environment.NewNumber(1)
					} else {
						row[j] = environment.NewNumber(0)
					}
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("diag", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("diag() expects 1 argument")
			}
			v := args[0]
			if v.Type == environment.ArrayType && v.Arr != nil {
				// check if it's 1D of numbers -> produce 2D diagonal matrix
				is1D := true
				for _, e := range *v.Arr {
					if e.Type == environment.ArrayType {
						is1D = false
						break
					}
				}
				if is1D {
					n := len(*v.Arr)
					out := make([]environment.Value, n)
					for i := 0; i < n; i++ {
						row := make([]environment.Value, n)
						for j := 0; j < n; j++ {
							if i == j {
								row[j] = (*v.Arr)[i]
							} else {
								t := environment.NewNumber(0)
								row[j] = t
							}
						}
						out[i] = environment.NewArray(row)
					}
					return environment.NewArray(out), nil
				}
				// otherwise assume 2D matrix -> extract diagonal
				n := len(*v.Arr)
				diag := make([]environment.Value, 0, n)
				for i := 0; i < n; i++ {
					row := (*v.Arr)[i]
					if row.Type != environment.ArrayType || row.Arr == nil || i >= len(*row.Arr) {
						diag = append(diag, environment.NewNil())
						continue
					}
					diag = append(diag, (*row.Arr)[i])
				}
				return environment.NewArray(diag), nil
			}
			return environment.NewNil(), fmt.Errorf("diag() expects an array")
		}),

		fn("diagflat", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("diagflat() expects 1 argument")
			}
			v := args[0]
			// flatten input
			var flat []environment.Value
			if v.Type == environment.ArrayType && v.Arr != nil {
				for _, e := range *v.Arr {
					if e.Type == environment.ArrayType && e.Arr != nil {
						for _, s := range *e.Arr {
							flat = append(flat, s)
						}
					} else {
						flat = append(flat, e)
					}
				}
			} else {
				flat = []environment.Value{v}
			}
			// then diag
			n := len(flat)
			out := make([]environment.Value, n)
			for i := 0; i < n; i++ {
				row := make([]environment.Value, n)
				for j := 0; j < n; j++ {
					if i == j {
						row[j] = flat[i]
					} else {
						row[j] = environment.NewNumber(0)
					}
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("fromfunction", func(args []environment.Value) (environment.Value, error) {
			// fromfunction(shape, fn)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("fromfunction() expects 2 arguments (shape, fn)")
			}
			shape := args[0]
			fnv := args[1]
			if !isCallable(fnv) {
				return environment.NewNil(), fmt.Errorf("fromfunction() second argument must be a function")
			}
			if shape.Type == environment.NumberType {
				n := int(shape.Num)
				out := make([]environment.Value, n)
				for i := 0; i < n; i++ {
					res, err := invokeFn(fnv, []environment.Value{environment.NewNumber(float64(i))})
					if err != nil {
						return environment.NewNil(), err
					}
					out[i] = res
				}
				return environment.NewArray(out), nil
			}
			if shape.Type == environment.ArrayType && shape.Arr != nil && len(*shape.Arr) == 2 {
				r := int((*shape.Arr)[0].Num)
				c := int((*shape.Arr)[1].Num)
				out := make([]environment.Value, r)
				for i := 0; i < r; i++ {
					row := make([]environment.Value, c)
					for j := 0; j < c; j++ {
						res, err := invokeFn(fnv, []environment.Value{environment.NewNumber(float64(i)), environment.NewNumber(float64(j))})
						if err != nil {
							return environment.NewNil(), err
						}
						row[j] = res
					}
					out[i] = environment.NewArray(row)
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("fromfunction() unsupported shape")
		}),

		fn("fromiter", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("fromiter() expects iterable and optional count")
			}
			it := args[0]
			if it.Type == environment.ArrayType && it.Arr != nil {
				// return shallow copy
				cp := make([]environment.Value, len(*it.Arr))
				copy(cp, *it.Arr)
				return environment.NewArray(cp), nil
			}
			if it.Type == environment.StringType {
				// produce array of characters as single-character strings
				runes := []rune(it.Str)
				out := make([]environment.Value, len(runes))
				for i, r := range runes {
					out[i] = environment.NewString(string(r))
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("fromiter() argument must be array or string")
		}),

		fn("frombuffer", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("frombuffer() expects 1 argument")
			}
			if args[0].Type == environment.StringType {
				b := []byte(args[0].Str)
				out := make([]environment.Value, len(b))
				for i := range b {
					out[i] = environment.NewNumber(float64(b[i]))
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("frombuffer() expects a string buffer")
		}),

		// shape(arr)
		fn("shape", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("shape() expects 1 argument")
			}
			sh := shapeOf(args[0])
			out := make([]environment.Value, len(sh))
			for i := range sh {
				out[i] = environment.NewNumber(float64(sh[i]))
			}
			return environment.NewArray(out), nil
		}),

		// reshape(arr, shapeArray)
		fn("reshape", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("reshape() expects 2 arguments")
			}
			shapeArg := args[1]
			if shapeArg.Type != environment.ArrayType || shapeArg.Arr == nil {
				return environment.NewNil(), fmt.Errorf("reshape() shape must be an array")
			}
			shape := make([]int, len(*shapeArg.Arr))
			for i, s := range *shapeArg.Arr {
				if s.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("reshape() shape elements must be numbers")
				}
				shape[i] = int(s.Num)
			}
			return reshapeValue(args[0], shape)
		}),

		// ravel / flatten
		fn("ravel", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ravel() expects 1 argument")
			}
			flat := flattenValues(args[0])
			return environment.NewArray(flat), nil
		}),
		fn("flatten", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["ravel"].Builtin(args)
		}),

		// transpose
		fn("transpose", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("transpose() expects 1 argument")
			}
			return transpose2D(args[0])
		}),

		// swapaxes (limited: supports swapping 0 and 1 for 2D)
		fn("swapaxes", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("swapaxes() expects 3 arguments (array, axis1, axis2)")
			}
			axis1, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("swapaxes: axis must be numeric")
			}
			axis2, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("swapaxes: axis must be numeric")
			}
			if int(axis1) == 0 && int(axis2) == 1 {
				return transpose2D(args[0])
			}
			return environment.NewNil(), fmt.Errorf("swapaxes: only 2D swap supported")
		}),

		// moveaxis (limited: 2D move)
		fn("moveaxis", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("moveaxis() expects 3 arguments (array, src, dst)")
			}
			src, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveaxis: src must be numeric")
			}
			dst, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("moveaxis: dst must be numeric")
			}
			// only support moving 0 <-> 1
			if int(src) == 0 && int(dst) == 1 {
				return transpose2D(args[0])
			}
			if int(src) == int(dst) {
				return args[0], nil
			}
			return environment.NewNil(), fmt.Errorf("moveaxis: unsupported axes for moveaxis")
		}),

		// expand_dims(array, axis)
		fn("expand_dims", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("expand_dims() expects 2 arguments (array, axis)")
			}
			axis, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("expand_dims: axis must be numeric")
			}
			ax := int(axis)
			// axis == 0 -> wrap as outer array
			if ax == 0 {
				return environment.NewArray([]environment.Value{args[0]}), nil
			}
			// axis == 1 and 1D array -> convert [a,b] -> [[a],[b]]
			if ax == 1 && args[0].Type == environment.ArrayType && args[0].Arr != nil {
				n := len(*args[0].Arr)
				out := make([]environment.Value, n)
				for i := 0; i < n; i++ {
					out[i] = environment.NewArray([]environment.Value{(*args[0].Arr)[i]})
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("expand_dims: unsupported axis or input")
		}),

		// squeeze(array) - remove dimensions of size 1
		fn("squeeze", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("squeeze() expects 1 argument")
			}
			v := args[0]
			// if top-level array of length 1, return its only element; do this recursively
			for v.Type == environment.ArrayType && v.Arr != nil && len(*v.Arr) == 1 {
				v = (*v.Arr)[0]
			}
			return v, nil
		}),

		// broadcast_to(value, shapeArray)
		fn("broadcast_to", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("broadcast_to() expects 2 arguments")
			}
			shapeArg := args[1]
			if shapeArg.Type != environment.ArrayType || shapeArg.Arr == nil {
				return environment.NewNil(), fmt.Errorf("broadcast_to() shape must be an array")
			}
			shape := make([]int, len(*shapeArg.Arr))
			for i, s := range *shapeArg.Arr {
				if s.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("broadcast_to() shape elements must be numbers")
				}
				shape[i] = int(s.Num)
			}
			return broadcastToSimple(args[0], shape)
		}),

		// broadcast_arrays(a1, a2, ...)
		fn("broadcast_arrays", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewNil(), fmt.Errorf("broadcast_arrays() expects at least one argument")
			}
			res, err := broadcastArraysSimple(args)
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewArray(res), nil
		}),

		// tile(array, reps)
		fn("tile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("tile() expects 2 arguments (array, reps)")
			}
			repsF, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("tile: reps must be numeric")
			}
			return tile1D(args[0], int(repsF))
		}),

		// repeat(array, repeats)
		fn("repeat", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("repeat() expects 2 arguments (array, repeats)")
			}
			repsF, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("repeat: repeats must be numeric")
			}
			return repeat1D(args[0], int(repsF))
		}),

		// concatenate(...arrays)
		fn("concatenate", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("concatenate() expects at least 1 argument")
			}
			return concatAlongAxis(args, 0)
		}),

		// stack(...arrays)
		fn("stack", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("stack() expects at least 1 argument")
			}
			return stackArrays(args)
		}),

		// vstack/hstack/dstack
		fn("vstack", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			// if inputs are 1D, stack as rows
			all1D := true
			for _, a := range args {
				if a.Type != environment.ArrayType || a.Arr == nil {
					all1D = false
					break
				}
				if len(*a.Arr) > 0 && (*a.Arr)[0].Type == environment.ArrayType {
					all1D = false
					break
				}
			}
			if all1D {
				rows := make([]environment.Value, len(args))
				for i := range args {
					rows[i] = args[i]
				}
				return environment.NewArray(rows), nil
			}
			// otherwise vertical concat
			return concatAlongAxis(args, 0)
		}),

		fn("hstack", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			// if all 1D: concatenate
			all1D := true
			for _, a := range args {
				if a.Type != environment.ArrayType || a.Arr == nil {
					all1D = false
					break
				}
				if len(*a.Arr) > 0 && (*a.Arr)[0].Type == environment.ArrayType {
					all1D = false
					break
				}
			}
			if all1D {
				return concatAlongAxis(args, 0)
			}
			return concatAlongAxis(args, 1)
		}),

		fn("dstack", func(args []environment.Value) (environment.Value, error) {
			// simple implementation: for two 2D arrays of same shape produce 3D where each element is [a,b]
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("dstack() expects at least 1 argument")
			}
			if len(args) == 1 {
				return args[0], nil
			}
			base := args[0]
			if base.Type != environment.ArrayType || base.Arr == nil {
				return environment.NewNil(), fmt.Errorf("dstack: expects 2D arrays")
			}
			r := len(*base.Arr)
			if r == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			c := 0
			if (*base.Arr)[0].Type == environment.ArrayType {
				c = len(*(*base.Arr)[0].Arr)
			}
			// ensure all same shape
			for _, a := range args {
				if a.Type != environment.ArrayType || a.Arr == nil || len(*a.Arr) != r {
					return environment.NewNil(), fmt.Errorf("dstack: inputs must have same shape")
				}
			}
			out := make([]environment.Value, r)
			for i := 0; i < r; i++ {
				row := make([]environment.Value, c)
				for j := 0; j < c; j++ {
					cell := make([]environment.Value, len(args))
					for k, a := range args {
						cell[k] = (*(*a.Arr)[i].Arr)[j]
					}
					row[j] = environment.NewArray(cell)
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("column_stack", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("column_stack() expects at least 1 argument")
			}
			// all args must be 1D of same length
			n := -1
			for _, a := range args {
				if a.Type != environment.ArrayType || a.Arr == nil {
					return environment.NewNil(), fmt.Errorf("column_stack: expects 1D arrays")
				}
				if n == -1 {
					n = len(*a.Arr)
				} else if len(*a.Arr) != n {
					return environment.NewNil(), fmt.Errorf("column_stack: all arrays must have same length")
				}
			}
			out := make([]environment.Value, n)
			for i := 0; i < n; i++ {
				row := make([]environment.Value, len(args))
				for j := 0; j < len(args); j++ {
					row[j] = (*args[j].Arr)[i]
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("row_stack", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["vstack"].Builtin(args)
		}),

		fn("split", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("split() expects 2 arguments (array, sections)")
			}
			sections, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("split: sections must be numeric")
			}
			parts, err := split1D(args[0], int(sections))
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewArray(parts), nil
		}),

		fn("array_split", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("array_split() expects 2 arguments (array, sections)")
			}
			sections, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("array_split: sections must be numeric")
			}
			parts, err := arraySplit1D(args[0], int(sections))
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewArray(parts), nil
		}),

		fn("hsplit", func(args []environment.Value) (environment.Value, error) {
			// split columns of 2D array into 'sections' parts
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("hsplit() expects 2 arguments (matrix, sections)")
			}
			mat := args[0]
			sections, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("hsplit: sections must be numeric")
			}
			if mat.Type != environment.ArrayType || mat.Arr == nil {
				return environment.NewNil(), fmt.Errorf("hsplit: expects 2D array")
			}
			if len(*mat.Arr) == 0 {
				return environment.NewNil(), fmt.Errorf("hsplit: empty matrix")
			}
			// number of columns
			firstRow := (*mat.Arr)[0]
			if firstRow.Type != environment.ArrayType || firstRow.Arr == nil {
				return environment.NewNil(), fmt.Errorf("hsplit: expects 2D array")
			}
			ncols := len(*firstRow.Arr)
			parts, err := arraySplit1D(environment.NewArray(func() []environment.Value {
				out := make([]environment.Value, ncols)
				for i := 0; i < ncols; i++ {
					out[i] = environment.NewNumber(float64(i))
				}
				return out
			}()), int(sections))
			if err != nil {
				return environment.NewNil(), err
			}
			// build column slices according to parts
			result := make([]environment.Value, len(parts))
			for pi, p := range parts {
				idxs := *p.Arr
				cols := make([]environment.Value, len(*mat.Arr))
				for r := 0; r < len(*mat.Arr); r++ {
					row := (*mat.Arr)[r]
					newRow := make([]environment.Value, len(idxs))
					for ci := 0; ci < len(idxs); ci++ {
						colIdx := int(idxs[ci].Num)
						newRow[ci] = (*row.Arr)[colIdx]
					}
					cols[r] = environment.NewArray(newRow)
				}
				result[pi] = environment.NewArray(cols)
			}
			return environment.NewArray(result), nil
		}),

		fn("vsplit", func(args []environment.Value) (environment.Value, error) {
			// split rows of 2D array
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("vsplit() expects 2 arguments (matrix, sections)")
			}
			mat := args[0]
			sections, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("vsplit: sections must be numeric")
			}
			if mat.Type != environment.ArrayType || mat.Arr == nil {
				return environment.NewNil(), fmt.Errorf("vsplit: expects 2D array")
			}
			parts, err := arraySplit1D(mat, int(sections))
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewArray(parts), nil
		}),

		fn("dsplit", func(args []environment.Value) (environment.Value, error) {
			// limited: for 2D behaves like hsplit
			return registry["mathx"].Entries["hsplit"].Builtin(args)
		}),

		// ----------------- Indexing / selection helpers -----------------
		fn("take", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("take() expects 2 arguments (array, indices)")
			}
			arr := args[0]
			if arr.Type != environment.ArrayType || arr.Arr == nil {
				return environment.NewNil(), fmt.Errorf("take() first argument must be an array")
			}
			idx := args[1]
			if idx.Type == environment.NumberType {
				i := int(idx.Num)
				if i < 0 {
					i += len(*arr.Arr)
				}
				if i < 0 || i >= len(*arr.Arr) {
					return environment.NewNil(), fmt.Errorf("take() index out of range")
				}
				return (*arr.Arr)[i], nil
			}
			if idx.Type == environment.ArrayType && idx.Arr != nil {
				var out []environment.Value
				for _, e := range *idx.Arr {
					if e.Type != environment.NumberType {
						return environment.NewNil(), fmt.Errorf("take() indices must be numbers")
					}
					i := int(e.Num)
					if i < 0 {
						i += len(*arr.Arr)
					}
					if i < 0 || i >= len(*arr.Arr) {
						return environment.NewNil(), fmt.Errorf("take() index out of range")
					}
					out = append(out, (*arr.Arr)[i])
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("take() invalid indices argument")
		}),

		fn("put", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("put() expects 3 arguments (array, indices, values)")
			}
			arr := args[0]
			if arr.Type != environment.ArrayType || arr.Arr == nil {
				return environment.NewNil(), fmt.Errorf("put() first argument must be an array")
			}
			idx := args[1]
			vals := args[2]

			if idx.Type == environment.NumberType {
				i := int(idx.Num)
				if i < 0 {
					i += len(*arr.Arr)
				}
				if i < 0 || i >= len(*arr.Arr) {
					return environment.NewNil(), fmt.Errorf("put() index out of range")
				}
				(*arr.Arr)[i] = vals
				return arr, nil
			}

			if idx.Type == environment.ArrayType && idx.Arr != nil {
				indices := *idx.Arr
				// values can be scalar or array
				if vals.Type == environment.ArrayType && vals.Arr != nil {
					vlist := *vals.Arr
					if len(vlist) == len(indices) {
						for k, idv := range indices {
							if idv.Type != environment.NumberType {
								return environment.NewNil(), fmt.Errorf("put() indices must be numbers")
							}
							i := int(idv.Num)
							if i < 0 {
								i += len(*arr.Arr)
							}
							if i < 0 || i >= len(*arr.Arr) {
								return environment.NewNil(), fmt.Errorf("put() index out of range")
							}
							(*arr.Arr)[i] = vlist[k]
						}
						return arr, nil
					}
					// if vlist length == 1, broadcast
					if len(vlist) == 1 {
						for _, idv := range indices {
							if idv.Type != environment.NumberType {
								return environment.NewNil(), fmt.Errorf("put() indices must be numbers")
							}
							i := int(idv.Num)
							if i < 0 {
								i += len(*arr.Arr)
							}
							if i < 0 || i >= len(*arr.Arr) {
								return environment.NewNil(), fmt.Errorf("put() index out of range")
							}
							(*arr.Arr)[i] = vlist[0]
						}
						return arr, nil
					}
					return environment.NewNil(), fmt.Errorf("put() values length mismatch")
				}
				// scalar values -> assign scalar to all indices
				for _, idv := range indices {
					if idv.Type != environment.NumberType {
						return environment.NewNil(), fmt.Errorf("put() indices must be numbers")
					}
					i := int(idv.Num)
					if i < 0 {
						i += len(*arr.Arr)
					}
					if i < 0 || i >= len(*arr.Arr) {
						return environment.NewNil(), fmt.Errorf("put() index out of range")
					}
					(*arr.Arr)[i] = vals
				}
				return arr, nil
			}
			return environment.NewNil(), fmt.Errorf("put() invalid indices argument")
		}),

		fn("where", func(args []environment.Value) (environment.Value, error) {
			// where(cond) -> indices; where(cond, x, y) -> elementwise choose
			if len(args) == 1 {
				cond := args[0]
				// 1D or 2D
				if cond.Type != environment.ArrayType || cond.Arr == nil {
					return environment.NewNil(), fmt.Errorf("where() expects an array condition")
				}
				// if 1D
				if len(*cond.Arr) == 0 || ((*cond.Arr)[0].Type != environment.ArrayType) {
					var idxs []environment.Value
					for i, e := range *cond.Arr {
						if e.IsTruthy() {
							idxs = append(idxs, environment.NewNumber(float64(i)))
						}
					}
					return environment.NewArray(idxs), nil
				}
				// 2D -> return pairs
				var pairs []environment.Value
				for r, row := range *cond.Arr {
					if row.Type != environment.ArrayType || row.Arr == nil {
						continue
					}
					for c, cell := range *row.Arr {
						if cell.IsTruthy() {
							pairs = append(pairs, environment.NewArray([]environment.Value{environment.NewNumber(float64(r)), environment.NewNumber(float64(c))}))
						}
					}
				}
				return environment.NewArray(pairs), nil
			}
			if len(args) == 3 {
				cond, x, y := args[0], args[1], args[2]
				if cond.Type != environment.ArrayType || cond.Arr == nil {
					return environment.NewNil(), fmt.Errorf("where() first argument must be condition array")
				}
				if x.Type != environment.ArrayType || x.Arr == nil || y.Type != environment.ArrayType || y.Arr == nil {
					return environment.NewNil(), fmt.Errorf("where() x and y must be arrays")
				}
				if len(*cond.Arr) != len(*x.Arr) || len(*x.Arr) != len(*y.Arr) {
					return environment.NewNil(), fmt.Errorf("where(): arrays must have same length")
				}
				out := make([]environment.Value, len(*cond.Arr))
				for i := range *cond.Arr {
					if (*cond.Arr)[i].IsTruthy() {
						out[i] = (*x.Arr)[i]
					} else {
						out[i] = (*y.Arr)[i]
					}
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("where() invalid arguments")
		}),

		fn("nonzero", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("nonzero() expects 1 argument")
			}
			v := args[0]
			// 1D -> return array of indices
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewNil(), fmt.Errorf("nonzero() expects an array")
			}
			if len(*v.Arr) == 0 || ((*v.Arr)[0].Type != environment.ArrayType) {
				var idxs []environment.Value
				for i, e := range *v.Arr {
					if e.IsTruthy() {
						idxs = append(idxs, environment.NewNumber(float64(i)))
					}
				}
				return environment.NewArray(idxs), nil
			}
			// 2D: return [rows_array, cols_array]
			var rows []environment.Value
			var cols []environment.Value
			for r, row := range *v.Arr {
				if row.Type != environment.ArrayType || row.Arr == nil {
					continue
				}
				for c, cell := range *row.Arr {
					if cell.IsTruthy() {
						rows = append(rows, environment.NewNumber(float64(r)))
						cols = append(cols, environment.NewNumber(float64(c)))
					}
				}
			}
			return environment.NewArray([]environment.Value{environment.NewArray(rows), environment.NewArray(cols)}), nil
		}),

		fn("argwhere", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("argwhere() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewNil(), fmt.Errorf("argwhere() expects an array")
			}
			// 1D -> return [[idx], ...]
			if len(*v.Arr) == 0 || ((*v.Arr)[0].Type != environment.ArrayType) {
				var out []environment.Value
				for i, e := range *v.Arr {
					if e.IsTruthy() {
						out = append(out, environment.NewArray([]environment.Value{environment.NewNumber(float64(i))}))
					}
				}
				return environment.NewArray(out), nil
			}
			// 2D -> return [[r,c], ...]
			var out []environment.Value
			for r, row := range *v.Arr {
				if row.Type != environment.ArrayType || row.Arr == nil {
					continue
				}
				for c, cell := range *row.Arr {
					if cell.IsTruthy() {
						out = append(out, environment.NewArray([]environment.Value{environment.NewNumber(float64(r)), environment.NewNumber(float64(c))}))
					}
				}
			}
			return environment.NewArray(out), nil
		}),

		fn("extract", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("extract() expects 2 arguments (condition, array)")
			}
			cond, arr := args[0], args[1]
			if cond.Type != environment.ArrayType || cond.Arr == nil || arr.Type != environment.ArrayType || arr.Arr == nil {
				return environment.NewNil(), fmt.Errorf("extract() expects arrays")
			}
			if len(*cond.Arr) != len(*arr.Arr) {
				return environment.NewNil(), fmt.Errorf("extract(): arrays must have same length")
			}
			var out []environment.Value
			for i := range *cond.Arr {
				if (*cond.Arr)[i].IsTruthy() {
					out = append(out, (*arr.Arr)[i])
				}
			}
			return environment.NewArray(out), nil
		}),

		fn("select", func(args []environment.Value) (environment.Value, error) {
			// select(condList, choiceList, default)
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("select() expects 2 or 3 arguments")
			}
			condList := args[0]
			choiceList := args[1]
			var def environment.Value = environment.NewNumber(0)
			if len(args) == 3 {
				def = args[2]
			}
			// condList and choiceList must be arrays of same length
			if condList.Type != environment.ArrayType || condList.Arr == nil || choiceList.Type != environment.ArrayType || choiceList.Arr == nil {
				return environment.NewNil(), fmt.Errorf("select() expects two lists of arrays")
			}
			if len(*condList.Arr) != len(*choiceList.Arr) {
				return environment.NewNil(), fmt.Errorf("select(): condList and choiceList must have same number of entries")
			}
			if len(*condList.Arr) == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			// determine length from first condition array
			firstCond := (*condList.Arr)[0]
			if firstCond.Type != environment.ArrayType || firstCond.Arr == nil {
				return environment.NewNil(), fmt.Errorf("select(): condition entries must be arrays")
			}
			n := len(*firstCond.Arr)
			out := make([]environment.Value, n)
			for i := 0; i < n; i++ {
				chosen := def
				for j := 0; j < len(*condList.Arr); j++ {
					c := (*condList.Arr)[j]
					if c.Type != environment.ArrayType || c.Arr == nil || i >= len(*c.Arr) {
						continue
					}
					if (*c.Arr)[i].IsTruthy() {
						ch := (*choiceList.Arr)[j]
						if ch.Type == environment.ArrayType && ch.Arr != nil {
							if i < len(*ch.Arr) {
								chosen = (*ch.Arr)[i]
							} else {
								chosen = environment.NewNil()
							}
						} else {
							chosen = ch
						}
						break
					}
				}
				out[i] = chosen
			}
			return environment.NewArray(out), nil
		}),

		fn("choose", func(args []environment.Value) (environment.Value, error) {
			// choose(indices, choices)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("choose() expects 2 arguments")
			}
			indices := args[0]
			choices := args[1]
			if choices.Type != environment.ArrayType || choices.Arr == nil {
				return environment.NewNil(), fmt.Errorf("choose() choices must be an array")
			}
			// scalar index
			if indices.Type == environment.NumberType {
				i := int(indices.Num)
				if i < 0 {
					i += len(*choices.Arr)
				}
				if i < 0 || i >= len(*choices.Arr) {
					return environment.NewNil(), fmt.Errorf("choose() index out of range")
				}
				return (*choices.Arr)[i], nil
			}
			// array of indices
			if indices.Type == environment.ArrayType && indices.Arr != nil {
				out := make([]environment.Value, len(*indices.Arr))
				for k, idv := range *indices.Arr {
					if idv.Type != environment.NumberType {
						return environment.NewNil(), fmt.Errorf("choose() indices must be numbers")
					}
					i := int(idv.Num)
					if i < 0 {
						i += len(*choices.Arr)
					}
					if i < 0 || i >= len(*choices.Arr) {
						return environment.NewNil(), fmt.Errorf("choose() index out of range")
					}
					chosen := (*choices.Arr)[i]
					// if chosen is array and has element at k, pick it, else pick scalar
					if chosen.Type == environment.ArrayType && chosen.Arr != nil {
						if k < len(*chosen.Arr) {
							out[k] = (*chosen.Arr)[k]
						} else {
							out[k] = environment.NewNil()
						}
					} else {
						out[k] = chosen
					}
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("choose() unsupported argument types")
		}),

		fn("compress", func(args []environment.Value) (environment.Value, error) {
			// alias of extract(condition, array)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("compress() expects 2 arguments (condition, array)")
			}
			return registry["mathx"].Entries["extract"].Builtin([]environment.Value{args[0], args[1]})
		}),

		// --- numeric elementwise / scalar ops ---
		fn("add", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("add", args, func(a, b float64) float64 { return a + b })
		}),

		fn("subtract", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("subtract", args, func(a, b float64) float64 { return a - b })
		}),

		fn("multiply", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("multiply", args, func(a, b float64) float64 { return a * b })
		}),

		fn("divide", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("divide", args, func(a, b float64) float64 { return a / b })
		}),

		fn("true_divide", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("true_divide", args, func(a, b float64) float64 { return a / b })
		}),

		fn("floor_divide", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("floor_divide", args, func(a, b float64) float64 { return math.Floor(a / b) })
		}),

		fn("power", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("power", args, func(a, b float64) float64 { return math.Pow(a, b) })
		}),

		fn("mod", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("mod", args, func(a, b float64) float64 { return math.Mod(a, b) })
		}),

		fn("remainder", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("remainder", args, func(a, b float64) float64 { return math.Remainder(a, b) })
		}),

		// --- polynomial helpers ---

		fn("polyval", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 { return environment.NewNil(), fmt.Errorf("polyval() expects 2 arguments (coeffs, x)") }
			coeffs, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), fmt.Errorf("polyval(): coeffs must be numeric array or number") }
			x := args[1]
			// Horner's method (coeffs in highest-first order)
			evalAt := func(xx float64) float64 {
				res := 0.0
				for _, c := range coeffs {
					res = res*xx + c
				}
				return res
			}
			if x.Type == environment.ArrayType && x.Arr != nil {
				out := make([]environment.Value, len(*x.Arr))
				for i, e := range *x.Arr {
					if e.Type != environment.NumberType { return environment.NewNil(), fmt.Errorf("polyval(): x array must be numeric") }
					out[i] = environment.NewNumber(evalAt(e.Num))
				}
				return environment.NewArray(out), nil
			}
			if x.Type != environment.NumberType { return environment.NewNil(), fmt.Errorf("polyval(): x must be number or array") }
			return environment.NewNumber(evalAt(x.Num)), nil
		}),

		fn("polyadd", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 { return environment.NewNil(), fmt.Errorf("polyadd() expects 2 arguments") }
			a, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), err }
			b, err := flattenNumbers(args[1]); if err != nil { return environment.NewNil(), err }
			// align to left (highest-first) by padding shorter with leading zeros
			la := len(a); lb := len(b); lm := la
			if lb > lm { lm = lb }
			r := make([]environment.Value, lm)
			pa := make([]float64, lm); pb := make([]float64, lm)
			copy(pa[lm-la:], a); copy(pb[lm-lb:], b)
			for i := 0; i < lm; i++ { r[i] = environment.NewNumber(pa[i] + pb[i]) }
			return environment.NewArray(r), nil
		}),

		fn("polysub", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 { return environment.NewNil(), fmt.Errorf("polysub() expects 2 arguments") }
			a, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), err }
			b, err := flattenNumbers(args[1]); if err != nil { return environment.NewNil(), err }
			la := len(a); lb := len(b); lm := la
			if lb > lm { lm = lb }
			r := make([]environment.Value, lm)
			pa := make([]float64, lm); pb := make([]float64, lm)
			copy(pa[lm-la:], a); copy(pb[lm-lb:], b)
			for i := 0; i < lm; i++ { r[i] = environment.NewNumber(pa[i] - pb[i]) }
			return environment.NewArray(r), nil
		}),

		fn("polymul", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 { return environment.NewNil(), fmt.Errorf("polymul() expects 2 arguments") }
			a, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), err }
			b, err := flattenNumbers(args[1]); if err != nil { return environment.NewNil(), err }
			la := len(a); lb := len(b)
			if la == 0 || lb == 0 { return environment.NewArray([]environment.Value{}), nil }
			// operate in reverse (constant-term first), then reverse result
			ra := make([]float64, la); rb := make([]float64, lb)
			for i := 0; i < la; i++ { ra[i] = a[la-1-i] }
			for i := 0; i < lb; i++ { rb[i] = b[lb-1-i] }
			rc := make([]float64, la+lb-1)
			for i := 0; i < la; i++ {
				for j := 0; j < lb; j++ {
					rc[i+j] += ra[i] * rb[j]
				}
			}
			out := make([]environment.Value, len(rc))
			for i := 0; i < len(rc); i++ { out[i] = environment.NewNumber(rc[len(rc)-1-i]) }
			return environment.NewArray(out), nil
		}),

		fn("polyder", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("polyder() expects 1 argument") }
			a, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), err }
			n := len(a) - 1
			if n <= 0 { return environment.NewArray([]environment.Value{environment.NewNumber(0)}), nil }
			out := make([]environment.Value, n)
			for i := 0; i < n; i++ {
				exp := float64(n - i)
				out[i] = environment.NewNumber(a[i] * exp)
			}
			return environment.NewArray(out), nil
		}),

		fn("polyint", func(args []environment.Value) (environment.Value, error) {
			// polyint(coeffs[, k]) -> integrated coefficients (highest-first), optional integration constant k (number)
			if len(args) < 1 || len(args) > 2 { return environment.NewNil(), fmt.Errorf("polyint() expects 1 or 2 arguments") }
			a, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), err }
			k := 0.0
			if len(args) == 2 {
				kv, err := args[1].AsNumber(); if err != nil { return environment.NewNil(), fmt.Errorf("polyint(): constant must be number") }
				k = kv
			}
			m := len(a)
			out := make([]environment.Value, m+1)
			for i := 0; i < m; i++ {
				exp := float64(m - i)
				out[i] = environment.NewNumber(a[i] / exp)
			}
			out[m] = environment.NewNumber(k)
			return environment.NewArray(out), nil
		}),

		fn("polyfit", func(args []environment.Value) (environment.Value, error) {
			// polyfit(x, y, deg)
			if len(args) != 3 { return environment.NewNil(), fmt.Errorf("polyfit() expects 3 arguments (x, y, deg)") }
			xv, err := flattenNumbers(args[0]); if err != nil { return environment.NewNil(), fmt.Errorf("polyfit(): x must be numeric array") }
			yv, err := flattenNumbers(args[1]); if err != nil { return environment.NewNil(), fmt.Errorf("polyfit(): y must be numeric array") }
			degF, err := args[2].AsNumber(); if err != nil { return environment.NewNil(), fmt.Errorf("polyfit(): deg must be number") }
			deg := int(degF)
			n := len(xv)
			if n != len(yv) { return environment.NewNil(), fmt.Errorf("polyfit(): x and y must have same length") }
			if n == 0 || deg < 0 { return environment.NewNil(), fmt.Errorf("polyfit(): invalid inputs") }
			if n <= deg { return environment.NewNil(), fmt.Errorf("polyfit(): number of points must be greater than deg") }
			// build design matrix A (n x (deg+1)) with powers x^(deg - j)
			p := deg + 1
			A := make([][]float64, n)
			for i := 0; i < n; i++ {
				row := make([]float64, p)
				for j := 0; j < p; j++ {
					row[j] = math.Pow(xv[i], float64(deg-j))
				}
				A[i] = row
			}
			// compute normal equations: ATA * c = ATy
			ATA := make([][]float64, p)
			for i := 0; i < p; i++ {
				ATA[i] = make([]float64, p)
				for j := 0; j < p; j++ {
					s := 0.0
					for k := 0; k < n; k++ { s += A[k][i] * A[k][j] }
					ATA[i][j] = s
				}
			}
			ATy := make([]float64, p)
			for i := 0; i < p; i++ {
				s := 0.0
				for k := 0; k < n; k++ { s += A[k][i] * yv[k] }
				ATy[i] = s
			}
			coeffs, err := solveLinear(ATA, ATy)
			if err != nil { return environment.NewNil(), fmt.Errorf("polyfit(): solve failed: %v", err) }
			out := make([]environment.Value, len(coeffs))
			for i := range coeffs { out[i] = environment.NewNumber(coeffs[i]) }
			return environment.NewArray(out), nil
		}),

		// --- utility helpers: ndim/size/itemsize/copyto/view and printoptions ---
		fn("ndim", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("ndim() expects 1 argument") }
			sh := shapeOf(args[0])
			return environment.NewNumber(float64(len(sh))), nil
		}),

		fn("size", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("size() expects 1 argument") }
			sh := shapeOf(args[0])
			return environment.NewNumber(float64(product(sh))), nil
		}),

		fn("itemsize", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("itemsize() expects 1 argument") }
			var infer func(environment.Value) int
			infer = func(v environment.Value) int {
				switch v.Type {
				case environment.NumberType:
					return 8
				case environment.BooleanType:
					return 1
				case environment.StringType:
					return len(v.Str)
				case environment.ArrayType:
					if v.Arr == nil || len(*v.Arr) == 0 { return 0 }
					for _, e := range *v.Arr {
						if e.Type != environment.NilType { return infer(e) }
					}
					return 0
				default:
					return 0
				}
			}
			sz := infer(args[0])
			return environment.NewNumber(float64(sz)), nil
		}),

		fn("copyto", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 { return environment.NewNil(), fmt.Errorf("copyto() expects 2 arguments (dest, src)") }
			dest := args[0]
			src := args[1]
			if dest.Type != environment.ArrayType || dest.Arr == nil { return environment.NewNil(), fmt.Errorf("copyto(): dest must be array") }
			// scalar src -> fill dest
			if src.Type != environment.ArrayType || src.Arr == nil {
				for i := range *dest.Arr { (*dest.Arr)[i] = src }
				return dest, nil
			}
			// both arrays -> lengths must match
			srcFlat := flattenValues(src)
			if len(srcFlat) != len(*dest.Arr) { return environment.NewNil(), fmt.Errorf("copyto(): source and destination must have same size") }
			for i := range srcFlat { (*dest.Arr)[i] = srcFlat[i] }
			return dest, nil
		}),

		fn("view", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("view() expects 1 argument") }
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil { return environment.NewNil(), fmt.Errorf("view(): expects array") }
			// return the same array value (shared backing)
			return v, nil
		}),

		fn("get_printoptions", func(args []environment.Value) (environment.Value, error) {
			m := map[string]environment.Value{
				"precision": environment.NewNumber(float64(mathxPrintOptions.Precision)),
				"linewidth": environment.NewNumber(float64(mathxPrintOptions.LineWidth)),
				"threshold": environment.NewNumber(float64(mathxPrintOptions.Threshold)),
			}
			keys := []string{"precision", "linewidth", "threshold"}
			return environment.NewObject(m, keys), nil
		}),

		fn("set_printoptions", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("set_printoptions() expects 1 argument (object)") }
			if args[0].Type != environment.ObjectType || args[0].Obj == nil { return environment.NewNil(), fmt.Errorf("set_printoptions(): argument must be object") }
			for _, k := range args[0].Obj.Keys {
				v := args[0].Obj.Entries[k]
				switch k {
				case "precision":
					if n, err := v.AsNumber(); err == nil { mathxPrintOptions.Precision = int(n) }
				case "linewidth":
					if n, err := v.AsNumber(); err == nil { mathxPrintOptions.LineWidth = int(n) }
				case "threshold":
					if n, err := v.AsNumber(); err == nil { mathxPrintOptions.Threshold = int(n) }
				}
			}
			return environment.NewNil(), nil
		}),

		fn("negative", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("negative() expects 1 argument") }
			return applyUnaryNumeric("negative", args[0], func(x float64) float64 { return -x })
		}),

		fn("sign", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 { return environment.NewNil(), fmt.Errorf("sign() expects 1 argument") }
			return applyUnaryNumeric("sign", args[0], func(x float64) float64 {
				if x > 0 { return 1 }
				if x < 0 { return -1 }
				return 0
			})
		}),

		fn("absolute", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("absolute() expects 1 argument")
			}
			return applyUnaryNumeric("absolute", args[0], math.Abs)
		}),

		fn("fabs", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fabs() expects 1 argument")
			}
			return applyUnaryNumeric("fabs", args[0], math.Abs)
		}),

		fn("sign", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sign() expects 1 argument")
			}
			return applyUnaryNumeric("sign", args[0], func(x float64) float64 {
				if math.IsNaN(x) {
					return math.NaN()
				}
				if x > 0 {
					return 1
				}
				if x < 0 {
					return -1
				}
				return 0
			})
		}),

		fn("sqrt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sqrt() expects 1 argument")
			}
			return applyUnaryNumeric("sqrt", args[0], math.Sqrt)
		}),

		fn("square", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("square() expects 1 argument")
			}
			return applyUnaryNumeric("square", args[0], func(x float64) float64 { return x * x })
		}),

		fn("reciprocal", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("reciprocal() expects 1 argument")
			}
			return applyUnaryNumeric("reciprocal", args[0], func(x float64) float64 { return 1.0 / x })
		}),

		fn("clip", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("clip() expects 3 arguments (array_or_scalar, min, max)")
			}
			// broadcast the three inputs
			b, err := broadcastArraysSimple(args)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("clip(): %v", err)
			}
			// scalar case
			if len(shapeOf(b[0])) == 0 {
				val, ok := b[0].Num, true
				minv, ok2 := b[1].Num, true
				maxv, ok3 := b[2].Num, true
				if !ok || !ok2 || !ok3 {
					return environment.NewNil(), fmt.Errorf("clip(): expects numeric arguments")
				}
				if minv > maxv {
					return environment.NewNil(), fmt.Errorf("clip(): min cannot be greater than max")
				}
				if val < minv {
					val = minv
				} else if val > maxv {
					val = maxv
				}
				return environment.NewNumber(val), nil
			}
			// array case
			flatV := flattenValues(b[0])
			flatMin := flattenValues(b[1])
			flatMax := flattenValues(b[2])
			if len(flatV) != len(flatMin) || len(flatV) != len(flatMax) {
				return environment.NewNil(), fmt.Errorf("clip(): broadcast produced mismatched shapes")
			}
			out := make([]environment.Value, len(flatV))
			for i := range flatV {
				if flatV[i].Type != environment.NumberType || flatMin[i].Type != environment.NumberType || flatMax[i].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("clip(): inputs must be numeric")
				}
				mn := flatMin[i].Num
				mx := flatMax[i].Num
				if mn > mx {
					return environment.NewNil(), fmt.Errorf("clip(): min cannot be greater than max")
				}
				v := flatV[i].Num
				if v < mn {
					v = mn
				} else if v > mx {
					v = mx
				}
				out[i] = environment.NewNumber(v)
			}
			return buildFromFlat(out, shapeOf(b[0]))
		}),

		fn("maximum", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("maximum", args, math.Max)
		}),

		fn("minimum", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("minimum", args, math.Min)
		}),

		// --- Comparisons & logicals ---
		fn("equal", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("equal() expects 2 arguments")
			}
			// both scalars
			if args[0].Type != environment.ArrayType && args[1].Type != environment.ArrayType {
				return environment.NewBool(valuesEqualBuiltin(args[0], args[1])), nil
			}
			bcast, err := broadcastArraysSimple([]environment.Value{args[0], args[1]})
			if err != nil {
				return environment.NewNil(), fmt.Errorf("equal(): %v", err)
			}
			if len(bcast) == 0 {
				return environment.NewNil(), fmt.Errorf("equal(): no inputs")
			}
			shape := shapeOf(bcast[0])
			if len(shape) == 0 {
				return environment.NewBool(valuesEqualBuiltin(bcast[0], bcast[1])), nil
			}
			f0 := flattenValues(bcast[0])
			f1 := flattenValues(bcast[1])
			out := make([]environment.Value, len(f0))
			for i := range f0 {
				out[i] = environment.NewBool(valuesEqualBuiltin(f0[i], f1[i]))
			}
			return buildFromFlat(out, shape)
		}),

		fn("not_equal", func(args []environment.Value) (environment.Value, error) {
			res, err := registry["mathx"].Entries["equal"].Builtin(args)
			if err != nil {
				return environment.NewNil(), err
			}
			// invert
			if res.Type != environment.ArrayType || res.Arr == nil {
				return environment.NewBool(!res.IsTruthy()), nil
			}
			flat := flattenValues(res)
			out := make([]environment.Value, len(flat))
			for i := range flat {
				out[i] = environment.NewBool(!flat[i].IsTruthy())
			}
			return buildFromFlat(out, shapeOf(res))
		}),

		fn("greater", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("greater() expects 2 arguments")
			}
			return applyBinaryCompare("greater", args[0], args[1], func(x, y float64) bool { return x > y })
		}),

		fn("greater_equal", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("greater_equal() expects 2 arguments")
			}
			return applyBinaryCompare("greater_equal", args[0], args[1], func(x, y float64) bool { return x >= y })
		}),

		fn("less", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("less() expects 2 arguments")
			}
			return applyBinaryCompare("less", args[0], args[1], func(x, y float64) bool { return x < y })
		}),

		fn("less_equal", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("less_equal() expects 2 arguments")
			}
			return applyBinaryCompare("less_equal", args[0], args[1], func(x, y float64) bool { return x <= y })
		}),

		fn("logical_and", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("logical_and() expects 2 arguments")
			}
			return applyBinaryTruth("logical_and", args[0], args[1], func(x, y bool) bool { return x && y })
		}),

		fn("logical_or", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("logical_or() expects 2 arguments")
			}
			return applyBinaryTruth("logical_or", args[0], args[1], func(x, y bool) bool { return x || y })
		}),

		fn("logical_xor", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("logical_xor() expects 2 arguments")
			}
			return applyBinaryTruth("logical_xor", args[0], args[1], func(x, y bool) bool { return x != y })
		}),

		fn("logical_not", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("logical_not() expects 1 argument")
			}
			return applyUnaryTruth("logical_not", args[0], func(x bool) bool { return !x })
		}),

		fn("all", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("all() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewBool(v.IsTruthy()), nil
			}
			flat := flattenValues(v)
			for _, e := range flat {
				if !e.IsTruthy() {
					return environment.NewBool(false), nil
				}
			}
			return environment.NewBool(true), nil
		}),

		fn("any", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("any() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewBool(v.IsTruthy()), nil
			}
			flat := flattenValues(v)
			for _, e := range flat {
				if e.IsTruthy() {
					return environment.NewBool(true), nil
				}
			}
			return environment.NewBool(false), nil
		}),

		fn("isfinite", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isfinite() expects 1 argument")
			}
			return applyUnaryPredicate("isfinite", args[0], func(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) })
		}),

		fn("isinf", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isinf() expects 1 argument")
			}
			return applyUnaryPredicate("isinf", args[0], func(x float64) bool { return math.IsInf(x, 0) })
		}),

		fn("isnan", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isnan() expects 1 argument")
			}
			return applyUnaryPredicate("isnan", args[0], func(x float64) bool { return math.IsNaN(x) })
		}),

		fn("isclose", func(args []environment.Value) (environment.Value, error) {
			// isclose(a, b [, rtol, atol])
			if len(args) < 2 || len(args) > 4 {
				return environment.NewNil(), fmt.Errorf("isclose() expects 2-4 arguments")
			}
			a := args[0]
			b := args[1]
			rtol := 1e-9
			atol := 0.0
			if len(args) >= 3 {
				if args[2].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("isclose(): rtol must be numeric")
				}
				rtol = args[2].Num
			}
			if len(args) == 4 {
				if args[3].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("isclose(): atol must be numeric")
				}
				atol = args[3].Num
			}
			// broadcast
			bcast, err := broadcastArraysSimple([]environment.Value{a, b})
			if err != nil {
				return environment.NewNil(), fmt.Errorf("isclose(): %v", err)
			}
			if len(bcast) == 0 {
				return environment.NewNil(), fmt.Errorf("isclose(): no inputs")
			}
			shape := shapeOf(bcast[0])
			if len(shape) == 0 {
				if bcast[0].Type != environment.NumberType || bcast[1].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("isclose(): numeric args expected")
				}
				d := math.Abs(bcast[0].Num - bcast[1].Num)
				return environment.NewBool(d <= (atol + rtol*math.Abs(bcast[1].Num))), nil
			}
			f0 := flattenValues(bcast[0])
			f1 := flattenValues(bcast[1])
			out := make([]environment.Value, len(f0))
			for i := range f0 {
				if f0[i].Type != environment.NumberType || f1[i].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("isclose(): numeric array elements expected")
				}
				d := math.Abs(f0[i].Num - f1[i].Num)
				out[i] = environment.NewBool(d <= (atol + rtol*math.Abs(f1[i].Num)))
			}
			return buildFromFlat(out, shape)
		}),

		// --- Sorting / set operations ---
		fn("sort", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sort() expects 1 argument")
			}
			// scalar -> return scalar
			if args[0].Type != environment.ArrayType {
				if args[0].Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("sort(): expects number or array")
				}
				return args[0], nil
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			sort.Float64s(nums)
			return from1DSlice(nums), nil
		}),

		fn("argsort", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("argsort() expects 1 argument")
			}
			if args[0].Type != environment.ArrayType && args[0].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("argsort(): expects number or array")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			n := len(nums)
			idx := make([]int, n)
			for i := 0; i < n; i++ {
				idx[i] = i
			}
			sort.SliceStable(idx, func(i, j int) bool { return nums[idx[i]] < nums[idx[j]] })
			out := make([]float64, n)
			for i := 0; i < n; i++ {
				out[i] = float64(idx[i])
			}
			return from1DSlice(out), nil
		}),

		fn("lexsort", func(args []environment.Value) (environment.Value, error) {
			// lexsort(keys_array) — keys_array is sequence/array of 1D key arrays
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("lexsort() expects 1 argument (array of keys)")
			}
			keysArr := args[0]
			if keysArr.Type != environment.ArrayType || keysArr.Arr == nil {
				return environment.NewNil(), fmt.Errorf("lexsort(): expects array of key arrays")
			}
			keys := *keysArr.Arr
			if len(keys) == 0 {
				return from1DSlice([]float64{}), nil
			}
			// convert keys to [][]float64
			ks := make([][]float64, len(keys))
			n := -1
			for i, k := range keys {
				arr, err := flattenNumbers(k)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("lexsort(): keys must be numeric 1D arrays: %v", err)
				}
				if n == -1 {
					n = len(arr)
				} else if len(arr) != n {
					return environment.NewNil(), fmt.Errorf("lexsort(): all keys must have same length")
				}
				ks[i] = arr
			}
			idx := make([]int, n)
			for i := 0; i < n; i++ {
				idx[i] = i
			}
			// compare by keys from last to first (numpy behaviour)
			sort.SliceStable(idx, func(i, j int) bool {
				for k := len(ks) - 1; k >= 0; k-- {
					if ks[k][idx[i]] < ks[k][idx[j]] {
						return true
					}
					if ks[k][idx[i]] > ks[k][idx[j]] {
						return false
					}
				}
				return false
			})
			out := make([]float64, n)
			for i := 0; i < n; i++ {
				out[i] = float64(idx[i])
			}
			return from1DSlice(out), nil
		}),

		fn("partition", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("partition() expects 2 arguments (array, kth)")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			kf, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("partition(): kth must be numeric")
			}
			k := int(kf)
			n := len(nums)
			if k < 0 || k >= n {
				return environment.NewNil(), fmt.Errorf("partition(): kth out of range")
			}
			// simple implementation: return fully sorted array (satisfies kth property)
			sorted := make([]float64, n)
			copy(sorted, nums)
			sort.Float64s(sorted)
			return from1DSlice(sorted), nil
		}),

		fn("argpartition", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("argpartition() expects 2 arguments (array, kth)")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			kf, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("argpartition(): kth must be numeric")
			}
			k := int(kf)
			n := len(nums)
			if k < 0 || k >= n {
				return environment.NewNil(), fmt.Errorf("argpartition(): kth out of range")
			}
			idx := make([]int, n)
			for i := 0; i < n; i++ {
				idx[i] = i
			}
			sort.SliceStable(idx, func(i, j int) bool { return nums[idx[i]] < nums[idx[j]] })
			// return indices (simple: full argsort)
			out := make([]float64, n)
			for i := 0; i < n; i++ {
				out[i] = float64(idx[i])
			}
			return from1DSlice(out), nil
		}),

		fn("unique", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("unique() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			set := make(map[float64]struct{})
			for _, v := range nums {
				set[v] = struct{}{}
			}
			out := make([]float64, 0, len(set))
			for v := range set {
				out = append(out, v)
			}
			sort.Float64s(out)
			return from1DSlice(out), nil
		}),

		fn("setdiff1d", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("setdiff1d() expects 2 arguments")
			}
			a, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b, err := flattenNumbers(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			bset := make(map[float64]struct{})
			for _, v := range b {
				bset[v] = struct{}{}
			}
			outSet := make(map[float64]struct{})
			for _, v := range a {
				if _, ok := bset[v]; !ok {
					outSet[v] = struct{}{}
				}
			}
			out := make([]float64, 0, len(outSet))
			for v := range outSet {
				out = append(out, v)
			}
			sort.Float64s(out)
			return from1DSlice(out), nil
		}),

		fn("intersect1d", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("intersect1d() expects 2 arguments")
			}
			a, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b, err := flattenNumbers(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			aset := make(map[float64]struct{})
			for _, v := range a {
				aset[v] = struct{}{}
			}
			out := make([]float64, 0)
			seen := make(map[float64]struct{})
			for _, v := range b {
				if _, ok := aset[v]; ok {
					if _, s := seen[v]; !s {
						out = append(out, v)
						seen[v] = struct{}{}
					}
				}
			}
			sort.Float64s(out)
			return from1DSlice(out), nil
		}),

		fn("union1d", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("union1d() expects 2 arguments")
			}
			a, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b, err := flattenNumbers(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			set := make(map[float64]struct{})
			for _, v := range a {
				set[v] = struct{}{}
			}
			for _, v := range b {
				set[v] = struct{}{}
			}
			out := make([]float64, 0, len(set))
			for v := range set {
				out = append(out, v)
			}
			sort.Float64s(out)
			return from1DSlice(out), nil
		}),

		fn("in1d", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("in1d() expects 2 arguments")
			}
			aVals, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			bVals, err := flattenNumbers(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			bset := make(map[float64]struct{})
			for _, v := range bVals {
				bset[v] = struct{}{}
			}
			out := make([]environment.Value, len(aVals))
			for i, v := range aVals {
				_, ok := bset[v]
				out[i] = environment.NewBool(ok)
			}
			return buildFromFlat(out, shapeOf(args[0]))
		}),

		// --- Types & casting helpers ---
		fn("dtype", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("dtype() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewString(typeNameForValue(v)), nil
			}
			// array: inspect element types
			flat := flattenValues(v)
			if len(flat) == 0 {
				return environment.NewString("null"), nil
			}
			first := typeNameForValue(flat[0])
			for _, e := range flat[1:] {
				if typeNameForValue(e) != first {
					return environment.NewString("mixed"), nil
				}
			}
			return environment.NewString(first), nil
		}),

		fn("astype", func(args []environment.Value) (environment.Value, error) {
			// astype(value, dtype)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("astype() expects 2 arguments")
			}
			v := args[0]
			dt, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("astype(): dtype must be string")
			}
			dt = strings.ToLower(strings.TrimSpace(dt))
			if v.Type != environment.ArrayType || v.Arr == nil {
				return castToType(v, dt)
			}
			flat := flattenValues(v)
			out := make([]environment.Value, len(flat))
			for i, e := range flat {
				cv, err := castToType(e, dt)
				if err != nil {
					return environment.NewNil(), err
				}
				out[i] = cv
			}
			return buildFromFlat(out, shapeOf(v))
		}),

		fn("issubdtype", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("issubdtype() expects 2 arguments")
			}
			// allow (value, dtype) or (dtype, dtype)
			var leftType string
			if args[0].Type == environment.StringType {
				leftType = strings.ToLower(strings.TrimSpace(args[0].Str))
			} else {
				leftType = typeNameForValue(args[0])
			}
			var rightType string
			if args[1].Type == environment.StringType {
				rightType = strings.ToLower(strings.TrimSpace(args[1].Str))
			} else {
				rightType = typeNameForValue(args[1])
			}
			// simple hierarchy: same -> true
			return environment.NewBool(leftType == rightType), nil
		}),

		fn("can_cast", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("can_cast() expects 2 arguments")
			}
			// two common call patterns:
			// 1) can_cast(dtypeName, dtypeName)  -- static rules
			// 2) can_cast(valueString, dtypeName) -- try casting the value
			if args[0].Type == environment.StringType && args[1].Type == environment.StringType {
				fromRaw := strings.ToLower(strings.TrimSpace(args[0].Str))
				to := strings.ToLower(strings.TrimSpace(args[1].Str))
				// if first string looks like a dtype name, treat both as dtype-names
				if isSimpleTypeName(fromRaw) && isSimpleTypeName(to) {
					if fromRaw == to {
						return environment.NewBool(true), nil
					}
					// allow number<->string, boolean <-> number/string
					if (fromRaw == "number" && to == "string") || (fromRaw == "string" && to == "number") {
						return environment.NewBool(true), nil
					}
					if (fromRaw == "boolean" && (to == "number" || to == "string")) || ((fromRaw == "number" || fromRaw == "string") && to == "boolean") {
						return environment.NewBool(true), nil
					}
					return environment.NewBool(false), nil
				}
				// otherwise treat args[0] as a literal string value to be cast to 'to'
				_, err := castToType(environment.NewString(fromRaw), to)
				return environment.NewBool(err == nil), nil
			}
			// if first arg is a non-string value, try casting the actual value
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("can_cast(): target dtype must be string")
			}
			dt := strings.ToLower(strings.TrimSpace(args[1].Str))
			_, err := castToType(args[0], dt)
			return environment.NewBool(err == nil), nil
		}),

		fn("result_type", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewString("null"), nil
			}
			// accept mix of dtype-strings or values
			hasString, hasNumber, hasBool := false, false, false
			for _, a := range args {
				var tn string
				if a.Type == environment.StringType {
					// if the string argument itself is a dtype name (e.g. "number"),
					// treat it as that dtype. Otherwise it's a runtime string value.
					s := strings.ToLower(strings.TrimSpace(a.Str))
					if isSimpleTypeName(s) {
						tn = s
					} else {
						tn = "string"
					}
				} else {
					tn = typeNameForValue(a)
				}
				switch tn {
				case "string":
					hasString = true
				case "number":
					hasNumber = true
				case "boolean":
					hasBool = true
				}
			}
			if hasString {
				return environment.NewString("string"), nil
			}
			if hasNumber {
				return environment.NewString("number"), nil
			}
			if hasBool {
				return environment.NewString("boolean"), nil
			}
			return environment.NewString("null"), nil
		}),

		// --- Random / sampling (np.random-like, limited) ---
		fn("rand", func(args []environment.Value) (environment.Value, error) {
			// rand() -> scalar; rand(n) -> 1D array; rand(d0, d1, ...) -> shaped array
			if len(args) == 0 {
				return environment.NewNumber(rand.Float64()), nil
			}
			// all args must be numbers (dimensions)
			shape := []int{}
			for _, a := range args {
				n, err := a.AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("rand(): shape dims must be numbers")
				}
				shape = append(shape, int(n))
			}
			return buildRandomArray(shape, func() float64 { return rand.Float64() })
		}),

		fn("random", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["rand"].Builtin(args)
		}),

		fn("randn", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewNumber(rand.NormFloat64()), nil
			}
			shape := []int{}
			for _, a := range args {
				n, err := a.AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("randn(): shape dims must be numbers")
				}
				shape = append(shape, int(n))
			}
			return buildRandomArray(shape, func() float64 { return rand.NormFloat64() })
		}),

		fn("randint", func(args []environment.Value) (environment.Value, error) {
			// randint(high) or randint(low, high) or randint(low, high, size)
			if len(args) < 1 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("randint() expects 1-3 arguments")
			}
			if len(args) == 1 {
				hi, err := args[0].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("randint(): numeric expected")
				}
				return environment.NewNumber(float64(rand.Intn(int(hi)))), nil
			}
			// len >=2
			low, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("randint(): numeric expected")
			}
			hi, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("randint(): numeric expected")
			}
			if len(args) == 2 {
				return environment.NewNumber(float64(int(low) + rand.Intn(int(hi)-int(low)))), nil
			}
			// size provided
			sz, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("randint(): size must be numeric")
			}
			n := int(sz)
			out := make([]environment.Value, n)
			for i := 0; i < n; i++ {
				out[i] = environment.NewNumber(float64(int(low) + rand.Intn(int(hi)-int(low))))
			}
			return environment.NewArray(out), nil
		}),

		fn("choice", func(args []environment.Value) (environment.Value, error) {
			// choice(array_or_int, size?)
			if len(args) == 0 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("choice() expects 1 or 2 arguments")
			}
			src := args[0]
			var size int = 0
			if len(args) == 2 {
				sz, err := args[1].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("choice(): size must be numeric")
				}
				size = int(sz)
			}
			// source numeric -> choose from 0..n-1
			if src.Type == environment.NumberType {
				n := int(src.Num)
				if n <= 0 {
					return environment.NewNil(), fmt.Errorf("choice(): range must be > 0")
				}
				if size == 0 {
					return environment.NewNumber(float64(rand.Intn(n))), nil
				}
				out := make([]environment.Value, size)
				for i := 0; i < size; i++ {
					out[i] = environment.NewNumber(float64(rand.Intn(n)))
				}
				return environment.NewArray(out), nil
			}
			// source array
			if src.Type != environment.ArrayType || src.Arr == nil {
				return environment.NewNil(), fmt.Errorf("choice(): first arg must be array or number")
			}
			srcArr := *src.Arr
			if len(srcArr) == 0 {
				return environment.NewNil(), fmt.Errorf("choice(): empty array")
			}
			if size == 0 {
				idx := rand.Intn(len(srcArr))
				return srcArr[idx], nil
			}
			out := make([]environment.Value, size)
			for i := 0; i < size; i++ {
				out[i] = srcArr[rand.Intn(len(srcArr))]
			}
			return environment.NewArray(out), nil
		}),

		fn("shuffle", func(args []environment.Value) (environment.Value, error) {
			// delegate to arrays.shuffle (in-place)
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("shuffle() expects 1 argument")
			}
			return registry["arrays"].Entries["shuffle"].Builtin(args)
		}),

		fn("permutation", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("permutation() expects 1 argument")
			}
			if args[0].Type == environment.NumberType {
				n := int(args[0].Num)
				idx := make([]int, n)
				for i := 0; i < n; i++ {
					idx[i] = i
				}
				rand.Shuffle(n, func(i, j int) { idx[i], idx[j] = idx[j], idx[i] })
				out := make([]environment.Value, n)
				for i := 0; i < n; i++ {
					out[i] = environment.NewNumber(float64(idx[i]))
				}
				return environment.NewArray(out), nil
			}
			// array -> return permuted copy
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("permutation(): expects number or array")
			}
			arr := make([]environment.Value, len(*args[0].Arr))
			copy(arr, *args[0].Arr)
			rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
			return environment.NewArray(arr), nil
		}),

		fn("normal", func(args []environment.Value) (environment.Value, error) {
			// normal([loc, scale, size?])
			loc := 0.0
			scale := 1.0
			if len(args) >= 1 {
				if args[0].Type == environment.NumberType {
					loc = args[0].Num
				} else {
					return environment.NewNil(), fmt.Errorf("normal(): loc must be numeric")
				}
			}
			if len(args) >= 2 {
				if args[1].Type == environment.NumberType {
					scale = args[1].Num
				} else {
					return environment.NewNil(), fmt.Errorf("normal(): scale must be numeric")
				}
			}
			if len(args) == 3 {
				sz, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("normal(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, func() float64 { return loc + scale*rand.NormFloat64() })
			}
			if len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("normal(): too many arguments")
			}
			return environment.NewNumber(loc + scale*rand.NormFloat64()), nil
		}),

		fn("uniform", func(args []environment.Value) (environment.Value, error) {
			low := 0.0
			high := 1.0
			if len(args) >= 1 {
				if args[0].Type == environment.NumberType {
					low = args[0].Num
				} else {
					return environment.NewNil(), fmt.Errorf("uniform(): low must be numeric")
				}
			}
			if len(args) >= 2 {
				if args[1].Type == environment.NumberType {
					high = args[1].Num
				} else {
					return environment.NewNil(), fmt.Errorf("uniform(): high must be numeric")
				}
			}
			if len(args) == 3 {
				sz, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("uniform(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, func() float64 { return low + rand.Float64()*(high-low) })
			}
			if len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("uniform(): too many arguments")
			}
			return environment.NewNumber(low + rand.Float64()*(high-low)), nil
		}),

		fn("binomial", func(args []environment.Value) (environment.Value, error) {
			// binomial(n, p [, size]) — simple summation of Bernoullis (works for modest n)
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("binomial() expects 2-3 arguments")
			}
			nf, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("binomial(): n must be numeric")
			}
			n := int(nf)
			p, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("binomial(): p must be numeric")
			}
			if len(args) == 3 {
				sz, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("binomial(): size must be numeric")
				}
				out := make([]environment.Value, int(sz))
				for i := range out {
					count := 0
					for k := 0; k < n; k++ {
						if rand.Float64() < p {
							count++
						}
					}
					out[i] = environment.NewNumber(float64(count))
				}
				return environment.NewArray(out), nil
			}
			count := 0
			for k := 0; k < n; k++ {
				if rand.Float64() < p {
					count++
				}
			}
			return environment.NewNumber(float64(count)), nil
		}),

		fn("poisson", func(args []environment.Value) (environment.Value, error) {
			// poisson(lambda [, size]) — Knuth's algorithm
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("poisson() expects 1-2 arguments")
			}
			lam, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("poisson(): lambda must be numeric")
			}
			gen := func() float64 {
				L := math.Exp(-lam)
				k := 0
				p := 1.0
				for p > L {
					k++
					p *= rand.Float64()
				}
				return float64(k - 1)
			}
			if len(args) == 2 {
				sz, err := args[1].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("poisson(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, gen)
			}
			return environment.NewNumber(gen()), nil
		}),

		fn("exponential", func(args []environment.Value) (environment.Value, error) {
			// exponential(scale=1 [, size])
			scale := 1.0
			if len(args) >= 1 {
				if args[0].Type == environment.NumberType {
					scale = args[0].Num
				} else {
					return environment.NewNil(), fmt.Errorf("exponential(): scale must be numeric")
				}
			}
			gen := func() float64 { return rand.ExpFloat64() * scale }
			if len(args) == 2 {
				sz, err := args[1].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("exponential(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, gen)
			}
			return environment.NewNumber(gen()), nil
		}),

		fn("gamma", func(args []environment.Value) (environment.Value, error) {
			// limited: support integer shape k by summing exponentials; gamma(k, scale)
			if len(args) < 1 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("gamma() expects 1-3 arguments")
			}
			kf, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("gamma(): shape must be numeric")
			}
			k := int(kf)
			scale := 1.0
			if len(args) >= 2 {
				if args[1].Type == environment.NumberType {
					scale = args[1].Num
				} else {
					return environment.NewNil(), fmt.Errorf("gamma(): scale must be numeric")
				}
			}
			genInt := func() float64 {
				s := 0.0
				for i := 0; i < k; i++ {
					s += rand.ExpFloat64() * scale
				}
				return s
			}
			if len(args) == 3 {
				sz, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("gamma(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, genInt)
			}
			return environment.NewNumber(genInt()), nil
		}),

		fn("beta", func(args []environment.Value) (environment.Value, error) {
			// limited: support integer a,b by sampling Gamma(a,1)/[Gamma(a,1)+Gamma(b,1)]
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("beta() expects 2-3 arguments")
			}
			af, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("beta(): a must be numeric")
			}
			bf, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("beta(): b must be numeric")
			}
			a := int(af)
			b := int(bf)
			gen := func() float64 {
				x := 0.0
				for i := 0; i < a; i++ {
					x += rand.ExpFloat64()
				}
				y := 0.0
				for i := 0; i < b; i++ {
					y += rand.ExpFloat64()
				}
				if x+y == 0 {
					return 0.0
				}
				return x / (x + y)
			}
			if len(args) == 3 {
				sz, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("beta(): size must be numeric")
				}
				return buildRandomArray([]int{int(sz)}, gen)
			}
			return environment.NewNumber(gen()), nil
		}),

		// ---------------- Trigonometry ----------------
		fn("sin", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sin() expects 1 argument")
			}
			return applyUnaryNumeric("sin", args[0], math.Sin)
		}),

		fn("cos", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cos() expects 1 argument")
			}
			return applyUnaryNumeric("cos", args[0], math.Cos)
		}),

		fn("tan", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("tan() expects 1 argument")
			}
			return applyUnaryNumeric("tan", args[0], math.Tan)
		}),

		fn("arcsin", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("arcsin() expects 1 argument")
			}
			return applyUnaryNumeric("arcsin", args[0], math.Asin)
		}),

		fn("arccos", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("arccos() expects 1 argument")
			}
			return applyUnaryNumeric("arccos", args[0], math.Acos)
		}),

		fn("arctan", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("arctan() expects 1 argument")
			}
			return applyUnaryNumeric("arctan", args[0], math.Atan)
		}),

		fn("arctan2", func(args []environment.Value) (environment.Value, error) {
			return applyBinaryNumeric("arctan2", args, math.Atan2)
		}),

		fn("hypot", func(args []environment.Value) (environment.Value, error) {
			// supports 2+ args via pairwise reduction
			return applyBinaryNumeric("hypot", args, math.Hypot)
		}),

		fn("deg2rad", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("deg2rad() expects 1 argument")
			}
			return applyUnaryNumeric("deg2rad", args[0], func(x float64) float64 { return x * math.Pi / 180.0 })
		}),

		fn("rad2deg", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("rad2deg() expects 1 argument")
			}
			return applyUnaryNumeric("rad2deg", args[0], func(x float64) float64 { return x * 180.0 / math.Pi })
		}),

		// ---------------- Exponentials & Logarithms ----------------
		fn("exp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("exp() expects 1 argument")
			}
			return applyUnaryNumeric("exp", args[0], math.Exp)
		}),

		fn("exp2", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("exp2() expects 1 argument")
			}
			return applyUnaryNumeric("exp2", args[0], math.Exp2)
		}),

		fn("expm1", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("expm1() expects 1 argument")
			}
			return applyUnaryNumeric("expm1", args[0], math.Expm1)
		}),

		fn("log", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log() expects 1 argument")
			}
			return applyUnaryNumeric("log", args[0], math.Log)
		}),

		fn("log2", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log2() expects 1 argument")
			}
			return applyUnaryNumeric("log2", args[0], math.Log2)
		}),

		fn("log10", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log10() expects 1 argument")
			}
			return applyUnaryNumeric("log10", args[0], math.Log10)
		}),

		fn("log1p", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log1p() expects 1 argument")
			}
			return applyUnaryNumeric("log1p", args[0], math.Log1p)
		}),

		// ---------------- Statistics ----------------
		fn("sum", func(args []environment.Value) (environment.Value, error) {
			if len(args) == 0 {
				return environment.NewNil(), fmt.Errorf("sum() expects at least 1 argument")
			}
			// single array or number
			if len(args) == 1 {
				if args[0].Type == environment.NumberType {
					return args[0], nil
				}
				nums, err := flattenNumbers(args[0])
				if err != nil {
					return environment.NewNil(), err
				}
				s := 0.0
				for _, v := range nums {
					s += v
				}
				return environment.NewNumber(s), nil
			}
			// multiple numeric args
			s := 0.0
			for _, a := range args {
				if a.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("sum(): all arguments must be numbers or a single array")
				}
				s += a.Num
			}
			return environment.NewNumber(s), nil
		}),

		fn("mean", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("mean() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("mean(): empty array")
			}
			s := 0.0
			for _, v := range nums {
				s += v
			}
			return environment.NewNumber(s / float64(len(nums))), nil
		}),

		// alias
		fn("average", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["mean"].Builtin(args)
		}),

		fn("var", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("var() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("var(): empty array")
			}
			m := 0.0
			for _, v := range nums {
				m += v
			}
			m /= float64(len(nums))
			v := 0.0
			for _, x := range nums {
				d := x - m
				v += d * d
			}
			return environment.NewNumber(v / float64(len(nums))), nil
		}),

		fn("std", func(args []environment.Value) (environment.Value, error) {
			res, err := registry["mathx"].Entries["var"].Builtin(args)
			if err != nil {
				return environment.NewNil(), err
			}
			varv := res.Num
			return environment.NewNumber(math.Sqrt(varv)), nil
		}),

		fn("median", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("median() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			n := len(nums)
			if n == 0 {
				return environment.NewNil(), fmt.Errorf("median(): empty array")
			}
			sort.Float64s(nums)
			if n%2 == 1 {
				return environment.NewNumber(nums[n/2]), nil
			}
			mid := n / 2
			return environment.NewNumber((nums[mid-1] + nums[mid]) / 2.0), nil
		}),

		fn("percentile", func(args []environment.Value) (environment.Value, error) {
			// percentile(array, p) where p in [0,100]
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("percentile() expects 2 arguments")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("percentile(): empty array")
			}
			p, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("percentile(): p must be a number")
			}
			if p < 0 || p > 100 {
				return environment.NewNil(), fmt.Errorf("percentile(): p must be in [0,100]")
			}
			sort.Float64s(nums)
			if p == 0 {
				return environment.NewNumber(nums[0]), nil
			}
			if p == 100 {
				return environment.NewNumber(nums[len(nums)-1]), nil
			}
			idx := (p / 100.0) * float64(len(nums)-1)
			lo := int(math.Floor(idx))
			hi := int(math.Ceil(idx))
			w := idx - float64(lo)
			if lo == hi {
				return environment.NewNumber(nums[lo]), nil
			}
			return environment.NewNumber(nums[lo]*(1-w) + nums[hi]*w), nil
		}),

		fn("quantile", func(args []environment.Value) (environment.Value, error) {
			// quantile(array, q) where q in [0,1]
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("quantile() expects 2 arguments")
			}
			qv, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("quantile(): q must be a number")
			}
			return registry["mathx"].Entries["percentile"].Builtin([]environment.Value{args[0], environment.NewNumber(qv * 100)})
		}),

		// min/max helpers (array -> scalar)
		fn("min", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("min() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("min(): empty array")
			}
			m := nums[0]
			for _, v := range nums {
				if v < m {
					m = v
				}
			}
			return environment.NewNumber(m), nil
		}),
		fn("max", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("max() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("max(): empty array")
			}
			m := nums[0]
			for _, v := range nums {
				if v > m {
					m = v
				}
			}
			return environment.NewNumber(m), nil
		}),

		fn("ptp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ptp() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("ptp(): empty array")
			}
			minv, maxv := nums[0], nums[0]
			for _, v := range nums {
				if v < minv {
					minv = v
				}
				if v > maxv {
					maxv = v
				}
			}
			return environment.NewNumber(maxv - minv), nil
		}),

		fn("argmin", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("argmin() expects 1 argument")
			}
			vals := flattenValues(args[0])
			if len(vals) == 0 {
				return environment.NewNil(), fmt.Errorf("argmin(): empty array")
			}
			minIdx := -1
			minVal := math.Inf(1)
			for i, v := range vals {
				if v.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("argmin(): array elements must be numbers")
				}
				if v.Num < minVal {
					minVal = v.Num
					minIdx = i
				}
			}
			return environment.NewNumber(float64(minIdx)), nil
		}),

		fn("argmax", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("argmax() expects 1 argument")
			}
			vals := flattenValues(args[0])
			if len(vals) == 0 {
				return environment.NewNil(), fmt.Errorf("argmax(): empty array")
			}
			maxIdx := -1
			maxVal := math.Inf(-1)
			for i, v := range vals {
				if v.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("argmax(): array elements must be numbers")
				}
				if v.Num > maxVal {
					maxVal = v.Num
					maxIdx = i
				}
			}
			return environment.NewNumber(float64(maxIdx)), nil
		}),

		fn("cumsum", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cumsum() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			out := make([]environment.Value, len(nums))
			running := 0.0
			for i, v := range nums {
				running += v
				out[i] = environment.NewNumber(running)
			}
			return environment.NewArray(out), nil
		}),

		fn("cumprod", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cumprod() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			out := make([]environment.Value, len(nums))
			running := 1.0
			for i, v := range nums {
				running *= v
				out[i] = environment.NewNumber(running)
			}
			return environment.NewArray(out), nil
		}),

		fn("histogram", func(args []environment.Value) (environment.Value, error) {
			// histogram(array, bins=10)
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("histogram() expects 1 or 2 arguments")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(nums) == 0 {
				return environment.NewNil(), fmt.Errorf("histogram(): empty array")
			}
			nBins := 10
			if len(args) == 2 {
				if args[1].Type == environment.NumberType {
					nBins = int(args[1].Num)
					if nBins <= 0 {
						return environment.NewNil(), fmt.Errorf("histogram(): bins must be > 0")
					}
				} else {
					return environment.NewNil(), fmt.Errorf("histogram(): bins must be a number (count)")
				}
			}
			minv, maxv := nums[0], nums[0]
			for _, v := range nums {
				if v < minv {
					minv = v
				}
				if v > maxv {
					maxv = v
				}
			}
			counts := make([]environment.Value, nBins)
			if minv == maxv {
				counts[0] = environment.NewNumber(float64(len(nums)))
				for i := 1; i < nBins; i++ {
					counts[i] = environment.NewNumber(0)
				}
				return environment.NewArray(counts), nil
			}
			step := (maxv - minv) / float64(nBins)
			for _, v := range nums {
				idx := int((v - minv) / step)
				if idx < 0 {
					idx = 0
				}
				if idx >= nBins {
					idx = nBins - 1
				}
				if counts[idx].Type != environment.NumberType {
					counts[idx] = environment.NewNumber(0)
				}
				counts[idx] = environment.NewNumber(counts[idx].Num + 1)
			}
			return environment.NewArray(counts), nil
		}),

		fn("bincount", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bincount() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			maxIdx := -1
			for _, v := range nums {
				if v < 0 || math.Trunc(v) != v {
					return environment.NewNil(), fmt.Errorf("bincount(): expects non-negative integer values")
				}
				if int(v) > maxIdx {
					maxIdx = int(v)
				}
			}
			if maxIdx < 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			counts := make([]environment.Value, maxIdx+1)
			for _, v := range nums {
				idx := int(v)
				counts[idx] = environment.NewNumber(counts[idx].Num + 1)
			}
			return environment.NewArray(counts), nil
		}),

		fn("cov", func(args []environment.Value) (environment.Value, error) {
			// cov(x, y) -> scalar covariance; cov(matrix) -> covariance matrix between rows
			if len(args) == 1 {
				mat := args[0]
				if mat.Type != environment.ArrayType || mat.Arr == nil {
					return environment.NewNil(), fmt.Errorf("cov(): expects an array or two arrays")
				}
				// expect 2D array: rows = variables
				if len(*mat.Arr) == 0 || (*mat.Arr)[0].Type != environment.ArrayType {
					return environment.NewNil(), fmt.Errorf("cov(): expects 2D array")
				}
				nVars := len(*mat.Arr)
				// compute means
				means := make([]float64, nVars)
				for i := 0; i < nVars; i++ {
					row := (*mat.Arr)[i]
					nums, err := flattenNumbers(row)
					if err != nil {
						return environment.NewNil(), err
					}
					if len(nums) == 0 {
						return environment.NewNil(), fmt.Errorf("cov(): empty row")
					}
					s := 0.0
					for _, v := range nums {
						s += v
					}
					means[i] = s / float64(len(nums))
				}
				// build covariance matrix
				out := make([]environment.Value, nVars)
				for i := 0; i < nVars; i++ {
					rowVals := make([]environment.Value, nVars)
					numsI, _ := flattenNumbers((*mat.Arr)[i])
					for j := 0; j < nVars; j++ {
						numsJ, _ := flattenNumbers((*mat.Arr)[j])
						if len(numsI) != len(numsJ) {
							return environment.NewNil(), fmt.Errorf("cov(): rows must have same length")
						}
						n := len(numsI)
						acc := 0.0
						for k := 0; k < n; k++ {
							acc += (numsI[k] - means[i]) * (numsJ[k] - means[j])
						}
						rowVals[j] = environment.NewNumber(acc / float64(n))
					}
					out[i] = environment.NewArray(rowVals)
				}
				return environment.NewArray(out), nil
			}
			// two arrays
			if len(args) == 2 {
				xs, err := flattenNumbers(args[0])
				if err != nil {
					return environment.NewNil(), err
				}
				ys, err := flattenNumbers(args[1])
				if err != nil {
					return environment.NewNil(), err
				}
				if len(xs) != len(ys) {
					return environment.NewNil(), fmt.Errorf("cov(): arrays must have same length")
				}
				n := len(xs)
				mx, my := 0.0, 0.0
				for i := 0; i < n; i++ {
					mx += xs[i]
					my += ys[i]
				}
				mx /= float64(n)
				my /= float64(n)
				acc := 0.0
				for i := 0; i < n; i++ {
					acc += (xs[i] - mx) * (ys[i] - my)
				}
				return environment.NewNumber(acc / float64(n)), nil
			}
			return environment.NewNil(), fmt.Errorf("cov(): invalid arguments")
		}),

		fn("corrcoef", func(args []environment.Value) (environment.Value, error) {
			// corrcoef(x,y) -> scalar corr; corrcoef(matrix) -> correlation matrix between rows
			if len(args) == 2 {
				covv, err := registry["mathx"].Entries["cov"].Builtin([]environment.Value{args[0], args[1]})
				if err != nil {
					return environment.NewNil(), err
				}
				cov := covv.Num
				stdx, err := registry["mathx"].Entries["std"].Builtin([]environment.Value{args[0]})
				if err != nil {
					return environment.NewNil(), err
				}
				stdy, err := registry["mathx"].Entries["std"].Builtin([]environment.Value{args[1]})
				if err != nil {
					return environment.NewNil(), err
				}
				if stdx.Num == 0 || stdy.Num == 0 {
					return environment.NewNumber(0), nil
				}
				return environment.NewNumber(cov / (stdx.Num * stdy.Num)), nil
			}
			if len(args) == 1 {
				mat := args[0]
				covM, err := registry["mathx"].Entries["cov"].Builtin([]environment.Value{mat})
				if err != nil {
					return environment.NewNil(), err
				}
				// convert cov matrix to corr
				if covM.Type != environment.ArrayType || covM.Arr == nil {
					return environment.NewNil(), fmt.Errorf("corrcoef(): internal error")
				}
				n := len(*covM.Arr)
				out := make([]environment.Value, n)
				stds := make([]float64, n)
				for i := 0; i < n; i++ {
					stds[i] = math.Sqrt((*(*covM.Arr)[i].Arr)[i].Num)
				}
				for i := 0; i < n; i++ {
					row := make([]environment.Value, n)
					for j := 0; j < n; j++ {
						covij := (*(*covM.Arr)[i].Arr)[j].Num
						den := stds[i] * stds[j]
						if den == 0 {
							row[j] = environment.NewNumber(0)
						} else {
							row[j] = environment.NewNumber(covij / den)
						}
					}
					out[i] = environment.NewArray(row)
				}
				return environment.NewArray(out), nil
			}
			return environment.NewNil(), fmt.Errorf("corrcoef(): invalid arguments")
		}),

		// ---------------- Linear algebra (linalg) ----------------
		fn("dot", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("dot() expects 2 arguments")
			}
			a, b := args[0], args[1]
			sa := shapeOf(a)
			sb := shapeOf(b)
			// scalar * scalar
			if len(sa) == 0 && len(sb) == 0 {
				return environment.NewNumber(a.Num * b.Num), nil
			}
			// both 1D
			if len(sa) <= 1 && len(sb) <= 1 {
				a1, err := to1DFloatSlice(a)
				if err != nil {
					return environment.NewNil(), err
				}
				b1, err := to1DFloatSlice(b)
				if err != nil {
					return environment.NewNil(), err
				}
				if len(a1) != len(b1) {
					return environment.NewNil(), fmt.Errorf("dot(): length mismatch")
				}
				s := 0.0
				for i := range a1 {
					s += a1[i] * b1[i]
				}
				return environment.NewNumber(s), nil
			}
			// matrix-vector / vector-matrix / matrix-matrix
			if len(sa) == 2 && len(sb) == 1 {
				A, err := to2DFloatSlice(a)
				if err != nil {
					return environment.NewNil(), err
				}
				v, err := to1DFloatSlice(b)
				if err != nil {
					return environment.NewNil(), err
				}
				res, err := matVecMulFloat(A, v)
				if err != nil {
					return environment.NewNil(), err
				}
				return from1DSlice(res), nil
			}
			if len(sa) == 1 && len(sb) == 2 {
				v, err := to1DFloatSlice(a)
				if err != nil {
					return environment.NewNil(), err
				}
				B, err := to2DFloatSlice(b)
				if err != nil {
					return environment.NewNil(), err
				}
				res, err := vecMatMulFloat(v, B)
				if err != nil {
					return environment.NewNil(), err
				}
				return from1DSlice(res), nil
			}
			if len(sa) == 2 && len(sb) == 2 {
				A, err := to2DFloatSlice(a)
				if err != nil {
					return environment.NewNil(), err
				}
				B, err := to2DFloatSlice(b)
				if err != nil {
					return environment.NewNil(), err
				}
				C, err := matMulFloat(A, B)
				if err != nil {
					return environment.NewNil(), err
				}
				return from2DFloatSlice(C), nil
			}
			return environment.NewNil(), fmt.Errorf("dot(): unsupported shapes")
		}),

		fn("matmul", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["dot"].Builtin(args)
		}),

		fn("inner", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("inner() expects 2 arguments")
			}
			a1, err := to1DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b1, err := to1DFloatSlice(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(a1) != len(b1) {
				return environment.NewNil(), fmt.Errorf("inner(): length mismatch")
			}
			s := 0.0
			for i := range a1 {
				s += a1[i] * b1[i]
			}
			return environment.NewNumber(s), nil
		}),

		fn("outer", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("outer() expects 2 arguments")
			}
			a1, err := to1DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b1, err := to1DFloatSlice(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			mat := make([][]float64, len(a1))
			for i := range a1 {
				row := make([]float64, len(b1))
				for j := range b1 {
					row[j] = a1[i] * b1[j]
				}
				mat[i] = row
			}
			return from2DFloatSlice(mat), nil
		}),

		fn("vdot", func(args []environment.Value) (environment.Value, error) {
			return registry["mathx"].Entries["inner"].Builtin(args)
		}),

		fn("tensordot", func(args []environment.Value) (environment.Value, error) {
			// tensordot(a,b) -> contract last axis of a with first of b (limited support)
			if len(args) < 2 || len(args) > 3 {
				return environment.NewNil(), fmt.Errorf("tensordot() expects 2 or 3 arguments")
			}
			axes := 1
			if len(args) == 3 {
				av, err := args[2].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("tensordot(): axes must be numeric")
				}
				axes = int(av)
			}
			if axes != 1 {
				return environment.NewNil(), fmt.Errorf("tensordot(): only axes=1 supported")
			}
			// support 2D x 2D -> matmul
			return registry["mathx"].Entries["matmul"].Builtin([]environment.Value{args[0], args[1]})
		}),

		fn("trace", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("trace() expects 1 argument")
			}
			if shape := shapeOf(args[0]); len(shape) == 2 {
				A, err := to2DFloatSlice(args[0])
				if err != nil {
					return environment.NewNil(), err
				}
				s := 0.0
				n := int(math.Min(float64(len(A)), float64(len(A[0]))))
				for i := 0; i < n; i++ {
					s += A[i][i]
				}
				return environment.NewNumber(s), nil
			}
			// fallback: sum of flattened
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			s := 0.0
			for _, v := range nums {
				s += v
			}
			return environment.NewNumber(s), nil
		}),

		fn("norm", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("norm() expects 1 argument")
			}
			nums, err := flattenNumbers(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			s := 0.0
			for _, v := range nums {
				s += v * v
			}
			return environment.NewNumber(math.Sqrt(s)), nil
		}),

		fn("det", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("det() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			d, err := detFloat(A)
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewNumber(d), nil
		}),

		fn("inv", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("inv() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			X, err := invFloat(A)
			if err != nil {
				return environment.NewNil(), err
			}
			return from2DFloatSlice(X), nil
		}),

		fn("pinv", func(args []environment.Value) (environment.Value, error) {
			// limited: fallback to inverse for square invertible
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("pinv() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			X, err := invFloat(A)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("pinv(): only square invertible matrices supported")
			}
			return from2DFloatSlice(X), nil
		}),

		fn("solve", func(args []environment.Value) (environment.Value, error) {
			// solve(A, b)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("solve() expects 2 arguments")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			b, err := to1DFloatSlice(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			x, err := solveLinear(A, b)
			if err != nil {
				return environment.NewNil(), err
			}
			return from1DSlice(x), nil
		}),

		fn("eig", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("eig() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			vals, vecs, err := eig2x2(A)
			if err != nil {
				return environment.NewNil(), err
			}
			// return [eigenvalues_array, eigenvectors_matrix]
			valsArr := from1DSlice(vals)
			vecRows := make([]environment.Value, len(vecs))
			for i := range vecs {
				vecRows[i] = from1DSlice(vecs[i])
			}
			vecMat := environment.NewArray(vecRows)
			return environment.NewArray([]environment.Value{valsArr, vecMat}), nil
		}),

		fn("eigvals", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("eigvals() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			vals, _, err := eig2x2(A)
			if err != nil {
				return environment.NewNil(), err
			}
			return from1DSlice(vals), nil
		}),

		fn("svd", func(args []environment.Value) (environment.Value, error) {
			// limited: support 1x1 and 2x2 (returns [U, S, V]) where U/V are identities
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("svd() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			if len(A) == 1 && len(A[0]) == 1 {
				s := math.Abs(A[0][0])
				U := from2DFloatSlice([][]float64{{1}})
				V := from2DFloatSlice([][]float64{{1}})
				S := from1DSlice([]float64{s})
				return environment.NewArray([]environment.Value{U, S, V}), nil
			}
			if len(A) == 2 && len(A[0]) == 2 {
				// compute singular values = sqrt(eigvals(A^T A))
				AtA := make([][]float64, 2)
				AtA[0] = []float64{A[0][0]*A[0][0] + A[1][0]*A[1][0], A[0][0]*A[0][1] + A[1][0]*A[1][1]}
				AtA[1] = []float64{AtA[0][1], A[0][1]*A[0][1] + A[1][1]*A[1][1]}
				vals, _, err := eig2x2(AtA)
				if err != nil {
					return environment.NewNil(), err
				}
				s1 := math.Sqrt(math.Max(0, vals[0]))
				s2 := math.Sqrt(math.Max(0, vals[1]))
				// sort desc
				if s1 < s2 {
					s1, s2 = s2, s1
				}
				U := from2DFloatSlice([][]float64{{1, 0}, {0, 1}})
				V := from2DFloatSlice([][]float64{{1, 0}, {0, 1}})
				S := from1DSlice([]float64{s1, s2})
				return environment.NewArray([]environment.Value{U, S, V}), nil
			}
			return environment.NewNil(), fmt.Errorf("svd(): only 1x1 or 2x2 supported")
		}),

		fn("qr", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("qr() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			Q, R, err := qrGramSchmidt(A)
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewArray([]environment.Value{from2DFloatSlice(Q), from2DFloatSlice(R)}), nil
		}),

		fn("cholesky", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cholesky() expects 1 argument")
			}
			A, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			L, err := choleskyDecomp(A)
			if err != nil {
				return environment.NewNil(), err
			}
			return from2DFloatSlice(L), nil
		}),

		fn("matrix_rank", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("matrix_rank() expects 1 argument")
			}
			mat, err := to2DFloatSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			r, err := matrixRankFloat(mat)
			if err != nil {
				return environment.NewNil(), err
			}
			return environment.NewNumber(float64(r)), nil
		}),

		// --- FFT / spectral builtins ---
		fn("fft", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fft() expects 1 argument")
			}
			cs, err := envArrayToComplexSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			res := fftComplex(cs)
			return complexSliceToEnvArray(res), nil
		}),

		fn("ifft", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ifft() expects 1 argument")
			}
			cs, err := envArrayToComplexSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			res := ifftComplex(cs)
			return complexSliceToEnvArray(res), nil
		}),

		fn("rfft", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("rfft() expects 1 argument")
			}
			cs, err := envArrayToComplexSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			res := fftComplex(cs)
			m := len(res)/2 + 1
			return complexSliceToEnvArray(res[:m]), nil
		}),

		fn("irfft", func(args []environment.Value) (environment.Value, error) {
			// irfft(rspec [, n]) -> real sequence
			if len(args) < 1 || len(args) > 2 {
				return environment.NewNil(), fmt.Errorf("irfft() expects 1 or 2 arguments")
			}
			// parse input (array of complex [re,im])
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("irfft(): expects array of complex pairs")
			}
			m := len(*args[0].Arr)
			if m == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			// determine N
			var N int
			if len(args) == 2 {
				nn, err := args[1].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("irfft(): n must be number")
				}
				N = int(nn)
			} else {
				if m == 1 {
					N = 1
				} else {
					N = 2 * (m - 1)
				}
			}
			// build full spectrum
			spec := make([]complex128, N)
			// fill first m elements from input
			for i := 0; i < m; i++ {
				cell := (*args[0].Arr)[i]
				if cell.Type != environment.ArrayType || cell.Arr == nil || len(*cell.Arr) != 2 {
					return environment.NewNil(), fmt.Errorf("irfft(): spectrum must be array of [re,im] pairs")
				}
				re := (*cell.Arr)[0]
				im := (*cell.Arr)[1]
				if re.Type != environment.NumberType || im.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("irfft(): spectrum entries must be numbers")
				}
				spec[i] = complex(re.Num, im.Num)
			}
			// mirror conjugates for remaining bins (except Nyquist/0)
			for i := 1; i < m-1; i++ {
				spec[N-i] = cmplx.Conj(spec[i])
			}
			// compute inverse
			res := ifftComplex(spec)
			// return real parts
			out := make([]environment.Value, len(res))
			for i := range res {
				out[i] = environment.NewNumber(real(res[i]))
			}
			return environment.NewArray(out), nil
		}),

		fn("fft2", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fft2() expects 1 argument")
			}
			A, err := to2DComplexSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			// FFT rows
			nrows := len(A)
			if nrows == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			ncols := len(A[0])
			B := make([][]complex128, nrows)
			for i := 0; i < nrows; i++ {
				B[i] = fftComplex(A[i])
			}
			// FFT columns
			// transpose B
			T := make([][]complex128, ncols)
			for j := 0; j < ncols; j++ {
				col := make([]complex128, nrows)
				for i := 0; i < nrows; i++ {
					col[i] = B[i][j]
				}
				Tc := fftComplex(col)
				T[j] = Tc
			}
			// transpose back
			out := make([]environment.Value, nrows)
			for i := 0; i < nrows; i++ {
				row := make([]environment.Value, ncols)
				for j := 0; j < ncols; j++ {
					row[j] = environment.NewArray([]environment.Value{environment.NewNumber(real(T[j][i])), environment.NewNumber(imag(T[j][i]))})
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("ifft2", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ifft2() expects 1 argument")
			}
			A, err := to2DComplexSlice(args[0])
			if err != nil {
				return environment.NewNil(), err
			}
			// inverse FFT rows
			nrows := len(A)
			if nrows == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			ncols := len(A[0])
			B := make([][]complex128, nrows)
			for i := 0; i < nrows; i++ {
				B[i] = ifftComplex(A[i])
			}
			// inverse FFT columns
			T := make([][]complex128, ncols)
			for j := 0; j < ncols; j++ {
				col := make([]complex128, nrows)
				for i := 0; i < nrows; i++ {
					col[i] = B[i][j]
				}
				Tc := ifftComplex(col)
				T[j] = Tc
			}
			// transpose back and return
			out := make([]environment.Value, nrows)
			for i := 0; i < nrows; i++ {
				row := make([]environment.Value, ncols)
				for j := 0; j < ncols; j++ {
					row[j] = environment.NewArray([]environment.Value{environment.NewNumber(real(T[j][i])), environment.NewNumber(imag(T[j][i]))})
				}
				out[i] = environment.NewArray(row)
			}
			return environment.NewArray(out), nil
		}),

		fn("fftn", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fftn() expects 1 argument")
			}
			sh := shapeOf(args[0])
			switch len(sh) {
			case 0:
				// scalar -> return scalar complex
				res, err := registry["mathx"].Entries["fft"].Builtin([]environment.Value{args[0]})
				return res, err
			case 1:
				return registry["mathx"].Entries["fft"].Builtin([]environment.Value{args[0]})
			case 2:
				return registry["mathx"].Entries["fft2"].Builtin([]environment.Value{args[0]})
			default:
				return environment.NewNil(), fmt.Errorf("fftn(): only up to 2D supported")
			}
		}),

		fn("ifftn", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ifftn() expects 1 argument")
			}
			sh := shapeOf(args[0])
			switch len(sh) {
			case 0:
				return registry["mathx"].Entries["ifft"].Builtin([]environment.Value{args[0]})
			case 1:
				return registry["mathx"].Entries["ifft"].Builtin([]environment.Value{args[0]})
			case 2:
				return registry["mathx"].Entries["ifft2"].Builtin([]environment.Value{args[0]})
			default:
				return environment.NewNil(), fmt.Errorf("ifftn(): only up to 2D supported")
			}
		}),

		fn("fftshift", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fftshift() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewNil(), fmt.Errorf("fftshift(): expects array")
			}
			// 1D
			first := (*v.Arr)[0]
			if first.Type != environment.ArrayType || first.Arr == nil {
				return environment.NewArray(fftshift1D(*v.Arr)), nil
			}
			// 2D: roll rows and columns
			nrows := len(*v.Arr)
			if nrows == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			ncols := len(*(*v.Arr)[0].Arr)
			rows := roll1D(*v.Arr, nrows/2)
			out := make([]environment.Value, nrows)
			for i := 0; i < nrows; i++ {
				rowVals := *rows[i].Arr
				out[i] = environment.NewArray(roll1D(rowVals, ncols/2))
			}
			return environment.NewArray(out), nil
		}),

		fn("ifftshift", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ifftshift() expects 1 argument")
			}
			v := args[0]
			if v.Type != environment.ArrayType || v.Arr == nil {
				return environment.NewNil(), fmt.Errorf("ifftshift(): expects array")
			}
			first := (*v.Arr)[0]
			if first.Type != environment.ArrayType || first.Arr == nil {
				return environment.NewArray(ifftshift1D(*v.Arr)), nil
			}
			nrows := len(*v.Arr)
			if nrows == 0 {
				return environment.NewArray([]environment.Value{}), nil
			}
			ncols := len(*(*v.Arr)[0].Arr)
			rows := roll1D(*v.Arr, (nrows+1)/2)
			out := make([]environment.Value, nrows)
			for i := 0; i < nrows; i++ {
				rowVals := *rows[i].Arr
				out[i] = environment.NewArray(roll1D(rowVals, (ncols+1)/2))
			}
			return environment.NewArray(out), nil
		}),

		// --- I/O helpers for mathx (simple, pragmatic implementations) ---
		fn("save", func(args []environment.Value) (environment.Value, error) {
			// save(filename, value) -> writes JSON representation
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("save() expects 2 arguments")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("save(): filename must be string")
			}
			v := args[1]
			obj, err := valueToInterface(v)
			if err != nil {
				return environment.NewNil(), err
			}
			b, err := json.MarshalIndent(obj, "", "  ")
			if err != nil {
				return environment.NewNil(), fmt.Errorf("save(): json marshal: %v", err)
			}
			if err := os.WriteFile(path, b, 0644); err != nil {
				return environment.NewNil(), fmt.Errorf("save(): write error: %v", err)
			}
			return environment.NewNil(), nil
		}),

		fn("load", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("load() expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("load(): filename must be string")
			}
			b, err := os.ReadFile(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("load(): read error: %v", err)
			}
			var obj interface{}
			if err := json.Unmarshal(b, &obj); err == nil {
				return interfaceToValue(obj), nil
			}
			// fallback: return file contents as string
			return environment.NewString(string(b)), nil
		}),

		fn("savez", func(args []environment.Value) (environment.Value, error) {
			// savez(filename, object) OR savez(filename, name1, val1, name2, val2...)
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("savez() expects at least 2 arguments")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("savez(): filename must be string")
			}
			var m map[string]interface{} = make(map[string]interface{})
			if args[1].Type == environment.ObjectType {
				for _, k := range args[1].Obj.Keys {
					iv, _ := valueToInterface(args[1].Obj.Entries[k])
					m[k] = iv
				}
			} else {
				if (len(args)-1)%2 != 0 {
					return environment.NewNil(), fmt.Errorf("savez(): expected pairs of (name, value)")
				}
				for i := 1; i < len(args); i += 2 {
					name, err := args[i].AsString()
					if err != nil {
						return environment.NewNil(), fmt.Errorf("savez(): expected string name")
					}
					iv, err := valueToInterface(args[i+1])
					if err != nil {
						return environment.NewNil(), err
					}
					m[name] = iv
				}
			}
			b, err := json.MarshalIndent(m, "", "  ")
			if err != nil {
				return environment.NewNil(), fmt.Errorf("savez(): json marshal: %v", err)
			}
			if err := os.WriteFile(path, b, 0644); err != nil {
				return environment.NewNil(), fmt.Errorf("savez(): write error: %v", err)
			}
			return environment.NewNil(), nil
		}),

		fn("savetxt", func(args []environment.Value) (environment.Value, error) {
			// savetxt(filename, array)
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("savetxt() expects 2 arguments")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("savetxt(): filename must be string")
			}
			arr := args[1]
			if arr.Type != environment.ArrayType || arr.Arr == nil {
				return environment.NewNil(), fmt.Errorf("savetxt(): expects array")
			}
			f, err := os.Create(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("savetxt(): create error: %v", err)
			}
			defer f.Close()
			w := csv.NewWriter(f)
			for _, r := range *arr.Arr {
				// if row is array -> multiple columns, else single column
				if r.Type == environment.ArrayType && r.Arr != nil {
					rec := make([]string, len(*r.Arr))
					for i, c := range *r.Arr {
						rec[i] = c.String()
					}
					w.Write(rec)
				} else {
					w.Write([]string{r.String()})
				}
			}
			w.Flush()
			if err := w.Error(); err != nil {
				return environment.NewNil(), fmt.Errorf("savetxt(): csv write: %v", err)
			}
			return environment.NewNil(), nil
		}),

		fn("loadtxt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("loadtxt() expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("loadtxt(): filename must be string")
			}
			f, err := os.Open(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("loadtxt(): open error: %v", err)
			}
			defer f.Close()
			r := csv.NewReader(bufio.NewReader(f))
			records, err := r.ReadAll()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("loadtxt(): csv read: %v", err)
			}
			// build array
			rows := make([]environment.Value, len(records))
			for i, rec := range records {
				if len(rec) == 1 {
					// single column -> number when possible
					if s := strings.TrimSpace(rec[0]); s == "" {
						rows[i] = environment.NewArray([]environment.Value{})
					} else if v, err := strconv.ParseFloat(s, 64); err == nil {
						rows[i] = environment.NewNumber(v)
					} else {
						rows[i] = environment.NewString(s)
					}
				} else {
					cols := make([]environment.Value, len(rec))
					for j, tok := range rec {
						t := strings.TrimSpace(tok)
						if v, err := strconv.ParseFloat(t, 64); err == nil {
							cols[j] = environment.NewNumber(v)
						} else {
							cols[j] = environment.NewString(t)
						}
					}
					rows[i] = environment.NewArray(cols)
				}
			}
			return environment.NewArray(rows), nil
		}),

		fn("genfromtxt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("genfromtxt() expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("genfromtxt(): filename must be string")
			}
			f, err := os.Open(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("genfromtxt(): open error: %v", err)
			}
			defer f.Close()
			sc := bufio.NewScanner(f)
			var rows []environment.Value
			for sc.Scan() {
				line := strings.TrimSpace(sc.Text())
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				parts := strings.Fields(line)
				if len(parts) == 0 {
					continue
				}
				if len(parts) == 1 {
					p := parts[0]
					if p == "" {
						rows = append(rows, environment.NewNil())
						continue
					}
					if v, err := strconv.ParseFloat(p, 64); err == nil {
						rows = append(rows, environment.NewNumber(v))
						continue
					}
					rows = append(rows, environment.NewString(p))
					continue
				}
				cols := make([]environment.Value, len(parts))
				for i, p := range parts {
					if p == "" {
						cols[i] = environment.NewNil()
						continue
					}
					if v, err := strconv.ParseFloat(p, 64); err == nil {
						cols[i] = environment.NewNumber(v)
					} else {
						cols[i] = environment.NewString(p)
					}
				}
				rows = append(rows, environment.NewArray(cols))
			}
			if err := sc.Err(); err != nil {
				return environment.NewNil(), fmt.Errorf("genfromtxt(): scan error: %v", err)
			}
			return environment.NewArray(rows), nil
		}),

		fn("tofile", func(args []environment.Value) (environment.Value, error) {
			// tofile(filename, array) -- binary float64 little-endian
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("tofile() expects 2 arguments")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("tofile(): filename must be string")
			}
			nums, err := flattenNumbers(args[1])
			if err != nil {
				return environment.NewNil(), err
			}
			f, err := os.Create(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("tofile(): create error: %v", err)
			}
			defer f.Close()
			for _, v := range nums {
				if err := binary.Write(f, binary.LittleEndian, v); err != nil {
					return environment.NewNil(), fmt.Errorf("tofile(): write error: %v", err)
				}
			}
			return environment.NewNil(), nil
		}),

		fn("fromfile", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fromfile() expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("fromfile(): filename must be string")
			}
			b, err := os.ReadFile(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("fromfile(): read error: %v", err)
			}
			if len(b)%8 != 0 {
				return environment.NewNil(), fmt.Errorf("fromfile(): file size not multiple of 8 bytes")
			}
			cnt := len(b) / 8
			out := make([]environment.Value, cnt)
			buf := bytes.NewReader(b)
			for i := 0; i < cnt; i++ {
				var val float64
				if err := binary.Read(buf, binary.LittleEndian, &val); err != nil {
					return environment.NewNil(), fmt.Errorf("fromfile(): read error: %v", err)
				}
				out[i] = environment.NewNumber(val)
			}
			return environment.NewArray(out), nil
		}),
	))
}

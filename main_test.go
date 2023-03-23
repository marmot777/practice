package main

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSub(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Sub(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相减乘", t, func() {
		So(Multiply(1, 2), ShouldEqual, 2)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {
		Convey("被除数为0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("被除数不为0", func() {
			num, err := Division(10, 2)
			So(num, ShouldEqual, 5)
			So(err, ShouldBeNil)
		})
	})
}

func TestSplit(t *testing.T) {

	got := Split("a:b:c", ":") // 程序输出的结果
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}


func TestSplitWithComplexSep(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}



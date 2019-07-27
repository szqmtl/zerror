package zerror

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var errMsg = "new error"

func TestNew(t *testing.T) {

	s := GetDefaultSeverity()
	z := New(errMsg)

	msg := z.String()

	if !strings.Contains(msg, errMsg) {
		t.Errorf("default severity error: msg(%s), s(%s)", msg, s.String())
	}
	if !strings.Contains(msg, s.String()) {
		t.Errorf("default severity error: msg(%s), s(%s)", msg, s.String())
	}

	s = SeverityFatal
	z = NewFatal(errMsg)
	msg = z.String()
	if !strings.Contains(msg, s.String()) {
		t.Errorf("severity fatal error: msg(%s), s(%s)", msg, s.String())
	}

	s = SeverityWarn
	z = NewWarn(errMsg)
	msg = z.String()
	if !strings.Contains(msg, s.String()) {
		t.Errorf("severity warn error: msg(%s), s(%s)", msg, s.String())
	}

	s = SeverityInfo
	z = NewInfo(errMsg)
	msg = z.String()
	if !strings.Contains(msg, s.String()) {
		t.Errorf("severity info error: msg(%s), s(%s)", msg, s.String())
	}
}

func TestSetDefaultSeverity(t *testing.T) {
	o := GetDefaultSeverity()
	defer SetDefaultSeverity(o)

	n := SeverityWarn
	SetDefaultSeverity(n)

	if GetDefaultSeverity() != n {
		t.Errorf("changing severity error: get(%s), n(%s)", GetDefaultSeverity().String(), n.String())
	}
	z := New(errMsg)
	msg := z.String()
	zs := z.GetSeverity()

	if zs != n || !strings.Contains(msg, n.String()) {
		t.Errorf("changing severity error: msg(%s), n(%s)", msg, n.String())
	}
}

func TestSetTimeFormat(t *testing.T) {
	o := GetTimeFormat()
	defer SetTimeFormat(o)

	n := time.RFC822
	SetTimeFormat(n)

	if !strings.EqualFold(GetTimeFormat(), n) {
		t.Errorf("changing time format error: get(%s), n(%s)", GetTimeFormat(), n)
	}

	z := New(errMsg)
	msg := z.String()
	zc := z.GetCreated()

	if !strings.Contains(msg, zc.Format(n)) {
		t.Errorf("changing time format error: get(%s), n(%s)", msg, zc.Format(n))
	}
}

func TestSetMessageFormat(t *testing.T) {
	o := GetMessageFormat()
	defer SetMessageFormat(o)

	n := fmt.Sprintf("%s: %s", NotationTime, NotationMessage)
	SetMessageFormat(n)

	z := New(errMsg)
	msg := z.String()
	zs := z.GetSeverity()

	if strings.Contains(msg, zs.String()) {
		t.Errorf("changing string format error: get(%s), n(%s)", msg, zs.String())
	}
}

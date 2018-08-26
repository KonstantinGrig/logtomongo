package parser

import "testing"

func TestParserFirstLogTimeParse(t *testing.T) {
	parser := ParserFirst{}
	logTime, err := parser.logTimeParse("Feb 1, 2018 at 3:04:05pm (UTC) | This is log message1")
	if err == nil {
		y := logTime.Year()
		if y != 2018 {
			t.Error("Expected: 2018, got ", y)
		}

		m := logTime.Month()
		if m != 2 {
			t.Error("Expected: 2, got ", m)
		}

		d := logTime.Day()
		if d != 1 {
			t.Error("Expected: 1, got ", d)
		}

		h := logTime.Hour()
		if h != 15 {
			t.Error("Expected: 15, got ", h)
		}

		min := logTime.Minute()
		if min != 4 {
			t.Error("Expected: 4, got ", min)
		}

		sec := logTime.Second()
		if sec != 5 {
			t.Error("Expected: 5, got ", sec)
		}
	} else {
		t.Error("Error ParserFirst: logTimeParse")
	}
}

func TestParserFirstLogMsgParse(t *testing.T) {
	parser := ParserFirst{}
	logMsg, err := parser.logMsgParse("Feb 1, 2018 at 3:04:05pm (UTC) | This is log message1")
	t.Log(logMsg)
	if err == nil {
		if logMsg != "This is log message1" {
			t.Error("Expected: \"This is log message1\", got ", logMsg)
		}
	} else {
		t.Error(err)
	}
}

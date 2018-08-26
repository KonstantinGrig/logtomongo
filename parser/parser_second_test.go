package parser

import "testing"

func TestParserSecondLogTimeParse(t *testing.T) {
	parser := ParserSecond{}
	logTime, err := parser.logTimeParse("2018-02-01T15:04:05Z | This is log message")
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

func TestParserSecondLogMsgParse(t *testing.T) {
	parser := ParserSecond{}
	logMsg, err := parser.logMsgParse("2018-02-01T15:04:05Z | This is log message1")
	t.Log(logMsg)
	if err == nil {
		if logMsg != "This is log message1" {
			t.Error("Expected: \"This is log message1\", got ", logMsg)
		}
	} else {
		t.Error(err)
	}
}

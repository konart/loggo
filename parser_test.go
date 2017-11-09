package main

import "testing"

func TestGetLogRecord(t *testing.T) {
	s := "2017-10-30 17:28:13,781 [DEBUG] -- Канал ETRAN_NSI_DOC. Записей прочитано: 1, добавлено в очередь: 1"
	record := getLogRecord(s)
	if level := "[DEBUG]"; record.level != level {
		t.Errorf("Expected level '[DEBUG]', got '%s' instead", level, record.level)
	}
	if message := "2017-10-30 17:28:13,781 Канал ETRAN_NSI_DOC. Записей прочитано: 1, добавлено в очередь: 1"; record.message != message {
		t.Errorf("Expected message '%s', got '%s' instead", message, record.message)
	}
}

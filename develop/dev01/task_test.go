package main

import (
	"testing"
	"time"
)

type fakeWriter struct {
	content string
}

func (f *fakeWriter) Write(p []byte) (n int, err error) {
	f.content += string(p)
	return 0, nil
}

func Test_main(t *testing.T) {
	var testTable = []struct {
		legalTimeHost     bool
		timeHost          string
		expectedCurTime   time.Time
		expectedExactTime time.Time
	}{
		{
			legalTimeHost: true,
			timeHost:      "0.beevik-ntp.pool.ntp.org",
		},
		{
			legalTimeHost: false,
			timeHost:      "",
		},
	}

	for _, test := range testTable {
		errFakeWriter := fakeWriter{}
		Stderr = &errFakeWriter

		outFakeWriter := fakeWriter{}
		Stdout = &outFakeWriter

		TimeHost = test.timeHost

		ExitFunc = func(n int) {}

		main()

		if test.legalTimeHost && errFakeWriter.content != "" {
			t.Error("Cannot get time from ntp server")
		} else if !test.legalTimeHost && errFakeWriter.content == "" {
			t.Error("Error code not returned to OS")
		}
	}
}

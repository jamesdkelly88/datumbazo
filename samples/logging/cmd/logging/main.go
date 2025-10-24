package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"

	"example.com/logging/internal/sub"
)

// custom levels
const (
	LevelTrace = slog.Level(-8)
	LevelFatal = slog.Level(12)
)

func main() {
	// built in log library
	log.Println("This is log.println")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// using a buffer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	// slog
	log.SetFlags(log.LstdFlags)
	slog.Info("info level")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Info("context", "key", "value", "another key", "another value")

	logger2 := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: LevelTrace,
	}))
	logger2.Debug("debug")
	logger2.Info("info")
	logger2.Warn("warn")
	logger2.Error("error")

	logger2.Info("context", "key", "value", "another key", "another value")

	slog.SetDefault(logger2)

	logFromElsewhere("via function")
	sub.LogFromPackage("via package")

}

func logFromElsewhere(message string) {
	slog.Debug(message)
}

// output:
// 2025/10/24 18:13:31 This is log.println
// 2025/10/24 18:13:31.692580 with micro
// 2025/10/24 18:13:31 main.go:21: with file/line
// my:2025/10/24 18:13:31 from mylog
// ohmy:2025/10/24 18:13:31 from mylog
// from buflog:buf:2025/10/24 18:13:31 hello
// 2025/10/24 18:13:31 INFO info level
// {"time":"2025-10-24T18:13:31.692722+01:00","level":"INFO","msg":"info"}
// {"time":"2025-10-24T18:13:31.692744+01:00","level":"WARN","msg":"warn"}
// {"time":"2025-10-24T18:13:31.692751+01:00","level":"ERROR","msg":"error"}
// {"time":"2025-10-24T18:13:31.692758+01:00","level":"INFO","msg":"context","key":"value","another key":"another value"}
// time=2025-10-24T18:13:31.692+01:00 level=DEBUG msg=debug
// time=2025-10-24T18:13:31.692+01:00 level=INFO msg=info
// time=2025-10-24T18:13:31.692+01:00 level=WARN msg=warn
// time=2025-10-24T18:13:31.692+01:00 level=ERROR msg=error
// time=2025-10-24T18:13:31.692+01:00 level=INFO msg=context key=value "another key"="another value"
// time=2025-10-24T18:13:31.692+01:00 level=DEBUG msg="via function"
// time=2025-10-24T18:13:31.692+01:00 level=DEBUG msg="via package"

// setting the default logger in main.main changes the behaviour (showing DEBUG) in both the func in main and the sub package
// - logging can be defined at the start of the application and be used everywhere else via the default!

// level is just an integer (default 0 = INFO)
// anything above that level is displayed

// TODO: custom levels
// https://betterstack.com/community/guides/logging/logging-in-go/#customizing-slog-levels

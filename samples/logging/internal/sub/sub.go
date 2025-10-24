package sub

import "log/slog"

func LogFromPackage(message string) {
	slog.Debug(message)
}

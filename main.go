package logger

import "log"

func LogInfo(message string) {
	log.Printf("Log Information - %v", message)
}

func LogWarning(message string) {
	log.Printf("Log Warning - %v", message)
}

func LogError(message string) {
	log.Printf("Log Error - %v", message)
}

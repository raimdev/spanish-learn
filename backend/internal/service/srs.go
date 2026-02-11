package service

import (
	"math"
	"time"
)

// SM2Result holds the output of the SM-2 algorithm.
type SM2Result struct {
	Repetitions    int64
	EasinessFactor float64
	IntervalDays   int64
	NextReviewDate string
}

// SM2 implements the SuperMemo SM-2 spaced repetition algorithm.
// quality: 0-5 rating (0-2 = incorrect, 3-5 = correct with varying ease)
func SM2(quality int, repetitions int64, easinessFactor float64, intervalDays int64) SM2Result {
	var newReps int64
	var newInterval int64
	newEF := easinessFactor

	if quality >= 3 {
		switch repetitions {
		case 0:
			newInterval = 1
		case 1:
			newInterval = 6
		default:
			newInterval = int64(math.Round(float64(intervalDays) * easinessFactor))
		}
		newReps = repetitions + 1
	} else {
		newReps = 0
		newInterval = 1
	}

	// Update easiness factor
	q := float64(quality)
	newEF = newEF + (0.1 - (5.0-q)*(0.08+(5.0-q)*0.02))
	if newEF < 1.3 {
		newEF = 1.3
	}

	nextReview := time.Now().Add(time.Duration(newInterval) * 24 * time.Hour).UTC().Format("2006-01-02 15:04:05")

	return SM2Result{
		Repetitions:    newReps,
		EasinessFactor: newEF,
		IntervalDays:   newInterval,
		NextReviewDate: nextReview,
	}
}

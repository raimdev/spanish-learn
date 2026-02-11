package service

import (
	"testing"
	"time"
)

func TestSM2_CorrectFirstTime(t *testing.T) {
	result := SM2(4, 0, 2.5, 0)

	if result.Repetitions != 1 {
		t.Errorf("expected repetitions=1, got %d", result.Repetitions)
	}
	if result.IntervalDays != 1 {
		t.Errorf("expected interval=1, got %d", result.IntervalDays)
	}
	if result.EasinessFactor < 2.3 || result.EasinessFactor > 2.7 {
		t.Errorf("expected EF around 2.5, got %f", result.EasinessFactor)
	}
}

func TestSM2_CorrectSecondTime(t *testing.T) {
	result := SM2(4, 1, 2.5, 1)

	if result.Repetitions != 2 {
		t.Errorf("expected repetitions=2, got %d", result.Repetitions)
	}
	if result.IntervalDays != 6 {
		t.Errorf("expected interval=6, got %d", result.IntervalDays)
	}
}

func TestSM2_CorrectThirdTime(t *testing.T) {
	result := SM2(4, 2, 2.5, 6)

	if result.Repetitions != 3 {
		t.Errorf("expected repetitions=3, got %d", result.Repetitions)
	}
	// interval = round(6 * 2.5) = 15
	if result.IntervalDays != 15 {
		t.Errorf("expected interval=15, got %d", result.IntervalDays)
	}
}

func TestSM2_IncorrectResetsRepetitions(t *testing.T) {
	result := SM2(1, 5, 2.5, 30)

	if result.Repetitions != 0 {
		t.Errorf("expected repetitions=0 after incorrect, got %d", result.Repetitions)
	}
	if result.IntervalDays != 1 {
		t.Errorf("expected interval=1 after incorrect, got %d", result.IntervalDays)
	}
}

func TestSM2_EasinessFactorNeverBelowMin(t *testing.T) {
	// Repeated failures should not drop EF below 1.3
	result := SM2(0, 0, 1.3, 0)

	if result.EasinessFactor < 1.3 {
		t.Errorf("EF should not go below 1.3, got %f", result.EasinessFactor)
	}
}

func TestSM2_QualityThreeIsPassingThreshold(t *testing.T) {
	result := SM2(3, 0, 2.5, 0)

	if result.Repetitions != 1 {
		t.Errorf("quality=3 should be passing, expected reps=1, got %d", result.Repetitions)
	}
}

func TestSM2_QualityTwoIsFailure(t *testing.T) {
	result := SM2(2, 3, 2.5, 15)

	if result.Repetitions != 0 {
		t.Errorf("quality=2 should be failure, expected reps=0, got %d", result.Repetitions)
	}
}

func TestSM2_EasyResponseIncreasesEF(t *testing.T) {
	result := SM2(5, 0, 2.5, 0)

	if result.EasinessFactor <= 2.5 {
		t.Errorf("quality=5 should increase EF, got %f", result.EasinessFactor)
	}
}

func TestSM2_HardResponseDecreasesEF(t *testing.T) {
	result := SM2(3, 0, 2.5, 0)

	if result.EasinessFactor >= 2.5 {
		t.Errorf("quality=3 should decrease EF, got %f", result.EasinessFactor)
	}
}

func TestSM2_NextReviewDateIsFuture(t *testing.T) {
	result := SM2(4, 0, 2.5, 0)

	nextDate, err := time.Parse("2006-01-02 15:04:05", result.NextReviewDate)
	if err != nil {
		t.Fatalf("could not parse next review date: %v", err)
	}

	if !nextDate.After(time.Now()) {
		t.Errorf("next review date should be in the future, got %s", result.NextReviewDate)
	}
}

func TestSM2_ProgressionSequence(t *testing.T) {
	// Simulate a card being reviewed correctly multiple times
	var reps int64
	var ef float64 = 2.5
	var interval int64

	expectedIntervals := []int64{1, 6}

	for i := 0; i < 5; i++ {
		result := SM2(4, reps, ef, interval)
		if i < len(expectedIntervals) {
			if result.IntervalDays != expectedIntervals[i] {
				t.Errorf("review %d: expected interval=%d, got %d", i, expectedIntervals[i], result.IntervalDays)
			}
		}
		// Intervals should be non-decreasing for consistent correct answers
		if i > 0 && result.IntervalDays < interval {
			t.Errorf("review %d: interval decreased from %d to %d", i, interval, result.IntervalDays)
		}
		reps = result.Repetitions
		ef = result.EasinessFactor
		interval = result.IntervalDays
	}
}

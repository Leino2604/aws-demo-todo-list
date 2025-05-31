package handlers

import (
	"net/http"
	"time"
	"sync"
)

type PomodoroHandler struct {
	mu          sync.Mutex
	isRunning   bool
	workTime    time.Duration
	breakTime   time.Duration
	cycles      int
	currentCycle int
	timer       *time.Timer
}

func NewPomodoroHandler(workTime, breakTime time.Duration, cycles int) *PomodoroHandler {
	return &PomodoroHandler{
		workTime:  workTime,
		breakTime: breakTime,
		cycles:    cycles,
	}
}

func (h *PomodoroHandler) StartPomodoro(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.isRunning {
		http.Error(w, "Pomodoro is already running", http.StatusConflict)
		return
	}

	h.isRunning = true
	h.currentCycle = 0
	h.startTimer(w)
}

func (h *PomodoroHandler) startTimer(w http.ResponseWriter) {
	h.timer = time.AfterFunc(h.workTime, func() {
		h.isRunning = false
		h.currentCycle++
		if h.currentCycle < h.cycles {
			h.startBreak(w)
		} else {
			http.Error(w, "Pomodoro completed", http.StatusOK)
		}
	})
}

func (h *PomodoroHandler) startBreak(w http.ResponseWriter) {
	h.timer = time.AfterFunc(h.breakTime, func() {
		h.isRunning = false
		h.startTimer(w)
	})
}

func (h *PomodoroHandler) StopPomodoro(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.isRunning {
		http.Error(w, "Pomodoro is not running", http.StatusConflict)
		return
	}

	h.timer.Stop()
	h.isRunning = false
	http.Error(w, "Pomodoro stopped", http.StatusOK)
}
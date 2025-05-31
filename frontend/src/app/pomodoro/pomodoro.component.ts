import { Component } from '@angular/core';

@Component({
  selector: 'app-pomodoro',
  templateUrl: './pomodoro.component.html',
  styleUrls: ['./pomodoro.component.css']
})
export class PomodoroComponent {
  workTime: number = 25; // default work time in minutes
  restTime: number = 5; // default rest time in minutes
  cycles: number = 4; // default number of cycles
  currentCycle: number = 0;
  timer: any;
  isWorking: boolean = false;
  timeLeft: number = this.workTime * 60; // time left in seconds

  startTimer() {
    this.isWorking = true;
    this.timeLeft = this.workTime * 60;
    this.timer = setInterval(() => {
      if (this.timeLeft > 0) {
        this.timeLeft--;
      } else {
        this.completeCycle();
      }
    }, 1000);
  }

  completeCycle() {
    clearInterval(this.timer);
    this.currentCycle++;
    if (this.currentCycle < this.cycles) {
      this.isWorking = false;
      this.timeLeft = this.restTime * 60;
      this.timer = setInterval(() => {
        if (this.timeLeft > 0) {
          this.timeLeft--;
        } else {
          this.startTimer();
        }
      }, 1000);
    } else {
      this.resetTimer();
    }
  }

  resetTimer() {
    clearInterval(this.timer);
    this.isWorking = false;
    this.currentCycle = 0;
    this.timeLeft = this.workTime * 60;
  }

  updateSettings(work: number, rest: number, cycles: number) {
    this.workTime = work;
    this.restTime = rest;
    this.cycles = cycles;
    this.resetTimer();
  }
}
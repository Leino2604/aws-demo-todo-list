import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-pomodoro',
  templateUrl: './pomodoro.component.html',
  styleUrls: ['./pomodoro.component.css']
})
export class PomodoroComponent implements OnInit {
  workDuration: number = 25;
  restDuration: number = 5;
  cycles: number = 4;
  timer: any;
  isActive: boolean = false;
  remainingTime: number = 0;
  isWorkTime: boolean = true;
  currentCycle: number = 1;

  constructor() { }

  ngOnInit(): void {
  }

  startPomodoro(): void {
    if (!this.isActive) {
      this.isActive = true;
      this.isWorkTime = true;
      this.remainingTime = this.workDuration * 60;
      this.startTimer();
    }
  }

  stopPomodoro(): void {
    if (this.isActive) {
      this.isActive = false;
      clearInterval(this.timer);
    }
  }

  private startTimer(): void {
    this.timer = setInterval(() => {
      if (this.remainingTime > 0) {
        this.remainingTime--;
      } else {
        if (this.isWorkTime) {
          this.isWorkTime = false;
          this.remainingTime = this.restDuration * 60;
        } else {
          this.isWorkTime = true;
          this.currentCycle++;

          if (this.currentCycle > this.cycles) {
            this.stopPomodoro();
            this.currentCycle = 1;
          } else {
            this.remainingTime = this.workDuration * 60;
          }
        }
      }
    }, 1000);
  }
}
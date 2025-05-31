import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css']
})
export class TodoComponent implements OnInit {
  tasks: { name: string; estimatedTime: number; status: string }[] = [];
  newTaskName: string = '';
  newTaskEstimatedTime: number | null = null;

  constructor() {}

  ngOnInit(): void {
    // Load tasks from local storage or initialize empty
    const storedTasks = localStorage.getItem('tasks');
    this.tasks = storedTasks ? JSON.parse(storedTasks) : [];
  }

  addTask(): void {
    if (this.newTaskName && this.newTaskEstimatedTime) {
      this.tasks.push({
        name: this.newTaskName,
        estimatedTime: this.newTaskEstimatedTime,
        status: 'pending'
      });
      this.saveTasks();
      this.newTaskName = '';
      this.newTaskEstimatedTime = null;
    }
  }

  startTask(task: { name: string; estimatedTime: number; status: string }): void {
    task.status = 'in-progress';
    this.saveTasks();
    this.startTimer(task);
  }

  startTimer(task: { estimatedTime: number }): void {
    let timeLeft = task.estimatedTime;
    const timerInterval = setInterval(() => {
      if (timeLeft > 0) {
        timeLeft--;
      } else {
        clearInterval(timerInterval);
        task.status = 'completed';
        this.saveTasks();
      }
    }, 1000);
  }

  saveTasks(): void {
    localStorage.setItem('tasks', JSON.stringify(this.tasks));
  }
}
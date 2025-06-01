export class TodoItemComponent {
  description: string;
  estimatedTime: number;
  remainingTime: number;
  isActive: boolean;

  constructor(description: string, estimatedTime: number) {
    this.description = description;
    this.estimatedTime = estimatedTime;
    this.remainingTime = estimatedTime;
    this.isActive = false;
  }

  startTask() {
    this.isActive = true;
    this.countdown();
  }

  countdown() {
    const interval = setInterval(() => {
      if (this.remainingTime > 0 && this.isActive) {
        this.remainingTime--;
      } else {
        clearInterval(interval);
        this.isActive = false;
      }
    }, 1000);
  }

  resetTask() {
    this.remainingTime = this.estimatedTime;
    this.isActive = false;
  }
}
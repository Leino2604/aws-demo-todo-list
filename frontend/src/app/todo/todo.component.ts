import { Component, OnInit } from '@angular/core';

interface TodoItem {
  title: string;
  estimatedTime: number;
  status: string;
}

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css']
})
export class TodoComponent implements OnInit {
  todos: TodoItem[] = [];
  newTodo: TodoItem = { title: '', estimatedTime: 0, status: 'pending' };

  constructor() { }

  ngOnInit(): void {
    // Load tasks from local storage or initialize empty
    const storedTasks = localStorage.getItem('tasks');
    this.todos = storedTasks ? JSON.parse(storedTasks) : [];
  }

  addTodo(): void {
    if (this.newTodo.title.trim()) {
      this.todos.push({ ...this.newTodo });
      this.saveTasks();
      this.newTodo = { title: '', estimatedTime: 0, status: 'pending' };
    }
  }

  removeTodo(index: number): void {
    this.todos.splice(index, 1);
    this.saveTasks();
  }

  markAsCompleted(todo: TodoItem): void {
    todo.status = 'completed';
    this.saveTasks();
  }

  saveTasks(): void {
    localStorage.setItem('tasks', JSON.stringify(this.todos));
  }
}
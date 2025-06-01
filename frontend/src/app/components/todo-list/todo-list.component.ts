export class TodoListComponent {
  todos: Todo[] = [];
  newTodoDescription: string = '';
  newTodoEstimatedTime: number | null = null;

  constructor(private todoService: TodoService) {}

  ngOnInit() {
    this.loadTodos();
  }

  loadTodos() {
    this.todoService.getTodos().subscribe((todos) => {
      this.todos = todos;
    });
  }

  addTodo() {
    if (this.newTodoDescription && this.newTodoEstimatedTime) {
      const newTodo: Todo = {
        description: this.newTodoDescription,
        estimatedTime: this.newTodoEstimatedTime,
        status: 'pending',
      };
      this.todoService.addTodo(newTodo).subscribe(() => {
        this.loadTodos();
        this.newTodoDescription = '';
        this.newTodoEstimatedTime = null;
      });
    }
  }

  removeTodo(todoId: string) {
    this.todoService.removeTodo(todoId).subscribe(() => {
      this.loadTodos();
    });
  }

  startTask(todoId: string) {
    this.todoService.startTask(todoId).subscribe(() => {
      this.loadTodos();
    });
  }
}
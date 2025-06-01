export class TodoService {
    private apiUrl: string;

    constructor() {
        this.apiUrl = 'http://localhost:8080/api/todos'; // Change this for production
    }

    getTodos(): Promise<Todo[]> {
        return fetch(this.apiUrl)
            .then(response => response.json());
    }

    addTodo(todo: Todo): Promise<Todo> {
        return fetch(this.apiUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo)
        }).then(response => response.json());
    }

    updateTodo(todo: Todo): Promise<Todo> {
        return fetch(`${this.apiUrl}/${todo.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo)
        }).then(response => response.json());
    }

    deleteTodo(id: string): Promise<void> {
        return fetch(`${this.apiUrl}/${id}`, {
            method: 'DELETE'
        }).then(() => {});
    }
}
export interface Todo {
    id: string;
    description: string;
    estimatedTime: number; // in minutes
    status: 'pending' | 'in-progress' | 'completed';
    startTime?: Date; // optional, only when the task is in progress
}
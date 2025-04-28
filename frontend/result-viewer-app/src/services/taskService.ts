import { Task } from '../types/task';

const API_URL = 'http://localhost:8081';

export const taskService = {
  async getTasks(): Promise<Task[]> {
    const response = await fetch(`${API_URL}/tasks`);

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch tasks');
    }

    return response.json();
  },

  async getTask(id: string): Promise<Task> {
    const response = await fetch(`${API_URL}/tasks/${id}`);

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch task');
    }

    return response.json();
  },
};

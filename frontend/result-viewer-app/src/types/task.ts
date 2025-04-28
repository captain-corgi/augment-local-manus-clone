export type TaskStatus = 'pending' | 'running' | 'completed' | 'failed';

export interface Task {
  id: string;
  title: string;
  description: string;
  status: TaskStatus;
  input: string;
  result?: string;
  created_at: string;
  updated_at: string;
}

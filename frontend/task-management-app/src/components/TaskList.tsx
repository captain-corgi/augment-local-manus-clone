import { useEffect } from 'react';
import { Task } from '../types/task';
import { taskService } from '../services/taskService';
import TaskStatusBadge from './TaskStatusBadge';

interface TaskListProps {
  tasks: Task[];
  isLoading: boolean;
  setTasks: (tasks: Task[]) => void;
  setIsLoading: (isLoading: boolean) => void;
  setError: (error: string | null) => void;
}

function TaskList({ tasks, isLoading, setTasks, setIsLoading, setError }: TaskListProps) {
  useEffect(() => {
    const fetchTasks = async () => {
      setIsLoading(true);
      setError(null);
      
      try {
        const fetchedTasks = await taskService.getTasks();
        setTasks(fetchedTasks);
      } catch (error) {
        setError(error instanceof Error ? error.message : 'An unknown error occurred');
      } finally {
        setIsLoading(false);
      }
    };
    
    fetchTasks();
    
    // Poll for task updates every 5 seconds
    const interval = setInterval(fetchTasks, 5000);
    
    return () => clearInterval(interval);
  }, [setTasks, setIsLoading, setError]);

  if (isLoading && tasks.length === 0) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (tasks.length === 0) {
    return (
      <div className="bg-white shadow-md rounded px-8 py-6 mb-4">
        <p className="text-gray-500 text-center">No tasks found. Create a new task to get started.</p>
      </div>
    );
  }

  return (
    <div className="bg-white shadow-md rounded overflow-hidden">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Title
            </th>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Status
            </th>
            <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Created
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {tasks.map((task) => (
            <tr key={task.id} className="hover:bg-gray-50">
              <td className="px-6 py-4 whitespace-nowrap">
                <div className="text-sm font-medium text-gray-900">{task.title}</div>
                <div className="text-sm text-gray-500">{task.description}</div>
              </td>
              <td className="px-6 py-4 whitespace-nowrap">
                <TaskStatusBadge status={task.status} />
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {new Date(task.created_at).toLocaleString()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default TaskList;

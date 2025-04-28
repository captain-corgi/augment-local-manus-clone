import { Task } from '../types/task';
import TaskStatusBadge from './TaskStatusBadge';

interface TaskResultListProps {
  tasks: Task[];
  isLoading: boolean;
  selectedTaskId: string | undefined;
  onTaskSelect: (task: Task) => void;
}

function TaskResultList({ tasks, isLoading, selectedTaskId, onTaskSelect }: TaskResultListProps) {
  if (isLoading && tasks.length === 0) {
    return (
      <div className="bg-white shadow-md rounded px-8 py-6 mb-4">
        <div className="flex justify-center items-center h-64">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
        </div>
      </div>
    );
  }

  if (tasks.length === 0) {
    return (
      <div className="bg-white shadow-md rounded px-8 py-6 mb-4">
        <p className="text-gray-500 text-center">No completed tasks found.</p>
      </div>
    );
  }

  return (
    <div className="bg-white shadow-md rounded overflow-hidden">
      <div className="px-4 py-3 bg-gray-50 border-b border-gray-200">
        <h2 className="text-lg font-medium text-gray-900">Completed Tasks</h2>
      </div>
      <ul className="divide-y divide-gray-200">
        {tasks.map((task) => (
          <li 
            key={task.id}
            className={`cursor-pointer hover:bg-gray-50 ${selectedTaskId === task.id ? 'bg-blue-50' : ''}`}
            onClick={() => onTaskSelect(task)}
          >
            <div className="px-4 py-4 sm:px-6">
              <div className="flex items-center justify-between">
                <p className="text-sm font-medium text-gray-900 truncate">{task.title}</p>
                <TaskStatusBadge status={task.status} />
              </div>
              <div className="mt-2 sm:flex sm:justify-between">
                <div className="sm:flex">
                  <p className="flex items-center text-sm text-gray-500 truncate">
                    {task.description ? task.description.substring(0, 50) + (task.description.length > 50 ? '...' : '') : 'No description'}
                  </p>
                </div>
                <div className="mt-2 flex items-center text-sm text-gray-500 sm:mt-0">
                  <p>
                    {new Date(task.updated_at).toLocaleString()}
                  </p>
                </div>
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default TaskResultList;

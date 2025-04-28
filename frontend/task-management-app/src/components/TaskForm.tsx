import { useForm } from 'react-hook-form';
import { CreateTaskInput, Task } from '../types/task';
import { taskService } from '../services/taskService';

interface TaskFormProps {
  onTaskCreated: (task: Task) => void;
  setIsLoading: (isLoading: boolean) => void;
  setError: (error: string | null) => void;
}

function TaskForm({ onTaskCreated, setIsLoading, setError }: TaskFormProps) {
  const { register, handleSubmit, reset, formState: { errors } } = useForm<CreateTaskInput>();

  const onSubmit = async (data: CreateTaskInput) => {
    setIsLoading(true);
    setError(null);
    
    try {
      const task = await taskService.createTask(data);
      onTaskCreated(task);
      reset();
    } catch (error) {
      setError(error instanceof Error ? error.message : 'An unknown error occurred');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
      <div className="mb-4">
        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="title">
          Title
        </label>
        <input
          className={`shadow appearance-none border ${errors.title ? 'border-red-500' : 'border-gray-300'} rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline`}
          id="title"
          type="text"
          placeholder="Task title"
          {...register('title', { required: 'Title is required' })}
        />
        {errors.title && <p className="text-red-500 text-xs italic">{errors.title.message}</p>}
      </div>
      
      <div className="mb-4">
        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="description">
          Description
        </label>
        <textarea
          className={`shadow appearance-none border ${errors.description ? 'border-red-500' : 'border-gray-300'} rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline`}
          id="description"
          placeholder="Task description"
          rows={3}
          {...register('description')}
        />
      </div>
      
      <div className="mb-6">
        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="input">
          Input
        </label>
        <textarea
          className={`shadow appearance-none border ${errors.input ? 'border-red-500' : 'border-gray-300'} rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline`}
          id="input"
          placeholder="Task input (e.g., code, prompt, etc.)"
          rows={5}
          {...register('input', { required: 'Input is required' })}
        />
        {errors.input && <p className="text-red-500 text-xs italic">{errors.input.message}</p>}
      </div>
      
      <div className="flex items-center justify-end">
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          type="submit"
        >
          Create Task
        </button>
      </div>
    </form>
  );
}

export default TaskForm;

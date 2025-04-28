import { useState } from 'react';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';
import { Task } from './types/task';

function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const addTask = (task: Task) => {
    setTasks([task, ...tasks]);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-6">Task Management</h1>
      
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div>
          <h2 className="text-xl font-semibold mb-4">Create New Task</h2>
          <TaskForm 
            onTaskCreated={addTask} 
            setIsLoading={setIsLoading} 
            setError={setError} 
          />
          {error && (
            <div className="mt-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
              <strong className="font-bold">Error!</strong>
              <span className="block sm:inline"> {error}</span>
            </div>
          )}
        </div>
        
        <div>
          <h2 className="text-xl font-semibold mb-4">Task List</h2>
          <TaskList 
            tasks={tasks} 
            isLoading={isLoading} 
            setTasks={setTasks} 
            setIsLoading={setIsLoading} 
            setError={setError} 
          />
        </div>
      </div>
    </div>
  );
}

export default App;

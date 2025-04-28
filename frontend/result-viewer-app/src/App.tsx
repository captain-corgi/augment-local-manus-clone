import { useState, useEffect } from 'react';
import TaskResultList from './components/TaskResultList';
import TaskResultDetail from './components/TaskResultDetail';
import { Task } from './types/task';
import { taskService } from './services/taskService';

function App() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchTasks = async () => {
      setIsLoading(true);
      setError(null);
      
      try {
        const fetchedTasks = await taskService.getTasks();
        // Filter tasks that have results or are completed/failed
        const completedTasks = fetchedTasks.filter(
          task => task.result || task.status === 'completed' || task.status === 'failed'
        );
        setTasks(completedTasks);
        
        // Select the first task by default if available
        if (completedTasks.length > 0 && !selectedTask) {
          setSelectedTask(completedTasks[0]);
        }
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
  }, [selectedTask]);

  const handleTaskSelect = (task: Task) => {
    setSelectedTask(task);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-6">Task Results</h1>
      
      {error && (
        <div className="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
          <strong className="font-bold">Error!</strong>
          <span className="block sm:inline"> {error}</span>
        </div>
      )}
      
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-1">
          <TaskResultList 
            tasks={tasks} 
            isLoading={isLoading} 
            selectedTaskId={selectedTask?.id} 
            onTaskSelect={handleTaskSelect} 
          />
        </div>
        
        <div className="lg:col-span-2">
          {selectedTask ? (
            <TaskResultDetail task={selectedTask} />
          ) : (
            <div className="bg-white shadow-md rounded px-8 py-6 mb-4">
              <p className="text-gray-500 text-center">Select a task to view its result</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;

import { useState } from 'react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';
import { Task } from '../types/task';
import TaskStatusBadge from './TaskStatusBadge';

interface TaskResultDetailProps {
  task: Task;
}

function TaskResultDetail({ task }: TaskResultDetailProps) {
  const [activeTab, setActiveTab] = useState<'result' | 'input'>('result');

  const detectLanguage = (code: string): string => {
    // Simple language detection based on content
    if (code.includes('function') || code.includes('const') || code.includes('let') || code.includes('var')) {
      return 'javascript';
    } else if (code.includes('def ') || code.includes('import ') || code.includes('class ') || code.includes('print(')) {
      return 'python';
    } else if (code.includes('package ') || code.includes('func ') || code.includes('import (')) {
      return 'go';
    } else if (code.includes('public class') || code.includes('public static void main')) {
      return 'java';
    } else if (code.includes('<html') || code.includes('<!DOCTYPE')) {
      return 'html';
    } else if (code.includes('SELECT ') || code.includes('FROM ') || code.includes('WHERE ')) {
      return 'sql';
    }
    return 'text';
  };

  return (
    <div className="bg-white shadow-md rounded overflow-hidden">
      <div className="px-4 py-5 sm:px-6">
        <div className="flex justify-between items-center">
          <h3 className="text-lg leading-6 font-medium text-gray-900">{task.title}</h3>
          <TaskStatusBadge status={task.status} />
        </div>
        <p className="mt-1 max-w-2xl text-sm text-gray-500">{task.description}</p>
      </div>
      
      <div className="border-t border-gray-200">
        <div className="flex border-b">
          <button
            className={`px-4 py-2 text-sm font-medium ${
              activeTab === 'result' ? 'border-b-2 border-blue-500 text-blue-600' : 'text-gray-500 hover:text-gray-700'
            }`}
            onClick={() => setActiveTab('result')}
          >
            Result
          </button>
          <button
            className={`px-4 py-2 text-sm font-medium ${
              activeTab === 'input' ? 'border-b-2 border-blue-500 text-blue-600' : 'text-gray-500 hover:text-gray-700'
            }`}
            onClick={() => setActiveTab('input')}
          >
            Input
          </button>
        </div>
        
        <div className="px-4 py-5 sm:p-0">
          {activeTab === 'result' ? (
            <div className="sm:p-6">
              {task.status === 'completed' && task.result ? (
                <SyntaxHighlighter 
                  language={detectLanguage(task.result)} 
                  style={vscDarkPlus}
                  customStyle={{ borderRadius: '0.375rem' }}
                  showLineNumbers
                >
                  {task.result}
                </SyntaxHighlighter>
              ) : task.status === 'failed' ? (
                <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative">
                  <strong className="font-bold">Error:</strong>
                  <span className="block sm:inline"> {task.result || 'Task execution failed'}</span>
                </div>
              ) : (
                <div className="text-gray-500 text-center py-8">
                  {task.status === 'pending' ? 'Task is pending execution' : 'Task is currently running'}
                </div>
              )}
            </div>
          ) : (
            <div className="sm:p-6">
              <SyntaxHighlighter 
                language={detectLanguage(task.input)} 
                style={vscDarkPlus}
                customStyle={{ borderRadius: '0.375rem' }}
                showLineNumbers
              >
                {task.input}
              </SyntaxHighlighter>
            </div>
          )}
        </div>
      </div>
      
      <div className="px-4 py-3 bg-gray-50 text-right sm:px-6 border-t border-gray-200">
        <div className="text-sm text-gray-500">
          <span className="font-medium">Created:</span> {new Date(task.created_at).toLocaleString()}
          <span className="mx-2">|</span>
          <span className="font-medium">Updated:</span> {new Date(task.updated_at).toLocaleString()}
        </div>
      </div>
    </div>
  );
}

export default TaskResultDetail;

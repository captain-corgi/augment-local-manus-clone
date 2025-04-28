function LoadingFallback() {
  return (
    <div className="flex items-center justify-center h-full">
      <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      <p className="ml-3 text-lg font-medium text-gray-700">Loading...</p>
    </div>
  );
}

export default LoadingFallback;

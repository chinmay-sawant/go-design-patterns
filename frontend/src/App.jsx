import { useState, useCallback, useEffect } from 'react';
import { Sidebar } from './components/Sidebar';
import { CodeViewer } from './components/CodeViewer';
import { Code, Github } from 'lucide-react';
import data from './data.json'; 

function App() {
  const [width, setWidth] = useState(260);
  const [isResizing, setIsResizing] = useState(false);
  const [activeFile, setActiveFile] = useState(null);
  const [activePath, setActivePath] = useState(null);

  const startResizing = useCallback(() => {
    setIsResizing(true);
  }, []);

  const stopResizing = useCallback(() => {
    setIsResizing(false);
  }, []);

  const resize = useCallback(
    (mouseMoveEvent) => {
      if (isResizing) {
        const newWidth = mouseMoveEvent.clientX;
        if (newWidth > 180 && newWidth < 500) {
          setWidth(newWidth);
        }
      }
    },
    [isResizing]
  );

  useEffect(() => {
    window.addEventListener("mousemove", resize);
    window.addEventListener("mouseup", stopResizing);
    return () => {
      window.removeEventListener("mousemove", resize);
      window.removeEventListener("mouseup", stopResizing);
    };
  }, [resize, stopResizing]);

  const handleSelect = (item) => {
    setActiveFile(item);
    setActivePath(item.path);
  };

  const getGithubUrl = (path) => {
    const cleanPath = path.replace(/^\.\//, ''); 
    return `https://github.com/chinmay-sawant/go-design-patterns/blob/master/${cleanPath}`;
  };

  return (
    <div className="flex h-screen w-full bg-black text-neutral-200 font-sans overflow-hidden">
      {/* Sidebar */}
      <div 
        style={{ width: width }} 
        className="flex-shrink-0 relative flex flex-col h-full border-r border-neutral-800/50 select-none"
      >
        <Sidebar 
          data={data} 
          onSelect={handleSelect} 
          activePath={activePath} 
        />

        {/* Resizer */}
        <div
          className={`absolute right-0 top-0 w-1 h-full cursor-col-resize transition-colors z-50 ${isResizing ? 'bg-sky-500' : 'bg-transparent hover:bg-neutral-700'}`}
          onMouseDown={startResizing}
        />
      </div>

      {/* Main Content */}
      <div className="flex-1 flex flex-col h-full bg-neutral-950 relative min-w-0">
        {activeFile ? (
          <div className="flex flex-col h-full">
            {/* File Header */}
            <div className="h-10 bg-black border-b border-neutral-800/50 flex items-center px-4 justify-between flex-shrink-0">
              <span className="text-xs text-neutral-500 truncate font-mono">
                {activeFile.path}
              </span>
              <a 
                href={getGithubUrl(activeFile.path)} 
                target="_blank" 
                rel="noopener noreferrer"
                className="flex items-center text-xs text-neutral-400 hover:text-white transition-colors"
              >
                <Github size={12} className="mr-1" />
                GitHub
              </a>
            </div>
            {/* Code */}
            <div className="flex-1 overflow-hidden">
              <CodeViewer 
                code={activeFile.content} 
                language="go" 
                filename={activeFile.name} 
              />
            </div>
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center h-full">
            <div className="text-center max-w-sm">
              <Code className="w-10 h-10 mx-auto mb-4 text-neutral-700" strokeWidth={1.5} />
              <h2 className="text-lg font-medium text-neutral-300 mb-2">Go Design Patterns</h2>
              <p className="text-sm text-neutral-600">
                Select a file from the sidebar to view its contents.
              </p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;

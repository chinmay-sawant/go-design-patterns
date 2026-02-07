import { useState, useCallback, useEffect } from 'react';
import { Sidebar } from './components/Sidebar';
import { CodeViewer } from './components/CodeViewer';
import { motion as Motion } from 'framer-motion';
import { Code, BookOpen, GitBranch, Github } from 'lucide-react';
import data from './data.json'; 

function App() {
  const [width, setWidth] = useState(300);
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
        if (newWidth > 200 && newWidth < 800) {
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

  // Convert local relative path to GitHub URL if needed
  const getGithubUrl = (path) => {
    // path example: ./design_patterns/behavioral/chain_of_responsibility/main.go
    // desired: https://github.com/chinmay-sawant/go-design-patterns/tree/master/design_patterns/behavioral/chain_of_responsibility/main.go
    // Adjust logic to strip leading './' and prepend base
    const cleanPath = path.replace(/^\.\//, ''); 
    return `https://github.com/chinmay-sawant/go-design-patterns/blob/master/${cleanPath}`;
  };

  return (
    <div className="flex h-screen w-full bg-bg-primary text-text-primary font-sans overflow-hidden">
      {/* Sidebar Area */}
      <div 
        style={{ width: width }} 
        className="flex-shrink-0 relative flex flex-col h-full bg-bg-secondary border-r border-border select-none"
      >
        <div className="flex-1 overflow-hidden flex flex-col">
            <Sidebar 
                data={data} 
                onSelect={handleSelect} 
                activePath={activePath} 
            />
        </div>

        {/* Footer/Status Bar */}
        <div className="h-8 border-t border-border flex items-center px-3 text-xs text-text-secondary bg-bg-tertiary flex-shrink-0 justify-between">
          <div className="flex items-center">
            <GitBranch className="w-3 h-3 mr-1" />
            <span>master</span>
          </div>
          <span className="opacity-50 text-[10px]">
            {activePath || 'No file selected'}
          </span>
        </div>

        {/* Resizer Handle */}
        <div
          className={`absolute right-0 top-0 w-1 h-full cursor-col-resize hover:bg-accent transition-colors z-50 ${isResizing ? 'bg-accent' : 'bg-transparent'}`}
          onMouseDown={startResizing}
        />
      </div>

      {/* Main Content Area */}
      <div className="flex-1 flex flex-col h-full bg-bg-primary relative min-w-0">
        {activeFile ? (
          <div className="flex flex-col h-full">
             <div className="h-10 bg-bg-tertiary border-b border-border flex items-center px-4 justify-between flex-shrink-0">
                <span className="text-sm text-text-secondary truncate font-mono opacity-80">
                  {activeFile.path}
                </span>
                <a 
                  href={getGithubUrl(activeFile.path)} 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="flex items-center text-xs text-accent hover:underline hover:text-accent-hover transition-colors"
                >
                  <Github size={12} className="mr-1" />
                  View on GitHub
                </a>
             </div>
             <div className="flex-1 overflow-hidden">
                <CodeViewer 
                    code={activeFile.content} 
                    language="go" 
                    filename={activeFile.name} 
                />
             </div>
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center h-full text-text-secondary">
            <Motion.div 
              initial={{ opacity: 0, scale: 0.95 }} 
              animate={{ opacity: 1, scale: 1 }}
              transition={{ duration: 0.3 }}
              className="text-center max-w-md p-8 border border-border rounded-xl bg-bg-secondary shadow-2xl"
            >
              <Code className="w-12 h-12 mx-auto mb-4 text-accent opacity-80" />
              <h2 className="text-xl font-semibold text-text-primary mb-2">Go Design Patterns</h2>
              <p className="text-sm opacity-60 mb-6">
                Explore idiomatic Go implementations of classic design patterns. 
                Select a pattern from the sidebar to begin.
              </p>
            </Motion.div>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;

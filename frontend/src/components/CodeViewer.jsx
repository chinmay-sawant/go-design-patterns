import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';

export const CodeViewer = ({ code, language = 'go', filename }) => {
  if (!code) {
    return (
      <div className="flex items-center justify-center h-full text-text-secondary bg-bg-primary">
        <div className="text-center">
          <div className="text-4xl mb-4 opacity-20">⌘</div>
          <p>Select a file to view content</p>
        </div>
      </div>
    );
  }

  return (
    <div className="h-full flex flex-col bg-bg-primary flex-1 min-w-0">
      <div className="h-14 border-b border-border flex items-center justify-between px-6 bg-bg-secondary flex-shrink-0">
        <span className="text-sm font-medium text-text-primary flex items-center gap-2 truncate">
          <span className="w-2 h-2 rounded-full bg-accent"></span>
          {filename}
        </span>
        <span className="text-xs text-text-secondary uppercase tracking-widest font-mono hidden sm:block">
          {language}
        </span>
      </div>
      <div className="flex-1 overflow-auto relative custom-scrollbar flex flex-col min-h-0">
        <SyntaxHighlighter
          language={language}
          style={vscDarkPlus}
          customStyle={{
            margin: 0,
            padding: '2rem',
            background: 'transparent',
            fontSize: '14px',
            lineHeight: '1.6',
            fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
            flex: 1,
          }}
          showLineNumbers={true}
          lineNumberStyle={{ 
            minWidth: "3em", 
            userSelect: "none", 
            color: "#ffffff", 
            opacity: 0.4, 
            paddingRight: "1.5rem", 
            textAlign: "right",
            display: "inline-block"
          }}
          wrapLines={true}
        >
          {code}
        </SyntaxHighlighter>
      </div>
    </div>
  );
};


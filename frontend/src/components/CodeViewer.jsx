import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';

export const CodeViewer = ({ code, language = 'go' }) => {
  if (!code) {
    return (
      <div className="flex items-center justify-center h-full text-neutral-600 bg-neutral-950">
        <p className="text-sm">Select a file to view</p>
      </div>
    );
  }

  return (
    <div className="h-full flex flex-col bg-neutral-950 flex-1 min-w-0">
      <div className="flex-1 overflow-auto custom-scrollbar">
        <SyntaxHighlighter
          language={language}
          style={vscDarkPlus}
          customStyle={{
            margin: 0,
            padding: '1.5rem',
            background: 'transparent',
            fontSize: '13px',
            lineHeight: '1.7',
            fontFamily: "'JetBrains Mono', 'Fira Code', 'SF Mono', Consolas, monospace",
          }}
          showLineNumbers={true}
          lineNumberStyle={{ 
            minWidth: "3em", 
            userSelect: "none", 
            color: "rgba(255,255,255,0.25)", 
            paddingRight: "1.5rem", 
            textAlign: "right",
          }}
          wrapLines={true}
        >
          {code}
        </SyntaxHighlighter>
      </div>
    </div>
  );
};

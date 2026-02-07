import { motion, AnimatePresence } from 'framer-motion';
import { ChevronRight, Folder, File, FileCode } from 'lucide-react';
import { useState } from 'react';
import clsx from 'clsx';

const TreeItem = ({ item, onSelect, activePath, depth = 0 }) => {
  const [isOpen, setIsOpen] = useState(false);
  const isFile = item.type === 'file';
  const isActive = activePath === item.path;

  const handleClick = (e) => {
    e.stopPropagation();
    if (isFile) {
      onSelect(item);
    } else {
      setIsOpen(!isOpen);
    }
  };

  return (
    <div className="select-none text-sm">
      <div
        className={clsx(
          "flex items-center px-3 py-1.5 cursor-pointer transition-colors border-l-2",
          isActive 
            ? "border-accent bg-accent/10 text-accent" 
            : "border-transparent text-text-secondary hover:text-text-primary hover:bg-white/5"
        )}
        style={{ paddingLeft: `${depth * 16 + 12}px` }}
        onClick={handleClick}
      >
        <span className="mr-2 opacity-70 flex-shrink-0">
          {!isFile ? (
            <motion.div
              animate={{ rotate: isOpen ? 90 : 0 }}
              transition={{ duration: 0.1 }}
            >
              <ChevronRight size={14} />
            </motion.div>
          ) : (
             <div className="w-[14px]" /> 
          )}
        </span>
        
        <span className={clsx("mr-2 flex-shrink-0", isFile ? "text-accent" : "text-blue-400/50")}>
          {isFile ? <FileCode size={16} /> : <Folder size={16} fill="currentColor" />}
        </span>

        <span className="truncate">{item.name}</span>
      </div>

      <AnimatePresence initial={false}>
        {isOpen && !isFile && item.children && (
          <motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: 'auto', opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.2 }}
            className="overflow-hidden"
          >
            {item.children.map((child) => (
              <TreeItem 
                key={child.path} 
                item={child} 
                onSelect={onSelect} 
                activePath={activePath} 
                depth={depth + 1} 
              />
            ))}
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export const Sidebar = ({ data, onSelect, activePath }) => {
  return (
    <div className="h-full bg-bg-secondary border-r border-border flex flex-col">
      <div className="h-14 flex items-center px-4 border-b border-border flex-shrink-0">
        <span className="font-semibold tracking-tight text-text-primary">Design Patterns</span>
      </div>
      <div className="flex-1 overflow-y-auto custom-scrollbar py-2">
        {data.map((item) => (
          <TreeItem 
            key={item.path} 
            item={item} 
            onSelect={onSelect} 
            activePath={activePath}
          />
        ))}
      </div>
    </div>
  );
};

import { motion as Motion, AnimatePresence } from 'framer-motion';
import { ChevronRight, Folder, FileCode } from 'lucide-react';
import { useState } from 'react';
import clsx from 'clsx';

// Helper to get/set opened folders from localStorage
const STORAGE_KEY = 'go-patterns-opened-folders';

const getOpenedFolders = () => {
  try {
    const stored = localStorage.getItem(STORAGE_KEY);
    return stored ? JSON.parse(stored) : {};
  } catch {
    return {};
  }
};

const setOpenedFolder = (path, isOpen) => {
  try {
    const current = getOpenedFolders();
    if (isOpen) {
      current[path] = true;
    } else {
      delete current[path];
    }
    localStorage.setItem(STORAGE_KEY, JSON.stringify(current));
  } catch {
    // Ignore storage errors
  }
};

const TreeItem = ({ item, onSelect, activePath, depth = 0, openedFolders }) => {
  const isFile = item.type === 'file';
  const [isOpen, setIsOpen] = useState(() => {
    return !isFile && openedFolders[item.path] === true;
  });
  const isActive = activePath === item.path;

  const handleClick = (e) => {
    e.stopPropagation();
    if (isFile) {
      onSelect(item);
    } else {
      const newState = !isOpen;
      setIsOpen(newState);
      setOpenedFolder(item.path, newState);
    }
  };

  return (
    <div className="select-none">
      <div
        className={clsx(
          "flex items-center py-1 cursor-pointer transition-all duration-150",
          isActive 
            ? "text-white bg-white/[0.06]" 
            : "text-neutral-400 hover:text-neutral-200 hover:bg-white/[0.03]"
        )}
        style={{ paddingLeft: `${depth * 12 + 16}px`, paddingRight: '16px' }}
        onClick={handleClick}
      >
        <span className="mr-1.5 flex-shrink-0 text-neutral-500">
          {!isFile ? (
            <Motion.div
              animate={{ rotate: isOpen ? 90 : 0 }}
              transition={{ duration: 0.1 }}
            >
              <ChevronRight size={12} />
            </Motion.div>
          ) : (
            <div className="w-3" /> 
          )}
        </span>
        
        <span className={clsx("mr-2 flex-shrink-0", isFile ? "text-sky-400" : "text-amber-500/70")}>
          {isFile ? <FileCode size={14} /> : <Folder size={14} />}
        </span>

        <span className="truncate text-[13px]">{item.name}</span>
      </div>

      <AnimatePresence initial={false}>
        {isOpen && !isFile && item.children && (
          <Motion.div
            initial={{ height: 0, opacity: 0 }}
            animate={{ height: 'auto', opacity: 1 }}
            exit={{ height: 0, opacity: 0 }}
            transition={{ duration: 0.15 }}
            className="overflow-hidden"
          >
            {item.children.map((child) => (
              <TreeItem 
                key={child.path} 
                item={child} 
                onSelect={onSelect} 
                activePath={activePath} 
                depth={depth + 1}
                openedFolders={openedFolders}
              />
            ))}
          </Motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export const Sidebar = ({ data, onSelect, activePath }) => {
  // Initialize from localStorage synchronously to avoid effect cascade
  const [openedFolders] = useState(() => getOpenedFolders());

  return (
    <div className="h-full flex flex-col bg-black">
      <div className="h-12 flex items-center px-4 border-b border-neutral-800/50 flex-shrink-0">
        <span className="text-sm font-medium text-neutral-300">Explorer</span>
      </div>
      <div className="flex-1 overflow-y-auto custom-scrollbar py-1">
        {data.map((item) => (
          <TreeItem 
            key={item.path} 
            item={item} 
            onSelect={onSelect} 
            activePath={activePath}
            openedFolders={openedFolders}
          />
        ))}
      </div>
    </div>
  );
};

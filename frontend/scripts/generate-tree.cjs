const fs = require('fs');
const path = require('path');

const rootDir = path.resolve(__dirname, '../../design_patterns');
const outputFile = path.resolve(__dirname, '../src/data.json');

// Base relative path as requested by user
const BASE_RELATIVE_PATH = './design_patterns';

function readDir(dir, relativePath = '') {
  let items;
  try {
    items = fs.readdirSync(dir, { withFileTypes: true });
  } catch (err) {
    console.error(`Error reading directory ${dir}:`, err);
    return [];
  }

  const result = [];

  // Sort: Directories first, then files
  items.sort((a, b) => {
    if (a.isDirectory() && !b.isDirectory()) return -1;
    if (!a.isDirectory() && b.isDirectory()) return 1;
    return a.name.localeCompare(b.name);
  });

  for (const item of items) {
    // Current absolute path
    const itemPath = path.join(dir, item.name);
    
    // Construct the relative path string that the user wants to see
    // e.g. ./design_patterns/behavioral/...
    const userVisiblePath = path.join(BASE_RELATIVE_PATH, relativePath, item.name);
    
    // Skip hidden files/dirs and git
    if (item.name.startsWith('.') || item.name === 'node_modules') continue;

    if (item.isDirectory()) {
      const children = readDir(itemPath, path.join(relativePath, item.name));
      // Only include non-empty directories if desired, but for now include all
      result.push({
        name: item.name,
        type: 'directory',
        path: userVisiblePath, // This is the ID used for selection
        children: children
      });
    } else if (item.isFile() && item.name.endsWith('.go')) {
      const content = fs.readFileSync(itemPath, 'utf-8');
      result.push({
        name: item.name,
        type: 'file',
        path: userVisiblePath,
        content: content
      });
    }
  }
  return result;
}

try {
  if (!fs.existsSync(rootDir)) {
    console.error(`Directory not found: ${rootDir}`);
    // Create empty array if not found to avoid crash
    fs.writeFileSync(outputFile, JSON.stringify([], null, 2));
    process.exit(0); 
  }
  
  console.log(`Scanning ${rootDir}...`);
  const tree = readDir(rootDir);
  
  // Wrap in a root if needed, but array is fine for the recursive component
  fs.writeFileSync(outputFile, JSON.stringify(tree, null, 2));
  console.log(`Generated structure to ${outputFile}`);
} catch (error) {
  console.error('Error generating tree:', error);
  process.exit(1);
}

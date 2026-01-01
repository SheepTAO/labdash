import { useState } from 'react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { oneDark } from 'react-syntax-highlighter/dist/esm/styles/prism';

// --- Code Block Component with Copy ---
const CodeBlock = ({ node, inline, className, children, ...props }) => {
  const [copied, setCopied] = useState(false);
  const match = /language-(\w+)/.exec(className || '');
  const language = match ? match[1] : 'text';

  const handleCopy = (e) => {
    e.stopPropagation();
    navigator.clipboard.writeText(String(children).replace(/\n$/, ''));
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  if (!inline && match) {
    return (
      <div className="relative group my-4 rounded-xl shadow-sm border border-slate-200/60 max-w-full bg-[#282c34] overflow-hidden">
        {/* Floating Language Label / Copy Button */}
        <div className="absolute right-2.5 top-2.5 z-20">
            <button
                onClick={handleCopy}
                className={`
                  group/btn relative flex items-center justify-center px-2.5 py-1.5 rounded-lg transition-all duration-200 border backdrop-blur-md
                  ${copied 
                    ? 'bg-emerald-500/20 border-emerald-500/30 text-emerald-400' 
                    : 'bg-white/5 border-white/10 text-slate-400 hover:bg-white/10 hover:text-slate-200'
                  }
                `}
            >
                <span className="text-xs font-bold uppercase tracking-wider font-mono">
                  {copied ? 'Copied' : language}
                </span>
                
                {/* Tooltip */}
                {!copied && (
                   <div className="absolute top-full right-0 mt-2 px-2 py-1 bg-black/80 text-white text-[10px] rounded opacity-0 group-hover/btn:opacity-100 pointer-events-none whitespace-nowrap transition-opacity backdrop-blur-sm shadow-lg">
                      Click to Copy
                   </div>
                )}
            </button>
        </div>

        <div className="overflow-x-auto">
          <SyntaxHighlighter
            {...props}
            style={oneDark}
            language={language}
            PreTag="div"
            customStyle={{
              margin: 0,
              padding: '0.75rem 1rem 0.75rem 1rem', // Minimal padding
              borderRadius: '0',
              background: 'transparent',
              fontSize: '1.125rem', // Matches prose-lg (18px)
              lineHeight: '1.4',
            }}
            codeTagProps={{
              style: { 
                fontFamily: "'JetBrains Mono', 'Fira Code', Consolas, monospace",
                color: '#abb2bf'
              }
            }}
          >
            {String(children).replace(/\n$/, '')}
          </SyntaxHighlighter>
        </div>
      </div>
    );
  }

  return (
    <code className={className} style={{ backgroundColor: '#e2e8f0', color: '#1e293b', padding: '0.15rem 0.25rem', borderRadius: '0.375rem', border: '1px solid #cbd5e1', fontWeight: '500', fontSize: '0.9em' }} {...props}>
      {children}
    </code>
  );
};

export default CodeBlock;

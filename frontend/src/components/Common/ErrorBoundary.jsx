import { Component } from 'react';
import { AlertTriangle } from 'lucide-react';

/**
 * Error Boundary - Catches JavaScript errors in child components
 * Prevents white screen errors and provides fallback UI
 * 
 * Usage:
 * <ErrorBoundary>
 *   <YourComponent />
 * </ErrorBoundary>
 */
class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error) {
    // Update state when error occurs
    return { hasError: true, error };
  }

  componentDidCatch(error, errorInfo) {
    // Log error to console (could send to logging service)
    console.error('ErrorBoundary caught:', error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      // Fallback UI when error occurs
      return (
        <div className="min-h-screen bg-gray-900 flex items-center justify-center p-4">
          <div className="bg-gray-800 border border-red-500/30 rounded-lg p-8 max-w-lg w-full">
            <div className="flex items-center gap-3 mb-4">
              <AlertTriangle className="text-red-500" size={32} />
              <h1 className="text-2xl font-bold text-red-500">Something Went Wrong</h1>
            </div>
            <p className="text-gray-300 mb-4">
              The application encountered an error. Please try refreshing the page.
            </p>
            <div className="bg-gray-900 p-4 rounded text-sm text-gray-400 font-mono mb-4 overflow-auto">
              {this.state.error?.toString()}
            </div>
            <button
              onClick={() => window.location.reload()}
              className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded transition"
            >
              Reload Page
            </button>
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;

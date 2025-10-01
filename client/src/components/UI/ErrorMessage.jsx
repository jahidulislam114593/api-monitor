import React from "react";

const ErrorMessage = ({ message, onClose }) => {
  if (!message) return null;

  return (
    <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-6">
      {message}
      {onClose && (
        <button
          onClick={onClose}
          className="float-right font-bold hover:text-red-900"
        >
          ×
        </button>
      )}
    </div>
  );
};

export default ErrorMessage;

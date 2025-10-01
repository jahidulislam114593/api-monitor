import React from "react";
import { Play, Trash2, CheckCircle, XCircle } from "lucide-react";

const EndpointCard = ({
  endpoint,
  latestCheck,
  onRunCheck,
  onDelete,
  loading,
}) => {
  const getStatusColor = (isUp) => (isUp ? "text-green-600" : "text-red-600");

  return (
    <div className="border border-gray-200 rounded-lg p-4 hover:shadow-md transition-shadow">
      <div className="flex justify-between items-start">
        <div className="flex-1">
          <div className="flex items-center gap-3 mb-2">
            <span className="font-medium text-lg">{endpoint.name}</span>
            <span
              className={`px-2 py-1 text-xs rounded-full ${
                endpoint.is_active
                  ? "bg-green-100 text-green-800"
                  : "bg-gray-100 text-gray-800"
              }`}
            >
              {endpoint.is_active ? "Active" : "Inactive"}
            </span>
            {latestCheck && (
              <span
                className={`flex items-center gap-1 text-sm ${getStatusColor(
                  latestCheck.is_up
                )}`}
              >
                {latestCheck.is_up ? (
                  <CheckCircle size={14} />
                ) : (
                  <XCircle size={14} />
                )}
                {latestCheck.is_up ? "UP" : "DOWN"}
              </span>
            )}
          </div>

          <div className="text-gray-600 mb-2">
            <span className="font-mono text-sm bg-gray-100 px-2 py-1 rounded">
              {endpoint.method}
            </span>
            <span className="ml-2">{endpoint.url}</span>
          </div>

          {endpoint.description && (
            <div className="text-gray-500 text-sm mb-2">
              {endpoint.description}
            </div>
          )}

          {latestCheck && (
            <div className="flex gap-4 text-sm text-gray-600">
              <span>Status: {latestCheck.status}</span>
              <span>Response: {latestCheck.response_time}ms</span>
              <span>Last Check: {latestCheck.checked_at}</span>
            </div>
          )}
        </div>

        <div className="flex gap-2">
          <button
            onClick={() => onRunCheck(endpoint.id)}
            disabled={loading}
            className="text-green-600 hover:text-green-800 p-1 disabled:opacity-50"
            title="Run check"
          >
            <Play size={16} />
          </button>
          <button
            onClick={() => onDelete(endpoint.id)}
            className="text-red-600 hover:text-red-800 p-1"
            title="Delete endpoint"
          >
            <Trash2 size={16} />
          </button>
        </div>
      </div>
    </div>
  );
};

export default EndpointCard;

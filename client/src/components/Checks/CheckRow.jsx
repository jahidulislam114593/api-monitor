import React from "react";
import { CheckCircle, XCircle } from "lucide-react";

const CheckRow = ({ check }) => {
  const getStatusColor = (isUp) => (isUp ? "text-green-600" : "text-red-600");
  const getStatusBg = (isUp) => (isUp ? "bg-green-50" : "bg-red-50");

  const getResponseTimeColor = (responseTime) => {
    if (responseTime < 200) return "bg-green-100 text-green-800";
    if (responseTime < 500) return "bg-yellow-100 text-yellow-800";
    return "bg-red-100 text-red-800";
  };

  const getStatusCodeColor = (status) => {
    if (status >= 200 && status < 300) return "bg-green-100 text-green-800";
    if (status >= 300 && status < 400) return "bg-blue-100 text-blue-800";
    if (status >= 400) return "bg-red-100 text-red-800";
    return "bg-gray-100 text-gray-800";
  };

  return (
    <tr className={`border-t ${getStatusBg(check.is_up)}`}>
      <td className="p-3">
        <div
          className={`flex items-center gap-2 ${getStatusColor(check.is_up)}`}
        >
          {check.is_up ? <CheckCircle size={16} /> : <XCircle size={16} />}
          <span className="font-medium">{check.is_up ? "UP" : "DOWN"}</span>
        </div>
      </td>
      <td className="p-3">
        <div className="font-mono text-sm max-w-xs truncate">{check.url}</div>
      </td>
      <td className="p-3">
        <span
          className={`px-2 py-1 rounded text-sm ${getResponseTimeColor(
            check.response_time
          )}`}
        >
          {check.response_time}ms
        </span>
      </td>
      <td className="p-3">
        <span
          className={`px-2 py-1 rounded text-sm ${getStatusCodeColor(
            check.status
          )}`}
        >
          {check.status || "N/A"}
        </span>
      </td>
      <td className="p-3 text-sm text-gray-600">{check.checked_at}</td>
    </tr>
  );
};

export default CheckRow;

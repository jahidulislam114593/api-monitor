import React from "react";
import CheckRow from "./CheckRow";

const ChecksTable = ({ checks }) => {
  if (checks.length === 0) {
    return (
      <div className="text-center py-8 text-gray-500">
        No checks performed yet. Run a check to see results!
      </div>
    );
  }

  return (
    <div>
      <h3 className="text-lg font-semibold mb-4">
        Recent Checks ({checks.length})
      </h3>

      <div className="overflow-x-auto">
        <table className="w-full">
          <thead>
            <tr className="bg-gray-50">
              <th className="text-left p-3 font-medium">Status</th>
              <th className="text-left p-3 font-medium">URL</th>
              <th className="text-left p-3 font-medium">Response Time</th>
              <th className="text-left p-3 font-medium">HTTP Status</th>
              <th className="text-left p-3 font-medium">Checked At</th>
            </tr>
          </thead>
          <tbody>
            {checks
              .sort((a, b) => new Date(b.checked_at) - new Date(a.checked_at))
              .slice(0, 20)
              .map((check) => (
                <CheckRow key={check.id} check={check} />
              ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ChecksTable;

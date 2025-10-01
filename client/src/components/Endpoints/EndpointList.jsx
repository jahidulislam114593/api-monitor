import React from "react";
import EndpointCard from "./EndpointCard";

const EndpointList = ({ endpoints, checks, onRunCheck, onDelete, loading }) => {
  if (endpoints.length === 0) {
    return (
      <div className="text-center py-8 text-gray-500">
        No endpoints configured. Add one to start monitoring!
      </div>
    );
  }

  return (
    <div className="mb-8">
      <h3 className="text-lg font-semibold mb-4">
        Monitored Endpoints ({endpoints.length})
      </h3>

      <div className="grid gap-4">
        {endpoints.map((endpoint) => {
          const latestCheck = checks
            .filter((check) => check.url === endpoint.url)
            .sort((a, b) => new Date(b.checked_at) - new Date(a.checked_at))[0];

          return (
            <EndpointCard
              key={endpoint.id}
              endpoint={endpoint}
              latestCheck={latestCheck}
              onRunCheck={onRunCheck}
              onDelete={onDelete}
              loading={loading}
            />
          );
        })}
      </div>
    </div>
  );
};

export default EndpointList;

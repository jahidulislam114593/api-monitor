import React, { useState } from "react";
import { Plus, Play } from "lucide-react";

import Dashboard from "./components/Dashboard/Dashboard";
import EndpointForm from "./components/Endpoints/EndpointForm";
import EndpointList from "./components/Endpoints/EndpointList";
import ChecksTable from "./components/Checks/ChecksTable";
import Button from "./components/UI/Button";
import ErrorMessage from "./components/UI/ErrorMessage";

import { useAPIMonitor } from "./hooks/useAPIMonitor";

function App() {
  const [showAddForm, setShowAddForm] = useState(false);

  const {
    endpoints,
    checks,
    stats,
    loading,
    error,
    addEndpoint,
    runCheck,
    runAllChecks,
    deleteEndpoint,
    clearError,
  } = useAPIMonitor();

  const handleAddEndpoint = async (endpointData) => {
    const success = await addEndpoint(endpointData);
    if (success) {
      setShowAddForm(false);
    }
  };

  return (
    <div className="max-w-6xl mx-auto p-6">
      <div className="bg-white rounded-lg shadow-lg">
        <Dashboard stats={stats} />

        <div className="p-6">
          {/* Action Buttons */}
          <div className="flex gap-4 mb-6">
            <Button
              onClick={() => setShowAddForm(!showAddForm)}
              variant="primary"
            >
              <Plus size={20} />
              Add Endpoint
            </Button>
            <Button onClick={runAllChecks} disabled={loading} variant="success">
              <Play size={20} />
              {loading ? "Running..." : "Check All"}
            </Button>
          </div>

          {/* Error Display */}
          <ErrorMessage message={error} onClose={clearError} />

          {/* Add Endpoint Form */}
          {showAddForm && (
            <EndpointForm
              onSubmit={handleAddEndpoint}
              onCancel={() => setShowAddForm(false)}
              loading={loading}
            />
          )}

          {/* Endpoints List */}
          <EndpointList
            endpoints={endpoints}
            checks={checks}
            onRunCheck={runCheck}
            onDelete={deleteEndpoint}
            loading={loading}
          />

          {/* Checks Table */}
          <ChecksTable checks={checks} />
        </div>
      </div>
    </div>
  );
}

export default App;

import { useState, useEffect, useCallback } from "react";
import { apiService } from "../services/api";

export const useAPIMonitor = () => {
  const [endpoints, setEndpoints] = useState([]);
  const [checks, setChecks] = useState([]);
  const [stats, setStats] = useState({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const fetchEndpoints = useCallback(async () => {
    try {
      const data = await apiService.getEndpoints();
      setEndpoints(data || []);
    } catch (err) {
      setError(`Failed to fetch endpoints: ${err.message}`);
    }
  }, []);

  const fetchChecks = useCallback(async (limit = 50) => {
    try {
      const data = await apiService.getChecks(limit);
      setChecks(data || []);
    } catch (err) {
      setError(`Failed to fetch checks: ${err.message}`);
    }
  }, []);

  const fetchStats = useCallback(async () => {
    try {
      const data = await apiService.getStats();
      setStats(data || {});
    } catch (err) {
      setError(`Failed to fetch stats: ${err.message}`);
    }
  }, []);

  const addEndpoint = async (endpointData) => {
    setLoading(true);
    try {
      await apiService.createEndpoint(endpointData);
      await fetchEndpoints();
      return true;
    } catch (err) {
      setError(`Failed to add endpoint: ${err.message}`);
      return false;
    } finally {
      setLoading(false);
    }
  };

  const runCheck = async (endpointId) => {
    setLoading(true);
    setError("");
    try {
      await apiService.runCheck(endpointId);
      await fetchChecks();
      await fetchStats();
    } catch (err) {
      setError(`Failed to run check: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  const runAllChecks = async () => {
    setLoading(true);
    setError("");
    try {
      await apiService.runAllChecks();
      await fetchChecks();
      await fetchStats();
    } catch (err) {
      setError(`Failed to run all checks: ${err.message}`);
    } finally {
      setLoading(false);
    }
  };

  const deleteEndpoint = async (id) => {
    if (!window.confirm("Are you sure you want to delete this endpoint?")) {
      return false;
    }

    setLoading(true);
    try {
      await apiService.deleteEndpoint(id);
      await fetchEndpoints();
      return true;
    } catch (err) {
      setError(`Failed to delete endpoint: ${err.message}`);
      return false;
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEndpoints();
    fetchChecks();
    fetchStats();
  }, [fetchEndpoints, fetchChecks, fetchStats]);

  return {
    endpoints,
    checks,
    stats,
    loading,
    error,
    addEndpoint,
    runCheck,
    runAllChecks,
    deleteEndpoint,
    refreshData: () => {
      fetchEndpoints();
      fetchChecks();
      fetchStats();
    },
    clearError: () => setError(""),
  };
};

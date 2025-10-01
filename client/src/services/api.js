const API_BASE_URL = "http://localhost:8080/api";

class APIService {
  async request(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const config = {
      headers: {
        "Content-Type": "application/json",
        ...options.headers,
      },
      ...options,
    };

    try {
      const response = await fetch(url, config);

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const contentType = response.headers.get("content-type");
      if (contentType && contentType.includes("application/json")) {
        return await response.json();
      }

      return null;
    } catch (error) {
      console.error("API Request failed:", error);
      throw error;
    }
  }

  async getEndpoints() {
    return this.request("/endpoints");
  }

  async createEndpoint(endpointData) {
    return this.request("/endpoints", {
      method: "POST",
      body: JSON.stringify(endpointData),
    });
  }

  async updateEndpoint(id, endpointData) {
    return this.request(`/endpoints/${id}`, {
      method: "PUT",
      body: JSON.stringify(endpointData),
    });
  }

  async deleteEndpoint(id) {
    return this.request(`/endpoints/${id}`, {
      method: "DELETE",
    });
  }

  async getChecks(limit) {
    const params = limit ? `?limit=${limit}` : "";
    return this.request(`/checks${params}`);
  }

  async runCheck(endpointId) {
    return this.request(`/checks/run/${endpointId}`, {
      method: "POST",
    });
  }

  async runAllChecks() {
    return this.request("/checks/run-all", {
      method: "POST",
    });
  }

  async getStats() {
    return this.request("/stats");
  }
}

export const apiService = new APIService();

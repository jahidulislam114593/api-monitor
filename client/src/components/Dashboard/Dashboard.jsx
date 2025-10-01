import React from "react";
import { Activity, BarChart3, CheckCircle, Clock } from "lucide-react";
import StatCard from "./StatCard";

const Dashboard = ({ stats }) => {
  const statCards = [
    {
      icon: BarChart3,
      title: "Endpoints",
      value: stats.total_endpoints || 0,
      bgColor: "bg-blue-50",
      textColor: "text-blue-600",
      iconColor: "text-blue-600",
    },
    {
      icon: CheckCircle,
      title: "Uptime",
      value: stats.uptime_percentage || "0%",
      bgColor: "bg-green-50",
      textColor: "text-green-600",
      iconColor: "text-green-600",
    },
    {
      icon: Clock,
      title: "Avg Response",
      value: stats.avg_response_time || "0 ms",
      bgColor: "bg-purple-50",
      textColor: "text-purple-600",
      iconColor: "text-purple-600",
    },
    {
      icon: Activity,
      title: "Total Checks",
      value: stats.total_checks || 0,
      bgColor: "bg-orange-50",
      textColor: "text-orange-600",
      iconColor: "text-orange-600",
    },
  ];

  return (
    <div className="p-6 border-b border-gray-200">
      <div className="flex items-center gap-3 mb-4">
        <Activity className="text-blue-600" size={32} />
        <div>
          <h1 className="text-2xl font-bold text-gray-800">
            API Response Time Monitor
          </h1>
          <p className="text-gray-600">
            Monitor your API endpoints and track response times
          </p>
        </div>
      </div>

      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        {statCards.map((card, index) => (
          <StatCard key={index} {...card} />
        ))}
      </div>
    </div>
  );
};

export default Dashboard;

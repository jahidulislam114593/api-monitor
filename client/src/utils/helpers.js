export const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString() + " " + date.toLocaleTimeString();
};

export const getStatusColor = (isUp) => {
  return isUp ? "text-green-600" : "text-red-600";
};

export const getStatusBadge = (isUp) => {
  return isUp ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800";
};

export const getResponseTimeColor = (responseTime) => {
  if (responseTime < 200) return "bg-green-100 text-green-800";
  if (responseTime < 500) return "bg-yellow-100 text-yellow-800";
  return "bg-red-100 text-red-800";
};

export const truncateUrl = (url, maxLength = 50) => {
  return url.length > maxLength ? url.substring(0, maxLength) + "..." : url;
};
